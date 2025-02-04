package rpc

import (
	"context"
	"testing"

	"bitsnake-server/internal/solana/client_test"
)

func TestGetAccountInfo(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getAccountInfo", "params":["RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7"]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":77317716},"value":{"data":"","executable":false,"lamports":21474700400,"owner":"11111111111111111111111111111111","rentEpoch":178}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetAccountInfo(
						context.Background(),
						"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
					)
				},
				ExpectedValue: JsonRpcResponse[ValueWithContext[AccountInfo]]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: ValueWithContext[AccountInfo]{
						Context: Context{
							Slot: 77317716,
						},
						Value: AccountInfo{
							Lamports:   21474700400,
							Owner:      "11111111111111111111111111111111",
							Executable: false,
							RentEpoch:  178,
							Data:       "",
						},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getAccountInfo", "params":["FaTGhPTgKeZZzQwLenoxn2VZXPWV1FpjQ1AQe77JUeJw"]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":77382573},"value":null},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetAccountInfo(
						context.Background(),
						"FaTGhPTgKeZZzQwLenoxn2VZXPWV1FpjQ1AQe77JUeJw",
					)
				},
				ExpectedValue: JsonRpcResponse[ValueWithContext[AccountInfo]]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: ValueWithContext[AccountInfo]{
						Context: Context{
							Slot: 77382573,
						},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getAccountInfo", "params":["F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb"]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":77317716},"value":{"data":"DK9MyTraxAdzd5fQ2Cvpbb2CuDd3VHxAiXuVi3E9Swzr9urV1kwxJonAiZK2zQ5xyy2FqiguDwNUGtofpzWwz3UxafwMgjFS6jx82g1B7Z2tAAj","executable":false,"lamports":1461600,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":178}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetAccountInfo(
						context.Background(),
						"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
					)
				},
				ExpectedValue: JsonRpcResponse[ValueWithContext[AccountInfo]]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: ValueWithContext[AccountInfo]{
						Context: Context{
							Slot: 77317716,
						},
						Value: AccountInfo{
							Lamports:   1461600,
							Owner:      "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
							Executable: false,
							RentEpoch:  178,
							Data:       "DK9MyTraxAdzd5fQ2Cvpbb2CuDd3VHxAiXuVi3E9Swzr9urV1kwxJonAiZK2zQ5xyy2FqiguDwNUGtofpzWwz3UxafwMgjFS6jx82g1B7Z2tAAj",
						},
					},
				},
				ExpectedError: nil,
			},
		},
	)
}
