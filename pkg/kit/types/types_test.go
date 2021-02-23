package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHash_String(t *testing.T) {
	b := Bytes2Hash([]byte("0002288331189898"))
	require.Equal(t, "000000...383938", b.ShortString())
}

func TestString2Hash(t *testing.T) {
	s := "0x9f41dd84524bf8a42f8ab58ecfca6e1752d6fd93fe8dc00af4c71963c97db59f"
	require.Equal(t, "9f41dd...7db59f", String2Hash(s).ShortString())
}
