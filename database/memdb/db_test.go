package memdb

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ldmtam/raft-auto-increment/database"
)

var (
	testDB database.AutoIncrement
)

func TestMain(m *testing.M) {
	var err error

	testDB, err = New(nil)
	if err != nil {
		panic(err)
	}

	exitCode := m.Run()

	testDB.Close()

	os.Exit(exitCode)
}

func Test_SetAndGetLastInserted(t *testing.T) {
	err := testDB.Set("foo", 1)
	assert.Nil(t, err)

	value, err := testDB.GetLastInserted("foo")
	assert.Nil(t, err)
	assert.EqualValues(t, 1, value)
}

func Test_Backup(t *testing.T) {
	err := testDB.Set("foo", 1)
	assert.Nil(t, err)

	err = testDB.Set("bar", 2)
	assert.Nil(t, err)

	backup, err := testDB.Backup()
	assert.Nil(t, err)

	newDB, err := New(backup)
	assert.Nil(t, err)

	val, err := newDB.GetLastInserted("foo")
	assert.Nil(t, err)
	assert.EqualValues(t, 1, val)

	val, err = newDB.GetLastInserted("bar")
	assert.Nil(t, err)
	assert.EqualValues(t, 2, val)
}
