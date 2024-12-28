package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"os/exec"
)

type CommandInjectVuln1Controlleraajvp struct {
	beego.Controller
}

func (c *CommandInjectVuln1Controlleraajvp) Get() {
	dir := c.GetString("dir")
	var a12341 BinaryOpInterface = &ImplBinaryOpInterfaceClass2{}
	dir = a12341.InterfaceCall("", dir)
	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}
