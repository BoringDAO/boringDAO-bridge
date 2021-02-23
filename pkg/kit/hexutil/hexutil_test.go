package hexutil

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncode(t *testing.T) {
	cases := []struct {
		input  []byte
		output string
	}{
		{
			input:  []byte("112233445566"),
			output: "0x313132323333343435353636",
		},
		{
			input:  []byte("667788990022"),
			output: "0x363637373838393930303232",
		},
	}

	for _, c := range cases {
		require.Equal(t, c.output, Encode(c.input))
	}
}

func TestDecode(t *testing.T) {
	cases := []struct {
		input  string
		output []byte
	}{
		{
			input:  "0x313132323333343435353636",
			output: []byte("112233445566"),
		},
		{
			input:  "313132323333343435353636",
			output: []byte("112233445566"),
		},
	}

	for _, c := range cases {
		require.Equal(t, c.output, Decode(c.input))
	}
}
