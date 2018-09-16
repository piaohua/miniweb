package models

import (
	"sort"
	"sync"
	"time"

	"github.com/globalsign/mgo/bson"
)

//Rank 每日分享配置
type Rank struct {
	ID    string     `bson:"_id" json:"id"`      //ID
	Info  []RankInfo `bson:"info" json:"info"`   //info
	Utime time.Time  `bson:"utime" json:"utime"` //更新时间
	Ctime time.Time  `bson:"ctime" json:"ctime"` //创建时间
}

//RankInfo rank info
type RankInfo struct {
	Gateid    int32  `bson:"gateid,omitempty" json:"gateid,omitempty"` //id
	Type      int32  `bson:"type,omitempty" json:"type,omitempty"`     //type
	Userid    string `bson:"userid,omitempty", json:"userid,omitempty"`
	NickName  string `bson:"nick_name,omitempty", json:"nick_name,omitempty"`
	AvatarUrl string `bson:"avatar_url,omitempty", json:"avatar_url,omitempty"`
	Score     int32  `bson:"score,omitempty" json:"score,omitempty,omitempty"` //score
}

//Get 加载
func (t *Rank) Get() {
	Get(Ranks, t.ID, t)
}

//Save 写入数据库
func (t *Rank) Save() bool {
	t.Ctime = bson.Now()
	return Insert(Ranks, t)
}

//Upsert 更新数据库
func (t *Rank) Upsert() bool {
	t.Utime = bson.Now()
	return Upsert(Ranks, bson.M{"_id": t.ID}, t)
}

//HasRank key是否存在
func HasRank(key string) bool {
	return Has(Ranks, bson.M{"_id": key})
}

//RankUpsert 更新数据库
func RankUpsert(key string, info []RankInfo) bool {
	if HasRank(key) {
		return Update(Ranks, bson.M{"_id": key},
			bson.M{"$set": bson.M{"info": info, "utime": bson.Now()}})
	}
	t := &Rank{ID: key, Info: info}
	return t.Save()
}

//RankKey cache rank unique key
func RankKey(Type, Gateid int32) string {
	return "rank" + GateKey(Type, Gateid)
}

//GetRankInfo get rank from cache by id
func GetRankInfo(key string) (list []RankInfo) {
	//key := RankKey(Type, Gateid)
	if !Cache.IsExist(key) {
		rank := new(Rank)
		rank.ID = key
		rank.Get()
		Cache.Put(key, rank.Info, (300 * time.Second))
		list = rank.Info
		return
	}
	if v := Cache.Get(key); v != nil {
		if val, ok := v.([]RankInfo); ok {
			list = val
		}
	}
	return
}

var (
	lock sync.Mutex
)

//NewRankInfo new rank info
func NewRankInfo(Type, Gateid, score int32, user *User) RankInfo {
	return RankInfo{
		Type:      Type,
		Gateid:    Gateid,
		Score:     score,
		Userid:    user.ID,
		NickName:  user.NickName,
		AvatarUrl: user.AvatarUrl,
	}
}

//SetRankInfo set rank TODO 优化
func SetRankInfo(info RankInfo) (list []RankInfo) {
	lock.Lock()
	defer lock.Unlock()
	key := RankKey(info.Type, info.Gateid)
	list = GetRankInfo(key)
	for k, v := range list {
		if v.Userid == info.Userid && v.Score >= info.Score {
			return
		}
		if v.Userid == info.Userid && v.Score < info.Score {
			list[k] = info
			SortRank(list)
			Cache.Put(key, list, (300 * time.Second))
			RankUpsert(key, list)
			return
		}
	}
	if len(list) == 0 {
		list = append(list, info)
		Cache.Put(key, list, (300 * time.Second))
		RankUpsert(key, list)
		return
	}
	if list[len(list)-1].Score >= info.Score {
		return
	}
	if len(list) < 10 {
		list = append(list, info)
		SortRank(list)
		Cache.Put(key, list, (300 * time.Second))
		RankUpsert(key, list)
		return
	}
	list = append(list, info)
	SortRank(list)
	list = list[:10]
	Cache.Put(key, list, (300 * time.Second))
	RankUpsert(key, list)
	return
}

//SortRank sort rank
func SortRank(list []RankInfo) {
	sort.Slice(list, func(i, j int) bool {
		return list[i].Score >= list[j].Score
	})
}
