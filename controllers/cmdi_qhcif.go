package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"os/exec"
	"regexp"
)

type CommandInjectVuln1Controllerqhcif struct {
	beego.Controller
}

func (c *CommandInjectVuln1Controllerqhcif) Get() {
	dir := c.GetString("dir")

	regex42 := "[^ab]+"
	str42 := "aaa"
	matcher, _ := regexp.Compile(regex42)
	if matcher.MatchString(str42) {
		dir = ""
	}
	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}