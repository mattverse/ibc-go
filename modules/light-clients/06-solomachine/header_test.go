package solomachine_test

import (
	ibctesting "github.com/cosmos/ibc-go/testing"
	"github.com/cosmos/ibc-go/v5/modules/core/exported"
	solomachine "github.com/cosmos/ibc-go/v5/modules/light-clients/06-solomachine"
)

func (suite *SoloMachineTestSuite) TestHeaderValidateBasic() {
	// test singlesig and multisig public keys
	for _, sm := range []*ibctesting.Solomachine{suite.solomachine, suite.solomachineMulti} {

		header := sm.CreateHeader(sm.Diversifier)

		cases := []struct {
			name    string
			header  *solomachine.Header
			expPass bool
		}{
			{
				"valid header",
				header,
				true,
			},
			{
				"sequence is zero",
				&solomachine.Header{
					Sequence:       0,
					Timestamp:      header.Timestamp,
					Signature:      header.Signature,
					NewPublicKey:   header.NewPublicKey,
					NewDiversifier: header.NewDiversifier,
				},
				false,
			},
			{
				"timestamp is zero",
				&solomachine.Header{
					Sequence:       header.Sequence,
					Timestamp:      0,
					Signature:      header.Signature,
					NewPublicKey:   header.NewPublicKey,
					NewDiversifier: header.NewDiversifier,
				},
				false,
			},
			{
				"signature is empty",
				&solomachine.Header{
					Sequence:       header.Sequence,
					Timestamp:      header.Timestamp,
					Signature:      []byte{},
					NewPublicKey:   header.NewPublicKey,
					NewDiversifier: header.NewDiversifier,
				},
				false,
			},
			{
				"diversifier contains only spaces",
				&solomachine.Header{
					Sequence:       header.Sequence,
					Timestamp:      header.Timestamp,
					Signature:      header.Signature,
					NewPublicKey:   header.NewPublicKey,
					NewDiversifier: " ",
				},
				false,
			},
			{
				"public key is nil",
				&solomachine.Header{
					Sequence:       header.Sequence,
					Timestamp:      header.Timestamp,
					Signature:      header.Signature,
					NewPublicKey:   nil,
					NewDiversifier: header.NewDiversifier,
				},
				false,
			},
		}

		suite.Require().Equal(exported.Solomachine, header.ClientType())

		for _, tc := range cases {
			tc := tc

			suite.Run(tc.name, func() {
				err := tc.header.ValidateBasic()

				if tc.expPass {
					suite.Require().NoError(err)
				} else {
					suite.Require().Error(err)
				}
			})
		}
	}
}
