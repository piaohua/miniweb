package models

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
)

//healthcheck

//DatabaseCheck db conn check
type DatabaseCheck struct {
}

//Check check
func (dc *DatabaseCheck) Check() error {
	//if dc.isConnected() {
	if isConnected() {
		return nil
	}
	return errors.New("can't connect database")
}

//HealthCheck healthcheck
func HealthCheck() {
	toolbox.AddHealthCheck("database", &DatabaseCheck{})
}

//profile

//RunStatics statistics
func RunStatics() error {
	toolbox.StatisticsMap.AddStatistics("POST", "/code", "&admin.user", time.Duration(2000))
	toolbox.StatisticsMap.AddStatistics("GET", "/ws", "&admin.user", time.Duration(13000))
	toolbox.StatisticsMap.AddStatistics("GET", "/ws/login", "&admin.user", time.Duration(14000))

	beego.Info(toolbox.StatisticsMap.GetMap())

	data := toolbox.StatisticsMap.GetMapData()
	b, err := json.Marshal(data)
	if err != nil {
		beego.Error(err.Error())
	}

	beego.Info(string(b))
	return nil
}

//task

//RunTask task
func RunTask() {
	tk1 := toolbox.NewTask("tk1", "3 * * * * *", task1)
	tk2 := toolbox.NewTask("tk2", "3 * * * * *", RunStatics)
	toolbox.AddTask("tk1", tk1)
	toolbox.AddTask("tk2", tk2)
	toolbox.StartTask()
	//defer StopTask()
}

func task1() error {
	beego.Debug("tk1")
	return nil
}
