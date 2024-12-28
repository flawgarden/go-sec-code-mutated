package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"os/exec"
	"regexp"
)

type CommandInjectVuln1Controllerixuul struct {
	beego.Controller
}

func (c *CommandInjectVuln1Controllerixuul) Get() {
	dir := c.GetString("dir")

	regex42 := "(/|-)"
	input42 := "y-" + dir + "/x"
	pattern, _ := regexp.Compile(regex42)
	splitResult := pattern.Split(input42, -1)

	if len(splitResult) > 2 {
		dir = splitResult[1]
	}
	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}
