package types

import (
	"fmt"
	"testing"

	"bitsnake-server/internal/solana/common"

	"github.com/stretchr/testify/assert"
)

func TestMessage_DecompileInstructions(t *testing.T) {
	type fields struct {
		Version             MessageVersion
		Header              MessageHeader
		Accounts            []common.PublicKey
		RecentBlockHash     string
		Instructions        []CompiledInstruction
		AddressLookupTables []CompiledAddressLookupTable
	}
	tests := []struct {
		name   string
		fields fields
		want   []Instruction
		panic  string
	}{
		{
			fields: fields{
				Header: MessageHeader{
					NumRequireSignatures:        1,
					NumReadonlySignedAccounts:   0,
					NumReadonlyUnsignedAccounts: 1,
				},
				Accounts: []common.PublicKey{
					common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					common.PublicKeyFromString("A4iUVr5KjmsLymUcv4eSKPedUtoaBceiPeGipKMYc69b"),
					common.SystemProgramID,
				},
				RecentBlockHash: "FwRYtTPRk5N4wUeP87rTw9kQVSwigB6kbikGzzeCMrW5",
				Instructions: []CompiledInstruction{
					{
						ProgramIDIndex: 2,
						Accounts:       []int{0, 1},
						Data:           []byte{2, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
					},
				},
			},
			want: []Instruction{
				{
					Accounts: []AccountMeta{
						{PubKey: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"), IsSigner: true, IsWritable: true},
						{PubKey: common.PublicKeyFromString("A4iUVr5KjmsLymUcv4eSKPedUtoaBceiPeGipKMYc69b"), IsSigner: false, IsWritable: true},
					},
					ProgramID: common.SystemProgramID,
					Data:      []byte{2, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
				},
			},
		},
		{
			fields: fields{
				Version: MessageVersionV0,
				Header: MessageHeader{
					NumRequireSignatures:        1,
					NumReadonlySignedAccounts:   0,
					NumReadonlyUnsignedAccounts: 1,
				},
				Accounts: []common.PublicKey{
					common.PublicKeyFromString("9aE476sH92Vz7DMPyq5WLPkrKWivxeuTKEFKd2sZZcde"),
					common.SystemProgramID,
				},
				RecentBlockHash: "5EvWPqKeYfN2P7SAQZ2TLnXhV3Ltjn6qEhK1F279dUUW",
				Instructions: []CompiledInstruction{
					{
						ProgramIDIndex: 1,
						Accounts:       []int{0, 2},
						Data:           []byte{2, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
					},
				},
				AddressLookupTables: []CompiledAddressLookupTable{
					{
						AccountKey:      common.PublicKeyFromString("HEhDGuxaxGr9LuNtBdvbX2uggyAKoxYgHFaAiqxVu8UY"),
						WritableIndexes: []uint8{1},
						ReadonlyIndexes: []uint8{},
					},
				},
			},
			panic: "hasn't supported",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				Version:             tt.fields.Version,
				Header:              tt.fields.Header,
				Accounts:            tt.fields.Accounts,
				RecentBlockHash:     tt.fields.RecentBlockHash,
				Instructions:        tt.fields.Instructions,
				AddressLookupTables: tt.fields.AddressLookupTables,
			}
			if len(tt.panic) == 0 {
				assert.Equal(t, tt.want, m.DecompileInstructions())
			} else {
				assert.PanicsWithValue(t, tt.panic, func() {
					m.DecompileInstructions()
				})
			}
		})
	}
}

func TestNewMessage(t *testing.T) {
	type args struct {
		param NewMessageParam
	}
	tests := []struct {
		name string
		args args
		want Message
	}{
		{
			args: args{
				NewMessageParam{
					FeePayer: common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					Instructions: []Instruction{
						{
							ProgramID: common.SystemProgramID,
							Accounts: []AccountMeta{
								{
									PubKey:     common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
									IsSigner:   true,
									IsWritable: true,
								},
								{
									PubKey:     common.PublicKeyFromString("A4iUVr5KjmsLymUcv4eSKPedUtoaBceiPeGipKMYc69b"),
									IsSigner:   false,
									IsWritable: true,
								},
							},
							Data: []byte{2, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
						},
					},
					RecentBlockhash: "FwRYtTPRk5N4wUeP87rTw9kQVSwigB6kbikGzzeCMrW5",
				},
			},
			want: Message{
				Version: MessageVersionLegacy,
				Header: MessageHeader{
					NumRequireSignatures:        1,
					NumReadonlySignedAccounts:   0,
					NumReadonlyUnsignedAccounts: 1,
				},
				Accounts: []common.PublicKey{
					common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					common.PublicKeyFromString("A4iUVr5KjmsLymUcv4eSKPedUtoaBceiPeGipKMYc69b"),
					common.SystemProgramID,
				},
				RecentBlockHash: "FwRYtTPRk5N4wUeP87rTw9kQVSwigB6kbikGzzeCMrW5",
				Instructions: []CompiledInstruction{
					{
						ProgramIDIndex: 2,
						Accounts:       []int{0, 1},
						Data:           []byte{2, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
					},
				},
				AddressLookupTables: []CompiledAddressLookupTable{},
			},
		},
		{
			args: args{
				NewMessageParam{
					FeePayer: common.PublicKeyFromString("FUarP2p5EnxD66vVDL4PWRoWMzA56ZVHG24hpEDFShEz"),
					Instructions: []Instruction{
						{
							ProgramID: common.TokenProgramID,
							Accounts: []AccountMeta{
								{
									PubKey:     common.PublicKeyFromString("8YNmYW9rWwpmLxUDycqHj1JMAMdm1v2VBB55tXqt7jej"),
									IsSigner:   false,
									IsWritable: true,
								},
								{
									PubKey:     common.PublicKeyFromString("5XaEXmAEiA4t3EdFWADixN9537Nct5Y5PMRz391eD9N1"),
									IsSigner:   false,
									IsWritable: false,
								},
								{
									PubKey:     common.PublicKeyFromString("CPaB3EuV5qJK25stSWzH3815BspeyGgYvaR1Z8B72hbp"),
									IsSigner:   false,
									IsWritable: true,
								},
								{
									PubKey:     common.PublicKeyFromString("FUarP2p5EnxD66vVDL4PWRoWMzA56ZVHG24hpEDFShEz"),
									IsSigner:   true,
									IsWritable: false,
								},
							},
							Data: []byte{12, 1, 0, 0, 0, 0, 0, 0, 0, 9},
						},
					},
					RecentBlockhash: "5YjqMBZNwqmoUXkpoL4isLNwkaa2zuqxpRMBob47Bjxd",
					AddressLookupTableAccounts: []AddressLookupTableAccount{
						{
							Key: common.PublicKeyFromString("4jBXhGD8X8i2MCkunSDnqvyzQrGcfV6rqy5A4ETJBtaA"),
							Addresses: []common.PublicKey{
								common.PublicKeyFromString("5XaEXmAEiA4t3EdFWADixN9537Nct5Y5PMRz391eD9N1"),
								common.PublicKeyFromString("CPaB3EuV5qJK25stSWzH3815BspeyGgYvaR1Z8B72hbp"),
							},
						},
						{
							Key: common.PublicKeyFromString("F5wakDtup2KKx1SACvLyYDJn2r6eMGRwQDTw7ZKBWATb"),
							Addresses: []common.PublicKey{
								common.PublicKeyFromString("5XaEXmAEiA4t3EdFWADixN9537Nct5Y5PMRz391eD9N1"),
								common.PublicKeyFromString("8YNmYW9rWwpmLxUDycqHj1JMAMdm1v2VBB55tXqt7jej"),
							},
						},
					},
				},
			},
			want: Message{
				Version: MessageVersionV0,
				Header: MessageHeader{
					NumRequireSignatures:        1,
					NumReadonlySignedAccounts:   0,
					NumReadonlyUnsignedAccounts: 1,
				},
				Accounts: []common.PublicKey{
					common.PublicKeyFromString("FUarP2p5EnxD66vVDL4PWRoWMzA56ZVHG24hpEDFShEz"),
					common.TokenProgramID,
				},
				RecentBlockHash: "5YjqMBZNwqmoUXkpoL4isLNwkaa2zuqxpRMBob47Bjxd",
				Instructions: []CompiledInstruction{
					{
						ProgramIDIndex: 1,
						Accounts:       []int{3, 4, 2, 0},
						Data:           []byte{12, 1, 0, 0, 0, 0, 0, 0, 0, 9},
					},
				},
				AddressLookupTables: []CompiledAddressLookupTable{
					{
						AccountKey:      common.PublicKeyFromString("4jBXhGD8X8i2MCkunSDnqvyzQrGcfV6rqy5A4ETJBtaA"),
						WritableIndexes: []uint8{1},
						ReadonlyIndexes: []uint8{0},
					},
					{
						AccountKey:      common.PublicKeyFromString("F5wakDtup2KKx1SACvLyYDJn2r6eMGRwQDTw7ZKBWATb"),
						WritableIndexes: []uint8{1},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, NewMessage(tt.args.param), tt.want)
		})
	}
}

func TestMessageSerialize(t *testing.T) {
	type fields struct {
		Version             MessageVersion
		Header              MessageHeader
		Accounts            []common.PublicKey
		RecentBlockHash     string
		Instructions        []CompiledInstruction
		AddressLookupTables []CompiledAddressLookupTable
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
		err    error
	}{
		{
			fields: fields{
				Header: MessageHeader{
					NumRequireSignatures:        1,
					NumReadonlySignedAccounts:   0,
					NumReadonlyUnsignedAccounts: 1,
				},
				Accounts: []common.PublicKey{
					common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					common.PublicKeyFromString("A4iUVr5KjmsLymUcv4eSKPedUtoaBceiPeGipKMYc69b"),
					common.SystemProgramID,
				},
				RecentBlockHash: "FwRYtTPRk5N4wUeP87rTw9kQVSwigB6kbikGzzeCMrW5",
				Instructions: []CompiledInstruction{
					{
						ProgramIDIndex: 2,
						Accounts:       []int{0, 1},
						Data:           []byte{2, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
					},
				},
			},
			want: []byte{1, 0, 1, 3, 206, 211, 135, 230, 195, 111, 87, 254, 147, 239, 143, 81, 110, 159, 49, 140, 109, 137, 224, 197, 24, 49, 223, 61, 123, 8, 78, 109, 110, 136, 228, 240, 134, 172, 209, 213, 227, 137, 61, 108, 116, 171, 205, 124, 54, 68, 61, 110, 80, 31, 240, 117, 108, 137, 97, 222, 38, 242, 68, 156, 27, 65, 29, 142, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 221, 244, 189, 59, 8, 252, 7, 91, 129, 169, 22, 151, 32, 104, 208, 131, 64, 75, 232, 201, 77, 13, 187, 220, 103, 232, 190, 100, 35, 210, 17, 42, 1, 2, 2, 0, 1, 12, 2, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
			err:  nil,
		},
		{
			fields: fields{
				Version: MessageVersionV0,
				Header: MessageHeader{
					NumRequireSignatures:        1,
					NumReadonlySignedAccounts:   0,
					NumReadonlyUnsignedAccounts: 1,
				},
				Accounts: []common.PublicKey{
					common.PublicKeyFromString("FUarP2p5EnxD66vVDL4PWRoWMzA56ZVHG24hpEDFShEz"),
					common.TokenProgramID,
				},
				RecentBlockHash: "8QYt53pDt3jMhgFKWWeGPkbpPprGBp7mTx68q6vv5JW1",
				Instructions: []CompiledInstruction{
					{
						ProgramIDIndex: 1,
						Accounts:       []int{3, 4, 2, 0},
						Data:           []byte{12, 1, 0, 0, 0, 0, 0, 0, 0, 9},
					},
				},
				AddressLookupTables: []CompiledAddressLookupTable{
					{
						AccountKey:      common.PublicKeyFromString("4jBXhGD8X8i2MCkunSDnqvyzQrGcfV6rqy5A4ETJBtaA"),
						WritableIndexes: []uint8{1},
						ReadonlyIndexes: []uint8{0},
					},
					{
						AccountKey:      common.PublicKeyFromString("F5wakDtup2KKx1SACvLyYDJn2r6eMGRwQDTw7ZKBWATb"),
						WritableIndexes: []uint8{1},
					},
				},
			},
			want: []byte{128, 1, 0, 1, 2, 215, 20, 147, 30, 186, 106, 25, 168, 244, 220, 108, 1, 154, 255, 38, 79, 95, 191, 104, 197, 162, 142, 224, 179, 185, 135, 85, 206, 57, 214, 73, 211, 6, 221, 246, 225, 215, 101, 161, 147, 217, 203, 225, 70, 206, 235, 121, 172, 28, 180, 133, 237, 95, 91, 55, 145, 58, 140, 245, 133, 126, 255, 0, 169, 110, 10, 54, 233, 206, 100, 206, 20, 210, 67, 23, 247, 30, 228, 82, 91, 213, 121, 154, 103, 2, 244, 121, 216, 91, 51, 89, 238, 234, 100, 201, 70, 1, 1, 4, 3, 4, 2, 0, 10, 12, 1, 0, 0, 0, 0, 0, 0, 0, 9, 2, 55, 97, 92, 52, 123, 217, 238, 66, 226, 228, 18, 46, 33, 216, 61, 49, 147, 61, 56, 53, 154, 58, 97, 207, 99, 252, 242, 109, 33, 155, 109, 79, 1, 1, 1, 0, 209, 71, 167, 243, 125, 4, 5, 3, 86, 158, 20, 79, 26, 218, 111, 112, 201, 138, 90, 45, 166, 173, 184, 149, 122, 87, 238, 41, 150, 221, 227, 178, 1, 1, 0},
			err:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				Version:             tt.fields.Version,
				Header:              tt.fields.Header,
				Accounts:            tt.fields.Accounts,
				RecentBlockHash:     tt.fields.RecentBlockHash,
				Instructions:        tt.fields.Instructions,
				AddressLookupTables: tt.fields.AddressLookupTables,
			}
			got, err := m.Serialize()
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestMessageDeserialize(t *testing.T) {
	type args struct {
		messageData []byte
	}
	tests := []struct {
		name string
		args args
		want Message
		err  error
	}{
		{
			args: args{messageData: []byte{1, 0, 1, 3, 206, 211, 135, 230, 195, 111, 87, 254, 147, 239, 143, 81, 110, 159, 49, 140, 109, 137, 224, 197, 24, 49, 223, 61, 123, 8, 78, 109, 110, 136, 228, 240, 134, 172, 209, 213, 227, 137, 61, 108, 116, 171, 205, 124, 54, 68, 61, 110, 80, 31, 240, 117, 108, 137, 97, 222, 38, 242, 68, 156, 27, 65, 29, 142, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 221, 244, 189, 59, 8, 252, 7, 91, 129, 169, 22, 151, 32, 104, 208, 131, 64, 75, 232, 201, 77, 13, 187, 220, 103, 232, 190, 100, 35, 210, 17, 42, 1, 2, 2, 0, 1, 12, 2, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0}},
			want: Message{
				Version: MessageVersionLegacy,
				Header: MessageHeader{
					NumRequireSignatures:        1,
					NumReadonlySignedAccounts:   0,
					NumReadonlyUnsignedAccounts: 1,
				},
				Accounts: []common.PublicKey{
					common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
					common.PublicKeyFromString("A4iUVr5KjmsLymUcv4eSKPedUtoaBceiPeGipKMYc69b"),
					common.SystemProgramID,
				},
				RecentBlockHash: "FwRYtTPRk5N4wUeP87rTw9kQVSwigB6kbikGzzeCMrW5",
				Instructions: []CompiledInstruction{
					{
						ProgramIDIndex: 2,
						Accounts:       []int{0, 1},
						Data:           []byte{2, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
					},
				},
				AddressLookupTables: []CompiledAddressLookupTable{},
			},
			err: nil,
		},
		{
			args: args{messageData: []byte{128}},
			want: Message{},
			err:  fmt.Errorf("message header #1 parse error: data is empty"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MessageDeserialize(tt.args.messageData)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.err, err)
		})
	}
}
