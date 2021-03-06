package jobs

import (
	"encoding/json"
	"fmt"
	"time"

	bolt "github.com/coreos/bbolt"
	"github.com/gomeeseeks/meeseeks-box/meeseeks"
	"github.com/gomeeseeks/meeseeks-box/meeseeks/metrics"
	"github.com/gomeeseeks/meeseeks-box/persistence/db"
	"github.com/sirupsen/logrus"
)

var jobsBucketKey = []byte("jobs")
var runningJobsBucketKey = []byte("running-jobs")

// Jobs creates a new Jobs object
type Jobs struct{}

// Get returns an existing job by id
func (Jobs) Get(id uint64) (meeseeks.Job, error) {
	return get(id)
}

// Null returns a null job that will not be tracked
func (Jobs) Null(r meeseeks.Request) meeseeks.Job {
	return null(r)
}

// Create records a request in the DB and hands off a new job
func (Jobs) Create(r meeseeks.Request) (meeseeks.Job, error) {
	return create(r)
}

// Fail accounds for the job ending and sets the status.
func (Jobs) Fail(jobID uint64) error {
	return finish(jobID, meeseeks.JobFailedStatus)
}

// Succeed accounds for the job ending and sets the status.
func (Jobs) Succeed(jobID uint64) error {
	return finish(jobID, meeseeks.JobSuccessStatus)
}

// FailRunningJobs flags as failed any jobs that is still in running state
func (Jobs) FailRunningJobs() error {
	return failRunningJobs()
}

// Find will walk through the values on the jobs bucket and will apply the Match function
// to determine if the job matches a search criteria.
//
// Returns a list of jobs in descending order that match the filter
func (Jobs) Find(filter meeseeks.JobFilter) ([]meeseeks.Job, error) {
	return find(filter)
}

func null(req meeseeks.Request) meeseeks.Job {
	return meeseeks.Job{
		ID:        0,
		Request:   req,
		StartTime: time.Now().UTC(),
		Status:    meeseeks.JobRunningStatus,
	}
}

func create(req meeseeks.Request) (meeseeks.Job, error) {
	var job *meeseeks.Job
	err := db.Update(func(tx *bolt.Tx) error {
		jobID, bucket, err := db.NextSequenceFor(jobsBucketKey, tx)
		if err != nil {
			return fmt.Errorf("could not get next sequence for %s: %s", string(jobsBucketKey), err)
		}

		job = &meeseeks.Job{
			ID:        jobID,
			Request:   req,
			StartTime: time.Now().UTC(),
			Status:    meeseeks.JobRunningStatus,
		}
		logrus.Debugf("Creating job %#v", job)

		runningJobsBucket, err := tx.CreateBucketIfNotExists(runningJobsBucketKey)
		if err != nil {
			return fmt.Errorf("could not create running jobs bucket: %s", err)
		}
		if err = runningJobsBucket.Put(db.IDToBytes(job.ID), []byte(meeseeks.JobRunningStatus)); err != nil {
			return fmt.Errorf("could not save running job ID %d: %s", jobID, err)
		}

		return save(*job, bucket)
	})
	if err != nil {
		return meeseeks.Job{}, fmt.Errorf("failed to create a job %s", err)
	}
	return *job, nil
}

func get(id uint64) (meeseeks.Job, error) {
	job := &meeseeks.Job{}
	err := db.View(func(tx *bolt.Tx) error {
		jobsBucket := tx.Bucket(jobsBucketKey)
		if jobsBucket == nil {
			return meeseeks.ErrNoJobWithID
		}
		payload := jobsBucket.Get(db.IDToBytes(id))
		if payload == nil {
			return meeseeks.ErrNoJobWithID
		}
		return json.Unmarshal(payload, job)
	})
	logrus.Debugf("Returning job %#v for ID %d, err: %s", *job, id, err)
	return *job, err
}

// Finish sets the status of a job to whatever end state if it's current status is running
//
// It also sets the end time of the job
func finish(jobID uint64, status string) error {
	if !(status == meeseeks.JobSuccessStatus || status == meeseeks.JobFailedStatus) {
		return fmt.Errorf("invalid status %s", status)
	}
	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(jobsBucketKey)
		job, err := get(jobID)
		if err != nil {
			return fmt.Errorf("could not get job with id %d: %s", jobID, err)
		}
		if job.Status != meeseeks.JobRunningStatus {
			return fmt.Errorf("job is not in running status but %s", job.Status)
		}
		runningJobsBucket := tx.Bucket(runningJobsBucketKey)
		if err = runningJobsBucket.Delete(db.IDToBytes(jobID)); err != nil {
			return fmt.Errorf("could not remove job %d from running list: %s", jobID, err)
		}

		job.EndTime = time.Now().UTC()
		job.Status = status

		difference := job.EndTime.Sub(job.StartTime)
		metrics.TaskDurations.WithLabelValues(job.Request.Command, status).Observe(difference.Seconds())

		return save(job, bucket)
	})
}

func find(filter meeseeks.JobFilter) ([]meeseeks.Job, error) {
	latest := make([]meeseeks.Job, 0)
	matcher := func(job meeseeks.Job) bool {
		return true
	}
	if filter.Match != nil {
		matcher = filter.Match
	}
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(jobsBucketKey)
		if bucket == nil {
			return nil
		}
		cur := bucket.Cursor()
		_, payload := cur.Last()
		for len(latest) < filter.Limit {
			if payload == nil {
				break
			}

			job := meeseeks.Job{}
			if err := json.Unmarshal(payload, &job); err != nil {
				return fmt.Errorf("failed to load Job payload %s", err)
			}
			if matcher(job) {
				latest = append(latest, job)
			}
			_, payload = cur.Prev()
		}
		return nil
	})
	return latest, err
}

func failRunningJobs() error {
	return db.Update(func(tx *bolt.Tx) error {
		runningJobsBucket := tx.Bucket(runningJobsBucketKey)
		if runningJobsBucket == nil {
			return nil
		}

		jobsBucket := tx.Bucket(jobsBucketKey)
		if jobsBucket == nil {
			return nil
		}

		c := runningJobsBucket.Cursor()
		jobIDKey, _ := c.First()
		for {
			if jobIDKey == nil {
				break
			}
			jobID := db.IDFromBytes(jobIDKey)
			logrus.Warnf("Found job %d in running state, marking as killed", jobID)

			j := meeseeks.Job{}
			if err := json.Unmarshal(jobsBucket.Get(jobIDKey), &j); err != nil {
				return fmt.Errorf("could not read job %d from bucket: %s", jobID, err)
			}

			j.Status = meeseeks.JobKilledStatus
			j.EndTime = time.Now().UTC()
			if err := save(j, jobsBucket); err != nil {
				return fmt.Errorf("could not save killed job %d: %s", jobID, err)
			}

			if err := runningJobsBucket.Delete(jobIDKey); err != nil {
				return fmt.Errorf("could not delete running job %d: %s", jobID, err)
			}

			jobIDKey, _ = c.Next()
		}
		return nil
	})
}

func save(job meeseeks.Job, bucket *bolt.Bucket) error {
	buffer, err := json.Marshal(job)
	if err != nil {
		return err
	}
	return bucket.Put(db.IDToBytes(job.ID), buffer)
}
