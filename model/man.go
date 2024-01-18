package model

type Man struct {
	Woman
}

func (m *Man) Breath()     { m.breathing = true }
func (m *Man) Think()      { m.thinking = true }
func (m *Man) Eat()        { m.eating = true }
func (m *Man) Sex() string { return "Man" }
