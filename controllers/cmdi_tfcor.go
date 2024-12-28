package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"os/exec"
)

type CommandInjectVuln1Controllertfcor struct {
	beego.Controller
}

func (c *CommandInjectVuln1Controllertfcor) Get() {
	dir := c.GetString("dir")

	var multinterface12 = MultipleInterfaceClass_2Pos{}
	dir = multinterface12.Interface2Call(dir)

	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}
