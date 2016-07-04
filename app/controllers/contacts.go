package controllers

import "github.com/revel/revel"

type Contacts struct {
	*revel.Controller
}

type ContactsInfo struct {
	Active bool
	Intro  string
}

func (c Contacts) Index() revel.Result {
	contacts := &ContactsInfo{Active: true, Intro: "TBC"}
	return c.Render(contacts)
}
