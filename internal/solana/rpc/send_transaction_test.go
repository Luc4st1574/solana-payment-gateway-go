package rpc

import (
	"context"
	"testing"

	"bitsnake-server/internal/solana/client_test"
)

func TestSendTransaction(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0","id":1,"method":"sendTransaction","params":["37u9WtQpcm6ULa3Vmu7ySnANv"]}`,
				ResponseBody: `{"error":{"code":-32602,"message":"io error: failed to fill whole buffer"},"id":1,"jsonrpc":"2.0"}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.SendTransaction(context.Background(), "37u9WtQpcm6ULa3Vmu7ySnANv")
				},
				ExpectedValue: JsonRpcResponse[string]{
					JsonRpc: "2.0",
					Id:      1,
					Error: &JsonRpcError{
						Code:    -32602,
						Message: `io error: failed to fill whole buffer`,
					},
					Result: "",
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0","id":1,"method":"sendTransaction","params":["CayPuTQXGJTdD36zymjYnGPDovnMr9ZaTvqN4W4JDe3bNs6WySjg4qut7AQYHbMxf38f95qd7cQ3GtzHM1CWjoui6qPPkaMMAu9fyCvfsGXFkVjeTczjSrBCWz6t74m3voiTaLpVEG8WHosKfSVUUC1UMHdgHKp63ZZeA1k9ZH2hgwfByfnEftgkMTEGyQ8mMx1q8MZVbbQGs2eNeTKxbUupCp8WotHZ9YrtqwJfLXF8HMHqGHZ8VdpMV",{"skipPreflight":true}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":"2F53DggXYWLzczigoMr7smSEZtWSKmsWr7HMJQiNbTBdjjcN54LUMWdvTLj46MH7rAnJVPjJEjRjjXKeG7mssmZb","id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.SendTransactionWithConfig(
						context.Background(),
						"CayPuTQXGJTdD36zymjYnGPDovnMr9ZaTvqN4W4JDe3bNs6WySjg4qut7AQYHbMxf38f95qd7cQ3GtzHM1CWjoui6qPPkaMMAu9fyCvfsGXFkVjeTczjSrBCWz6t74m3voiTaLpVEG8WHosKfSVUUC1UMHdgHKp63ZZeA1k9ZH2hgwfByfnEftgkMTEGyQ8mMx1q8MZVbbQGs2eNeTKxbUupCp8WotHZ9YrtqwJfLXF8HMHqGHZ8VdpMV",
						SendTransactionConfig{
							SkipPreflight: true,
						},
					)
				},
				ExpectedValue: JsonRpcResponse[string]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result:  "2F53DggXYWLzczigoMr7smSEZtWSKmsWr7HMJQiNbTBdjjcN54LUMWdvTLj46MH7rAnJVPjJEjRjjXKeG7mssmZb",
				},
				ExpectedError: nil,
			},
		},
	)
}
