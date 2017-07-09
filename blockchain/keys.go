package blockchain

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"math/big"
)

const (
	keySize     = 1
	keyByteSize = 256
)

//Keypair holds the public and private keys
type Keypair struct {
	Public  []byte `json:"public"`
	Private []byte `json:"private"`
}

//NewKeypair generates a new keypair
func NewKeypair() (*Keypair, error) {
	public := make([]byte, keyByteSize)
	private := make([]byte, keyByteSize)
	pk, err := ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("error generating key, %s", err.Error())
	}

	b := serializeWithLength(keySize, pk.PublicKey.X, pk.PublicKey.Y)
	base64.StdEncoding.Encode(public, b)
	base64.StdEncoding.Encode(private, pk.D.Bytes())

	kp := Keypair{Public: public, Private: private}
	return &kp, nil
}

//Sign signs the keypair with proof of work
func (k *Keypair) Sign(hash []byte) ([]byte, error) {
	var public, private, signature []byte
	_, err := base64.StdEncoding.Decode(private, k.Private)
	if err != nil {
		return nil, err
	}

	_, err = base64.StdEncoding.Decode(public, k.Public)
	if err != nil {
		return nil, err
	}

	pub := deserializeByParts(public, 2)
	x, y := pub[0], pub[1]

	privateBigInt := new(big.Int).SetBytes(private)
	key := ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{Curve: elliptic.P224(), X: x, Y: y},
		D:         privateBigInt}

	r, s, _ := ecdsa.Sign(rand.Reader, &key, hash)

	base64.StdEncoding.Encode(signature, serializeWithLength(keySize, r, s))
	return signature, nil
}

func serializeWithLength(expectedLen int, bigValues ...*big.Int) []byte {
	var result []byte
	for i, b := range bigValues {
		byteValue := b.Bytes()
		diff := expectedLen - len(byteValue)
		if diff > 0 && i != 0 {
			byteValue = append(padBytes(diff, 0), byteValue...)
		}
		result = append(result, byteValue...)
	}
	return result
}

func deserializeByParts(blob []byte, parts int) []*big.Int {
	if len(blob)%2 != 0 {
		blob = append([]byte{0}, blob...)
	}

	l := len(blob) / parts
	as := make([]*big.Int, parts)

	for i := range as {
		as[i] = new(big.Int).SetBytes(blob[i*l : (i+1)*l])
	}
	return as
}
func padBytes(length int, byteValue byte) []byte {
	var byteSlice []byte

	for i := 0; i < length; i++ {
		byteSlice = append(byteSlice, byteValue)
	}
	return byteSlice

}
