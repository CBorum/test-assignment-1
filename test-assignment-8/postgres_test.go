package postgres

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testDB dbi

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func setup() {
	var err error
	testDB, err = NewDB("appdev", "appdev")
	if err != nil {
		panic(err)
	}
}

func tearDown() {
	testDB.Close()
}

func TestDBPing(t *testing.T) {
	err := testDB.Ping()
	assert.Nil(t, err)
}

func TestDBQuery1(t *testing.T) {
	res, err := Genre18And20(testDB)
	assert.Nil(t, err)

	last := res[len(res)-1]
	assert.Equal(t, int64(3237), last.trackid)
	assert.Equal(t, 39, len(res))
}

func TestDBQuery2(t *testing.T) {
	res, err := Invoices(testDB)
	assert.Nil(t, err)

	last := res[len(res)-1]
	assert.Equal(t, int64(319), last.invoiceid)
	assert.Equal(t, 115, len(res))
}

func TestDBQuery3(t *testing.T) {
	res, err := USACustomers(testDB)
	assert.Nil(t, err)

	last := res[len(res)-1]
	assert.Equal(t, int64(20), last.customerid)
	assert.Equal(t, 11, len(res))
}

func TestDBQuery4(t *testing.T) {
	res, err := MozartAndBach(testDB)
	assert.Nil(t, err)

	last := res[len(res)-1]
	assert.Equal(t, int64(327), last.albumid)
	assert.Equal(t, 12, len(res))
}
