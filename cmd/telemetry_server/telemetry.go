package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"telemetry/internal/telemetry/collect"
	_ "telemetry/internal/telemetry/collect"

	"telemetry/internal/telemetry/server"
	"telemetry/internal/telemetry/service"
	"telemetry/internal/telemetry/store"
	"time"

	"github.com/marmotedu/iam/pkg/log"
)

type sensorGroup struct {
	TelemetrySensorGroupID   [64]byte  /* <attr=1,key,persist> */
	TelemetrySensorPath      [128]byte /* <attr=2,key,persist> */
	TelemetrySensorRowStatus int32
}

type destGroup struct {
	TelemetryDestGroupId        [63 + 1]byte
	TelemetryDestIp             uint32
	TelemetryDestPort           uint32
	TelemetryDestGroupRowStatus int32
}

type subGroup struct {
	TelemetrySubscriptionId    [63 + 1]byte
	TelemetrySensorGroupId     [63 + 1]byte
	TelemetryDestGroupId       [63 + 1]byte
	TelemetryInterva           uint32
	TelemetrySubGroupRowStatus int32
}

const (
	SensorGroupTableID = iota + 1
	DestGroupTableID
	SubscriptionGroupTableID
)

func receiveSnmpMsg() {
	//
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 40000,
	})
	//	listen, err := net.ListenUDP("udp", udpaddr)
	if err != nil {
		fmt.Println("listen failed. err:", err)
		return
	}

	defer listen.Close()
	serv := service.NewService(store.StoreClient)
	log.Info("start listen1111111111111")
	for {
		var data [500]byte
		//		var sensorpathIndex int
		n, _, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read udp failed,err:", err)
			continue
		}

		fmt.Printf("data1:%v count:%v\n", data[:n], n)

		if data[0] == SensorGroupTableID {
			log.Info("SensorGroupTable")
			var sensorGrpEnt *store.SensorGroupEntry
			sersorGrp := &sensorGroup{}
			buf := bytes.NewReader(data[1:n])
			err = binary.Read(buf, binary.BigEndian, sersorGrp)
			if err != nil {
				log.Info("error")
				continue
			}
			//		buf = bytes.NewReader(data[85:n])
			//		binary.Read(buf, binary.BigEndian, &sersorGrp.TelemetrySensorRowStatus)
			log.Infof("group : %s ", sersorGrp.TelemetrySensorGroupID[:])
			log.Infof("sensorpath : %s ", sersorGrp.TelemetrySensorPath[:])
			log.Infof("rowstatus:%d ", sersorGrp.TelemetrySensorRowStatus)
			if sersorGrp.TelemetrySensorPath[0] != 0 {
				sensorGrpEnt = &store.SensorGroupEntry{
					SensorGroupId: string(sersorGrp.TelemetrySensorGroupID[:]),
					SensorPath:    string(sersorGrp.TelemetrySensorPath[:])}
			} else {
				sensorGrpEnt = &store.SensorGroupEntry{
					SensorGroupId: string(sersorGrp.TelemetrySensorGroupID[:]),
				}

			}

			if sersorGrp.TelemetrySensorRowStatus == 0 {
				err := serv.SensorGroup().Create(sensorGrpEnt)
				if err != nil {
					fmt.Printf("%+v", err)
				}
			} else if sersorGrp.TelemetrySensorRowStatus == 2 {
				err := serv.SensorGroup().Delete(sensorGrpEnt)
				if err != nil {
					fmt.Printf("%+v", err)
				}
			}
		} else if data[0] == DestGroupTableID {
			destGrp := &destGroup{}
			buf := bytes.NewReader(data[1:n])
			err = binary.Read(buf, binary.BigEndian, destGrp)
			if err != nil {
				continue
			}
			log.Infof("destGrp:%s %d %d\n", destGrp.TelemetryDestGroupId[:], destGrp.TelemetryDestIp,
				destGrp.TelemetryDestPort)
			log.Infof("%d\n", destGrp.TelemetryDestGroupRowStatus)
			destGrpEnt := &store.DestGroupEntry{
				DestGroupId: string(destGrp.TelemetryDestGroupId[:]),
				DestIp:      destGrp.TelemetryDestIp,
				DestPort:    destGrp.TelemetryDestPort}
			if destGrp.TelemetryDestGroupRowStatus == 0 {
				err := serv.DestGroup().Create(destGrpEnt)
				if err != nil {
					fmt.Printf("%+v", err)
				}
			} else if destGrp.TelemetryDestGroupRowStatus == 2 {
				err := serv.DestGroup().Delete(destGrpEnt)
				if err != nil {
					fmt.Printf("%+v", err)
				}
			}
		} else if data[0] == SubscriptionGroupTableID {
			subGrp := &subGroup{}
			buf := bytes.NewReader(data[1:n])
			err = binary.Read(buf, binary.BigEndian, subGrp)
			if err != nil {
				continue
			}
			var subGrpEnt *store.SubscriptionGroupEntry
			log.Infof("subGrp:%s %s %s %d\n", subGrp.TelemetrySubscriptionId, subGrp.TelemetrySensorGroupId[:],
				subGrp.TelemetryDestGroupId[:], subGrp.TelemetryInterva)
			log.Infof("%d\n", subGrp.TelemetrySubGroupRowStatus)

			if subGrp.TelemetrySensorGroupId[0] == 0 && subGrp.TelemetryDestGroupId[0] == 0 {
				subGrpEnt = &store.SubscriptionGroupEntry{
					SubscriptionId: fmt.Sprintf("%s-%s-%s", subGrp.TelemetrySubscriptionId, subGrp.TelemetrySensorGroupId, subGrp.TelemetryDestGroupId),
					Interval:       subGrp.TelemetryInterva}
			} else if subGrp.TelemetrySensorGroupId[0] == 0 {
				subGrpEnt = &store.SubscriptionGroupEntry{
					SubscriptionId: fmt.Sprintf("%s-%s-%s", subGrp.TelemetrySubscriptionId, subGrp.TelemetrySensorGroupId, subGrp.TelemetryDestGroupId),
					DestGroupId:    string(subGrp.TelemetryDestGroupId[:]),
					Interval:       subGrp.TelemetryInterva}

			} else if subGrp.TelemetryDestGroupId[0] == 0 {
				subGrpEnt = &store.SubscriptionGroupEntry{
					SubscriptionId: fmt.Sprintf("%s-%s-%s", subGrp.TelemetrySubscriptionId, subGrp.TelemetrySensorGroupId, subGrp.TelemetryDestGroupId),
					SensorGroupId:  string(subGrp.TelemetrySensorGroupId[:]),
					Interval:       subGrp.TelemetryInterva}
			} else {
				subGrpEnt = &store.SubscriptionGroupEntry{
					SubscriptionId: fmt.Sprintf("%s-%s-%s", subGrp.TelemetrySubscriptionId, subGrp.TelemetrySensorGroupId, subGrp.TelemetryDestGroupId),
					SensorGroupId:  string(subGrp.TelemetrySensorGroupId[:]),
					DestGroupId:    string(subGrp.TelemetryDestGroupId[:]),
					Interval:       subGrp.TelemetryInterva}
			}

			if subGrp.TelemetrySubGroupRowStatus == 0 {
				err := serv.SubGroup().Create(subGrpEnt)
				if err != nil {
					fmt.Printf("%+v", err)
				}
			} else if subGrp.TelemetrySubGroupRowStatus == 2 {
				err := serv.SubGroup().Delete(subGrpEnt)
				if err != nil {
					fmt.Printf("%+v", err)
				}
			}

		} else {
			err := serv.SensorGroup().GetALL()
			if err != nil {
				fmt.Printf("%+v", err)
			}
			err = serv.DestGroup().GetALL()
			if err != nil {
				fmt.Printf("%+v", err)
			}
			err = serv.SubGroup().GetALL()
			if err != nil {
				fmt.Printf("%+v", err)
			}
			err = serv.SensorData().GetALL()
			if err != nil {
				fmt.Printf("%+v", err)
			}
			err = serv.GrpcData().GetALL()
			if err != nil {
				fmt.Printf("%+v", err)
			}
		}
	}
}
func debugSub() {
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
	/* 	subGrpEnt1 := &store.SubscriptionGroupEntry{
	   		SubscriptionId: "name1--"}
	   	err = serv.SubGroup().Delete(subGrpEnt1)
	   	if err != nil {
	   		fmt.Printf("%+v", err)
	   		//	return
	   	} */
	serv.SensorData().GetALL()
	serv.GrpcData().GetALL()
}
func main() {
	//	log.Info("1111")
	go collect.ConnectToRedis()
	time.Sleep(time.Second)

	//接收上报的信息

	go collect.StartCollectDataLoop()
	go server.TelemetryDataTimer()
	receiveSnmpMsg()
	//debugSub()
	for {
		time.Sleep(time.Second)
	}
}
