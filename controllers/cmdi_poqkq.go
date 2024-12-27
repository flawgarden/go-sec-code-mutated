package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"os/exec"
)

type CommandInjectVuln1Controllerpoqkq struct {
	beego.Controller
}

func (c *CommandInjectVuln1Controllerpoqkq) Get() {
	dir := c.GetString("dir")

	tmpStr := dir
	tmpStr = "oKVVp"
	interpolatedStr := fmt.Sprintf("Some_prefix -> %s", tmpStr)
	dir = interpolatedStr

	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}
