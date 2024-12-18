package controllers
type Duck struct{}

func (d Duck) Quack(arg string) string {
    return arg
}

type NotADuck struct{}

func (n NotADuck) Quack(arg string) string {
    return "fixed_string"
}

type DefinitelyNotADuck struct{}

func (dn DefinitelyNotADuck) FakeQuack(arg string) string {
    return "fixed_string"
}

type DynamicDuck struct{}


type DuckWithAttribute struct {
    myArg   string
    constant int
}

func NewDuckWithAttribute(arg string) DuckWithAttribute {
    return DuckWithAttribute{
        myArg:   arg,
        constant: 42,
    }
}

func (d DuckWithAttribute) Quack(arg string) string {
    return d.myArg
}

type FakeDuckWithAttribute struct {
    myArg string
}

func NewFakeDuckWithAttribute(arg string) FakeDuckWithAttribute {
    return FakeDuckWithAttribute{
        myArg: arg,
    }
}

func (f FakeDuckWithAttribute) Quack(arg string) string {
    return f.myArg
}
