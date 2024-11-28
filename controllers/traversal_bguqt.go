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
//Mutation info: Insert template from templates-db/languages/go/sensitivity/collections/queue.tmt with name queue_poll_positive 
//Used extensions: 
//Original file region: 28, 35, null, null
//Mutated file region: 47, 67, null, null


package controllers

import (
"go-sec-code/utils"
"io/ioutil"
"path/filepath"
"strings"
beego "github.com/beego/beego/v2/server/web"
"container/list"
)

type PathTraversalVuln1ControllerBguqt struct {
	beego.Controller
}

type PathTraversalVuln2ControllerBguqt struct {
	beego.Controller
}

type PathTraversalSafe1ControllerBguqt struct {
	beego.Controller
}

type PathTraversalSafe2ControllerBguqt struct {
	beego.Controller
}

func (c *PathTraversalVuln1ControllerBguqt) Get() {
	file := c.GetString("file")
	

queue787231 := list.New()
queue787231.PushBack(file)
queue787231.PushBack("CtwGM")
queue787231.PushBack("WvWSY")
element := queue787231.Front()
if element != nil {
    file = element.Value.(string)
    queue787231.Remove(element)
}


	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalVuln2ControllerBguqt) Get() {
	file := c.GetString("file")
	file = filepath.Clean(file)
	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalSafe1ControllerBguqt) Get() {
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

func (c *PathTraversalSafe2ControllerBguqt) Get() {
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