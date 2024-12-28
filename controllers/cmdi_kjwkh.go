package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"os/exec"
	"regexp"
)

type CommandInjectVuln1Controllerkjwkh struct {
	beego.Controller
}

func (c *CommandInjectVuln1Controllerkjwkh) Get() {
	dir := c.GetString("dir")

	regex42 := "k+"
	str42 := "kkk"
	matcher, _ := regexp.Compile(regex42)
	matches := matcher.FindAllStringIndex(str42, 0)
	count42 := len(matches)

	if count42 != 0 {
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
