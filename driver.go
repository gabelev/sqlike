package sqlike

import (
	"database/sql/driver"
	"fmt"
	"sync"
)

type sqlikeDriver struct {
	sync.Mutex
	counter     int
	connections map[string]*Sqlike
}

func (d sqlikeDriver) Open(name string) (connection driver.Conn, err error) {
	d.Lock()
	defer d.Unlock()

	db, ok := d.connections[name]
	if !ok {
		return db, fmt.Errorf("There is no connection available")
	}
	db.opened++
	return db, err
}
