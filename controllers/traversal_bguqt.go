

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
