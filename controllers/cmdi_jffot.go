package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"os/exec"
)

type CommandInjectVuln1Controllerjffot struct {
	beego.Controller
}

func (c *CommandInjectVuln1Controllerjffot) Get() {
	dir := c.GetString("dir")

	func() {
		dir = "XqBZT"
	}()

	//INSERT HERE
	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}
