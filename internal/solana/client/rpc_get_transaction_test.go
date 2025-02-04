package client

import (
	"context"
	"testing"

	"bitsnake-server/internal/solana/client_test"
	"bitsnake-server/internal/solana/common"
	"bitsnake-server/internal/solana/pkg/pointer"
	"bitsnake-server/internal/solana/rpc"
	"bitsnake-server/internal/solana/types"

	"github.com/mr-tron/base58"
)

func TestClient_GetTransaction(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTransaction", "params":["25D9azGKNfJiKp4B5drSV1PjeePKaCreb9VAUFxAdm4qERDTMRjeKv4nfM1c1Wek879C9R2VT3x3hUdW5YCZ2hxp", {"encoding":"base64", "maxSupportedTransactionVersion": 0}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":null,"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetTransaction(
						context.Background(),
						"25D9azGKNfJiKp4B5drSV1PjeePKaCreb9VAUFxAdm4qERDTMRjeKv4nfM1c1Wek879C9R2VT3x3hUdW5YCZ2hxp",
					)
				},
				ExpectedValue: (*Transaction)(nil),
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getTransaction", "params": ["4fSTSDTTuYa1XXAFxFenuY3SoZWUwCzpMq7kUiya1zW6uqqh6C76GFqTQ3wvegEbZhbPJyr33iDAbieQVWCtVXmf", {"encoding": "base64", "maxSupportedTransactionVersion": 0}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"blockTime":1675511254,"meta":{"computeUnitsConsumed":12344,"err":null,"fee":5000,"innerInstructions":[],"loadedAddresses":{"readonly":["F1rcBbZB6tQZUTR2z8jKQxaAwUUkxnghSh941Q62hMi8","5jHeQFBSNxFqqkMF9YCYwtJbkzGarSGwGsmi2ZuPG6yw"],"writable":["3Yvq7e9UXLoFK4PKyxrpEA3y3TKmFK2Wb1f5tVFUgwPu","5McxjaxNKYLHtv9DqbMfoi6GNs7ZEMHGkJDrouPib4sW","GAXzq8BWdAWaS1kWFiL5tzV2h3AbRBtYGP5psNTWrM9g"]},"logMessages":["Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [1]","Program log: Instruction: TransferChecked","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 6172 of 400000 compute units","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [1]","Program log: Instruction: TransferChecked","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 6172 of 393828 compute units","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success"],"postBalances":[112595188235,2039280,934087680,2039280,2039280,2039280,1461600,1461600],"postTokenBalances":[{"accountIndex":1,"mint":"5jHeQFBSNxFqqkMF9YCYwtJbkzGarSGwGsmi2ZuPG6yw","owner":"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7","programId":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","uiTokenAmount":{"amount":"101","decimals":0,"uiAmount":101.0,"uiAmountString":"101"}},{"accountIndex":3,"mint":"F1rcBbZB6tQZUTR2z8jKQxaAwUUkxnghSh941Q62hMi8","owner":"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7","programId":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","uiTokenAmount":{"amount":"99","decimals":0,"uiAmount":99.0,"uiAmountString":"99"}},{"accountIndex":4,"mint":"5jHeQFBSNxFqqkMF9YCYwtJbkzGarSGwGsmi2ZuPG6yw","owner":"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7","programId":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","uiTokenAmount":{"amount":"99","decimals":0,"uiAmount":99.0,"uiAmountString":"99"}},{"accountIndex":5,"mint":"F1rcBbZB6tQZUTR2z8jKQxaAwUUkxnghSh941Q62hMi8","owner":"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7","programId":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","uiTokenAmount":{"amount":"101","decimals":0,"uiAmount":101.0,"uiAmountString":"101"}}],"preBalances":[112595193235,2039280,934087680,2039280,2039280,2039280,1461600,1461600],"preTokenBalances":[{"accountIndex":1,"mint":"5jHeQFBSNxFqqkMF9YCYwtJbkzGarSGwGsmi2ZuPG6yw","owner":"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7","programId":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","uiTokenAmount":{"amount":"100","decimals":0,"uiAmount":100.0,"uiAmountString":"100"}},{"accountIndex":3,"mint":"F1rcBbZB6tQZUTR2z8jKQxaAwUUkxnghSh941Q62hMi8","owner":"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7","programId":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","uiTokenAmount":{"amount":"100","decimals":0,"uiAmount":100.0,"uiAmountString":"100"}},{"accountIndex":4,"mint":"5jHeQFBSNxFqqkMF9YCYwtJbkzGarSGwGsmi2ZuPG6yw","owner":"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7","programId":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","uiTokenAmount":{"amount":"100","decimals":0,"uiAmount":100.0,"uiAmountString":"100"}},{"accountIndex":5,"mint":"F1rcBbZB6tQZUTR2z8jKQxaAwUUkxnghSh941Q62hMi8","owner":"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7","programId":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","uiTokenAmount":{"amount":"100","decimals":0,"uiAmount":100.0,"uiAmountString":"100"}}],"rewards":[],"status":{"Ok":null}},"slot":193487858,"transaction":["AbczATLXANCJ0Y2NoK0du6pwKuLSbYyG7YaFgJgQVtvjd7oKxHCE11YBK9DlyS2t2Fslh+oDT02oSJNGpJuCsQaAAQABAwY+cNmRV5jco+7bkTfPZMcP+vtizdOCgQUlC9drHWze+il9VuGydqFkeFhh/iremTB8Ngd13K3Xt+TOOJY8/QQG3fbh12Whk9nL4UbO63msHLSF7V9bN5E6jPWFfv8AqUTB7DdvVxpi/fsG318JDpL57X6sICK5kJnx/HugOWK7AgIEAwYFAAoMAQAAAAAAAAAAAgQEBwEACgwBAAAAAAAAAAACWt1BI7yRb9qO/G87o+tplZPL5F1W7UbkIFKWOJjtmUECAQIBAKMGCIabnF0TqEjGtz+67okLc/n3dwUqej+EGtkfc+eaAQIBAA==","base64"],"version":0},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetTransaction(
						context.Background(),
						"4fSTSDTTuYa1XXAFxFenuY3SoZWUwCzpMq7kUiya1zW6uqqh6C76GFqTQ3wvegEbZhbPJyr33iDAbieQVWCtVXmf",
					)
				},
				ExpectedValue: &Transaction{
					Slot:      193487858,
					BlockTime: pointer.Get[int64](1675511254),
					Meta: &TransactionMeta{
						Err: nil,
						Fee: 5000,
						PreBalances: []int64{
							112595193235,
							2039280,
							934087680,
							2039280,
							2039280,
							2039280,
							1461600,
							1461600,
						},
						PostBalances: []int64{
							112595188235,
							2039280,
							934087680,
							2039280,
							2039280,
							2039280,
							1461600,
							1461600,
						},
						LogMessages: []string{
							"Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [1]",
							"Program log: Instruction: TransferChecked",
							"Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 6172 of 400000 compute units",
							"Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success",
							"Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [1]",
							"Program log: Instruction: TransferChecked",
							"Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 6172 of 393828 compute units",
							"Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success",
						},
						LoadedAddresses: rpc.TransactionLoadedAddresses{
							Readonly: []string{
								"F1rcBbZB6tQZUTR2z8jKQxaAwUUkxnghSh941Q62hMi8",
								"5jHeQFBSNxFqqkMF9YCYwtJbkzGarSGwGsmi2ZuPG6yw",
							},
							Writable: []string{
								"3Yvq7e9UXLoFK4PKyxrpEA3y3TKmFK2Wb1f5tVFUgwPu",
								"5McxjaxNKYLHtv9DqbMfoi6GNs7ZEMHGkJDrouPib4sW",
								"GAXzq8BWdAWaS1kWFiL5tzV2h3AbRBtYGP5psNTWrM9g",
							},
						},
						PreTokenBalances: []rpc.TransactionMetaTokenBalance{
							{
								AccountIndex: 1,
								Mint:         "5jHeQFBSNxFqqkMF9YCYwtJbkzGarSGwGsmi2ZuPG6yw",
								Owner:        "RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
								ProgramId:    "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
								UITokenAmount: rpc.TokenAccountBalance{
									Amount:         "100",
									Decimals:       0,
									UIAmountString: "100",
								},
							},
							{
								AccountIndex: 3,
								Mint:         "F1rcBbZB6tQZUTR2z8jKQxaAwUUkxnghSh941Q62hMi8",
								Owner:        "RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
								ProgramId:    "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
								UITokenAmount: rpc.TokenAccountBalance{
									Amount:         "100",
									Decimals:       0,
									UIAmountString: "100",
								},
							},
							{
								AccountIndex: 4,
								Mint:         "5jHeQFBSNxFqqkMF9YCYwtJbkzGarSGwGsmi2ZuPG6yw",
								Owner:        "RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
								ProgramId:    "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
								UITokenAmount: rpc.TokenAccountBalance{
									Amount:         "100",
									Decimals:       0,
									UIAmountString: "100",
								},
							},
							{
								AccountIndex: 5,
								Mint:         "F1rcBbZB6tQZUTR2z8jKQxaAwUUkxnghSh941Q62hMi8",
								Owner:        "RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
								ProgramId:    "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
								UITokenAmount: rpc.TokenAccountBalance{
									Amount:         "100",
									Decimals:       0,
									UIAmountString: "100",
								},
							},
						},
						PostTokenBalances: []rpc.TransactionMetaTokenBalance{
							{
								AccountIndex: 1,
								Mint:         "5jHeQFBSNxFqqkMF9YCYwtJbkzGarSGwGsmi2ZuPG6yw",
								Owner:        "RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
								ProgramId:    "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
								UITokenAmount: rpc.TokenAccountBalance{
									Amount:         "101",
									Decimals:       0,
									UIAmountString: "101",
								},
							},
							{
								AccountIndex: 3,
								Mint:         "F1rcBbZB6tQZUTR2z8jKQxaAwUUkxnghSh941Q62hMi8",
								Owner:        "RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
								ProgramId:    "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
								UITokenAmount: rpc.TokenAccountBalance{
									Amount:         "99",
									Decimals:       0,
									UIAmountString: "99",
								},
							},
							{
								AccountIndex: 4,
								Mint:         "5jHeQFBSNxFqqkMF9YCYwtJbkzGarSGwGsmi2ZuPG6yw",
								Owner:        "RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
								ProgramId:    "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
								UITokenAmount: rpc.TokenAccountBalance{
									Amount:         "99",
									Decimals:       0,
									UIAmountString: "99",
								},
							},
							{
								AccountIndex: 5,
								Mint:         "F1rcBbZB6tQZUTR2z8jKQxaAwUUkxnghSh941Q62hMi8",
								Owner:        "RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
								ProgramId:    "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
								UITokenAmount: rpc.TokenAccountBalance{
									Amount:         "101",
									Decimals:       0,
									UIAmountString: "101",
								},
							},
						},
						InnerInstructions:    []InnerInstruction{},
						ComputeUnitsConsumed: pointer.Get[uint64](12344),
					},
					Transaction: types.Transaction{
						Signatures: []types.Signature{
							mustBase58Decode(t, "4fSTSDTTuYa1XXAFxFenuY3SoZWUwCzpMq7kUiya1zW6uqqh6C76GFqTQ3wvegEbZhbPJyr33iDAbieQVWCtVXmf"),
						},
						Message: types.Message{
							Version: types.MessageVersionV0,
							Header: types.MessageHeader{
								NumRequireSignatures:        1,
								NumReadonlySignedAccounts:   0,
								NumReadonlyUnsignedAccounts: 1,
							},
							Accounts: []common.PublicKey{
								common.PublicKeyFromString("RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7"),
								common.PublicKeyFromString("HqXcr9ja8jTZAfWN4YSSL8PPWFN3BFJsoxrCvSLaqww1"),
								common.PublicKeyFromString("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"),
							},
							RecentBlockHash: "5dQEKfLJt77vfrw2UxWrPrDFwFmxRui6Rk6FBjGnuZBg",
							Instructions: []types.CompiledInstruction{
								{
									ProgramIDIndex: 2,
									Accounts:       []int{3, 6, 5, 0},
									Data:           []byte{12, 1, 0, 0, 0, 0, 0, 0, 0, 0},
								},
								{
									ProgramIDIndex: 2,
									Accounts:       []int{4, 7, 1, 0},
									Data:           []byte{12, 1, 0, 0, 0, 0, 0, 0, 0, 0},
								},
							},
							AddressLookupTables: []types.CompiledAddressLookupTable{
								{
									AccountKey:      common.PublicKeyFromString("77hNYFDx74WFBD1jfM1gHFYk3naH8CxLzLG4KRJAHcRv"),
									ReadonlyIndexes: []uint8{0},
									WritableIndexes: []uint8{1, 2},
								},
								{
									AccountKey:      common.PublicKeyFromString("ByNnrePVpmJTXGiU3Nm9UxTN36tsbaahQcvUNFWmX2Do"),
									ReadonlyIndexes: []uint8{0},
									WritableIndexes: []uint8{2},
								},
							},
						},
					},
					AccountKeys: []common.PublicKey{
						common.PublicKeyFromString("RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7"),
						common.PublicKeyFromString("HqXcr9ja8jTZAfWN4YSSL8PPWFN3BFJsoxrCvSLaqww1"),
						common.PublicKeyFromString("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"),
						common.PublicKeyFromString("3Yvq7e9UXLoFK4PKyxrpEA3y3TKmFK2Wb1f5tVFUgwPu"),
						common.PublicKeyFromString("5McxjaxNKYLHtv9DqbMfoi6GNs7ZEMHGkJDrouPib4sW"),
						common.PublicKeyFromString("GAXzq8BWdAWaS1kWFiL5tzV2h3AbRBtYGP5psNTWrM9g"),
						common.PublicKeyFromString("F1rcBbZB6tQZUTR2z8jKQxaAwUUkxnghSh941Q62hMi8"),
						common.PublicKeyFromString("5jHeQFBSNxFqqkMF9YCYwtJbkzGarSGwGsmi2ZuPG6yw"),
					},
				},
				ExpectedError: nil,
			},
		},
	)
}

func mustBase58Decode(t *testing.T, s string) []byte {
	b, err := base58.Decode(s)
	if err != nil {
		t.Fatalf("failed to base58 decode %v", s)
	}
	return b
}
