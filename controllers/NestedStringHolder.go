package controllers

type InnerStringHolder struct {
	innerValue string
}

type NestedStringHolder struct {
	innerObject *InnerStringHolder
}

// Конструктор для NestedStringHolder с заданным значением
func NewNestedStringHolder(value string) *NestedStringHolder {
	return &NestedStringHolder{
		innerObject: &InnerStringHolder{innerValue: value},
	}
}

// Конструктор для NestedStringHolder с пустым значением
func NewNestedStringHolderEmpty() *NestedStringHolder {
	return &NestedStringHolder{
		innerObject: &InnerStringHolder{innerValue: ""},
	}
}

// Метод для получения значения
func (nsh *NestedStringHolder) GetValue() string {
	return nsh.innerObject.innerValue
}
