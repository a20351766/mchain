/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package historyleveldb

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/hyperledger/mchain/common/ledger/blkstorage"
	"github.com/hyperledger/mchain/common/ledger/blkstorage/fsblkstorage"
	"github.com/hyperledger/mchain/common/ledger/testutil"
	"github.com/hyperledger/mchain/core/ledger/kvledger/history/historydb"
	"github.com/hyperledger/mchain/core/ledger/kvledger/txmgmt/privacyenabledstate"
	"github.com/hyperledger/mchain/core/ledger/kvledger/txmgmt/txmgr"
	"github.com/hyperledger/mchain/core/ledger/kvledger/txmgmt/txmgr/lockbasedtxmgr"
	"github.com/hyperledger/mchain/core/ledger/ledgerconfig"
	"github.com/spf13/viper"
)

/////// levelDBLockBasedHistoryEnv //////

type levelDBLockBasedHistoryEnv struct {
	t                   testing.TB
	testBlockStorageEnv *testBlockStoreEnv

	testDBEnv privacyenabledstate.TestEnv
	txmgr     txmgr.TxMgr

	testHistoryDBProvider historydb.HistoryDBProvider
	testHistoryDB         historydb.HistoryDB
}

func newTestHistoryEnv(t *testing.T) *levelDBLockBasedHistoryEnv {
	viper.Set("ledger.history.enableHistoryDatabase", "true")
	testLedgerID := "TestLedger"

	blockStorageTestEnv := newBlockStorageTestEnv(t)

	testDBEnv := &privacyenabledstate.LevelDBCommonStorageTestEnv{}
	testDBEnv.Init(t)
	testDB := testDBEnv.GetDBHandle(testLedgerID)

	txMgr := lockbasedtxmgr.NewLockBasedTxMgr(testLedgerID, testDB, nil)
	testHistoryDBProvider := NewHistoryDBProvider()
	testHistoryDB, err := testHistoryDBProvider.GetDBHandle("TestHistoryDB")
	testutil.AssertNoError(t, err, "")

	return &levelDBLockBasedHistoryEnv{t,
		blockStorageTestEnv, testDBEnv,
		txMgr, testHistoryDBProvider, testHistoryDB}
}

func (env *levelDBLockBasedHistoryEnv) cleanup() {
	defer env.txmgr.Shutdown()
	defer env.testDBEnv.Cleanup()
	defer env.testBlockStorageEnv.cleanup()

	// clean up history
	env.testHistoryDBProvider.Close()
	removeDBPath(env.t)
}

func removeDBPath(t testing.TB) {
	removePath(t, ledgerconfig.GetHistoryLevelDBPath())
}

func removePath(t testing.TB, path string) {
	if err := os.RemoveAll(path); err != nil {
		t.Fatalf("Err: %s", err)
		t.FailNow()
	}
}

/////// testBlockStoreEnv//////

type testBlockStoreEnv struct {
	t               testing.TB
	provider        *fsblkstorage.FsBlockstoreProvider
	blockStorageDir string
}

func newBlockStorageTestEnv(t testing.TB) *testBlockStoreEnv {

	testPath, err := ioutil.TempDir("", "historyleveldb-")
	if err != nil {
		panic(err)
	}
	conf := fsblkstorage.NewConf(testPath, 0)

	attrsToIndex := []blkstorage.IndexableAttr{
		blkstorage.IndexableAttrBlockHash,
		blkstorage.IndexableAttrBlockNum,
		blkstorage.IndexableAttrTxID,
		blkstorage.IndexableAttrBlockNumTranNum,
	}
	indexConfig := &blkstorage.IndexConfig{AttrsToIndex: attrsToIndex}

	blockStorageProvider := fsblkstorage.NewProvider(conf, indexConfig).(*fsblkstorage.FsBlockstoreProvider)

	return &testBlockStoreEnv{t, blockStorageProvider, testPath}
}

func (env *testBlockStoreEnv) cleanup() {
	env.provider.Close()
	env.removeFSPath()
}

func (env *testBlockStoreEnv) removeFSPath() {
	fsPath := env.blockStorageDir
	os.RemoveAll(fsPath)
}
