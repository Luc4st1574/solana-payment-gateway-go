package rpc

import (
	"context"
	"testing"

	"bitsnake-server/internal/solana/client_test"
	"bitsnake-server/internal/solana/pkg/pointer"
)

func TestGetTransaction(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTransaction", "params":["4Dj8Xbs7L6z7pbNp5eGZXLmYZLwePPRVTfunjx2EWDc4nwtVYRq4YqduiFKXR23cGqmbF6LHoubGnKa7gCozstGF"]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"blockTime":1631380624,"meta":{"err":null,"fee":5000,"innerInstructions":[{"index":0,"instructions":[{"accounts":[0,1],"data":"3Bxs4h24hBtQy9rw","programIdIndex":3},{"accounts":[1],"data":"9krTDU2LzCSUJuVZ","programIdIndex":3},{"accounts":[1],"data":"SYXsBSQy3GeifSEQSGvTbrPNposbSAiSoh1YA85wcvGKSnYg","programIdIndex":3},{"accounts":[1,2,0,5],"data":"2","programIdIndex":4}]}],"logMessages":["Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL invoke [1]","Program log: Transfer 2039280 lamports to the associated token account","Program 11111111111111111111111111111111 invoke [2]","Program 11111111111111111111111111111111 success","Program log: Allocate space for the associated token account","Program 11111111111111111111111111111111 invoke [2]","Program 11111111111111111111111111111111 success","Program log: Assign the associated token account to the SPL Token program","Program 11111111111111111111111111111111 invoke [2]","Program 11111111111111111111111111111111 success","Program log: Initialize the associated token account","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [2]","Program log: Instruction: InitializeAccount","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 3412 of 177045 compute units","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success","Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL consumed 27016 of 200000 compute units","Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL success"],"postBalances":[38024615601,2039280,1461600,1,1089991680,1,898174080],"postTokenBalances":[{"accountIndex":1,"mint":"4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3","uiTokenAmount":{"amount":"0","decimals":9,"uiAmount":null,"uiAmountString":"0"}}],"preBalances":[38026659881,0,1461600,1,1089991680,1,898174080],"preTokenBalances":[],"rewards":[],"status":{"Ok":null}},"slot":80218681,"transaction":{"message":{"accountKeys":["27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ","AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ","4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3","11111111111111111111111111111111","TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","SysvarRent111111111111111111111111111111111","ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL"],"header":{"numReadonlySignedAccounts":0,"numReadonlyUnsignedAccounts":5,"numRequiredSignatures":1},"instructions":[{"accounts":[0,1,0,2,3,4,5],"data":"","programIdIndex":6}],"recentBlockhash":"Gpemb2whtMogoSGVe5KMjuoueeqNNkQ1kKnw7fsYKZHj"},"signatures":["4Dj8Xbs7L6z7pbNp5eGZXLmYZLwePPRVTfunjx2EWDc4nwtVYRq4YqduiFKXR23cGqmbF6LHoubGnKa7gCozstGF"]}},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetTransaction(
						context.TODO(),
						"4Dj8Xbs7L6z7pbNp5eGZXLmYZLwePPRVTfunjx2EWDc4nwtVYRq4YqduiFKXR23cGqmbF6LHoubGnKa7gCozstGF",
					)
				},
				ExpectedValue: JsonRpcResponse[*GetTransaction]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: &GetTransaction{
						Slot:      80218681,
						BlockTime: pointer.Get[int64](1631380624),
						Meta: &TransactionMeta{
							Fee: 5000,
							InnerInstructions: []TransactionMetaInnerInstruction{
								{
									Index: 0,
									Instructions: []any{
										map[string]any{
											"programIdIndex": 3.,
											"data":           "3Bxs4h24hBtQy9rw",
											"accounts":       []any{0., 1.},
										},
										map[string]any{
											"programIdIndex": 3.,
											"data":           "9krTDU2LzCSUJuVZ",
											"accounts":       []any{1.},
										},
										map[string]any{
											"programIdIndex": 3.,
											"data":           "SYXsBSQy3GeifSEQSGvTbrPNposbSAiSoh1YA85wcvGKSnYg",
											"accounts":       []any{1.},
										},
										map[string]any{
											"programIdIndex": 4.,
											"data":           "2",
											"accounts":       []any{1., 2., 0., 5.},
										},
									},
								},
							},
							LogMessages: []string{
								"Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL invoke [1]",
								"Program log: Transfer 2039280 lamports to the associated token account",
								"Program 11111111111111111111111111111111 invoke [2]",
								"Program 11111111111111111111111111111111 success",
								"Program log: Allocate space for the associated token account",
								"Program 11111111111111111111111111111111 invoke [2]",
								"Program 11111111111111111111111111111111 success",
								"Program log: Assign the associated token account to the SPL Token program",
								"Program 11111111111111111111111111111111 invoke [2]",
								"Program 11111111111111111111111111111111 success",
								"Program log: Initialize the associated token account",
								"Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [2]",
								"Program log: Instruction: InitializeAccount",
								"Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 3412 of 177045 compute units",
								"Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success",
								"Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL consumed 27016 of 200000 compute units",
								"Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL success",
							},
							Rewards: []Reward{},
							PreBalances: []int64{
								38026659881,
								0,
								1461600,
								1,
								1089991680,
								1,
								898174080,
							},
							PostBalances: []int64{
								38024615601,
								2039280,
								1461600,
								1,
								1089991680,
								1,
								898174080,
							},
							PreTokenBalances: []TransactionMetaTokenBalance{},
							PostTokenBalances: []TransactionMetaTokenBalance{
								{
									AccountIndex: 1,
									Mint:         "4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3",
									UITokenAmount: TokenAccountBalance{
										Amount:         "0",
										Decimals:       9,
										UIAmountString: "0",
									},
								},
							},
						},
						Transaction: map[string]any{
							"signatures": []any{
								"4Dj8Xbs7L6z7pbNp5eGZXLmYZLwePPRVTfunjx2EWDc4nwtVYRq4YqduiFKXR23cGqmbF6LHoubGnKa7gCozstGF",
							},
							"message": map[string]any{
								"header": map[string]any{
									"numReadonlySignedAccounts":   0.,
									"numReadonlyUnsignedAccounts": 5.,
									"numRequiredSignatures":       1.,
								},
								"accountKeys": []any{
									"27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ",
									"AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ",
									"4UyUTBdhPkFiu7ZE8zfxnE6hbbzf8LKo1uR5wSi5MYE3",
									"11111111111111111111111111111111",
									"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
									"SysvarRent111111111111111111111111111111111",
									"ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL",
								},
								"instructions": []any{
									map[string]any{
										"accounts":       []any{0., 1., 0., 2., 3., 4., 5.},
										"data":           "",
										"programIdIndex": 6.,
									},
								},
								"recentBlockhash": "Gpemb2whtMogoSGVe5KMjuoueeqNNkQ1kKnw7fsYKZHj",
							},
						},
					},
				},
				ExpectedError: nil,
			},

			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTransaction", "params":["4Dj8Xbs7L6z7pbNp5eGZXLmYZLwePPRVTfunjx2EWDc4nwtVYRq4YqduiFKXR23cGqmbF6LHoubGnKa7gCozstGF", {"encoding":"base64"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":null,"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetTransactionWithConfig(
						context.TODO(),
						"4Dj8Xbs7L6z7pbNp5eGZXLmYZLwePPRVTfunjx2EWDc4nwtVYRq4YqduiFKXR23cGqmbF6LHoubGnKa7gCozstGF",
						GetTransactionConfig{
							Encoding: TransactionEncodingBase64,
						},
					)
				},
				ExpectedValue: JsonRpcResponse[*GetTransaction]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result:  nil,
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0","id":1,"method":"getTransaction","params":["2vFRWHPwbxFkZMtgcyRPx8Hx1v4TTyvmhYZgYejfwWXQehiaNijP18FEQMpMFZHyJFmTXrWxncZsS6yKYESDBucQ",{"encoding":"base64", "maxSupportedTransactionVersion": 0}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"blockTime":1666897943,"meta":{"computeUnitsConsumed":185,"err":null,"fee":5000,"innerInstructions":[],"loadedAddresses":{"readonly":[],"writable":[]},"logMessages":["Program 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP invoke [1]","Program 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP consumed 185 of 200000 compute units","Program return: 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP AQIDBAU=","Program 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP success"],"postBalances":[34358789287,1141440],"postTokenBalances":[],"preBalances":[34358794287,1141440],"preTokenBalances":[],"returnData":{"data":["AQIDBAU=","base64"],"programId":"35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP"},"rewards":[],"status":{"Ok":null}},"slot":159780566,"transaction":["AV/vxqOrdrGio45xsX7l9jdCcQDy3VuY/wHlleHEuBuwDUIYO2ce/YpjkRfZHCq7tYSNEwFCRdolqNg2oibR5wUBAAECBj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4e0EmQh0otX6HS7HumAryrMtxCzacgpjtG6MY9cJWYYGK9azWH/heD6vSj5deOv9pmPQoZfCmIFJqrW8ixgJLtAQEAAA==","base64"],"version":"legacy"},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetTransactionWithConfig(
						context.TODO(),
						"2vFRWHPwbxFkZMtgcyRPx8Hx1v4TTyvmhYZgYejfwWXQehiaNijP18FEQMpMFZHyJFmTXrWxncZsS6yKYESDBucQ",
						GetTransactionConfig{
							Encoding:                       TransactionEncodingBase64,
							MaxSupportedTransactionVersion: pointer.Get[uint8](0),
						},
					)
				},
				ExpectedValue: JsonRpcResponse[*GetTransaction]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: &GetTransaction{
						Slot:      159780566,
						BlockTime: pointer.Get[int64](1666897943),
						Version:   "legacy",
						Meta: &TransactionMeta{
							Err: nil,
							Fee: 5000,
							PreBalances: []int64{
								34358794287,
								1141440,
							},
							PostBalances: []int64{
								34358789287,
								1141440,
							},
							LogMessages: []string{
								"Program 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP invoke [1]",
								"Program 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP consumed 185 of 200000 compute units",
								"Program return: 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP AQIDBAU=",
								"Program 35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP success",
							},
							Rewards: []Reward{},
							LoadedAddresses: TransactionLoadedAddresses{
								Writable: []string{},
								Readonly: []string{},
							},
							PreTokenBalances:  []TransactionMetaTokenBalance{},
							PostTokenBalances: []TransactionMetaTokenBalance{},
							InnerInstructions: []TransactionMetaInnerInstruction{},
							ReturnData: &ReturnData{
								ProgramId: "35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP",
								Data:      []any{"AQIDBAU=", "base64"},
							},
							ComputeUnitsConsumed: pointer.Get[uint64](185),
						},
						Transaction: []any{
							"AV/vxqOrdrGio45xsX7l9jdCcQDy3VuY/wHlleHEuBuwDUIYO2ce/YpjkRfZHCq7tYSNEwFCRdolqNg2oibR5wUBAAECBj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4e0EmQh0otX6HS7HumAryrMtxCzacgpjtG6MY9cJWYYGK9azWH/heD6vSj5deOv9pmPQoZfCmIFJqrW8ixgJLtAQEAAA==",
							"base64",
						},
					},
				},
				ExpectedError: nil,
			},
		},
	)
}
