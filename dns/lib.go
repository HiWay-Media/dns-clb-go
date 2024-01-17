package dns

import (
	"fmt"
)

type Address struct {
	Address string
	Port    uint16
}

func (a Address) String() string {
	return fmt.Sprintf("%s:%d", a.Address, a.Port)
}
