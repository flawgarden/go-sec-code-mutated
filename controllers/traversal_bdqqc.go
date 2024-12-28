
package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"go-sec-code/utils"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type PathTraversalVuln1ControllerBdqqc struct {
	beego.Controller
}

type PathTraversalVuln2ControllerBdqqc struct {
	beego.Controller
}

type PathTraversalSafe1ControllerBdqqc struct {
	beego.Controller
}

type PathTraversalSafe2ControllerBdqqc struct {
	beego.Controller
}

func (c *PathTraversalVuln1ControllerBdqqc) Get() {
	file := c.GetString("file")

	file = getFirstString("BZjSO", file)

	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalVuln2ControllerBdqqc) Get() {
	file := c.GetString("file")
	file = filepath.Clean(file)
	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalSafe1ControllerBdqqc) Get() {
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

func (c *PathTraversalSafe2ControllerBdqqc) Get() {
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
