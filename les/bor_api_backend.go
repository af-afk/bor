package les

import (
	"context"
	"errors"

	"github.com/maticnetwork/bor/core"
	"github.com/maticnetwork/bor/event"
)

// GetRootHash returns root hash for given start and end block
func (b *LesApiBackend) GetRootHash(ctx context.Context, starBlockNr uint64, endBlockNr uint64) (string, error) {
	return "", errors.New("Not implemented")
}

// SubscribeStateSyncEvent subscribe state sync event
func (b *LesApiBackend) SubscribeStateSyncEvent(ch chan<- core.StateSyncEvent) event.Subscription {
	return b.eth.blockchain.SubscribeStateSyncEvent(ch)
}
