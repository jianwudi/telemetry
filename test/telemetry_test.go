package test

import (
	"fmt"
	_ "telemetry/internal/telemetry/collect"
	"telemetry/internal/telemetry/server"
	_ "telemetry/internal/telemetry/server"
	"telemetry/internal/telemetry/service"
	"telemetry/internal/telemetry/store"
	"testing"
	"time"

	"github.com/marmotedu/iam/pkg/log"
)

func TestSenorPath(t *testing.T) {
	serv := service.NewService(store.StoreClient)
	/* 	sensorGrpEnt := &store.SensorGroupEntry{
	   		SensorGroupId: "group1",
	   		SensorPath:    "an_gpon_pm_olt_traffic:GponPmOltTraffics.2.1",
	   	}
	   	err := serv.SensorGroup().Create(sensorGrpEnt)
	   	if err != nil {
	   		fmt.Printf("%+v", err)
	   	} */
	sensorGrpEnt1 := &store.SensorGroupEntry{
		SensorGroupId: "group1",
		SensorPath:    "an_gpon_pm_olt_traffic:GponPmOltChannelTraffics.2.1.0",
	}
	err := serv.SensorGroup().Create(sensorGrpEnt1)
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	sensorGrpEnt2 := &store.SensorGroupEntry{
		SensorGroupId: "group1",
		SensorPath:    "an_gpon_pm_olt_traffic:GponPmOltChannelTraffics.2.1.1",
	}
	err = serv.SensorGroup().Create(sensorGrpEnt2)
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	serv.SensorGroup().GetALL()
	log.Info("after delete")
	sensorGrpEnt2 = &store.SensorGroupEntry{
		SensorGroupId: "group1",
		//		SensorPath:    "an_gpon_pm_olt_traffic:GponPmOltChannelTraffics.2.1.1",
	}
	err = serv.SensorGroup().Delete(sensorGrpEnt2)
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	serv.SensorGroup().GetALL()
}

func TestSub(t *testing.T) {
	go server.TelemetryDataTimer()
	serv := service.NewService(store.StoreClient)
	/* 	sensorGrpEnt := &store.SensorGroupEntry{
	   		   SensorGroupId: "group1",
	   		   SensorPath:    "an_gpon_pm_olt_traffic:GponPmOltTraffics.2.1",
	   	   }
	   	   err := serv.SensorGroup().Create(sensorGrpEnt)
	   	   if err != nil {
	   		   fmt.Printf("%+v", err)
	   	   } */
	sensorGrpEnt1 := &store.SensorGroupEntry{
		SensorGroupId: "group1",
		SensorPath:    "an_gpon_pm_olt_traffic:GponPmOltChannelTraffics.2.1.0",
	}
	err := serv.SensorGroup().Create(sensorGrpEnt1)
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	sensorGrpEnt2 := &store.SensorGroupEntry{
		SensorGroupId: "group2",
		SensorPath:    "an_gpon_pm_olt_traffic:GponPmOltTraffics.2.2.0",
	}
	err = serv.SensorGroup().Create(sensorGrpEnt2)
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	destGrpEnt := &store.DestGroupEntry{
		DestGroupId: "group1",
		DestIp:      0, //0xC0A803EA, //192.168.3.234
		DestPort:    50051}

	err = serv.DestGroup().Create(destGrpEnt)
	if err != nil {
		fmt.Printf("%+v", err)
		//		return

	}
	subGrpEnt := &store.SubscriptionGroupEntry{
		SubscriptionId: "name1-group1-group1",
		SensorGroupId:  "group1",
		DestGroupId:    "group1",
		Interval:       500}
	err = serv.SubGroup().Create(subGrpEnt)
	if err != nil {
		fmt.Printf("%+v", err)
		//	return
	}

	subGrpEnt2 := &store.SubscriptionGroupEntry{
		SubscriptionId: "name1-group2-group1",
		SensorGroupId:  "group2",
		DestGroupId:    "group1",
		Interval:       500}
	err = serv.SubGroup().Create(subGrpEnt2)
	if err != nil {
		fmt.Printf("%+v", err)
		//	return
	}
	serv.SensorData().GetALL()
	subGrpEnt = &store.SubscriptionGroupEntry{
		SubscriptionId: "name1-group1-",
		SensorGroupId:  "group1"}
	//	serv.GrpcData().GetALL()
	time.Sleep(1 * time.Second)
	err = serv.SubGroup().Delete(subGrpEnt)
	if err != nil {
		fmt.Printf("%+v", err)
		//	return
	}
	for {
		time.Sleep(time.Second)
	}
}
