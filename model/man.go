package model

type Man struct {
	Woman
}

// Each of this * means a pointer reference to the name of the structs,
// this is how go can attached this functions to the struct (POO)
func (m *Man) Breath()       { m.breathing = true }
func (m *Man) Think()        { m.thinking = true }
func (m *Man) Eat()          { m.eating = true }
func (m *Man) Sex() string   { return "Man" }
func (m *Man) IsAlive() bool { return m.isAlive }
