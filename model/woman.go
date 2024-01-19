package model

type Woman struct {
	age       int
	name      string
	high      float32
	weight    float32
	breathing bool
	thinking  bool
	eating    bool
	isAlive   bool
}

// Each of this * means a pointer reference to the name of the structs,
// this is how go can attached this functions to the struct (POO)
func (m *Woman) Breath()       { m.breathing = true }
func (m *Woman) Think()        { m.thinking = true }
func (m *Woman) Eat()          { m.eating = true }
func (m *Woman) Sex() string   { return "Woman" }
func (m *Woman) IsAlive() bool { return m.isAlive }
