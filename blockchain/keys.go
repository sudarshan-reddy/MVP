package blockchain

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/tv42/base58"
)

const (
	keySize = 100
)

//Keypair holds the public and private keys
type Keypair struct {
	Public  []byte `json:"public"`
	Private []byte `json:"private"`
}

//NewKeypair generates a new keypair
func NewKeypair() (*Keypair, error) {
	pk, err := ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("error generating key, %s", err.Error())
	}

	b := serializeWithLength(keySize, pk.PublicKey.X, pk.PublicKey.Y)
	public := base58.EncodeBig([]byte{}, b)
	private := base58.EncodeBig([]byte{}, pk.D)

	kp := Keypair{Public: public, Private: private}
	return &kp, nil
}

//Sign signs the keypair with proof of work
func (k *Keypair) Sign(hash []byte) ([]byte, error) {
	d, err := base58.DecodeToBig(k.Private)
	if err != nil {
		return nil, err
	}

	b, _ := base58.DecodeToBig(k.Public)

	pub := deserializeByParts(b, 2)
	x, y := pub[0], pub[1]

	key := ecdsa.PrivateKey{ecdsa.PublicKey{elliptic.P224(), x, y}, d}

	r, s, _ := ecdsa.Sign(rand.Reader, &key, hash)

	return base58.EncodeBig([]byte{}, serializeWithLength(keySize, r, s)), nil
}

func serializeWithLength(expectedLen int, bigValues ...*big.Int) *big.Int {
	byteSetter := []byte{}
	for i, b := range bigValues {
		byteValue := b.Bytes()
		diff := expectedLen - len(byteValue)
		if diff > 0 && i != 0 {
			byteValue = append(padBytes(diff, 0), byteValue...)
		}
		byteSetter = append(byteSetter, byteValue...)
	}
	return new(big.Int).SetBytes(byteSetter)
}

func deserializeByParts(blob *big.Int, parts int) []*big.Int {
	bs := blob.Bytes()
	if len(bs)%2 != 0 {
		bs = append([]byte{0}, bs...)
	}

	l := len(bs) / parts
	as := make([]*big.Int, parts)

	for i := range as {
		as[i] = new(big.Int).SetBytes(bs[i*l : (i+1)*l])
	}
	return as
}
func padBytes(length int, byteValue byte) []byte {
	var byteSlice []byte

	for i := length; i > 0; i++ {
		byteSlice = append(byteSlice, byteValue)
	}
	return byteSlice

}
