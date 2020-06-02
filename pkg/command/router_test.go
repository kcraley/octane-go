package command_test

import (
	"testing"

	"github.com/kcraley/octane-go/pkg/command"

	"github.com/google/go-cmp/cmp"
)

func TestNewRouter(t *testing.T) {
	// Test cases for creating a new Router
	var routerTestCases = map[string]struct {
		prefix   string
		expected *command.Router
	}{
		"defaultPrefix": {prefix: "!octane", expected: &command.Router{Prefix: "!octane"}},
		"customPrefix":  {prefix: "!custom", expected: &command.Router{Prefix: "!custom"}},
	}

	for name, test := range routerTestCases {
		t.Run(name, func(t *testing.T) {
			actual := command.NewRouter(test.prefix)
			diff := cmp.Diff(test.expected, actual)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestRegisterCommand(t *testing.T) {
	// Test cases for registering a new Command
	var registerCommandTestCases = map[string]struct {
		prefix     string
		newCommand *command.Command
		expected   *command.Router
	}{
		"emptyRouter": {
			prefix:     "!octane",
			newCommand: &command.Command{},
			expected: &command.Router{
				Prefix:   "!octane",
				Commands: []*command.Command{{}},
			},
		},
	}

	for name, test := range registerCommandTestCases {
		t.Run(name, func(t *testing.T) {
			actual := command.NewRouter(test.prefix)
			actual.RegisterCommand(test.newCommand)
			diff := cmp.Diff(test.expected, actual)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
