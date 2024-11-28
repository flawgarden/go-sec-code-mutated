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
//Mutation info: Insert template from templates-db/languages/go/sensitivity/collections/queue.tmt with name two_queues_positive 
//Used extensions: 
//Original file region: 28, 35, null, null
//Mutated file region: 47, 72, null, null


package controllers

import (
"go-sec-code/utils"
"io/ioutil"
"path/filepath"
"strings"
beego "github.com/beego/beego/v2/server/web"
"container/list"
)

type PathTraversalVuln1ControllerYdtvh struct {
	beego.Controller
}

type PathTraversalVuln2ControllerYdtvh struct {
	beego.Controller
}

type PathTraversalSafe1ControllerYdtvh struct {
	beego.Controller
}

type PathTraversalSafe2ControllerYdtvh struct {
	beego.Controller
}

func (c *PathTraversalVuln1ControllerYdtvh) Get() {
	file := c.GetString("file")
	

sourceQueue := list.New()
targetQueue := list.New()
sourceQueue.PushBack("uJbzG")
sourceQueue.PushBack(string("rOusV"))
targetQueue.PushBack(file)
targetQueue.PushBack("hVPBb")
for sourceQueue.Len() > 0 {
    element := sourceQueue.Front()
    if element != nil {
        targetQueue.PushBack(element.Value)
        sourceQueue.Remove(element)
    }
}
file = targetQueue.Front().Value.(string)


	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalVuln2ControllerYdtvh) Get() {
	file := c.GetString("file")
	file = filepath.Clean(file)
	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalSafe1ControllerYdtvh) Get() {
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

func (c *PathTraversalSafe2ControllerYdtvh) Get() {
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