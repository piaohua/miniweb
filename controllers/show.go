package controllers

import (
	"miniweb/models"
)

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

// Gate show gate list
func (s *ShowController) Gate() {

	if s.isPost() {
		s.Redirect("/", 302)
		return
	}

	list := models.GetGates()

	s.jsonResult(list)
}

// Share show share list
func (s *ShowController) Share() {

	if s.isPost() {
		s.Redirect("/", 302)
		return
	}

	list := models.GetShares()

	s.jsonResult(list)
}

// Invite show invite list
func (s *ShowController) Invite() {

	if s.isPost() {
		s.Redirect("/", 302)
		return
	}

	list := models.GetInvites()

	s.jsonResult(list)
}

// Rank show rank list
func (s *ShowController) Rank() {

	if s.isPost() {
		s.Redirect("/", 302)
		return
	}

	Type, _ := s.GetInt32("type")
	Gateid, _ := s.GetInt32("gateid")

	key := models.RankKey(Type, Gateid)
	list := models.GetRankInfo(key)

	s.jsonResult(list)
}
