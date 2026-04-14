package mina

// AccountBalance represents the balance of a Mina account.
type AccountBalance struct {
	Total  Currency
	Liquid *Currency
	Locked *Currency
}

// AccountData represents a Mina account.
type AccountData struct {
	PublicKey string
	Nonce     int
	Balance   AccountBalance
	Delegate  string
	TokenID   string
}

// PeerInfo represents a connected peer.
type PeerInfo struct {
	PeerID string
	Host   string
	Port   int
}

// DaemonStatus represents the status of the Mina daemon.
type DaemonStatus struct {
	SyncStatus                string
	BlockchainLength          *int
	HighestBlockLengthReceived *int
	UptimeSecs                *int
	Peers                     []PeerInfo
	CommitID                  string
	StateHash                 string
}

// BlockInfo represents a block in the best chain.
type BlockInfo struct {
	StateHash                string
	Height                   int
	GlobalSlotSinceHardFork  int
	GlobalSlotSinceGenesis   int
	CreatorPK                string
	CommandTransactionCount  int
}

// SendPaymentResult is the result of a send_payment mutation.
type SendPaymentResult struct {
	ID    string
	Hash  string
	Nonce int
}

// SendDelegationResult is the result of a send_delegation mutation.
type SendDelegationResult struct {
	ID    string
	Hash  string
	Nonce int
}

// PooledUserCommand represents a pending transaction in the mempool.
type PooledUserCommand struct {
	ID     string
	Hash   string
	Kind   string
	Nonce  string
	Amount string
	Fee    string
	From   string
	To     string
}
