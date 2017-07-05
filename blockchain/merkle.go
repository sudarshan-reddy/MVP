package blockchain

import "crypto/sha256"

func merkleRoot256(hashes [][]byte) []byte {
	hashLength := len(hashes)
	switch hashLength {
	case 0:
		return nil
	case 1:
		return hashes[0]
	default:
		if hashLength%2 == 1 {
			return merkleRoot256([][]byte{merkleRoot256(
				hashes[:hashLength-1]), hashes[hashLength-1]})
		}
		bs := make([][]byte, hashLength/2)
		for i := range bs {
			j, k := i*2, (i*2)+1
			bs[i] = getSHA256(append(hashes[j], hashes[k]...))
		}
		return merkleRoot256(bs)
	}
}

func getSHA256(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	return h.Sum(nil)
}
