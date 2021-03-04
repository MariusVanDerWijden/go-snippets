package txpool

import (
	"sort"

	"github.com/ethereum/go-ethereum/core/types"
)

type wrappedTx struct {
	nonce uint64
	tx    *types.Transaction
}

type wrappedTxHeap []wrappedTx

func (h wrappedTxHeap) Len() int           { return len(h) }
func (h wrappedTxHeap) Less(i, j int) bool { return h[i].nonce < h[j].nonce }
func (h wrappedTxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *wrappedTxHeap) Push(x interface{}) {
	*h = append(*h, x.(wrappedTx))
}

func (h *wrappedTxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h wrappedTxHeap) getIdxByNonce(nonce uint64) int {
	idx := sort.Search(len(h), func(i int) bool {
		return h[i].nonce >= nonce
	})
	return idx
}

func (h wrappedTxHeap) at(index int) *types.Transaction {
	return h[index].tx
}

// Unsafe, please check h.len() beforehand
func (h wrappedTxHeap) head() wrappedTx {
	return h[0]
}

func (h wrappedTxHeap) flatten()

type account struct {
	txs *wrappedTxHeap
}

func (a *account) Get(nonce uint64) *types.Transaction {
	return a.txs.at(a.txs.getIdxByNonce(nonce))
}

func (a *account) Put(tx *types.Transaction) {
	if tx2 := a.Get(tx.Nonce()); tx2 != nil {
		*tx2 = *tx
	}
	*a.txs = append(*a.txs, wrappedTx{tx.Nonce(), tx})
	sort.Sort(a.txs)
}

func (a *account) Forward(threshold uint64) types.Transactions {
	var removed types.Transactions
	for a.txs.Len() > 0 && a.txs.head().nonce < threshold {
		var tx wrappedTx
		tx, *a.txs = (*a.txs)[0], (*a.txs)[1:]
		removed = append(removed, tx.tx)
	}
	return removed
}

func (a *account) Filter(filter func(*types.Transaction) bool) types.Transactions {
	var removed types.Transactions
	for idx, tx := range *a.txs {
		if filter(tx.tx) {
			*a.txs = append((*a.txs)[:idx], (*a.txs)[idx+1:]...)
			removed = append(removed, tx.tx)
		}
	}
	return removed
}

func (a *account) Cap(threshold int) types.Transactions {
	// Short circuit if the number of items is under the limit
	if a.txs.Len() <= threshold {
		return nil
	}
	// Otherwise gather and drop the highest nonce'd transactions
	var drops types.Transactions
	for index := a.txs.Len(); index > threshold; index-- {
		drops = append(drops, (*a.txs)[index-1].tx)
	}
	*a.txs = (*a.txs)[:threshold]
	return drops
}

func (a *account) Remove(nonce uint64) bool {
	idx := a.txs.getIdxByNonce(nonce)
	if idx == a.txs.Len() {
		return false
	}
	*a.txs = append((*a.txs)[:idx], (*a.txs)[idx+1:]...)
	return true
}

func (a *account) Ready(start uint64) types.Transactions {
	// Short circuit if no transactions are available
	if a.txs.Len() == 0 || a.txs.head().nonce > start {
		return nil
	}
	var ready types.Transactions
	for next := a.txs.head().nonce; a.txs.Len() > 0 && a.txs.head().nonce == next; next++ {
		var tx wrappedTx
		tx, *a.txs = (*a.txs)[0], (*a.txs)[1:]
		ready = append(ready, tx.tx)
	}

	return ready
}

func (a *account) Len() int {
	return a.txs.Len()
}

func (a *account) Flatten() types.Transactions {
	cache := make(types.Transactions, 0, a.txs.Len())
	for _, tx := range *a.txs {
		cache = append(cache, tx.tx)
	}
	return cache
}

func (a *account) LastElement() *types.Transaction {
	return (*a.txs)[len(*a.txs)-1].tx
}
