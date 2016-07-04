package controllers

import "github.com/revel/revel"

type Chat struct {
	*revel.Controller
}

type ChatInfo struct {
	Active bool
	Intro  string
}

func (c Chat) Index() revel.Result {
	chat := &ChatInfo{Active: true, Intro: "TBC"}
	return c.Render(chat)
}
