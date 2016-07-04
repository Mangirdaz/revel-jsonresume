package controllers

import "github.com/revel/revel"
import "mangirdaz/jsonresume"

type About struct {
	*revel.Controller
}

type AboutMe struct {
	Active bool
	Resume jsonresume.ResumeJson
}

func (c About) Index() revel.Result {

	resume := jsonresume.New()
	about := &AboutMe{Active: true, Resume: resume}


	return c.Render(about)
}
