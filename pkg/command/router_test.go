package command_test

import (
	"testing"

	"github.com/kcraley/octane-go/pkg/command"

	"github.com/google/go-cmp/cmp"
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
			diff := cmp.Diff(test.expected, actual)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
