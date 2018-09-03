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
		beego.Error("set shop err: ", err)
		jsonData := models.WxErr{
			ErrCode: int(pb.SetShopFailed),
			ErrMsg:  err.Error(),
		}
		s.jsonResult(jsonData)
		return
	}

	if !models.UpsertShop(shop) {
		beego.Error("set shop err: ", err)
		jsonData := models.WxErr{
			ErrCode: int(pb.SetShopFailed),
			ErrMsg:  "set shop failed",
		}
		s.jsonResult(jsonData)
		return
	}
	beego.Info("set shop success: ", shop)

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
		beego.Error("set prize err: ", err)
		jsonData := models.WxErr{
			ErrCode: int(pb.SetPrizeFailed),
			ErrMsg:  err.Error(),
		}
		s.jsonResult(jsonData)
		return
	}

	if !models.UpsertPrize(prize) {
		beego.Error("set prize err: ", err)
		jsonData := models.WxErr{
			ErrCode: int(pb.SetPrizeFailed),
			ErrMsg:  "set prize failed",
		}
		s.jsonResult(jsonData)
		return
	}
	beego.Info("set prize success: ", prize)

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
		beego.Error("set prop err: ", err)
		jsonData := models.WxErr{
			ErrCode: int(pb.SetPropFailed),
			ErrMsg:  err.Error(),
		}
		s.jsonResult(jsonData)
		return
	}

	if !models.UpsertProp(prop) {
		beego.Error("set prop err: ", err)
		jsonData := models.WxErr{
			ErrCode: int(pb.SetPropFailed),
			ErrMsg:  "set prop failed",
		}
		s.jsonResult(jsonData)
		return
	}
	beego.Info("set prop success: ", prop)

	s.jsonResult(prop)
}

// Gate set gate info
func (s *SetController) Gate() {

	if !s.isPost() {
		s.Redirect("/", 302)
		return
	}

	beego.Info("set prop json: ", string(s.Ctx.Input.RequestBody))
	var gate models.Gate
	var err error
	err = json.Unmarshal(s.Ctx.Input.RequestBody, &gate)
	if err != nil {
		beego.Error("set gate err: ", err)
		jsonData := models.WxErr{
			ErrCode: int(pb.SetGateFailed),
			ErrMsg:  err.Error(),
		}
		s.jsonResult(jsonData)
		return
	}

	if !models.UpsertGate(gate) {
		beego.Error("set gate err: ", err)
		jsonData := models.WxErr{
			ErrCode: int(pb.SetGateFailed),
			ErrMsg:  "set gate failed",
		}
		s.jsonResult(jsonData)
		return
	}
	beego.Info("set gate success: ", gate)

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

	msg := &pb.ChangeCurrency{
		Userid: userid,
		Coin:   num,
	}
	MSPid.Tell(msg)

	jsonData := &models.WxErr{}
	s.jsonResult(jsonData)
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

	msg := &pb.ChangeCurrency{
		Userid:  userid,
		Diamond: num,
	}
	MSPid.Tell(msg)

	jsonData := &models.WxErr{}
	s.jsonResult(jsonData)
}

// Close ws close
func (s *SetController) Close() {

	if !s.isPost() {
		s.Redirect("/", 302)
		return
	}

	//close
	Handler.Close()
	StopMS()

	jsonData := &models.WxErr{}
	s.jsonResult(jsonData)
}
