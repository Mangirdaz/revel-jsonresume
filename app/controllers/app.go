package controllers

import "github.com/revel/revel"
import "mangirdaz/jsonresume"

//App revel  struct  controler
type App struct {
	*revel.Controller
}

//Todo: abstract these somehow...
func (c App) Index() revel.Result {

	resume := jsonresume.New()
	c.RenderArgs["headerName"] = resume.Basics.Name
	c.RenderArgs["headerTitle"] = resume.Basics.Label
	c.RenderArgs["aboutNav"] = c.Message("about")
	c.RenderArgs["notesNav"] = c.Message("notes")
	c.RenderArgs["chatNav"] = c.Message("chat")
	c.RenderArgs["contactsNav"] = c.Message("contacts")
	return c.Render()
}

func (c *About) Before() revel.Result {
	resume := jsonresume.New()
	c.RenderArgs["headerName"] = resume.Basics.Name
	c.RenderArgs["headerTitle"] = resume.Basics.Label
	c.RenderArgs["aboutNav"] = c.Message("about")
	c.RenderArgs["notesNav"] = c.Message("notes")
	c.RenderArgs["chatNav"] = c.Message("chat")
	c.RenderArgs["contactsNav"] = c.Message("contacts")
	c.RenderArgs["email"] = c.Message("email")
	c.RenderArgs["contancts"] = c.Message("contancts")
	c.RenderArgs["website"] = c.Message("website")
	c.RenderArgs["phone"] = c.Message("phone")
	c.RenderArgs["aboutTitle"] = c.Message("aboutTitle")
	return nil
}

func (c *App) Before() revel.Result {
	resume := jsonresume.New()
	c.RenderArgs["headerName"] = resume.Basics.Name
	c.RenderArgs["headerTitle"] = resume.Basics.Label
	c.RenderArgs["aboutNav"] = c.Message("about")
	c.RenderArgs["notesNav"] = c.Message("notes")
	c.RenderArgs["chatNav"] = c.Message("chat")
	c.RenderArgs["contactsNav"] = c.Message("contacts")
	return nil
}

func (c *Notes) Before() revel.Result {
	// Rendering useful info here.
	resume := jsonresume.New()
	c.RenderArgs["headerName"] = resume.Basics.Name
	c.RenderArgs["headerTitle"] = resume.Basics.Label
	c.RenderArgs["aboutNav"] = c.Message("about")
	c.RenderArgs["notesNav"] = c.Message("notes")
	c.RenderArgs["chatNav"] = c.Message("chat")
	c.RenderArgs["contactsNav"] = c.Message("contacts")
	return nil
}

func (c *Chat) Before() revel.Result {
	// Rendering useful info here.
	resume := jsonresume.New()
	c.RenderArgs["headerName"] = resume.Basics.Name
	c.RenderArgs["headerTitle"] = resume.Basics.Label
	c.RenderArgs["aboutNav"] = c.Message("about")
	c.RenderArgs["notesNav"] = c.Message("notes")
	c.RenderArgs["chatNav"] = c.Message("chat")
	c.RenderArgs["contactsNav"] = c.Message("contacts")
	return nil
}

func (c *Contacts) Before() revel.Result {
	// Rendering useful info here.
	resume := jsonresume.New()
	c.RenderArgs["headerName"] = resume.Basics.Name
	c.RenderArgs["headerTitle"] = resume.Basics.Label
	c.RenderArgs["aboutNav"] = c.Message("about")
	c.RenderArgs["notesNav"] = c.Message("notes")
	c.RenderArgs["chatNav"] = c.Message("chat")
	c.RenderArgs["contactsNav"] = c.Message("contacts")
	return nil
}

func init() {
	revel.InterceptMethod((*App).Before, revel.BEFORE)
	revel.InterceptMethod((*About).Before, revel.BEFORE)
	revel.InterceptMethod((*Notes).Before, revel.BEFORE)
	revel.InterceptMethod((*Chat).Before, revel.BEFORE)
	revel.InterceptMethod((*Contacts).Before, revel.BEFORE)
}
