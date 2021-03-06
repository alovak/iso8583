package encoding

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestASCII(t *testing.T) {
	enc := &asciiEncoder{}

	t.Run("Decode", func(t *testing.T) {
		res, err := enc.Decode([]byte("hello"), 0)

		require.NoError(t, err)
		require.Equal(t, []byte("hello"), res)

		_, err = enc.Decode([]byte("hello, 世界!"), 0)
		require.Error(t, err)
	})

	t.Run("Encode", func(t *testing.T) {
		res, err := enc.Encode([]byte("hello"))

		require.NoError(t, err)
		require.Equal(t, []byte("hello"), res)

		res, err = enc.Encode([]byte("hello, 世界!"))
		require.Error(t, err)
	})
}
