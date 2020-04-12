package bni

import (
	"encoding/hex"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/const/trans_type"
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/HyperService-Consortium/go-ves/config"
	dep_uip "github.com/HyperService-Consortium/go-ves/dependency/uip"
	"github.com/HyperService-Consortium/go-ves/types"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"testing"
)

func TestBN_Route(t *testing.T) {
	type fields struct {
		dns    types.ChainDNSInterface
		signer uip.Signer
	}
	type args struct {
		intent  uip.TransactionIntent
		storage uip.Storage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"test_easy", fields{
				dns: config.ChainDNS,
				signer: passwordSigner{
					pb: sugar.HandlerError(hex.DecodeString("6d8c6cb9d26b5a21ae498a22385ae4265f494cfc")).([]byte),
					ps: "123456"},
			}, args{
				intent: &dep_uip.TransactionIntent{
					TransactionIntent: &opintent.TransactionIntent{
						TransType: trans_type.Payment,
						Src:       sugar.HandlerError(hex.DecodeString("6d8c6cb9d26b5a21ae498a22385ae4265f494cfc")).([]byte),
						Dst:       sugar.HandlerError(hex.DecodeString("6d8c6cb9d26b5a21ae498a22385ae4265f494cfc")).([]byte),
						Meta:      nil,
						Amt:       "03e8",
						ChainID:   7,
					},
				},
				storage: nil,
			}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bn := &BN{
				dns:    tt.fields.dns,
				signer: tt.fields.signer,
			}
			got, err := bn.Route(tt.args.intent, tt.args.storage)
			if (err != nil) != tt.wantErr {
				t.Errorf("Route() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			fmt.Println(got)
			fmt.Println(hex.EncodeToString(got))
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("Route() got = %v, want %v", got, tt.want)
			//}
		})
	}
}
