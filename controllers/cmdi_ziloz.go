package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"os/exec"
	"regexp"
)

type CommandInjectVuln1Controllerziloz struct {
	beego.Controller
}

func (c *CommandInjectVuln1Controllerziloz) Get() {
	dir := c.GetString("dir")

	regex42 := "^101"
	str42 := "10042"
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