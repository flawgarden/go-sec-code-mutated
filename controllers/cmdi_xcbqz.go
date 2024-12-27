package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"os/exec"
	"regexp"
)

type CommandInjectVuln1Controllerxcbqz struct {
	beego.Controller
}

func (c *CommandInjectVuln1Controllerxcbqz) Get() {
	dir := c.GetString("dir")

	regex42 := "aboba"
	str42 := "aboba"
	matcher, _ := regexp.Compile(regex42)
	dir = matcher.ReplaceAllString(dir+str42, "")

	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}
