package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"os/exec"
)

type CommandInjectVuln1Controllertbdpn struct {
	beego.Controller
}

func (c *CommandInjectVuln1Controllertbdpn) Get() {
	dir := c.GetString("dir")

	impl42 := &ImplBinaryOpInterfaceClass1{}
	constrained09143 := NewInterfaceConstrainedClass(impl42)
	dir = constrained09143.chooseNone(dir, "fixed_string")

	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}
