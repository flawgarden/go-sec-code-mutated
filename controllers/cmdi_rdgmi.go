package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"os/exec"
	"regexp"
)

type CommandInjectVuln1Controllerrdgmi struct {
	beego.Controller
}

func (c *CommandInjectVuln1Controllerrdgmi) Get() {
	dir := c.GetString("dir")

	regex42 := ".+(aboba)$"
	tmpVar42 := dir + "aboba"
	matcher, _ := regexp.Compile(regex42)
	matches := matcher.FindStringSubmatch(tmpVar42)

	if matches != nil {
		dir = matches[0]
	}

	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}
