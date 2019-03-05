package pkg1

type Foo struct {
}

func (*Foo) PublicFn() {}

type Fooable interface {
	PublicFn()
}

func Do(Fooable) {}