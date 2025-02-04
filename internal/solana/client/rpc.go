package client

import (
	"encoding/base64"
	"fmt"

	"bitsnake-server/internal/solana/common"
	"bitsnake-server/internal/solana/rpc"
)

func process[A any, B any](fetch func() (rpc.JsonRpcResponse[A], error), convert func(A) (B, error)) (B, error) {
	var output B
	res, err := fetch()
	if err != nil {
		return output, err
	}
	if err = res.GetError(); err != nil {
		return output, err
	}
	return convert(res.GetResult())
}

func value[T any](v rpc.ValueWithContext[T]) (T, error) {
	return v.Value, nil
}

func forward[T any](v T) (T, error) {
	return v, nil
}

type TokenAmount struct {
	Amount         uint64
	Decimals       uint8
	UIAmountString string
}

type ReturnData struct {
	ProgramId common.PublicKey
	Data      []byte
}

func convertReturnData(d rpc.ReturnData) (ReturnData, error) {
	programId := common.PublicKeyFromString(d.ProgramId)
	s, ok := d.Data.([]any)
	if !ok {
		return ReturnData{}, fmt.Errorf("failed to get data")
	}
	if len(s) != 2 {
		return ReturnData{}, fmt.Errorf("unexpected slice lentgh")
	}
	if s[1].(string) != "base64" {
		return ReturnData{}, fmt.Errorf("unexpected encoding method")
	}
	data, err := base64.StdEncoding.DecodeString(s[0].(string))
	if err != nil {
		return ReturnData{}, fmt.Errorf("failed to decode data")
	}

	return ReturnData{
		ProgramId: programId,
		Data:      data,
	}, nil
}

type Reward struct {
	Pubkey       common.PublicKey `json:"pubkey"`
	Lamports     int64            `json:"lamports"`
	PostBalances uint64           `json:"postBalance"`
	RewardType   *rpc.RewardType  `json:"rewardType"`
	Commission   *uint8           `json:"commission"`
}
