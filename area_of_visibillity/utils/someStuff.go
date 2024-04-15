package utils

type Area struct {
	Vars   map[string]int
	Parent *Area
}

func NewArea(parent *Area) *Area {
	return &Area{
		Vars:   make(map[string]int),
		Parent: parent,
	}
}

func (a *Area) SaveVar(name string, value int) {
	a.Vars[name] = value
}

func (a *Area) ShowVar(name string) (int, bool) {
	value, exists := a.Vars[name]
	if !exists {
		return 0, false
	}
	return value, exists
}
