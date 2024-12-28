package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"os/exec"
)

type CommandInjectVuln1Controllerynkhz struct {
	beego.Controller
}

func (c *CommandInjectVuln1Controllerynkhz) Get() {
	dir := c.GetString("dir")

	defer func() {
		dir = "Gxhzt"
	}()

	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}
