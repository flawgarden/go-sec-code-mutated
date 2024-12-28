package controllers

type A interface{}

type B struct{}

type C struct{}

type GenericClass[T any] struct {
	value T
}

func NewGenericClass[T any](value T) *GenericClass[T] {
	return &GenericClass[T]{value: value}
}

func (g *GenericClass[T]) GetObjectValue() interface{} {
	return g.value
}

func (g *GenericClass[T]) GetValue() T {
	return g.value
}

func ExtendsVariance[T A](genericClassObj *GenericClass[T]) bool {
	_, ok := genericClassObj.GetObjectValue().(B)
	return ok
}
