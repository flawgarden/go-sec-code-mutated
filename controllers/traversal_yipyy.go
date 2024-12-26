

package controllers

import (
"go-sec-code/utils"
"io/ioutil"
"path/filepath"
"strings"
beego "github.com/beego/beego/v2/server/web"
"sync"
)

type PathTraversalVuln1ControllerYipyy struct {
	beego.Controller
}

type PathTraversalVuln2ControllerYipyy struct {
	beego.Controller
}

type PathTraversalSafe1ControllerYipyy struct {
	beego.Controller
}

type PathTraversalSafe2ControllerYipyy struct {
	beego.Controller
}

func (c *PathTraversalVuln1ControllerYipyy) Get() {
	file := c.GetString("file")
	

w := &Wrapper[string]{Value: file}
task := NewSettingTask(w, "")
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    task.Run()
}()
wg.Wait()
file = w.Value


	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalVuln2ControllerYipyy) Get() {
	file := c.GetString("file")
	file = filepath.Clean(file)
	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalSafe1ControllerYipyy) Get() {
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

func (c *PathTraversalSafe2ControllerYipyy) Get() {
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
