package ohmygo

type mammal interface {
	canWalk() bool
	canFeed() bool
}

type man struct{}

func (man) canWalk() bool {
	return true
}
func (man) canFeed() bool {
	return true
}

type dog struct{}

func (dog) canWalk() bool {
	return true
}
func (dog) canFeed() bool {
	return true
}

func newMan() mammal {
	return man{}
}

func newDog() mammal {
	return dog{}
}
