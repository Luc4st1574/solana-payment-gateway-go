package system

import (
	"reflect"
	"testing"

	"bitsnake-server/internal/solana/common"

	"bitsnake-server/internal/solana/types"
)

func TestCreateAccount(t *testing.T) {
	type args struct {
		param CreateAccountParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: CreateAccountParam{
					From:     common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					New:      common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Owner:    common.StakeProgramID,
					Lamports: 1,
					Space:    200,
				},
			},
			want: types.Instruction{
				ProgramID: common.SystemProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: true},
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: true, IsWritable: true},
				},
				Data: []byte{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 200, 0, 0, 0, 0, 0, 0, 0, 6, 161, 216, 23, 145, 55, 84, 42, 152, 52, 55, 189, 254, 42, 122, 178, 85, 127, 83, 92, 138, 120, 114, 43, 104, 164, 157, 192, 0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateAccount(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAssign(t *testing.T) {
	type args struct {
		param AssignParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: AssignParam{
					From:  common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Owner: common.StakeProgramID,
				},
			},
			want: types.Instruction{
				ProgramID: common.SystemProgramID,
				Accounts: []types.AccountMeta{
					{PubKey: common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"), IsSigner: true, IsWritable: true},
				},
				Data: []byte{1, 0, 0, 0, 6, 161, 216, 23, 145, 55, 84, 42, 152, 52, 55, 189, 254, 42, 122, 178, 85, 127, 83, 92, 138, 120, 114, 43, 104, 164, 157, 192, 0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Assign(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Assign() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransfer(t *testing.T) {
	type args struct {
		param TransferParam
	}
	tests := []struct {
		name string
		args args
		want types.Instruction
	}{
		{
			args: args{
				param: TransferParam{
					From:   common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					To:     common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
					Amount: 1,
				},
			},
			want: types.Instruction{
				ProgramID: common.SystemProgramID,
				Accounts: []types.AccountMeta{
					{
						PubKey:     common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
						IsSigner:   true,
						IsWritable: true,
					},
					{
						PubKey:     common.PublicKeyFromString("BkXBQ9ThbQffhmG39c2TbXW94pEmVGJAvxWk6hfxRvUJ"),
						IsSigner:   false,
						IsWritable: true,
					},
				},
				Data: []byte{2, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Transfer(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transfer() = %v, want %v", got, tt.want)
			}
		})
	}
}
