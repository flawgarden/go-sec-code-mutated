package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"os/exec"
)

type CommandInjectVuln1Controlleraabvp struct {
	beego.Controller
}

func (c *CommandInjectVuln1Controlleraabvp) Get() {
	dir := c.GetString("dir")
	slice42 := [2]int{1, 3}
	for i, _ := range slice42 {
		if i == 0 {
			dir = "fixed_string"
		}
	}
	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}
