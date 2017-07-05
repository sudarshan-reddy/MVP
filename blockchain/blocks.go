package blockchain

type transaction struct {
	header    transactionHeader
	signature []byte
	payload   []byte
}

type transactionHeader struct {
	From          []byte
	To            []byte
	Timestamp     uint32
	PayloadHash   []byte
	PayloadLength uint32
	Nonce         uint32
}

type block struct {
	*blockHeader
	signature    []byte
	transactions []transaction
}

type blockHeader struct {
	origin     []byte
	prevBlock  []byte
	merkleRoot []byte
	timestamp  uint32
	nonce      uint32
}
