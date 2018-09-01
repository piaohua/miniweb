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

// Shop show shop list
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

// Prop show prize list
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
		beego.Error("set prize err: ", err)
		return
	}

	s.jsonResult(prop)
}
