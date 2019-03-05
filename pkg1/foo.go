package pkg1

type Foo struct {
}

func (*Foo) PublicFn() {}
func (*Foo) privateFn() {}

type Fooable interface {
	PublicFn()
	privateFn()
}

func Do(Fooable) {}