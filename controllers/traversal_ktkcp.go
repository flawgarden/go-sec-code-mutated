

package controllers

import (
	"go-sec-code/utils"
	"io/ioutil"
	"path/filepath"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type PathTraversalVuln1ControllerKtkcp struct {
	beego.Controller
}

type PathTraversalVuln2ControllerKtkcp struct {
	beego.Controller
}

type PathTraversalSafe1ControllerKtkcp struct {
	beego.Controller
}

type PathTraversalSafe2ControllerKtkcp struct {
	beego.Controller
}

func (c *PathTraversalVuln1ControllerKtkcp) Get() {
	file := c.GetString("file")
	

set111 := make(map[string]struct{})
set222 := make(map[string]struct{})
set111["YClYO"] = struct{}{}
set222[file] = struct{}{}
for k := range set222 {
    set111[k] = struct{}{}
}
if _, ok := set111[file]; ok {
    file = "pbAqu"
}


	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalVuln2ControllerKtkcp) Get() {
	file := c.GetString("file")
	file = filepath.Clean(file)
	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalSafe1ControllerKtkcp) Get() {
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

func (c *PathTraversalSafe2ControllerKtkcp) Get() {
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
