package controllers

type StringHolder struct {
	value string
}

// Создает новый экземпляр StringHolder с пустым значением
func NewStringHolder() *StringHolder {
	return &StringHolder{value: ""}
}

// Создает новый экземпляр StringHolder с заданным значением
func NewStringHolderWithValue(value string) *StringHolder {
	return &StringHolder{value: value}
}
