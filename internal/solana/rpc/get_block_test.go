package rpc

import (
	"context"
	"testing"

	"bitsnake-server/internal/solana/client_test"
	"bitsnake-server/internal/solana/pkg/pointer"
)

func TestGetBlock(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlock", "params":[33]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"blockHeight":33,"blockTime":1631803928,"blockhash":"HUonDijNaSHAPobKtAkg1ewJjy2wECpynbCq5wQ5dkCT","parentSlot":32,"previousBlockhash":"CXjZvhmFVa4ATW8Qq7XSXJFmB25aEqfHiEbCieujPd9q","rewards":[{"commission":null,"lamports":5000,"postBalance":499999840001,"pubkey":"9HvwukipCq1TVcSWoNQW7ajTUDFyC16KrARqnXppBdwX","rewardType":"Fee"}],"transactions":[{"meta":{"err":null,"fee":10000,"innerInstructions":[],"logMessages":["Program Vote111111111111111111111111111111111111111 invoke [1]","Program Vote111111111111111111111111111111111111111 success"],"postBalances":[499999835001,1000000000000000,143487360,1169280,1],"postTokenBalances":[],"preBalances":[499999845001,1000000000000000,143487360,1169280,1],"preTokenBalances":[],"rewards":[],"status":{"Ok":null}},"transaction":{"message":{"accountKeys":["9HvwukipCq1TVcSWoNQW7ajTUDFyC16KrARqnXppBdwX","3UbyTNpi3omt7hfEqQRB5844VANQFWiq8uEDNCrSwAVG","SysvarS1otHashes111111111111111111111111111","SysvarC1ock11111111111111111111111111111111","Vote111111111111111111111111111111111111111"],"header":{"numReadonlySignedAccounts":0,"numReadonlyUnsignedAccounts":3,"numRequiredSignatures":2},"instructions":[{"accounts":[1,2,3,1],"data":"2ZjTR1vUs2pHXyTLuZA9zjpNqav47YU1uqenSEcYn6xkrdmMkUJK8JDHd5TcEU7K5R9pbB2UxbY95zDzHio","programIdIndex":4}],"recentBlockhash":"CXjZvhmFVa4ATW8Qq7XSXJFmB25aEqfHiEbCieujPd9q"},"signatures":["3Me2gWFGDFwWnhugNt5u1fFvU2CyVtY4WcRzBXRKUWtgnYSxnt72p5fWiNrAkEoNTLL6FdLmk34kC41Ph91LKr6A","4cWqSVUcxTujZ6eHtNWESwCrBUfidbZ1J124VU2jY9TQpXxyHSDku1NiZhw95SzXe1mGihiP9AdQNEkLMAdvBYPQ"]}}]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetBlock(
						context.TODO(),
						33,
					)
				},
				ExpectedValue: JsonRpcResponse[*GetBlock]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: &GetBlock{
						ParentSlot:        32,
						BlockHeight:       pointer.Get[int64](33),
						BlockTime:         pointer.Get[int64](1631803928),
						PreviousBlockhash: "CXjZvhmFVa4ATW8Qq7XSXJFmB25aEqfHiEbCieujPd9q",
						Blockhash:         "HUonDijNaSHAPobKtAkg1ewJjy2wECpynbCq5wQ5dkCT",
						Rewards: []Reward{
							{
								Pubkey:       "9HvwukipCq1TVcSWoNQW7ajTUDFyC16KrARqnXppBdwX",
								Lamports:     5000,
								PostBalances: 499999840001,
								RewardType:   pointer.Get[RewardType](RewardTypeFee),
								Commission:   nil,
							},
						},
						Transactions: []GetBlockTransaction{
							{
								Meta: &TransactionMeta{
									Err: nil,
									Fee: 10000,
									PreBalances: []int64{
										499999845001,
										1000000000000000,
										143487360,
										1169280,
										1,
									},
									PostBalances: []int64{
										499999835001,
										1000000000000000,
										143487360,
										1169280,
										1,
									},
									PreTokenBalances:  []TransactionMetaTokenBalance{},
									PostTokenBalances: []TransactionMetaTokenBalance{},
									LogMessages: []string{
										"Program Vote111111111111111111111111111111111111111 invoke [1]",
										"Program Vote111111111111111111111111111111111111111 success",
									},
									Rewards:           []Reward{},
									InnerInstructions: []TransactionMetaInnerInstruction{},
								},
								Transaction: map[string]any{
									"signatures": []any{
										"3Me2gWFGDFwWnhugNt5u1fFvU2CyVtY4WcRzBXRKUWtgnYSxnt72p5fWiNrAkEoNTLL6FdLmk34kC41Ph91LKr6A",
										"4cWqSVUcxTujZ6eHtNWESwCrBUfidbZ1J124VU2jY9TQpXxyHSDku1NiZhw95SzXe1mGihiP9AdQNEkLMAdvBYPQ",
									},
									"message": map[string]any{
										"accountKeys": []any{
											"9HvwukipCq1TVcSWoNQW7ajTUDFyC16KrARqnXppBdwX",
											"3UbyTNpi3omt7hfEqQRB5844VANQFWiq8uEDNCrSwAVG",
											"SysvarS1otHashes111111111111111111111111111",
											"SysvarC1ock11111111111111111111111111111111",
											"Vote111111111111111111111111111111111111111",
										},
										"header": map[string]any{
											"numReadonlySignedAccounts":   0.,
											"numReadonlyUnsignedAccounts": 3.,
											"numRequiredSignatures":       2.,
										},
										"instructions": []any{
											map[string]any{
												"accounts":       []any{1., 2., 3., 1.},
												"data":           "2ZjTR1vUs2pHXyTLuZA9zjpNqav47YU1uqenSEcYn6xkrdmMkUJK8JDHd5TcEU7K5R9pbB2UxbY95zDzHio",
												"programIdIndex": 4.,
											},
										},
										"recentBlockhash": "CXjZvhmFVa4ATW8Qq7XSXJFmB25aEqfHiEbCieujPd9q",
									},
								},
							},
						},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBlock", "params":[33, {"encoding": "base64", "transactionDetails":"signatures"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"blockHeight":33,"blockTime":1631803928,"blockhash":"HUonDijNaSHAPobKtAkg1ewJjy2wECpynbCq5wQ5dkCT","parentSlot":32,"previousBlockhash":"CXjZvhmFVa4ATW8Qq7XSXJFmB25aEqfHiEbCieujPd9q","rewards":[{"commission":null,"lamports":5000,"postBalance":499999840001,"pubkey":"9HvwukipCq1TVcSWoNQW7ajTUDFyC16KrARqnXppBdwX","rewardType":"Fee"}],"signatures":["3Me2gWFGDFwWnhugNt5u1fFvU2CyVtY4WcRzBXRKUWtgnYSxnt72p5fWiNrAkEoNTLL6FdLmk34kC41Ph91LKr6A"]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetBlockWithConfig(
						context.TODO(),
						33,
						GetBlockConfig{
							Encoding:           GetBlockConfigEncodingBase64,
							TransactionDetails: GetBlockConfigTransactionDetailsSignatures,
						},
					)
				},
				ExpectedValue: JsonRpcResponse[*GetBlock]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: &GetBlock{
						ParentSlot:        32,
						BlockHeight:       pointer.Get[int64](33),
						BlockTime:         pointer.Get[int64](1631803928),
						PreviousBlockhash: "CXjZvhmFVa4ATW8Qq7XSXJFmB25aEqfHiEbCieujPd9q",
						Blockhash:         "HUonDijNaSHAPobKtAkg1ewJjy2wECpynbCq5wQ5dkCT",
						Rewards: []Reward{
							{
								Pubkey:       "9HvwukipCq1TVcSWoNQW7ajTUDFyC16KrARqnXppBdwX",
								Lamports:     5000,
								PostBalances: 499999840001,
								RewardType:   pointer.Get[RewardType](RewardTypeFee),
								Commission:   nil,
							},
						},
						Signatures: []string{
							"3Me2gWFGDFwWnhugNt5u1fFvU2CyVtY4WcRzBXRKUWtgnYSxnt72p5fWiNrAkEoNTLL6FdLmk34kC41Ph91LKr6A",
						},
					},
				},
				ExpectedError: nil,
			},
		},
	)
}
