package controllers

import "miniweb/models"

// ShowController show
type ShowController struct {
	baseController // Embed to use methods that are implemented in baseController.
}

// Shop show shop list
func (s *ShowController) Shop() {

	if s.isPost() {
		s.Redirect("/", 302)
		return
	}

	list := models.GetShops()

	s.jsonResult(list)
}

// Prize show prize list
func (s *ShowController) Prize() {

	if s.isPost() {
		s.Redirect("/", 302)
		return
	}

	list := models.GetLoginPrizes()

	s.jsonResult(list)
}

// Prop show prize list
func (s *ShowController) Prop() {

	if s.isPost() {
		s.Redirect("/", 302)
		return
	}

	list := models.GetProps()

	s.jsonResult(list)
}
