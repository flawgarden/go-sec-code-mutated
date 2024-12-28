package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"os/exec"
	"strings"
)

type CommandInjectVuln1Controllerxanak struct {
	beego.Controller
}

func (c *CommandInjectVuln1Controllerxanak) Get() {
	dir := c.GetString("dir")
	dir = simplePatternMatchingString1(dir)
	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func simplePatternMatchingString1(obj interface{}) string {
	if str, ok := obj.(string); ok {
		return strings.ToUpper(str)
	}
	return ""
}

func simplePatternMatchingString2(obj interface{}) string {
	if str, ok := obj.(string); ok && len(str) > 5 {
		return strings.ToUpper(str)
	}
	return ""
}
