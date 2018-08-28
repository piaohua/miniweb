package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
	"gohappy/data"
)

//Shop 商城
type Shop struct {
	Id     string    `bson:"_id" json:"id"`        //购买ID
	Status int32     `bson:"status" json:"status"` //物品状态
	Propid int32     `bson:"propid" json:"propid"` //兑换的物品
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
	//NewShop("1", 1, 2, 2, 10000, 100, "金币", "金币10000")
	//NewShop("2", 1, 2, 2, 20000, 200, "金币", "金币20000")
	//NewShop("3", 1, 2, 2, 50000, 450, "金币", "金币50000")
	//NewShop("4", 1, 1, 1, 100, 10, "钻石", "钻石100")
	//NewShop("5", 1, 1, 1, 200, 20, "钻石", "钻石200")
	//NewShop("6", 1, 1, 1, 500, 45, "钻石", "钻石500")
	NewShop("7", 1, 2, 1, 1200, 12, "金豆", "金豆1200")
	NewShop("8", 1, 2, 1, 3800, 38, "金豆", "金豆3800")
	NewShop("9", 1, 2, 1, 6800, 68, "金豆", "金豆6800")
	NewShop("10", 1, 2, 1, 12800, 128, "金豆", "金豆12800")
	NewShop("11", 1, 2, 1, 38000, 388, "金豆", "金豆38000")
	NewShop("12", 1, 2, 1, 59000, 599, "金豆", "金豆59000")
	NewShop("13", 1, 2, 1, 78900, 789, "金豆", "金豆78900")
	NewShop("14", 1, 2, 1, 99900, 999, "金豆", "金豆99900")
}

//NewShop 添加商品
func NewShop(id string, status, propid, payway int,
	number, price uint32, name, info string) {
	t := data.Shop{
		Id:     id,
		Status: status,
		Propid: propid,
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