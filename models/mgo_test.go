package models

import (
	"testing"

	"github.com/globalsign/mgo/bson"
)

func TestGroup(t *testing.T) {
	InitMgo("127.0.0.1", "27017", "", "", "test")
	var types []string
	ListByQ(Users, bson.M{"$group": bson.M{"userid": "$userid"}}, &types)
	t.Logf("%#v", types)
	m := bson.M{
		"$match": bson.M{
		//"diamond": bson.M{"$ne": 0},
		},
	}
	o := bson.M{
		"$project": bson.M{
			"_id":    1,
			"gender": 1,
		},
	}
	n := bson.M{
		"$group": bson.M{
			"_id": "$_id",
		},
	}
	operations := []bson.M{m, o, n}
	result := []bson.M{}
	pipe := Users.Pipe(operations)
	err2 := pipe.All(&result)
	t.Logf("%#v", err2)
	t.Logf("%#v", result)
}

func TestOr(t *testing.T) {
	InitMgo("127.0.0.1", "27017", "", "", "test")
	or := []bson.M{bson.M{"userid": "10030"}}
	m := bson.M{
		"$match": bson.M{
			"$or": or,
		},
	}
	operations := []bson.M{m}
	result := []bson.M{}
	pipe := Users.Pipe(operations)
	err2 := pipe.All(&result)
	t.Logf("%#v", err2)
	t.Logf("%#v", result)
}
