//Analyzer5 original results: [22]
//Analyzer1 original results: [22]
//Analyzer3 original results: []
//Analyzer2 original results: []
//Analyzer4 original results: []
//-------------
//Analyzer3 analysis results: []
//Analyzer2 analysis results: []
//Analyzer4 analysis results: []
//Analyzer5 analysis results: [22]
//Analyzer1 analysis results: [563]
//Original file name: controllers/traversal.go
//Original file CWE's: [22]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/collections/set.tmt with name set_add_simple_negative 
//Used extensions: MACRO_Create_Set -> ~[MACRO_SetName]~ := make(map[~[TYPE@1]~]struct{}) | MACRO_Add_Fixed_CONST_ToSet -> ~[MACRO_SetName]~[~[CONST_~[TYPE@1]~@1]~] = struct{}{} | MACRO_SetName -> set787231 | MACRO_SetName -> set787231 | MACRO_SetName -> set787231
//Program:
package controllers

import (
	"go-sec-code/utils"
	"io/ioutil"
	"path/filepath"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type PathTraversalVuln1ControllerWvfws struct {
	beego.Controller
}

type PathTraversalVuln2ControllerWvfws struct {
	beego.Controller
}

type PathTraversalSafe1ControllerWvfws struct {
	beego.Controller
}

type PathTraversalSafe2ControllerWvfws struct {
	beego.Controller
}

func (c *PathTraversalVuln1ControllerWvfws) Get() {
	file := c.GetString("file")

set787231 := make(map[string]struct{})
set787231["RSgbd"] = struct{}{}
file = func() string {
    for k := range set787231 {
        return k
    }
    return "ieBat"
}()

	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalVuln2ControllerWvfws) Get() {
	file := c.GetString("file")
	file = filepath.Clean(file)
	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalSafe1ControllerWvfws) Get() {
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

func (c *PathTraversalSafe2ControllerWvfws) Get() {
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
