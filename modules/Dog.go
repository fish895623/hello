package Dog

type Animal interface {
	Sounds() string
}

type Dog struct {
	name string
}

func NewDog(name string) Dog {
	var o Dog
	o.name = name
	return o
}

func (o Dog) Sounds() string {
	return "asdf"
}
