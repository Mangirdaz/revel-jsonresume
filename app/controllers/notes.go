package controllers

import "github.com/revel/revel"

type Notes struct {
	*revel.Controller
}

type NotesInfo struct {
	Active bool
	Intro  string
}

func (c Notes) Index() revel.Result {
	notes := &NotesInfo{Active: true, Intro: "TBC"}
	return c.Render(notes)
}
