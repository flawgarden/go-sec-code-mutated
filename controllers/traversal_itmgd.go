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
//Mutation info: Insert template from templates-db/languages/go/sensitivity/concurrency/concurrency.tmt with name two_set_threads_in_sequence_positive 
//Used extensions: 
//Original file region: 28, 35, null, null
//Mutated file region: 47, 71, null, null


package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"go-sec-code/utils"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"
)

type PathTraversalVuln1ControllerItmgd struct {
	beego.Controller
}

type PathTraversalVuln2ControllerItmgd struct {
	beego.Controller
}

type PathTraversalSafe1ControllerItmgd struct {
	beego.Controller
}

type PathTraversalSafe2ControllerItmgd struct {
	beego.Controller
}

func (c *PathTraversalVuln1ControllerItmgd) Get() {
	file := c.GetString("file")

	w := &Wrapper[string]{Value: file}
	task1 := NewSettingTask(w, "")
	task2 := NewSettingTask(w, file)
	var wg sync.WaitGroup
	wg.Add(2) //Добавляем 2 задачи в WaitGroup
	go func() {
		defer wg.Done()
		task1.Run()
	}()
	go func() {
		defer wg.Done()
		task2.Run()
	}()
	wg.Wait()
	file = w.Value

	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalVuln2ControllerItmgd) Get() {
	file := c.GetString("file")
	file = filepath.Clean(file)
	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalSafe1ControllerItmgd) Get() {
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

func (c *PathTraversalSafe2ControllerItmgd) Get() {
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