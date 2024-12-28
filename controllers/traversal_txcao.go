
package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"go-sec-code/utils"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type PathTraversalVuln1ControllerTxcao struct {
	beego.Controller
}

type PathTraversalVuln2ControllerTxcao struct {
	beego.Controller
}

type PathTraversalSafe1ControllerTxcao struct {
	beego.Controller
}

type PathTraversalSafe2ControllerTxcao struct {
	beego.Controller
}

func (c *PathTraversalVuln1ControllerTxcao) Get() {
	file := c.GetString("file")

generator := makeMessageGeneratorBroken(file)()
file = generator("fixed_string")

	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalVuln2ControllerTxcao) Get() {
	file := c.GetString("file")
	file = filepath.Clean(file)
	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalSafe1ControllerTxcao) Get() {
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

func (c *PathTraversalSafe2ControllerTxcao) Get() {
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
