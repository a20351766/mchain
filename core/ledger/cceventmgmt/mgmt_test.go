/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package cceventmgmt

import (
	"os"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/mchain/common/flogging"
	"github.com/hyperledger/mchain/core/common/ccprovider"
	"github.com/hyperledger/mchain/protos/ledger/rwset/kvrwset"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	flogging.SetModuleLevel("eventmgmt", "debug")
	os.Exit(m.Run())
}
func TestCCEventMgmt(t *testing.T) {
	cc1Def := &ChaincodeDefinition{Name: "cc1", Version: "v1", Hash: []byte("cc1")}
	cc1DBArtifactsTar := []byte("cc1DBArtifacts")

	cc2Def := &ChaincodeDefinition{Name: "cc2", Version: "v1", Hash: []byte("cc2")}
	cc2DBArtifactsTar := []byte("cc2DBArtifacts")

	cc3Def := &ChaincodeDefinition{Name: "cc3", Version: "v1", Hash: []byte("cc3")}
	cc3DBArtifactsTar := []byte("cc3DBArtifacts")

	// cc1 is deployed and installed. cc2 is deployed but not installed. cc3 is not deployed but installed
	mockProvider := newMockProvider()
	mockProvider.setChaincodeInstalled(cc1Def, cc1DBArtifactsTar)
	mockProvider.setChaincodeDeployed("channel1", cc1Def)
	mockProvider.setChaincodeDeployed("channel1", cc2Def)
	mockProvider.setChaincodeInstalled(cc3Def, cc3DBArtifactsTar)
	setEventMgrForTest(newMgr(mockProvider))
	defer clearEventMgrForTest()

	handler1, handler2 := &mockHandler{}, &mockHandler{}
	eventMgr := GetMgr()
	assert.NotNil(t, eventMgr)
	eventMgr.Register("channel1", handler1)
	eventMgr.Register("channel2", handler2)

	cc2ExpectedEvent := &mockEvent{cc2Def, cc2DBArtifactsTar}
	cc3ExpectedEvent := &mockEvent{cc3Def, cc3DBArtifactsTar}

	// Deploy cc3 on chain1 - only handler1 should recieve event because cc3 is being deployed only on chain1
	eventMgr.HandleChaincodeDeploy("channel1", []*ChaincodeDefinition{cc3Def})
	assert.Contains(t, handler1.eventsRecieved, cc3ExpectedEvent)
	assert.NotContains(t, handler2.eventsRecieved, cc3ExpectedEvent)

	// Deploy cc3 on chain2 as well and this time handler2 should also recieve event
	eventMgr.HandleChaincodeDeploy("channel2", []*ChaincodeDefinition{cc3Def})
	assert.Contains(t, handler2.eventsRecieved, cc3ExpectedEvent)

	// Install CC2 - only handler1 should receive event because cc2 is deployed only on chain1 and not on chain2
	eventMgr.HandleChaincodeInstall(cc2Def, cc2DBArtifactsTar)
	assert.Contains(t, handler1.eventsRecieved, cc2ExpectedEvent)
	assert.NotContains(t, handler2.eventsRecieved, cc2ExpectedEvent)
}

