package sqlm_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/jzs/sqlm"
)

// DBShopgunAuth struct to represent an auth token
type DBShopgunAuth struct {
	ID      uint64    `db:"id" type:"BIGSERIAL PRIMARY KEY NOT NULL"`
	Token   string    `db:"token" type:"text NOT NULL"`
	Expires time.Time `type:"timestamp NOT NULL"`
}

func (d DBShopgunAuth) Table() string {
	return "shopgun_auth"
}

func TestCreateTable(t *testing.T) {
	tbl := DBShopgunAuth{Token: "token"}
	str := sqlm.CreateTable(tbl)
	fmt.Println(str)

	str, vals := sqlm.Insert(tbl)
	fmt.Println(str)
	fmt.Println(vals)

	str, vals = sqlm.Update(tbl, "id = ?")
	vals = append(vals, 1) // WHERE value...
	fmt.Println(str)
	fmt.Println(vals)

	str = sqlm.Get(tbl, "WHERE id = ?")
	fmt.Println(str)
}

var expected string = "SELECT id, token, Expires FROM shopgun_auth"

func TestGetTableSlicePointers(t *testing.T) {
	tbl := []struct {
		In any
	}{
		{In: &SliceTestStruct{ID: 132}},
		{In: SliceTestStruct{ID: 132}},
		{In: []SliceTestStruct{}},
		{In: []*SliceTestStruct{}},
	}
	for _, test := range tbl {
		r := sqlm.Get(test.In, "")
		if r != expected {
			t.Fatalf("expect: %v, got %v", expected, r)
		}
	}
}

type SliceTestStruct struct {
	ID      uint64    `db:"id" type:"BIGSERIAL PRIMARY KEY NOT NULL"`
	Token   string    `db:"token" type:"text NOT NULL"`
	Expires time.Time `type:"timestamp NOT NULL"`
}

func (d SliceTestStruct) Table() string {
	return "shopgun_auth"
}
