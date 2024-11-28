//Analyzer1 original results: [22]
//Analyzer2 original results: [22]
//Analyzer3 original results: []
//Analyzer4 original results: []
//Analyzer5 original results: []
//-------------
//Analyzer3 analysis results: []
//Analyzer4 analysis results: []
//Analyzer1 analysis results: [22, 703]
//Analyzer2 analysis results: [563]
//Analyzer5 analysis results: []
//Original file name: controllers/traversal.go
//Original file CWE's: [22]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/varargs/varargs.tmt with name varargs_get_first_string_from_array_negative 
//Used extensions: 
//Original file region: 28, 35, null, null
//Mutated file region: 46, 56, null, null

package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"go-sec-code/utils"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type PathTraversalVuln1ControllerLwitm struct {
	beego.Controller
}

type PathTraversalVuln2ControllerLwitm struct {
	beego.Controller
}

type PathTraversalSafe1ControllerLwitm struct {
	beego.Controller
}

type PathTraversalSafe2ControllerLwitm struct {
	beego.Controller
}

func (c *PathTraversalVuln1ControllerLwitm) Get() {
	file := c.GetString("file")

file = getFirstStringFromArray("aggkq", "mZxRP")

	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalVuln2ControllerLwitm) Get() {
	file := c.GetString("file")
	file = filepath.Clean(file)
	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalSafe1ControllerLwitm) Get() {
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

func (c *PathTraversalSafe2ControllerLwitm) Get() {
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