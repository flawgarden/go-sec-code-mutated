package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"os/exec"
)

type CommandInjectVuln1Controllerjcdmg struct {
	beego.Controller
}

func (c *CommandInjectVuln1Controllerjcdmg) Get() {
	dir := c.GetString("dir")
	gc42 := NewGenericClass("fixed_string")
	if value, ok := gc42.GetObjectValue().(string); ok {
		dir = value
	}
	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}
