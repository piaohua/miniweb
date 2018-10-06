package controllers

import (
	"encoding/json"

	"miniweb/models"
	"miniweb/pb"

	"github.com/astaxie/beego"
)

// SetController show
type SetController struct {
	baseController      // Embed to use methods that are implemented in baseController.
	token          bool // have token
}

// Shop set shop list
func (s *SetController) Shop() {
	jsonData := &models.WxErr{}
	defer s.jsonResult(jsonData)

	if !s.isPost() {
		jsonData.ErrCode = int(pb.Failed)
		jsonData.ErrMsg = "method error"
		return
	}

	if !s.token {
		jsonData.ErrCode = int(pb.Failed)
		jsonData.ErrMsg = "token error"
		return
	}

	var shop models.Shop
	var err error
	err = json.Unmarshal(s.Ctx.Input.RequestBody, &shop)
	if err != nil {
		beego.Error("set shop err: ", err)
		jsonData.ErrCode = int(pb.SetShopFailed)
		jsonData.ErrMsg = err.Error()
		return
	}

	if !models.UpsertShop(shop) {
		beego.Error("set shop err: ", err)
		jsonData.ErrCode = int(pb.SetShopFailed)
		jsonData.ErrMsg = "set shop failed"
		return
	}
	beego.Info("set shop success: ", shop)
}

// Prize set prize list
func (s *SetController) Prize() {
	jsonData := &models.WxErr{}
	defer s.jsonResult(jsonData)

	if !s.isPost() {
		jsonData.ErrCode = int(pb.Failed)
		jsonData.ErrMsg = "method error"
		return
	}

	if !s.token {
		jsonData.ErrCode = int(pb.Failed)
		jsonData.ErrMsg = "token error"
		return
	}

	var prize models.LoginPrize
	var err error
	err = json.Unmarshal(s.Ctx.Input.RequestBody, &prize)
	if err != nil {
		beego.Error("set prize err: ", err)
		jsonData.ErrCode = int(pb.SetPrizeFailed)
		jsonData.ErrMsg = err.Error()
		return
	}

	if !models.UpsertPrize(prize) {
		beego.Error("set prize err: ", err)
		jsonData.ErrCode = int(pb.SetPropFailed)
		jsonData.ErrMsg = "set prize failed"
		return
	}
	beego.Info("set prize success: ", prize)
}

// Prop set prize list
func (s *SetController) Prop() {
	jsonData := &models.WxErr{}
	defer s.jsonResult(jsonData)

	if !s.isPost() {
		jsonData.ErrCode = int(pb.Failed)
		jsonData.ErrMsg = "method error"
		return
	}

	if !s.token {
		jsonData.ErrCode = int(pb.Failed)
		jsonData.ErrMsg = "token error"
		return
	}

	var prop models.Prop
	var err error
	err = json.Unmarshal(s.Ctx.Input.RequestBody, &prop)
	if err != nil {
		beego.Error("set prop err: ", err)
		jsonData.ErrCode = int(pb.SetPropFailed)
		jsonData.ErrMsg = err.Error()
		return
	}

	if !models.UpsertProp(prop) {
		beego.Error("set prop err: ", err)
		jsonData.ErrCode = int(pb.SetPropFailed)
		jsonData.ErrMsg = "set prop failed"
		return
	}
	beego.Info("set prop success: ", prop)
}

// Gate set gate info
func (s *SetController) Gate() {
	jsonData := &models.WxErr{}
	defer s.jsonResult(jsonData)

	if !s.isPost() {
		jsonData.ErrCode = int(pb.Failed)
		jsonData.ErrMsg = "method error"
		return
	}

	if !s.token {
		jsonData.ErrCode = int(pb.Failed)
		jsonData.ErrMsg = "token error"
		return
	}

	beego.Info("set prop json: ", string(s.Ctx.Input.RequestBody))
	var gate models.Gate
	var err error
	err = json.Unmarshal(s.Ctx.Input.RequestBody, &gate)
	if err != nil {
		beego.Error("set gate err: ", err)
		jsonData.ErrCode = int(pb.SetGateFailed)
		jsonData.ErrMsg = err.Error()
		return
	}

	if !models.UpsertGate(gate) {
		beego.Error("set gate err: ", err)
		jsonData.ErrCode = int(pb.SetGateFailed)
		jsonData.ErrMsg = "set gate failed"
		return
	}
	beego.Info("set gate success: ", gate)
}

