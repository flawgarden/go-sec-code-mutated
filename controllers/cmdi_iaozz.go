package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"os/exec"
	"regexp"
)

type CommandInjectVuln1Controlleriaozz struct {
	beego.Controller
}

func (c *CommandInjectVuln1Controlleriaozz) Get() {
	dir := c.GetString("dir")
	regex42 := "abc"
	str42 := "AbC"
	matcher, _ := regexp.Compile("(?i)" + regex42)
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
