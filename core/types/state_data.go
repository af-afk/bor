package types

import "github.com/maticnetwork/bor/common"

// StateSyncData represents state received from Ethereum Blockchain
type StateSyncData struct {
	ID       uint64
	Contract common.Address
	Data     string
	TxHash   common.Hash
}
