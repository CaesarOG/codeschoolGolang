package model

type gopher struct {
	name string
	age int
	isAdult bool
}

func (g gopher) Jump() string {
	if g.age < 65 {
		return g.name + " can jump HIGH"
	} else {
		return g.name + " can still jump"
	}
}

type horse struct {
	name string
	weight int
}

func (h horse) Jump() string {
	if h.weight > 2500 {
		return "I'm too heavy, can't jump..."
	} else {
		return "I will jump, neigh!"
	}
}

type jumper interface {
	Jump() string
}
//removed pointer, didn't need after started using interface jumper
func GetList() []jumper { //returns slice of pointers to gopher
	phil := &gopher{name: "Phil", age: 30}
	noodles := &gopher{name: "Noodles", age: 90}
	barbaro := &horse{name: "Barbaro", weight: 2000}
	list := []jumper{phil, noodles, barbaro}
	return list
}