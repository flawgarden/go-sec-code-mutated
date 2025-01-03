package controllers

type NestedFieldsBase struct {
	values []string
	value  string
}

// Конструктор для NestedFieldsBase с одним значением
func NewNestedFieldsBase(value string) *NestedFieldsBase {
	return &NestedFieldsBase{value: value}
}

// Конструктор для NestedFieldsBase с массивом значений
func NewNestedFieldsBaseFromArray(initialValues []string) *NestedFieldsBase {
	return &NestedFieldsBase{values: initialValues}
}

// Конструктор для NestedFieldsBase с массивом значений и дополнительным значением
func NewNestedFieldsBaseFromArrayAndValue(initialValues []string, value string) *NestedFieldsBase {
	return &NestedFieldsBase{values: initialValues, value: value}
}

type NestedFields1 struct {
	value   string
	values  []string
	nested1 *NestedFieldsBase
}

// Конструктор для NestedFields1 с одним значением
func NewNestedFields1(value string) *NestedFields1 {
	return &NestedFields1{
		value:   value,
		nested1: NewNestedFieldsBase(value),
	}
}

// Конструктор для NestedFields1 с массивом значений
func NewNestedFields1FromArray(initialValues []string) *NestedFields1 {
	return &NestedFields1{
		values:  initialValues,
		nested1: NewNestedFieldsBaseFromArray(initialValues),
	}
}

// Конструктор для NestedFields1 с массивом значений и дополнительным значением
func NewNestedFields1FromArrayAndValue(initialValues []string, value string) *NestedFields1 {
	return &NestedFields1{
		values:  initialValues,
		value:   value,
		nested1: NewNestedFieldsBaseFromArrayAndValue(initialValues, value),
	}
}

type NestedFields2 struct {
	value   string
	values  []string
	nested1 *NestedFields1
}

// Конструктор для NestedFields2 с одним значением
func NewNestedFields2(value string) *NestedFields2 {
	return &NestedFields2{
		value:   value,
		nested1: NewNestedFields1(value),
	}
}

// Конструктор для NestedFields2 с массивом значений
func NewNestedFields2FromArray(initialValues []string) *NestedFields2 {
	return &NestedFields2{
		values:  initialValues,
		nested1: NewNestedFields1FromArray(initialValues),
	}
}

// Конструктор для NestedFields2 с массивом значений и дополнительным значением
func NewNestedFields2FromArrayAndValue(initialValues []string, value string) *NestedFields2 {
	return &NestedFields2{
		values:  initialValues,
		value:   value,
		nested1: NewNestedFields1FromArrayAndValue(initialValues, value),
	}
}

type NestedFields3 struct {
	value   string
	values  []string
	nested1 *NestedFields2
}

// Конструктор для NestedFields3 с одним значением
func NewNestedFields3(value string) *NestedFields3 {
	return &NestedFields3{
		value:   value,
		nested1: NewNestedFields2(value),
	}
}

// Конструктор для NestedFields3 с массивом значений
func NewNestedFields3FromArray(initialValues []string) *NestedFields3 {
	return &NestedFields3{
		values:  initialValues,
		nested1: NewNestedFields2FromArray(initialValues),
	}
}

// Конструктор для NestedFields3 с массивом значений и дополнительным значением
func NewNestedFields3FromArrayAndValue(initialValues []string, value string) *NestedFields3 {
	return &NestedFields3{
		values:  initialValues,
		value:   value,
		nested1: NewNestedFields2FromArrayAndValue(initialValues, value),
	}
}

type NestedFields4 struct {
	value   string
	values  []string
	nested1 *NestedFields3
}

// Конструктор для NestedFields4 с одним значением
func NewNestedFields4(value string) *NestedFields4 {
	return &NestedFields4{
		value:   value,
		nested1: NewNestedFields3(value),
	}
}

// Конструктор для NestedFields4 с массивом значений
func NewNestedFields4FromArray(initialValues []string) *NestedFields4 {
	return &NestedFields4{
		values:  initialValues,
		nested1: NewNestedFields3FromArray(initialValues),
	}
}

// Конструктор для NestedFields4 с массивом значений и дополнительным значением
func NewNestedFields4FromArrayAndValue(initialValues []string, value string) *NestedFields4 {
	return &NestedFields4{
		values:  initialValues,
		value:   value,
		nested1: NewNestedFields3FromArrayAndValue(initialValues, value),
	}
}
