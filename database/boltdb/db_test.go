package boltdb

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/ldmtam/raft-auto-increment/database"
	"github.com/stretchr/testify/assert"
)

var (
	testDB database.AutoIncrement
)

func TestMain(m *testing.M) {
	var err error
	testDB, err = New("./test")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	exitCode := m.Run()

	testDB.Close()
	os.RemoveAll("./test")

	os.Exit(exitCode)
}

func Test_GetOne(t *testing.T) {
	result, err := testDB.GetOne("key1")
	assert.Nil(t, err)
	assert.EqualValues(t, 1, result)

	result, err = testDB.GetOne("key1")
	assert.Nil(t, err)
	assert.EqualValues(t, 2, result)

	result, err = testDB.GetOne("key2")
	assert.Nil(t, err)
	assert.EqualValues(t, 1, result)
}

func Test_GetMany(t *testing.T) {
	result, err := testDB.GetMany("key1", 5)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 5)
	fmt.Println(result)
}

func Test_GetLastInserted(t *testing.T) {
	value, err := testDB.GetOne("key1")
	assert.Nil(t, err)

	last, err := testDB.GetLastInserted("key1")
	assert.Nil(t, err)
	assert.EqualValues(t, value, last)
}

func Test_Backup(t *testing.T) {
	inserted, err := testDB.GetOne("key1")
	assert.Nil(t, err)

	snapshot, err := testDB.Backup()
	assert.Nil(t, err)

	err = os.MkdirAll("./backup", 0777)
	assert.Nil(t, err)
	defer os.RemoveAll("./backup")

	err = ioutil.WriteFile("./backup/data.db", snapshot, 0777)
	assert.Nil(t, err)

	db2, err := New("./backup")
	assert.Nil(t, err)
	defer db2.Close()

	value, err := db2.GetLastInserted("key1")
	assert.Nil(t, err)
	assert.EqualValues(t, inserted, value)
}
