//-------------
//
//Original file name: controllers/traversal.go
//Original file CWE's: [22]
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/exceptions/panic.tmt with name dereferencing_panic_positive
//Used extensions:
//Program:
package controllers

import (
	"errors"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"go-sec-code/utils"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type PathTraversalVuln1ControllerKcbdk struct {
	beego.Controller
}

type PathTraversalVuln2ControllerKcbdk struct {
	beego.Controller
}

type PathTraversalSafe1ControllerKcbdk struct {
	beego.Controller
}

type PathTraversalSafe2ControllerKcbdk struct {
	beego.Controller
}

func (c *PathTraversalVuln1ControllerKcbdk) Get() {
	file := c.GetString("file")

	var err error
	func() {
		defer handlePanic(&err) //Отложенный вызов для обработки паники
		dereferencing(new(int))
	}()
	if err != nil {
		file = file + "suffix"
	} else {
		file = "777"
	}

	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalVuln2ControllerKcbdk) Get() {
	file := c.GetString("file")
	file = filepath.Clean(file)
	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalSafe1ControllerKcbdk) Get() {
	file := c.GetString("file")
	pathTraversalFilter := utils.PathTraversalFilter{}
	if pathTraversalFilter.DoFilter(file) {
		c.Ctx.ResponseWriter.Write([]byte("evil input"))
	} else {
		output, err := ioutil.ReadFile("static/" + file)
		if err != nil {
			panic(err)
		}
		c.Ctx.ResponseWriter.Write(output)
	}
}

func (c *PathTraversalSafe2ControllerKcbdk) Get() {
	file := c.GetString("file")
	file = filepath.Join("static/", file)
	if !strings.HasPrefix(file, "static/") {
		c.Ctx.ResponseWriter.Write([]byte("evil input"))
	} else {
		output, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		c.Ctx.ResponseWriter.Write(output)
	}
}

func raisePanic() error {
	panic("PANIC")
}

func divide_simple(num int) int {
	return 1 / num
}

func dereferencing(num *int) int {
	return *num
}

func getElWithIndex(ind int) error {
	slice := []int{1, 2, 3}
	_ = slice[ind]
	return nil
}

func handlePanic(err *error) {
	if r := recover(); r != nil {
		//Преобразуем panic в error
		*err = errors.New(fmt.Sprint(r))
	}
}
