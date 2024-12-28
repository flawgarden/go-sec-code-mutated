

package controllers

import (
	"go-sec-code/utils"
	"io/ioutil"
	"path/filepath"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type PathTraversalVuln1ControllerJhwbm struct {
	beego.Controller
}

type PathTraversalVuln2ControllerJhwbm struct {
	beego.Controller
}

type PathTraversalSafe1ControllerJhwbm struct {
	beego.Controller
}

type PathTraversalSafe2ControllerJhwbm struct {
	beego.Controller
}

func (c *PathTraversalVuln1ControllerJhwbm) Get() {
	file := c.GetString("file")
	

tmpUnique42 := file
for {
    file = tmpUnique42
    if !(strings.Index("azBZF", "mGKHu") >= 0) && (strings.EqualFold(file, file)) {
        break
    }
}


	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalVuln2ControllerJhwbm) Get() {
	file := c.GetString("file")
	file = filepath.Clean(file)
	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalSafe1ControllerJhwbm) Get() {
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

func (c *PathTraversalSafe2ControllerJhwbm) Get() {
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
