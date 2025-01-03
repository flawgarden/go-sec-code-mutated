package controllers

type InstanceInitializer struct {
	value string
	list  []string
}

// Конструктор
func NewInstanceInitializer(value string) *InstanceInitializer {
	instance := &InstanceInitializer{
		value: value,
		list:  []string{""}, // Инициализация списка с пустой строкой
	}
	instance.list = append(instance.list, value) // Добавляем значение в список
	return instance
}
