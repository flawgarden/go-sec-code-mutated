package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"os/exec"
	"regexp"
)

type CommandInjectVuln1Controllerjssjy struct {
	beego.Controller
}

func (c *CommandInjectVuln1Controllerjssjy) Get() {
	dir := c.GetString("dir")
	regex42 := "/.+/.+/"
	input42 := "-/" + dir + "/x/"

	pattern, _ := regexp.Compile(regex42)

	matcher := pattern.FindString(input42)
	if matcher != "" {
		dir = matcher // Set fixedVar to the matched string
	}
	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}
