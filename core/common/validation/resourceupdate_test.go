/*
Copyright IBM Corp. 2016 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package validation

import (
	"testing"

	"github.com/docker/docker/pkg/testutil/assert"
	"github.com/hyperledger/mchain/common/mocks/config"
	"github.com/hyperledger/mchain/common/util"
	cb "github.com/hyperledger/mchain/protos/common"
	"github.com/hyperledger/mchain/protos/peer"
	"github.com/hyperledger/mchain/protos/utils"
)

func TestValidateResourceUpdateTx(t *testing.T) {
	chainID := util.GetTestChainID()

	updateResult := &cb.Envelope{
		Payload: utils.MarshalOrPanic(&cb.Payload{Header: &cb.Header{
			ChannelHeader: utils.MarshalOrPanic(&cb.ChannelHeader{
				Type:      int32(cb.HeaderType_PEER_RESOURCE_UPDATE),
				ChannelId: chainID,
			}),
			SignatureHeader: utils.MarshalOrPanic(&cb.SignatureHeader{
				Creator: signerSerialized,
				Nonce:   utils.CreateNonceOrPanic(),
			}),
		},
			Data: utils.MarshalOrPanic(&cb.ConfigEnvelope{LastUpdate: &cb.Envelope{}}),
		}),
	}
	updateResult.Signature, _ = signer.Sign(updateResult.Payload)
	_, txResult := ValidateTransaction(updateResult, &config.MockApplicationCapabilities{})
	assert.Equal(t, txResult, peer.TxValidationCode_UNSUPPORTED_TX_PAYLOAD)
	_, txResult = ValidateTransaction(updateResult, &config.MockApplicationCapabilities{ResourcesTreeRv: true})
	assert.Equal(t, txResult, peer.TxValidationCode_VALID)
}
