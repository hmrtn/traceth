# go-ethereum Tracing

```go
// TraceTransaction returns the structured logs created during the execution of EVM
// and returns them as a JSON object.
func (api *API) TraceTransaction(ctx context.Context, hash common.Hash, config *TraceConfig) (interface{}, error) {
	_, blockHash, blockNumber, index, err := api.backend.GetTransaction(ctx, hash)
	if err != nil {
		return nil, err
	}
	// It shouldn't happen in practice.
	if blockNumber == 0 {
		return nil, errors.New("genesis is not traceable")
	}
	reexec := defaultTraceReexec
	if config != nil && config.Reexec != nil {
		reexec = *config.Reexec
	}
	block, err := api.blockByNumberAndHash(ctx, rpc.BlockNumber(blockNumber), blockHash)
	if err != nil {
		return nil, err
	}
	msg, vmctx, statedb, err := api.backend.StateAtTransaction(ctx, block, int(index), reexec)
	if err != nil {
		return nil, err
	}
	txctx := &Context{
		BlockHash: blockHash,
		TxIndex:   int(index),
		TxHash:    hash,
	}
	return api.traceTx(ctx, msg, txctx, vmctx, statedb, config)
}
```
