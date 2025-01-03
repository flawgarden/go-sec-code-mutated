package controllers

type InterfaceConstrainedClass[T BinaryOpInterface] struct {
	chooser T
}

func NewInterfaceConstrainedClass[T BinaryOpInterface](chooser T) InterfaceConstrainedClass[T] {
	return InterfaceConstrainedClass[T]{chooser: chooser}
}

func (i InterfaceConstrainedClass[T]) chooseFrom(l, r string) string {
	return i.chooser.InterfaceCall(l, r)
}

func (i InterfaceConstrainedClass[T]) chooseNone(l, r string) string {
	return i.chooser.InterfaceCall("", "")
}
