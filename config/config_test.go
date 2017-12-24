package config_test

import (
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/renstrom/dedent"
	"gitlab.com/mr-meeseeks/meeseeks-box/config"
)

func Test_ConfigurationReading(t *testing.T) {
	tt := []struct {
		Name     string
		Content  string
		Expected config.Config
	}{
		{
			"Basic",
			"",
			config.Config{},
		},
		{
			"With messages",
			dedent.Dedent(`
				messages:
				  handshake: ["hallo"]
				`),
			config.Config{
				Messages: map[string][]string{
					"handshake": []string{"hallo"},
				},
			},
		},
		{
			"With messages",
			dedent.Dedent(`
				commands:
				  something:
				    command: "ssh"
				    authorized: ["someone"]
				    arguments: ["none"]
				    action: select
				    options:
				      message: pick
				      values: 
				        first: "one"
				        second: "two"
				`),
			config.Config{
				Commands: map[string]config.Command{
					"something": config.Command{
						Cmd:        "ssh",
						Authorized: []string{"someone"},
						Args:       []string{"none"},
						Action:     config.ActionSelect,
						Options: config.CommandOptions{
							Message: "pick",
							Values: map[string]string{
								"first":  "one",
								"second": "two",
							},
						},
					},
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			c, err := config.New(strings.NewReader(tc.Content))
			if err != nil {
				t.Fatalf("failed to parse configuration: %s", err)
			}
			if !reflect.DeepEqual(tc.Expected, c) {
				t.Fatalf("configuration is not as expected; got %+v instead of %+v", c, tc.Expected)
			}

		})
	}
}

func Test_Errors(t *testing.T) {
	tc := []struct {
		name     string
		reader   io.Reader
		expected string
	}{
		{
			"invalid configuration",
			strings.NewReader("	invalid"),
			"could not parse configuration: yaml: found character that cannot start any token",
		},
		{
			"bad reader",
			badReader{},
			"could not read configuration: bad reader",
		},
	}
	for _, tc := range tc {
		t.Run(tc.name, func(t *testing.T) {
			_, err := config.New(tc.reader)
			if err.Error() != tc.expected {
				t.Fatalf("wrong error, expected %s; got %s", tc.expected, err)
			}
		})
	}
}

type badReader struct {
}

func (badReader) Read(b []byte) (n int, err error) {
	return 0, fmt.Errorf("bad reader")
}