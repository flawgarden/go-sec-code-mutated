//Analyzer5 original results: [22]
//Analyzer1 original results: [22]
//Analyzer2 original results: []
//Analyzer4 original results: []
//-------------
//Analyzer2 analysis results: []
//Analyzer4 analysis results: []
//Analyzer5 analysis results: [22]
//Analyzer1 analysis results: [563]
//Original file name: controllers/traversal.go
//Original file CWE's: [22]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/types/alias.tmt with name type_alias_simple_negative 
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

type PathTraversalVuln1ControllerSyadh struct {
	beego.Controller
}

type PathTraversalVuln2ControllerSyadh struct {
	beego.Controller
}

type PathTraversalSafe1ControllerSyadh struct {
	beego.Controller
}

type PathTraversalSafe2ControllerSyadh struct {
	beego.Controller
}

func (c *PathTraversalVuln1ControllerSyadh) Get() {
	file := c.GetString("file")

type MyStringSyadh = string
var x MyStringSyadh = "UvIbi"
file = x

	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalVuln2ControllerSyadh) Get() {
	file := c.GetString("file")
	file = filepath.Clean(file)
	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalSafe1ControllerSyadh) Get() {
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

func (c *PathTraversalSafe2ControllerSyadh) Get() {
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
