package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"os/exec"
)

type CommandInjectVuln1Controllersovma struct {
	beego.Controller
}

func (c *CommandInjectVuln1Controllersovma) Get() {
	dir := c.GetString("dir")

	simpleStruct42 := NewSimpleStruct("fixed_string", dir, 5)
	dir = simpleStruct42.Str1

	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}
