package controllers

type MultipleInterfaceClass_2Neg struct{}

func (m MultipleInterfaceClass_2Neg) InterfaceCall(s string) string {
	return s
}

func (m MultipleInterfaceClass_2Neg) Interface2Call(s string) string {
	return ""
}

type MultipleInterfaceClass_2Pos struct{}

func (m MultipleInterfaceClass_2Pos) InterfaceCall(s string) string {
	return ""
}

func (m MultipleInterfaceClass_2Pos) Interface2Call(s string) string {
	return s
}
