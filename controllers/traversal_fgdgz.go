
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
queue787231.PushBack(file)
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
