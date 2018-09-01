package models

import (
	"time"

	"miniweb/libs"
	"miniweb/pb"

	"github.com/astaxie/beego"
	"github.com/globalsign/mgo/bson"
)

//Shop 商城
type Shop struct {
	ID     string          `bson:"_id" json:"id"`        //购买ID
	Status int32           `bson:"status" json:"status"` //物品状态
	Ptype  int32           `bson:"ptype" json:"ptype"`   //兑换的物品
	Payway int32           `bson:"payway" json:"payway"` //支付方式
	Number uint32          `bson:"number" json:"number"` //兑换的数量
	Price  uint32          `bson:"price" json:"price"`   //支付价格(单位元)
	Name   string          `bson:"name" json:"name"`     //物品名字
	Info   string          `bson:"info" json:"info"`     //物品信息
	Prize  []ShopPrizeProp `bson:"prize" json:"prize"`   //prize
	Del    int             `bson:"del" json:"del"`       //是否移除
	Etime  time.Time       `bson:"etime" json:"etime"`   //过期时间
	Ctime  time.Time       `bson:"ctime" json:"ctime"`   //创建时间
}

//ShopPrizeProp 奖励
type ShopPrizeProp struct {
	Type   int32 `bson:"type" json:"type"`     //物品类型
	Number int32 `bson:"number" json:"number"` //物品数量
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

//Upsert 更新数据库
func (t *Shop) Upsert() bool {
	return Upsert(Shops, bson.M{"_id": t.ID}, t)
}

//Delete 删除数据
func (t *Shop) Delete() bool {
	return Delete(Shops, bson.M{"_id": t.ID})
}

//SetShopList 添加商城物品
func SetShopList() {
	NewShop("1", int32(pb.SHOP_STATUS0), int32(pb.PROP_TYPE2),
		int32(pb.PAY_WAY1), 10000, 100, "金币", "100钻石兑换10000金币")
	NewShop("2", int32(pb.SHOP_STATUS1), int32(pb.PROP_TYPE1),
		int32(pb.PAY_WAY0), 10000, 100, "钻石", "100元兑换10000钻石")
	NewShop("3", int32(pb.SHOP_STATUS2), int32(pb.PROP_TYPE9),
		int32(pb.PAY_WAY1), 100, 10, "初级精力瓶", "10钻石兑换100初级精力瓶")
	NewShop("4", int32(pb.SHOP_STATUS2), int32(pb.PROP_TYPE10),
		int32(pb.PAY_WAY1), 100, 10, "中级精力瓶", "10钻石兑换100中级精力瓶")
	NewShop("5", int32(pb.SHOP_STATUS2), int32(pb.PROP_TYPE11),
		int32(pb.PAY_WAY1), 100, 10, "高级精力瓶", "10钻石兑换100高级精力瓶")
}

//LoadShopList load shop info by shop.json
func LoadShopList() []Shop {
	filePath := "static/shop.json"
	list := make([]Shop, 0)
	err := libs.Load(filePath, &list)
	if err != nil {
		beego.Error("load shop err ", err)
	}
	return list
}

//InitShopList init shop to cache
func InitShopList() {
	list := GetShopList()
	//test use
	if !RunMode() {
		if len(list) == 0 {
			//SetShopList()
			list = LoadShopList()
		}
	}
	Cache.Put("shop", list, 0)
	for k, v := range list {
		Cache.Put(ShopKey(v.ID), &list[k], 0)
	}
}

//UpsertShop upsert shop
func UpsertShop(shop Shop) bool {
	key := ShopKey(shop.ID)
	list := GetShops()
	if shop.Del != 0 {
		if !shop.Delete() {
			beego.Error("shop delete err: ", shop)
			return false
		}
		Cache.Delete(key)
		for k, v := range list {
			if v.ID == shop.ID {
				list = append(list[:k], list[k+1:]...)
				break
			}
		}
		Cache.Put("shop", list, 0)
		return true
	}
	if !shop.Upsert() {
		beego.Error("shop upsert err: ", shop)
		return false
	}
	Cache.Put(key, &shop, 0)
	for k, v := range list {
		if v.ID == shop.ID {
			list[k] = shop
			break
		}
	}
	Cache.Put("shop", list, 0)
	return true
}

//NewShop 添加商品
func NewShop(id string, status, proptype, payway int32,
	number, price uint32, name, info string) {
	t := Shop{
		ID:     id,
		Status: status,
		Ptype:  proptype,
		Payway: payway,
		Number: number,
		Price:  price,
		Name:   name,
		Info:   info,
		Etime:  time.Now().AddDate(0, 0, 100),
	}
	t.Save()
}

//ShopKey cache shop unique key
func ShopKey(id string) string {
	return "shop" + id
}

//GetShops from cache
func GetShops() (l []Shop) {
	if v := Cache.Get("shop"); v != nil {
		if val, ok := v.([]Shop); ok {
			l = val
		}
	}
	return
}

//GetShop get shop from cache by id
func GetShop(id string) (shop *Shop) {
	if v := Cache.Get(ShopKey(id)); v != nil {
		if val, ok := v.(*Shop); ok {
			shop = val
		}
	}
	return
}
