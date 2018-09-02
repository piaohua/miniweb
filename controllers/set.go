package controllers

import (
	"encoding/json"

	"miniweb/models"
	"miniweb/pb"

	"github.com/astaxie/beego"
)

// SetController show
type SetController struct {
	baseController // Embed to use methods that are implemented in baseController.
}

// Shop set shop list
func (s *SetController) Shop() {

	if !s.isPost() {
		s.Redirect("/", 302)
		return
	}

	var shop models.Shop
	var err error
	err = json.Unmarshal(s.Ctx.Input.RequestBody, &shop)
	if err != nil {
		jsonData := models.WxErr{
			ErrCode: int(pb.SetShopFailed),
			ErrMsg:  err.Error(),
		}
		s.jsonResult(jsonData)
		beego.Error("set shop err: ", err)
		return
	}

	if !models.UpsertShop(shop) {
		jsonData := models.WxErr{
			ErrCode: int(pb.SetShopFailed),
			ErrMsg:  "set shop failed",
		}
		s.jsonResult(jsonData)
		beego.Error("set shop err: ", err)
		return
	}

	s.jsonResult(shop)
}

// Prize show prize list
func (s *SetController) Prize() {

	if !s.isPost() {
		s.Redirect("/", 302)
		return
	}

	var prize models.LoginPrize
	var err error
	err = json.Unmarshal(s.Ctx.Input.RequestBody, &prize)
	if err != nil {
		jsonData := models.WxErr{
			ErrCode: int(pb.SetPrizeFailed),
			ErrMsg:  err.Error(),
		}
		s.jsonResult(jsonData)
		beego.Error("set prize err: ", err)
		return
	}

	if !models.UpsertPrize(prize) {
		jsonData := models.WxErr{
			ErrCode: int(pb.SetPrizeFailed),
			ErrMsg:  "set prize failed",
		}
		s.jsonResult(jsonData)
		beego.Error("set prize err: ", err)
		return
	}

	s.jsonResult(prize)
}

// Prop set prize list
func (s *SetController) Prop() {

	if !s.isPost() {
		s.Redirect("/", 302)
		return
	}

	var prop models.Prop
	var err error
	err = json.Unmarshal(s.Ctx.Input.RequestBody, &prop)
	if err != nil {
		jsonData := models.WxErr{
			ErrCode: int(pb.SetPropFailed),
			ErrMsg:  err.Error(),
		}
		s.jsonResult(jsonData)
		beego.Error("set prop err: ", err)
		return
	}

	if !models.UpsertProp(prop) {
		jsonData := models.WxErr{
			ErrCode: int(pb.SetPropFailed),
			ErrMsg:  "set prop failed",
		}
		s.jsonResult(jsonData)
		beego.Error("set prop err: ", err)
		return
	}

	s.jsonResult(prop)
}

// Gate set gate info
func (s *SetController) Gate() {

	if !s.isPost() {
		s.Redirect("/", 302)
		return
	}

	var gate models.Gate
	var err error
	err = json.Unmarshal(s.Ctx.Input.RequestBody, &gate)
	if err != nil {
		jsonData := models.WxErr{
			ErrCode: int(pb.SetGateFailed),
			ErrMsg:  err.Error(),
		}
		s.jsonResult(jsonData)
		beego.Error("set gate err: ", err)
		return
	}

	if !models.UpsertGate(gate) {
		jsonData := models.WxErr{
			ErrCode: int(pb.SetGateFailed),
			ErrMsg:  "set gate failed",
		}
		s.jsonResult(jsonData)
		beego.Error("set gate err: ", err)
		return
	}

	s.jsonResult(gate)
}

// Coin set coin
func (s *SetController) Coin() {

	if !s.isPost() {
		s.Redirect("/", 302)
		return
	}

	userid := s.GetString("userid")
	num, err := s.GetInt64("num")
	if err != nil {
		beego.Error("set coin err: ", err)
		s.Redirect("/", 302)
		return
	}
	beego.Debug("set coin userid: ", userid, ", num: ", num)

	//TODO

	//s.jsonResult(prop)
}

// Diamond set coin
func (s *SetController) Diamond() {

	if !s.isPost() {
		s.Redirect("/", 302)
		return
	}

	userid := s.GetString("userid")
	num, err := s.GetInt64("num")
	if err != nil {
		beego.Error("set diamond err: ", err)
		s.Redirect("/", 302)
		return
	}
	beego.Debug("set diamond userid: ", userid, ", num: ", num)

	//TODO

	//s.jsonResult(prop)
}

// Close ws close
func (s *SetController) Close() {

	if !s.isPost() {
		s.Redirect("/", 302)
		return
	}

	Handler.Close()

	jsonData := &models.WxErr{}
	s.jsonResult(jsonData)
}
