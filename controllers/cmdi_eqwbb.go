package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"os/exec"
)

type CommandInjectVuln1Controllereqwbb struct {
	beego.Controller
}

func (c *CommandInjectVuln1Controllereqwbb) Get() {
	dir := c.GetString("dir")
	genericWithB := NewGenericClass(B{})
	if ExtendsVariance(genericWithB) {
		dir = "fixed_string"
	}
	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}
