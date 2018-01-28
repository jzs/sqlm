package sqlm

import (
	"fmt"
	"reflect"
	"strings"
)

// Table interface that all structs must implement
type Table interface {
	Table() string
}

// CreateTable generates an sql string to create a table based on a struct
func CreateTable(i Table) string {
	t := reflect.TypeOf(i)
	str := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v (", i.Table())
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		fname := f.Tag.Get("db")
		if fname == "" {
			fname = f.Name
		}

		if i == 0 {
			str = fmt.Sprintf("%v%v %v", str, fname, f.Tag.Get("type"))
			continue
		}

		str = fmt.Sprintf("%v, %v %v", str, fname, f.Tag.Get("type"))
	}
	str = fmt.Sprintf("%v)", str)
	return str
}

// Insert creates an insert statement for updating all columns. (Ignores bigserial column for now)
// Returns an sql string and the associated values based on the contents on the struct.
func Insert(i Table) (string, []interface{}) {
	t := reflect.TypeOf(i)
	val := reflect.ValueOf(i)

	str := fmt.Sprintf("INSERT INTO %v", i.Table())

	values := ""
	keys := ""
	vals := []interface{}{}

	cnt := -1
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		typ := f.Tag.Get("type")
		if strings.Contains(strings.ToLower(typ), "bigserial") {
			continue
		}
		cnt++

		v := val.FieldByName(f.Name)
		vals = append(vals, v.Interface())

		fname := f.Tag.Get("db")
		if fname == "" {
			fname = f.Name
		}

		if cnt == 0 {
			keys = fmt.Sprintf("'%v'", fname)
			values = fmt.Sprintf("?")
			continue
		}
		keys = fmt.Sprintf("%v, '%v'", keys, fname)
		values = fmt.Sprintf("%v,?", values)
	}

	str = fmt.Sprintf("%v (%v) VALUES (%v)", str, keys, values)

	return str, vals
}

// Update creates an update sql string and a slice of values based on the given Table
func Update(i Table, where string) (string, []interface{}) {
	t := reflect.TypeOf(i)
	val := reflect.ValueOf(i)

	str := fmt.Sprintf("UPDATE %v SET ", i.Table())

	keys := ""
	vals := []interface{}{}

	cnt := -1
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		typ := f.Tag.Get("type")
		if strings.Contains(strings.ToLower(typ), "bigserial") {
			continue
		}
		cnt++

		v := val.FieldByName(f.Name)
		vals = append(vals, v.Interface())

		fname := f.Tag.Get("db")
		if fname == "" {
			fname = f.Name
		}

		if cnt == 0 {
			keys = fmt.Sprintf("'%v' = ?", fname)
			continue
		}
		keys = fmt.Sprintf("%v, '%v' = ?", keys, fname)
	}

	str = fmt.Sprintf("%v %v", str, keys)

	if where != "" {
		str = fmt.Sprintf("%v WHERE %v", str, where)
	}

	return str, vals

}

// DropTable generates an sql string to drop table based on a struct
func DropTable(i Table) string {
	str := fmt.Sprintf("DROP TABLE IF EXISTS %v", i.Table())
	return str
}
