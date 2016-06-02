package sqlike

import (
	//	"database/sql"
	"database/sql/driver"
)

type Sqlike struct {
	ordered bool
	opened  int
	dsn     string
	conn    string
	driver  *sqlikeDriver
}

func (s Sqlike) Open(dns string) (d driver.Conn, e error) {
	return d, e
}

func (s Sqlike) Begin() (t driver.Tx, e error) {
	return t, e
}

func (s Sqlike) Close() (e error) {
	return e
}

func (s Sqlike) Prepare(stmt string) (ds driver.Stmt, e error) {
	return ds, e
}
