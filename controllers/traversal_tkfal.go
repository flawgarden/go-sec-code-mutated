//Analyzer1 original results: [22]
//Analyzer2 original results: [22]
//Analyzer3 original results: []
//Analyzer4 original results: []
//Analyzer5 original results: []
//-------------
//Analyzer3 analysis results: []
//Analyzer4 analysis results: []
//Analyzer1 analysis results: [22, 703]
//Analyzer2 analysis results: [99, 23, 22, 73, 36]
//Analyzer5 analysis results: []
//Original file name: controllers/traversal.go
//Original file CWE's: [22]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/namedreturns.tmt with name named_return_with_defer_positive 
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

type PathTraversalVuln1ControllerTkfal struct {
	beego.Controller
}

type PathTraversalVuln2ControllerTkfal struct {
	beego.Controller
}

type PathTraversalSafe1ControllerTkfal struct {
	beego.Controller
}

type PathTraversalSafe2ControllerTkfal struct {
	beego.Controller
}

func (c *PathTraversalVuln1ControllerTkfal) Get() {
	file := c.GetString("file")
	

file = brokenConcat("prefix", file)


	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalVuln2ControllerTkfal) Get() {
	file := c.GetString("file")
	file = filepath.Clean(file)
	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalSafe1ControllerTkfal) Get() {
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

func (c *PathTraversalSafe2ControllerTkfal) Get() {
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

func concat(a string, b string) (res string) {
    res = a + b
    return
}

func swap(a string, b string) (first string, second string) {
	first, second = b, a
	return
}

func brokenConcat(a string, b string) (result string) {
	defer func() {
		result = b
	}()
	result = a + b
	return
}

func getZeroValues() (x string, y string) {
    return
}