
package controllers

import (
	"go-sec-code/utils"
	"io/ioutil"
	"path/filepath"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type PathTraversalVuln1ControllerUbnvv struct {
	beego.Controller
}

type PathTraversalVuln2ControllerUbnvv struct {
	beego.Controller
}

type PathTraversalSafe1ControllerUbnvv struct {
	beego.Controller
}

type PathTraversalSafe2ControllerUbnvv struct {
	beego.Controller
}

func (c *PathTraversalVuln1ControllerUbnvv) Get() {
	file := c.GetString("file")

var a12341 BinaryOpInterface
if -1816957902 & -1115493020 == -1474193144 {
    a12341 = &ImplBinaryOpInterfaceClass1{}
} else {
    a12341 = &ImplBinaryOpInterfaceClass2{}
}
file = a12341.InterfaceCall("", "")

	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalVuln2ControllerUbnvv) Get() {
	file := c.GetString("file")
	file = filepath.Clean(file)
	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalSafe1ControllerUbnvv) Get() {
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

func (c *PathTraversalSafe2ControllerUbnvv) Get() {
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
