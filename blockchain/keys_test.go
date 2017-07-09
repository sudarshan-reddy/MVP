package blockchain

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PadBytes(t *testing.T) {
	assert := assert.New(t)
	var testData = []struct {
		length     int
		char       byte
		expectedOp []byte
	}{
		{4, byte('x'), []byte("xxxx")},
	}
	for _, each := range testData {
		actualOp := padBytes(each.length, each.char)
		assert.Equal(each.expectedOp, actualOp)
	}
}

func Test_SerializeWithLength(t *testing.T) {
	assert := assert.New(t)
	var testData = []struct {
		description string
		length      int
		val1        *big.Int
		val2        *big.Int
		expectedOp  []byte
	}{
		{"basic positive case",
			3,
			new(big.Int).SetBytes([]byte("xyz")),
			new(big.Int).SetBytes([]byte("abc")),
			[]byte("xyzabc"),
		},
		{"if length is greater",
			5,
			new(big.Int).SetBytes([]byte("xyz")),
			new(big.Int).SetBytes([]byte("abc")),
			[]byte("xyz\x00\x00abc"),
		},
	}
	for _, each := range testData {
		actualOp := serializeWithLength(each.length, each.val1, each.val2)
		assert.Equal(string(each.expectedOp), string(actualOp))
	}
}
