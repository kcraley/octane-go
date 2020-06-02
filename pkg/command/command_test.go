package command_test

import (
	"reflect"
	"testing"

	"github.com/kcraley/octane-go/pkg/command"
)

var (
	routerTestCases = map[string]struct {
		prefix   string
		expected *command.Router
	}{
		"defaultPrefix": {prefix: "!octane", expected: &command.Router{Prefix: "!octane"}},
		"customPrefix":  {prefix: "!custom", expected: &command.Router{Prefix: "!custom"}},
	}
)

func TestNewRouter(t *testing.T) {
	for name, test := range routerTestCases {
		t.Run(name, func(t *testing.T) {
			actual := command.NewRouter(test.prefix)
			if !reflect.DeepEqual(test.expected, actual) {
				t.Fatalf("Expected: %v, Actual: %v", test.expected, actual)
			}
		})
	}
}
