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
//Analyzer1 analysis results: []
//Original file name: controllers/traversal.go
//Original file CWE's: [22]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/concurrency/channels.tmt with name channel_string_positive 
//Used extensions: 
//Program:
package controllers

import (
"go-sec-code/utils"
"io/ioutil"
"path/filepath"
"strings"
beego "github.com/beego/beego/v2/server/web"
"sync"
)

type PathTraversalVuln1ControllerDnwqa struct {
	beego.Controller
}

type PathTraversalVuln2ControllerDnwqa struct {
	beego.Controller
}

type PathTraversalSafe1ControllerDnwqa struct {
	beego.Controller
}

type PathTraversalSafe2ControllerDnwqa struct {
	beego.Controller
}

func (c *PathTraversalVuln1ControllerDnwqa) Get() {
	file := c.GetString("file")

message123 := make(chan string, 1)
message123 <- file

var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    rmsg := <-message123
    message123 <- rmsg + "constant_string"
}()

wg.Wait()

file = <-message123

	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalVuln2ControllerDnwqa) Get() {
	file := c.GetString("file")
	file = filepath.Clean(file)
	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalSafe1ControllerDnwqa) Get() {
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

func (c *PathTraversalSafe2ControllerDnwqa) Get() {
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
