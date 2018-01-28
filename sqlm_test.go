package sqlm_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/sketchground/sqlm"
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
