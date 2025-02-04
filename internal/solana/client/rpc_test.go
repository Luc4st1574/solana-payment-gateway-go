package client

import (
	"reflect"
	"testing"

	"bitsnake-server/internal/solana/common"
	"bitsnake-server/internal/solana/rpc"
)

func Test_convertReturnData(t *testing.T) {
	type args struct {
		d rpc.ReturnData
	}
	tests := []struct {
		name    string
		args    args
		want    ReturnData
		wantErr bool
	}{
		{
			args: args{
				d: rpc.ReturnData{
					ProgramId: "35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP",
					Data:      []any{"AQIDBAU=", "base64"},
				},
			},
			want: ReturnData{
				ProgramId: common.PublicKeyFromString("35HSbe2xiLfid5QJeETGnUsGhkAiJWRKPrEGdQQ5xXrP"),
				Data:      []byte{1, 2, 3, 4, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertReturnData(tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertReturnData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertReturnData() = %v, want %v", got, tt.want)
			}
		})
	}
}
