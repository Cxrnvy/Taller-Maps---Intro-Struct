package farmacia

import (
	"errors"
	"time"
)

type Medicine struct {
	name            string
	manufacturer    string
	manufactureDate time.Time
	shelfLife       int
}

func (m Medicine) GetName() string {
	return m.name
}

func (m *Medicine) SetName(name string) error {
	if name == "" {
		return errors.New("el nombre no puede estar vacio")
	}
	m.name = name
	return nil
}
func (m Medicine) CalcularCaducidad() time.Time {
	return m.manufactureDate.AddDate(0, m.shelfLife, 0)
}
