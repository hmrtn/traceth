# A comparison of client tracing APIs

| Client      | Usage                                                                                                                             | Trace Dump                                 |
| ----------- | --------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------ |
| go-ethereum | [View](https://github.com/ethereum/go-ethereum/blob/ee120ef865e9468fef0bbb0a0bcffba93e3e358e/eth/tracers/api.go#757)              |                                            |
| Nethermind  | [View](https://docs.nethermind.io/nethermind/nethermind-utilities/cli/debug#debug-tracetransaction)                               | [View](./nethermind/nethermind_trace.json) |
| Erigon      | [View](https://github.com/ledgerwatch/erigon/blob/5ca558c6670dd96afac472ef573152dda97fc866/cmd/rpcdaemon/commands/tracing.go#L23) | [View](./erigon/erigon_trace.json)         |

## Getting Started

A convienient way to quickly visualize the difference in traces is to use the included `trace.go` program with a transaction hash, say this one - `0xbe811e3a5aea163edfa38f19e2a15eafc943c3e31101fe415eedf5c1337c73ec`. If we run the `trace.go`, an output of a highlighted trace is shown and are seperately stored. You will need access to an archive node with an enabled JSON RPC API for each client comparison.

```bash
git clone https://github.com/hansmrtn/evm-tracing-comparison.git
cd evm-tracing-comparison
go run trace.go
```

### Details

Any transaction on an EVM compatible protocol that involves a contract address can execute arbitrarily complex instructions, including acting on previously stored data with other accounts and addresses. The receipt of a transaction contains a status code indicating the its success or failure, but no other information relevant to the executions context. Hence it can be necessary to ask a client/node to trace a transaction by re-executing the the contract instructions and parsing the transactions in a useful structure to understand what (if any) transactions, modifications, and external code were invoked. The issue with this is the tracing modes and format is fragmented amongst popular clients.

This project is meant to serve as a repository of information about each client and it's tracing APIs with the goal of eventually standardizing key formats.
