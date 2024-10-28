//
//-------------
//
//Original file name: controllers/traversal.go
//Original file CWE's: [22]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/virtuality/interface.tmt with name base_binary_op_interface_positive 
//Used extensions: 
//Program:
package controllers

import (
	"go-sec-code/utils"
	"io/ioutil"
	"path/filepath"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type PathTraversalVuln1ControllerGaxqp struct {
	beego.Controller
}

type PathTraversalVuln2ControllerGaxqp struct {
	beego.Controller
}

type PathTraversalSafe1ControllerGaxqp struct {
	beego.Controller
}

type PathTraversalSafe2ControllerGaxqp struct {
	beego.Controller
}

func (c *PathTraversalVuln1ControllerGaxqp) Get() {
	file := c.GetString("file")
	

var a12341 BinaryOpInterface
if -78458653 == 0 {
    a12341 = &ImplBinaryOpInterfaceClass1{}
} else {
    a12341 = &ImplBinaryOpInterfaceClass2{}
}
file = a12341.InterfaceCall(file, file)


	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalVuln2ControllerGaxqp) Get() {
	file := c.GetString("file")
	file = filepath.Clean(file)
	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalSafe1ControllerGaxqp) Get() {
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

func (c *PathTraversalSafe2ControllerGaxqp) Get() {
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
