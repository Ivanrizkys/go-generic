package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Employee interface {
	GetName() string
}

func GetName[T Employee](param T) string {
	return param.GetName()
}

// * Manager mengimplementasikan interface Employee
type Manager interface {
	GetName() string
	GetManagerName() string
}

// * MyManager mengimplementasikan interface Employee
type MyManager struct {
	Name string
}

func (m *MyManager) GetName() string {
	return m.Name
}

func (m *MyManager) GetManagerName() string {
	return m.Name
}

// * VicePrecident mengimplementasikan interface Employee
type VicePrecident interface {
	GetName() string
	GetVicePrecidentName() string
}

// * MyVicePrecident mengimplementasikan interface Employee
type MyVicePrecident struct {
	Name string
}

func (m *MyVicePrecident) GetName() string {
	return m.Name
}

func (m *MyVicePrecident) GetVicePrecidentName() string {
	return m.Name
}

func TestTypeInheritance(t *testing.T) {
	assert.Equal(t, "Ivan", GetName[Manager](&MyManager{Name: "Ivan"}))
	assert.Equal(t, "Aisyah", GetName[VicePrecident](&MyVicePrecident{Name: "Aisyah"}))
}
