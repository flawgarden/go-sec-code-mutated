package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"os/exec"
)

type CommandInjectVuln1Controllerirdlq struct {
	beego.Controller
}

func (c *CommandInjectVuln1Controllerirdlq) Get() {
	dir := c.GetString("dir")
	var obj interface{} = dir
	switch v := obj.(type) {
	case int:
		dir = "fixed_string"
	case string:
		dir = v
	default:
		dir = "fixed_string_2"
	}
	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}
