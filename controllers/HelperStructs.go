package controllers

type SimpleStruct struct {
	Str1  string
	Str2  string
	Count int
}

func NewSimpleStruct(s string, s1 string, n int) SimpleStruct {
	return SimpleStruct{Str1: s, Str2: s1, Count: n}
}
