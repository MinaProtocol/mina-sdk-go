# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.0] - 2026-04-14

### Added
- `Client` with functional options (`WithGraphQLURI`, `WithRetries`, `WithRetryDelay`, `WithTimeout`)
- Query methods: `GetSyncStatus`, `GetDaemonStatus`, `GetNetworkID`, `GetAccount`, `GetBestChain`, `GetPeers`, `GetPooledUserCommands`
- Mutation methods: `SendPayment`, `SendDelegation`, `SetSnarkWorker`, `SetSnarkWorkFee`
- `Currency` type with nanomina precision arithmetic (Add, Sub, Mul, comparisons)
- Typed error types: `GraphQLError`, `ConnectionError`, `AccountNotFoundError`, `CurrencyUnderflowError`
- Response types: `DaemonStatus`, `AccountData`, `BlockInfo`, `PeerInfo`, `PooledUserCommand`
- Unit tests with HTTP mocking (28 tests)
- Integration tests against live daemon (13 tests, skipped without MINA_GRAPHQL_URI)
- CI workflows: test (Go 1.21-1.23), integration, release, schema drift
- Schema drift detection script
- Apache 2.0 license
