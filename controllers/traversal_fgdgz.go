//Analyzer5 original results: [22]
//Analyzer1 original results: [22]
//Analyzer3 original results: []
//Analyzer2 original results: []
//Analyzer4 original results: []
//-------------
//Analyzer3 analysis results: []
//Analyzer2 analysis results: []
//Analyzer4 analysis results: []
//Analyzer5 analysis results: [22, 703]
//Analyzer1 analysis results: [563]
//Original file name: controllers/traversal.go
//Original file CWE's: [22]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/collections/queue.tmt with name queue_poll_all_positive 
//Used extensions: MACRO_Create_Queue -> ~[MACRO_QueueName]~ := list.New() | MACRO_Add_EXPR_ToQueue -> ~[MACRO_QueueName]~.PushBack(~[EXPR_~[TYPE@1]~]~) | MACRO_Add_EXPR_ToQueue -> ~[MACRO_QueueName]~.PushBack(~[EXPR_~[TYPE@1]~]~) | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | EXPR_string -> strings.Join([]string{~[EXPR_string]~, ~[EXPR_string]~}, "") | MACRO_QueueName -> queue787231
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

type PathTraversalVuln1ControllerFgdgz struct {
	beego.Controller
}

type PathTraversalVuln2ControllerFgdgz struct {
	beego.Controller
}

type PathTraversalSafe1ControllerFgdgz struct {
	beego.Controller
}

type PathTraversalSafe2ControllerFgdgz struct {
	beego.Controller
}

func (c *PathTraversalVuln1ControllerFgdgz) Get() {
	file := c.GetString("file")

queue787231 := list.New()
queue787231.PushBack(strings.Join([]string{"WjjDB", "wRAWN"}, ""))
queue787231.PushBack("EKczk")
value7847 := "hgDoA"
for queue787231.Len() > 0 {
    element := queue787231.Front()
    if element != nil {
        value7847 = element.Value.(string)
        queue787231.Remove(element)
    }
}
file = value7847

	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalVuln2ControllerFgdgz) Get() {
	file := c.GetString("file")
	file = filepath.Clean(file)
	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalSafe1ControllerFgdgz) Get() {
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

func (c *PathTraversalSafe2ControllerFgdgz) Get() {
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
