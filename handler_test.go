package kvstore

import (
	"testing"

	"github.com/golang-collections/collections/stack"
)

func TestHandle(t *testing.T) {
	var tests = []struct {
		cmds     []string
		expected string
	}{
		{
			cmds: []string{
				"write k v",
				"start",
				"delete k",
				"commit",
				"read k",
			},
			expected: "",
		},
		{
			cmds: []string{
				"write k v1",
				"start",
				"write k v2",
				"abort",
				"read k",
			},
			expected: "v1",
		},
		{
			cmds: []string{
				"write k v1",
				"start",
				"write k v2",
				"commit",
				"read k",
			},
			expected: "v2",
		},
		{
			cmds: []string{
				"write k v1",
				"start",
				"write k v2",
				"start",
				"delete k",
				"commit",
				"write k v3",
				"start",
				"write k v4",
				"abort",
				"read k",
			},
			expected: "v3",
		},
		{
			cmds: []string{
				"write k v1",
				"start",
				"write k v2",
				"start",
				"delete k",
				"write k v3",
				"commit",
				"read k",
			},
			expected: "v3",
		},
		{
			cmds: []string{
				"write k v1",
				"start",
				"write k v2",
				"start",
				"delete k",
				"commit",
				"write k v3",
				"abort",
				"read k",
			},
			expected: "v1",
		},
		{
			cmds: []string{
				"write k v1",
				"start",
				"write k v2",
				"start",
				"delete k",
				"write k v3",
				"abort",
				"read k",
			},
			expected: "v2",
		},
	}

	for _, test := range tests {
		actual := ""
		store := Store{
			KVStack: stack.New(),
			Depth:   0,
		}
		kvlayer := map[string]string{}
		store.KVStack.Push(kvlayer)
		for _, c := range test.cmds {
			// fmt.Println(c, " \nDepth: ", store.Depth)
			actual, _ = Handle(c, &store)
		}
		if actual != test.expected {
			t.Errorf(
				"failed to handle:\n\texpected: %v\n\t  actual: %v",
				test.expected,
				actual,
			)
		}
	}
}
