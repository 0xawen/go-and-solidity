package helper

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"reflect"
	"testing"
)

func TestCheckAccount(t *testing.T) {
	for i, account := range AccountMap {
		privateK, err := crypto.HexToECDSA(account.PrivateKey)
		if err != nil {
			fmt.Println("hex to private err:", privateK)
		}
		publicK := privateK.Public().(*ecdsa.PublicKey)
		if crypto.PubkeyToAddress(*publicK).Hex() != account.Address {
			t.Log("check error,index:", i)
		}
	}
}

func TestGetAccountAddress(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		args args
		want common.Address
	}{
		{
			name: "test 0,should passed",
			args: args{
				index: 0,
			},
			want: common.HexToAddress("0x3c2Ea2A42529C55e914420f75bB3c25136E2b876"),
		},
		{
			name: "test 10,out of range",
			args: args{
				index: 10,
			},
			want: common.HexToAddress("0x0"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAccountAddress(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccountAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
