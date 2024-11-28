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
//Mutation info: Insert template from templates-db/languages/go/sensitivity/types/alias.tmt with name type_alias_simple_negative 
//Used extensions: 
//Original file region: 28, 35, null, null
//Mutated file region: 47, 59, null, null

package controllers

import (
	"go-sec-code/utils"
	"io/ioutil"
	"path/filepath"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type PathTraversalVuln1ControllerQyvdh struct {
	beego.Controller
}

type PathTraversalVuln2ControllerQyvdh struct {
	beego.Controller
}

type PathTraversalSafe1ControllerQyvdh struct {
	beego.Controller
}

type PathTraversalSafe2ControllerQyvdh struct {
	beego.Controller
}

func (c *PathTraversalVuln1ControllerQyvdh) Get() {
	file := c.GetString("file")

type MyStringQyvdh = string
var x MyStringQyvdh = "FKqqs"
file = x

	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalVuln2ControllerQyvdh) Get() {
	file := c.GetString("file")
	file = filepath.Clean(file)
	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalSafe1ControllerQyvdh) Get() {
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

func (c *PathTraversalSafe2ControllerQyvdh) Get() {
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