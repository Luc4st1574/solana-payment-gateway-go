package hdwallet

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func TestVerifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{path: "m"}, want: true},
		{args: args{path: "m/"}, want: false},
		{args: args{path: "m/44"}, want: false},
		{args: args{path: "m/44/"}, want: false},
		{args: args{path: "m/44'"}, want: true},
		{args: args{path: "m/44'/"}, want: false},
		{args: args{path: "m/100000000'"}, want: true},
		{args: args{path: "m/44/501'/1'"}, want: false},
		{args: args{path: "m/44'/501/1'"}, want: false},
		{args: args{path: "m/44'/501/1'/0'"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidPath(tt.args.path); got != tt.want {
				t.Errorf("VerifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDerived(t *testing.T) {
	type args struct {
		seed []byte
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    Key
		wantErr bool
	}{
		{
			args: args{
				seed: mustDecodeHex("000102030405060708090a0b0c0d0e0f"),
				path: "m",
			},
			want: Key{
				ChainCode:  mustDecodeHex("90046a93de5380a72b5e45010748567d5ea02bbf6522f979e05c0d8d8ca9fffb"),
				PrivateKey: mustDecodeHex("2b4be7f19ee27bbf30c667b642d5f4aa69fd169872f8fc3059c08ebae2eb19e7"),
			},
		},
		{
			args: args{
				seed: mustDecodeHex("fffcf9f6f3f0edeae7e4e1dedbd8d5d2cfccc9c6c3c0bdbab7b4b1aeaba8a5a29f9c999693908d8a8784817e7b7875726f6c696663605d5a5754514e4b484542"),
				path: "m/0'/2147483647'/1'/2147483646'/2'",
			},
			want: Key{
				ChainCode:  mustDecodeHex("5d70af781f3a37b829f0d060924d5e960bdc02e85423494afc0b1a41bbe196d4"),
				PrivateKey: mustDecodeHex("551d333177df541ad876a60ea71f00447931c0a9da16f227c11ea080d7391b8d"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Derived(tt.args.path, tt.args.seed)
			if (err != nil) != tt.wantErr {
				t.Errorf("Derived() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Derived() = %v, want %v", got, tt.want)
			}
		})
	}
}

func mustDecodeHex(s string) []byte {
	h, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return h
}
