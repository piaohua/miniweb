package models

import (
	"time"

	"miniweb/pb"

	"github.com/globalsign/mgo/bson"
)

//Shop 商城
type Shop struct {
	Id     string    `bson:"_id" json:"id"`        //购买ID
	Status int32     `bson:"status" json:"status"` //物品状态
	Ptype  int32     `bson:"ptype" json:"ptype"`   //兑换的物品
	Payway int32     `bson:"payway" json:"payway"` //支付方式
	Number uint32    `bson:"number" json:"number"` //兑换的数量
	Price  uint32    `bson:"price" json:"price"`   //支付价格(单位元)
	Name   string    `bson:"name" json:"name"`     //物品名字
	Info   string    `bson:"info" json:"info"`     //物品信息
	Del    int       `bson:"del" json:"del"`       //是否移除
	Etime  time.Time `bson:"etime" json:"etime"`   //过期时间
	Ctime  time.Time `bson:"ctime" json:"ctime"`   //创建时间
}

//GetShopList 商城列表
func GetShopList() []Shop {
	var list []Shop
	ListByQ(Shops, bson.M{"del": 0, "etime": bson.M{"$gt": bson.Now()}}, &list)
	return list
}

//Save 写入数据库
func (t *Shop) Save() bool {
	t.Ctime = bson.Now()
	return Insert(Shops, t)
}

//SetShopList 添加商城物品
func SetShopList() {
	NewShop("1", int32(pb.SHOP_STATUS0), int32(pb.PROP_TYPE2),
		int32(pb.PAY_WAY1), 10000, 100, "金币", "100钻石兑换10000金币")
	NewShop("2", int32(pb.SHOP_STATUS1), int32(pb.PROP_TYPE1),
		int32(pb.PAY_WAY0), 10000, 100, "钻石", "100元兑换10000钻石")
	NewShop("3", int32(pb.SHOP_STATUS2), int32(pb.PROP_TYPE3),
		int32(pb.PAY_WAY1), 100, 10, "精力", "10钻石兑换100精力")
}

//NewShop 添加商品
func NewShop(id string, status, proptype, payway int32,
	number, price uint32, name, info string) {
	t := Shop{
		Id:     id,
		Status: status,
		Ptype:  proptype,
		Payway: payway,
		Number: number,
		Price:  price,
		Name:   name,
		Info:   info,
		Etime:  time.Now().AddDate(0, 0, 100),
	}
	//config.SetShop(t)
	t.Save()
}
