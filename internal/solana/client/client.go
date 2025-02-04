package client

import (
	"bitsnake-server/internal/solana/common"
	"bitsnake-server/internal/solana/rpc"
	"bitsnake-server/internal/solana/types"
)

type Client struct {
	RpcClient rpc.RpcClient
}

func New(opts ...rpc.Option) *Client {
	return &Client{
		RpcClient: rpc.New(opts...),
	}
}

func NewClient(endpoint string) *Client {
	return &Client{rpc.New(rpc.WithEndpoint(endpoint))}
}

type QuickSendTransactionParam struct {
	Instructions []types.Instruction
	Signers      []types.Account
	FeePayer     common.PublicKey
}

func checkJsonRpcResponse[T any](res rpc.JsonRpcResponse[T], err error) error {
	if err != nil {
		return err
	}
	if res.Error != nil {
		return res.Error
	}
	return nil
}
