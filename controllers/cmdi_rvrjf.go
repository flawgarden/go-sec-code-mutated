package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"os/exec"
)

type CommandInjectVuln1Controllerrvrjf struct {
	beego.Controller
}

func (c *CommandInjectVuln1Controllerrvrjf) Get() {
	dir := c.GetString("dir")

	gc42 := NewGenericClass(dir)
	dir = gc42.GetValue()

	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}
