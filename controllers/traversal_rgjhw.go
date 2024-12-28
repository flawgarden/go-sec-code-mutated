
package controllers

import (
	"go-sec-code/utils"
	"io/ioutil"
	"path/filepath"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type PathTraversalVuln1ControllerRgjhw struct {
	beego.Controller
}

type PathTraversalVuln2ControllerRgjhw struct {
	beego.Controller
}

type PathTraversalSafe1ControllerRgjhw struct {
	beego.Controller
}

type PathTraversalSafe2ControllerRgjhw struct {
	beego.Controller
}

func (c *PathTraversalVuln1ControllerRgjhw) Get() {
	file := c.GetString("file")

sh := NewStringHolder()
file = sh.value

	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalVuln2ControllerRgjhw) Get() {
	file := c.GetString("file")
	file = filepath.Clean(file)
	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalSafe1ControllerRgjhw) Get() {
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

func (c *PathTraversalSafe2ControllerRgjhw) Get() {
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
