# EVM Tracing Comparison

Any transaction on an EVM compatible protocol that involves a contract address can execute arbitrarily complex instructions, including acting on previously stored data with other accounts and addresses. The receipt of a transaction contains a status code indicating the its success or failure, but no other information relevant to the executions context. Hence it can be necessary to ask a client/node to trace a transaction by re-executing the the contract instructions and parsing the transactions in a useful structure to understand what (if any) transactions, modifications, and external code were invoked. The issue with this is the tracing modes and format is fragmented amongst popular clients. 

This project is meant to serve as a repository of information about each client and it's tracing APIs with the goal of eventually standardizing key formats. 
