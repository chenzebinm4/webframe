package main

import "github.com/chenzebinm4/webframe"

func UserLoginController(c *webframe.Context) error {
	c.Json(200, "ok, UserLoginController")
	return nil
}
