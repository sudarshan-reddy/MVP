package blockchain

import "sort"

// Block is the datatype that holds a single block in the blockchain
type Block struct {
	*blockHeader
	signature    []byte
	transactions []Transaction
}

type blockHeader struct {
	origin        []byte
	previousBlock []byte
	merkleRoot    []byte
	timestamp     uint32
	nonce         uint32
}

type byTimestamp []Transaction

func (a byTimestamp) Len() int      { return len(a) }
func (a byTimestamp) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byTimestamp) Less(i, j int) bool {
	return a[i].Header.Timestamp < a[j].Header.Timestamp
}

// Transaction is a single transaction taking place for the
// blockchain to record. It could be anything from a financial
// transaction to a count of votes. This is to be moved to a
// separate package along with the transactionHeader struct
type Transaction struct {
	Header    TransactionHeader
	Signature []byte
	Payload   []byte
}

// TransactionHeader contains the metainfo of the transaction file
// and all peers are required to send it. This is to be moved to a
// separate package along with the transaction struct
type TransactionHeader struct {
	From          []byte
	To            []byte
	Timestamp     uint32
	PayloadHash   []byte
	PayloadLength uint32
	Nonce         uint32
}

// NewBlock returns a new instance of block that has a previous
// state of block attached to it
func NewBlock(previousBlock []byte) Block {
	return Block{
		blockHeader: &blockHeader{previousBlock: previousBlock},
	}
}

// AddTransaction adds a single transaction to the block in a to be
// processed state
func (b *Block) AddTransaction(t *Transaction) {
	b.transactions = append(b.transactions, *t)
	sort.Sort(byTimestamp(b.transactions))
}
