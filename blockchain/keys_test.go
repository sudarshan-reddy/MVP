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
		{"if length is lesser",
			2,
			new(big.Int).SetBytes([]byte("xyz")),
			new(big.Int).SetBytes([]byte("abc")),
			[]byte("xyzabc"),
		},
	}
	for _, each := range testData {
		actualOp := serializeWithLength(each.length, each.val1, each.val2)
		assert.Equal(each.expectedOp, actualOp)
	}
}

func Test_DeserializeByParts(t *testing.T) {
	assert := assert.New(t)
	var testData = []struct {
		description string
		blob        []byte
		parts       int
		expectedOp  []*big.Int
	}{
		{"basic positive case",
			[]byte("xyzabc"),
			2,
			[]*big.Int{new(big.Int).SetBytes([]byte("xyz")),
				new(big.Int).SetBytes([]byte("abc"))},
		},
		{"odd length",
			[]byte("xyzabcd"),
			2,
			[]*big.Int{new(big.Int).SetBytes([]byte("\x00xyz")),
				new(big.Int).SetBytes([]byte("abcd"))},
		},
	}
	for _, each := range testData {
		actualOp := deserializeByParts(each.blob, each.parts)
		assert.Equal(each.expectedOp, actualOp)
	}
}
