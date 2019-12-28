package leveldb

import (
	"fmt"
	"os"
	"testing"

	"github.com/ldmtam/raft-auto-increment/auto_increment/database"
	"github.com/stretchr/testify/assert"
)

var (
	testDB database.AutoIncrement
)

func TestMain(m *testing.M) {
	testDB, _ = New("./test")

	exitCode := m.Run()

	testDB.Close()
	os.RemoveAll("./test")

	os.Exit(exitCode)
}

func Test_GetSingle(t *testing.T) {
	result, err := testDB.GetSingle("key1")
	assert.Nil(t, err)
	assert.EqualValues(t, 1, result)

	result, err = testDB.GetSingle("key1")
	assert.Nil(t, err)
	assert.EqualValues(t, 2, result)

	result, err = testDB.GetSingle("key2")
	assert.Nil(t, err)
	assert.EqualValues(t, 1, result)
}

func Test_GetMultiple(t *testing.T) {
	result, err := testDB.GetMultiple("key1", 5)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 5)
	fmt.Println(result)
}