func TestLSCCListener(t *testing.T) {
	channelName := "testChannel"

	cc1Def := &ChaincodeDefinition{Name: "testChaincode1", Version: "v1", Hash: []byte("hash_testChaincode")}
	cc2Def := &ChaincodeDefinition{Name: "testChaincode2", Version: "v1", Hash: []byte("hash_testChaincode")}
	cc3Def := &ChaincodeDefinition{Name: "testChaincode~collection", Version: "v1", Hash: []byte("hash_testChaincode")}

	ccDBArtifactsTar := []byte("ccDBArtifacts")

	// cc1, cc2, cc3 installed but not deployed
	mockProvider := newMockProvider()
	mockProvider.setChaincodeInstalled(cc1Def, ccDBArtifactsTar)
	mockProvider.setChaincodeInstalled(cc2Def, ccDBArtifactsTar)
	mockProvider.setChaincodeInstalled(cc3Def, ccDBArtifactsTar)

	setEventMgrForTest(newMgr(mockProvider))
	defer clearEventMgrForTest()
	handler1 := &mockHandler{}
	GetMgr().Register(channelName, handler1)
	lsccStateListener := &KVLedgerLSCCStateListener{}

	// test1 regular deploy lscc event gets sent to handler
	t.Run("DeployEvent", func(t *testing.T) {
		sampleChaincodeData1 := &ccprovider.ChaincodeData{Name: cc1Def.Name, Version: cc1Def.Version, Id: cc1Def.Hash}
		sampleChaincodeDataBytes1, err := proto.Marshal(sampleChaincodeData1)
		assert.NoError(t, err, "")
		lsccStateListener.HandleStateUpdates(channelName, []*kvrwset.KVWrite{
			{Key: cc1Def.Name, Value: sampleChaincodeDataBytes1},
		})
		assert.Contains(t, handler1.eventsRecieved, &mockEvent{cc1Def, ccDBArtifactsTar})
	})

	// test2 delete lscc event NOT sent to handler
	t.Run("DeleteEvent", func(t *testing.T) {
		sampleChaincodeData2 := &ccprovider.ChaincodeData{Name: cc2Def.Name, Version: cc2Def.Version, Id: cc2Def.Hash}
		sampleChaincodeDataBytes2, err := proto.Marshal(sampleChaincodeData2)
		assert.NoError(t, err, "")
		lsccStateListener.HandleStateUpdates(channelName, []*kvrwset.KVWrite{
			{Key: cc2Def.Name, Value: sampleChaincodeDataBytes2, IsDelete: true},
		})
		assert.NotContains(t, handler1.eventsRecieved, &mockEvent{cc2Def, ccDBArtifactsTar})
	})

	// test3 collection lscc event (with tilda separator in chaincode key) NOT sent to handler
	t.Run("CollectionEvent", func(t *testing.T) {
		sampleChaincodeData3 := &ccprovider.ChaincodeData{Name: cc3Def.Name, Version: cc3Def.Version, Id: cc3Def.Hash}
		sampleChaincodeDataBytes3, err := proto.Marshal(sampleChaincodeData3)
		assert.NoError(t, err, "")
		lsccStateListener.HandleStateUpdates(channelName, []*kvrwset.KVWrite{
			{Key: cc3Def.Name, Value: sampleChaincodeDataBytes3},
		})
		assert.NotContains(t, handler1.eventsRecieved, &mockEvent{cc3Def, ccDBArtifactsTar})
	})
}

type mockProvider struct {
	chaincodesDeployed  map[[3]string]bool
	chaincodesInstalled map[[2]string][]byte
}

type mockHandler struct {
	eventsRecieved []*mockEvent
}

type mockEvent struct {
	def            *ChaincodeDefinition
	dbArtifactsTar []byte
}

func (l *mockHandler) HandleChaincodeDeploy(chaincodeDefinition *ChaincodeDefinition, dbArtifactsTar []byte) error {
	l.eventsRecieved = append(l.eventsRecieved, &mockEvent{def: chaincodeDefinition, dbArtifactsTar: dbArtifactsTar})
	return nil
}

func newMockProvider() *mockProvider {
	return &mockProvider{
		make(map[[3]string]bool),
		make(map[[2]string][]byte),
	}
}

func (p *mockProvider) setChaincodeDeployed(chainid string, chaincodeDefinition *ChaincodeDefinition) {
	p.chaincodesDeployed[[3]string{chainid, chaincodeDefinition.Name, chaincodeDefinition.Version}] = true
}

func (p *mockProvider) setChaincodeInstalled(chaincodeDefinition *ChaincodeDefinition, dbArtifactsTar []byte) {
	p.chaincodesInstalled[[2]string{chaincodeDefinition.Name, chaincodeDefinition.Version}] = dbArtifactsTar
}

func (p *mockProvider) setChaincodeDeployAndInstalled(chainid string, chaincodeDefinition *ChaincodeDefinition, dbArtifactsTar []byte) {
	p.setChaincodeDeployed(chainid, chaincodeDefinition)
	p.setChaincodeInstalled(chaincodeDefinition, dbArtifactsTar)
}

func (p *mockProvider) IsChaincodeDeployed(chainid string, chaincodeDefinition *ChaincodeDefinition) (bool, error) {
	return p.chaincodesDeployed[[3]string{chainid, chaincodeDefinition.Name, chaincodeDefinition.Version}], nil
}

func (p *mockProvider) RetrieveChaincodeArtifacts(chaincodeDefinition *ChaincodeDefinition) (installed bool, dbArtifactsTar []byte, err error) {
	dbArtifactsTar, ok := p.chaincodesInstalled[[2]string{chaincodeDefinition.Name, chaincodeDefinition.Version}]
	if !ok {
		return false, nil, nil
	}
	return true, dbArtifactsTar, nil
}

func setEventMgrForTest(eventMgr *Mgr) {
	mgr = eventMgr
}

func clearEventMgrForTest() {
	mgr = nil
}