// Share set share list
func (s *SetController) Share() {
	jsonData := &models.WxErr{}
	defer s.jsonResult(jsonData)

	if !s.isPost() {
		jsonData.ErrCode = int(pb.Failed)
		jsonData.ErrMsg = "method error"
		return
	}

	if !s.token {
		jsonData.ErrCode = int(pb.Failed)
		jsonData.ErrMsg = "token error"
		return
	}

	var share models.Share
	var err error
	err = json.Unmarshal(s.Ctx.Input.RequestBody, &share)
	if err != nil {
		beego.Error("set share err: ", err)
		jsonData.ErrCode = int(pb.SetShareFailed)
		jsonData.ErrMsg = err.Error()
		return
	}

	if !models.UpsertShare(share) {
		beego.Error("set share err: ", err)
		jsonData.ErrCode = int(pb.SetPropFailed)
		jsonData.ErrMsg = "set share failed"
		return
	}
	beego.Info("set share success: ", share)
}

// Invite set invite list
func (s *SetController) Invite() {
	jsonData := &models.WxErr{}
	defer s.jsonResult(jsonData)

	if !s.isPost() {
		jsonData.ErrCode = int(pb.Failed)
		jsonData.ErrMsg = "method error"
		return
	}

	if !s.token {
		jsonData.ErrCode = int(pb.Failed)
		jsonData.ErrMsg = "token error"
		return
	}

	var invite models.Invite
	var err error
	err = json.Unmarshal(s.Ctx.Input.RequestBody, &invite)
	if err != nil {
		beego.Error("set invite err: ", err)
		jsonData.ErrCode = int(pb.SetInviteFailed)
		jsonData.ErrMsg = err.Error()
		return
	}

	if !models.UpsertInvite(invite) {
		beego.Error("set invite err: ", err)
		jsonData.ErrCode = int(pb.SetPropFailed)
		jsonData.ErrMsg = "set invite failed"
		return
	}
	beego.Info("set invite success: ", invite)
}

// Coin set coin
func (s *SetController) Coin() {
	jsonData := &models.WxErr{}
	defer s.jsonResult(jsonData)

	if !s.isPost() {
		jsonData.ErrCode = int(pb.Failed)
		jsonData.ErrMsg = "method error"
		return
	}

	if !s.token {
		jsonData.ErrCode = int(pb.Failed)
		jsonData.ErrMsg = "token error"
		return
	}

	userid := s.GetString("userid")
	num, err := s.GetInt64("num")
	if err != nil {
		beego.Error("set coin err: ", err)
		jsonData.ErrCode = int(pb.Failed)
		jsonData.ErrMsg = err.Error()
		return
	}
	beego.Debug("set coin userid: ", userid, ", num: ", num)

	msg := &pb.ChangeCurrency{
		Userid: userid,
		Coin:   num,
	}
	MSPid.Tell(msg)
}

// Diamond set coin
func (s *SetController) Diamond() {
	jsonData := &models.WxErr{}
	defer s.jsonResult(jsonData)

	if !s.isPost() {
		jsonData.ErrCode = int(pb.Failed)
		jsonData.ErrMsg = "method error"
		return
	}

	if !s.token {
		jsonData.ErrCode = int(pb.Failed)
		jsonData.ErrMsg = "token error"
		return
	}

	userid := s.GetString("userid")
	num, err := s.GetInt64("num")
	if err != nil {
		beego.Error("set diamond err: ", err)
		jsonData.ErrCode = int(pb.Failed)
		jsonData.ErrMsg = err.Error()
		return
	}
	beego.Debug("set diamond userid: ", userid, ", num: ", num)

	msg := &pb.ChangeCurrency{
		Userid:  userid,
		Diamond: num,
	}
	MSPid.Tell(msg)
}

// Close ws close
func (s *SetController) Close() {
	jsonData := &models.WxErr{}
	defer s.jsonResult(jsonData)

	if !s.isPost() {
		jsonData.ErrCode = int(pb.Failed)
		jsonData.ErrMsg = "method error"
		return
	}

	if !s.token {
		jsonData.ErrCode = int(pb.Failed)
		jsonData.ErrMsg = "token error"
		return
	}

	//close
	CloseNode()
	CloseMS()
	Handler.Close()
	StopMS()
	StopNode()
}

// Prepare implemented Prepare() method for baseController.
func (s *SetController) Prepare() {
	token := s.Ctx.Request.Header.Get("token")
	setToken := beego.AppConfig.String("set.token")
	beego.Debug("token: ", token, ", setToken: ", setToken)
	if token != "" && token == setToken {
		s.token = true
	}
}
