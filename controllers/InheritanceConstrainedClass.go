package controllers

type InheritanceConstrainedClass[T BaseBinaryOpClass] struct {
	chooser T
}

func NewInheritanceConstrainedClass[T BaseBinaryOpClass](chooser T) InheritanceConstrainedClass[T] {
	return InheritanceConstrainedClass[T]{chooser: chooser}
}

func (i InheritanceConstrainedClass[T]) chooseFrom(l, r string) string {
	return i.chooser.VirtualCall(l, r)
}

func (i InheritanceConstrainedClass[T]) chooseNone(l, r string) string {
	return i.chooser.VirtualCall("", "")
}
