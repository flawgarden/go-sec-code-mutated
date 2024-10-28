//
//-------------
//
//Original file name: controllers/traversal.go
//Original file CWE's: [22]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/collections/queue.tmt with name two_queues_positive 
//Used extensions: 
//Program:
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
