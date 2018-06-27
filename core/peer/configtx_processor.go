/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package peer

import (
	"fmt"

	"github.com/hyperledger/mchain/common/channelconfig"
	"github.com/hyperledger/mchain/common/resourcesconfig"
	"github.com/hyperledger/mchain/core/ledger"
	"github.com/hyperledger/mchain/core/ledger/customtx"
	"github.com/hyperledger/mchain/protos/common"
	"github.com/hyperledger/mchain/protos/utils"
)

const (
	resourcesConfigKey = "resourcesconfigtx.RESOURCES_CONFIG_KEY"
	channelConfigKey   = "resourcesconfigtx.CHANNEL_CONFIG_KEY"
	peerNamespace      = ""
)

// txProcessor implements the interface 'github.com/hyperledger/mchain/core/ledger/customtx/Processor'
type configtxProcessor struct {
}

// newTxProcessor constructs a new instance of txProcessor
func newConfigTxProcessor() customtx.Processor {
	return &configtxProcessor{}
}

// GenerateSimulationResults implements function in the interface 'github.com/hyperledger/mchain/core/ledger/customtx/Processor'
// This implemantation processes following two types of transactions.
// CONFIG  - simply stores the config in the statedb. Additionally, stores the resource config seed if the transaction is from the genesis block.
// PEER_RESOURCE_UPDATE - In a normal course, this validates the transaction against the current resource bundle,
// computes the full configuration, and stores the full configuration if the transaction is found valid.
// However, if 'initializingLedger' is true (i.e., either the ledger is being created from the genesis block
// or the ledger is synching the state with the blockchain, during start up), the full config is computed using
// the most recent configs from statedb
func (tp *configtxProcessor) GenerateSimulationResults(txEnv *common.Envelope, simulator ledger.TxSimulator, initializingLedger bool) error {
	payload := utils.UnmarshalPayloadOrPanic(txEnv.Payload)
	channelHdr := utils.UnmarshalChannelHeaderOrPanic(payload.Header.ChannelHeader)
	chainid := channelHdr.ChannelId
	txType := common.HeaderType(channelHdr.GetType())

	switch txType {
	case common.HeaderType_CONFIG:
		peerLogger.Debugf("Processing CONFIG")
		return processChannelConfigTx(chainid, txEnv, simulator)

	case common.HeaderType_PEER_RESOURCE_UPDATE:
		peerLogger.Debugf("Processing PEER_RESOURCE_UPDATE")
		if initializingLedger {
			return processResourceConfigTxDuringInitialization(chainid, txEnv, simulator)
		}
		return processResourceConfigTx(chainid, txEnv, simulator)
	default:
		return fmt.Errorf("tx type [%s] is not expected", txType)
	}
}

func processChannelConfigTx(chainid string, txEnv *common.Envelope, simulator ledger.TxSimulator) error {
	configEnvelope := &common.ConfigEnvelope{}
	if _, err := utils.UnmarshalEnvelopeOfType(txEnv, common.HeaderType_CONFIG, configEnvelope); err != nil {
		return err
	}
	channelConfig := configEnvelope.Config

	if err := persistConf(simulator, channelConfigKey, channelConfig); err != nil {
		return err
	}

	peerLogger.Debugf("channelConfig=%s", channelConfig)
	if channelConfig == nil {
		return fmt.Errorf("Channel config found nil")
	}
	resConfCapabilityOn, err := isResConfigCapabilityOn(chainid, channelConfig)
	if err != nil {
		return err
	}
	resourceConfigSeed, err := extractFullConfigFromSeedTx(configEnvelope)
	if err != nil {
		return err
	}

	if channelConfig.Sequence == 1 && resConfCapabilityOn {
		if resourceConfigSeed == nil {
			return fmt.Errorf("Resource config cannot be nil in the genesis ('CONFIG') transaction")
		}
		return persistConf(simulator, resourcesConfigKey, resourceConfigSeed)
	}

	return nil
}

func processResourceConfigTx(chainid string, txEnv *common.Envelope, simulator ledger.TxSimulator) error {
	fullResConf, err := validateAndApplyResourceConfig(chainid, txEnv)
	if err != nil {
		return err
	}
	return persistConf(simulator, resourcesConfigKey, fullResConf)
}

func processResourceConfigTxDuringInitialization(chainid string, txEnv *common.Envelope, simulator ledger.TxSimulator) error {
	var existingResConf, existingChanConf, updatedResConf *common.Config
	var err error

	if existingResConf, err = retrievePersistedConf(simulator, resourcesConfigKey); err != nil {
		return err
	}
	if existingChanConf, err = retrievePersistedConf(simulator, channelConfigKey); err != nil {
		return err
	}

	if existingResConf == nil || existingChanConf == nil {
		return fmt.Errorf("Channel config or resource config should not be nil")
	}

	chanConfigBundle, err := channelconfig.NewBundle(chainid, existingChanConf)
	if err != nil {
		return err
	}

	resConfigBundle, err := resourcesconfig.NewBundle(chainid, existingResConf, chanConfigBundle)
	if err != nil {
		return err
	}
	if updatedResConf, err = computeFullConfig(resConfigBundle, txEnv); err != nil {
		return err
	}
	return persistConf(simulator, resourcesConfigKey, updatedResConf)
}

func persistConf(simulator ledger.TxSimulator, key string, config *common.Config) error {
	serializedConfig, err := serialize(config)
	if err != nil {
		return err
	}
	return simulator.SetState(peerNamespace, key, serializedConfig)
}

func retrievePersistedConf(queryExecuter ledger.QueryExecutor, key string) (*common.Config, error) {
	serializedConfig, err := queryExecuter.GetState(peerNamespace, key)
	if err != nil {
		return nil, err
	}
	if serializedConfig == nil {
		return nil, nil
	}
	return deserialize(serializedConfig)
}
