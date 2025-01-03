

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
