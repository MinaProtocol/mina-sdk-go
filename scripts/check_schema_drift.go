//go:build ignore

// check_schema_drift validates that the GraphQL queries used by the SDK
// are compatible with a running Mina daemon's schema.
//
// Usage:
//
//	go run scripts/check_schema_drift.go --endpoint http://127.0.0.1:8080/graphql --branch master --strict
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var queries = map[string]string{
	"syncStatus": `query { syncStatus }`,
	"daemonStatus": `query {
		daemonStatus {
			syncStatus blockchainLength highestBlockLengthReceived
			uptimeSecs stateHash commitId
			peers { peerId host libp2pPort }
		}
	}`,
	"networkID":          `query { networkID }`,
	"getAccount":         `query ($publicKey: PublicKey!) { account(publicKey: $publicKey) { publicKey nonce delegate tokenId balance { total liquid locked } } }`,
	"bestChain":          `query { bestChain(maxLength: 1) { stateHash commandTransactionCount creatorAccount { publicKey } protocolState { consensusState { blockHeight slotSinceGenesis slot } } } }`,
	"getPeers":           `query { getPeers { peerId host libp2pPort } }`,
	"pooledUserCommands": `query { pooledUserCommands { id hash kind nonce amount fee from to } }`,
}

func main() {
	endpoint := flag.String("endpoint", "http://127.0.0.1:8080/graphql", "GraphQL endpoint")
	branch := flag.String("branch", "unknown", "Mina branch being tested")
	strict := flag.Bool("strict", false, "Fail on any drift")
	flag.Parse()

	fmt.Printf("Schema drift check against %s (%s)\n", *branch, *endpoint)

	failures := 0
	warnings := 0

	for name, query := range queries {
		payload, _ := json.Marshal(map[string]any{"query": query})
		resp, err := http.Post(*endpoint, "application/json", bytes.NewReader(payload))
		if err != nil {
			fmt.Printf("FAIL  %s: connection error: %v\n", name, err)
			failures++
			continue
		}
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		var result struct {
			Data   json.RawMessage  `json:"data"`
			Errors []map[string]any `json:"errors"`
		}
		if err := json.Unmarshal(body, &result); err != nil {
			fmt.Printf("FAIL  %s: invalid JSON response\n", name)
			failures++
			continue
		}

		if len(result.Errors) > 0 {
			msgs := make([]string, len(result.Errors))
			for i, e := range result.Errors {
				if m, ok := e["message"].(string); ok {
					msgs[i] = m
				}
			}
			fmt.Printf("DRIFT %s: %s\n", name, strings.Join(msgs, "; "))
			if *strict {
				failures++
			} else {
				warnings++
			}
			continue
		}

		fmt.Printf("OK    %s\n", name)
	}

	fmt.Printf("\nResults: %d ok, %d drift, %d failures\n",
		len(queries)-failures-warnings, warnings, failures)

	if failures > 0 {
		os.Exit(1)
	}
}
