package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"os/exec"
	"strings"
)

type CommandInjectVuln1Controllerllepy struct {
	beego.Controller
}

func (c *CommandInjectVuln1Controllerllepy) Get() {
	dir := c.GetString("dir")
	tmpStr := dir
	interpolatedStr := fmt.Sprintf("Some_prefix -> %s", strings.ToUpper(tmpStr))
	dir = interpolatedStr
	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}
