package yaml_test

import (
	"bytes"
	"testing"

	"github.com/VojtechVitek/yaml/pkg/cli"
	"github.com/google/go-cmp/cmp"
)

func TestGrep(t *testing.T) {
	tt := []struct {
		in  []byte
		cmd []string
		out []byte
	}{
		{
			in:  fooYAML,
			cmd: []string{"yaml", "grep", "foo: bar"},
			out: []byte(`foo: bar
`),
		},
		{
			in:  fooYAML,
			cmd: []string{"yaml", "grep", "foo: bar", "foo: baz"},
			out: []byte(`foo: bar
---
foo: baz
`),
		},
		{
			in:  fooYAML,
			cmd: []string{"yaml", "grep", "doesnt: exist"},
			out: []byte(``),
		},
	}

	for i, tc := range tt {
		var b bytes.Buffer

		if err := cli.Run(&b, bytes.NewReader(tc.in), tc.cmd); err != nil {
			t.Error(err)
		}

		if diff := cmp.Diff(tc.out, b.Bytes()); diff != "" {
			t.Errorf("tc[%v] mismatch (-want +got):\n%s", i, diff)
		}
	}
}

var fooYAML = []byte(`---
foo: bar
---
foo: baz
--
# empty
--
no-foo: nope
--
foo: nope
--
foo:
  - bar
--
foo:
  - bar
  - baz
--
foo:
  bar: baz
`)
