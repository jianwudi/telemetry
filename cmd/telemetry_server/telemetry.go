package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"telemetry/internal/telemetry/service"
	"telemetry/internal/telemetry/store"
	"time"

	"github.com/marmotedu/iam/pkg/log"
)

type sensorGroup struct {
	TelemetrySensorGroupID   [63 + 1]byte  /* <attr=1,key,persist> */
	TelemetrySensorPath      [100 + 1]byte /* <attr=2,key,persist> */
	TelemetrySensorRowStatus int32
}

type destGroup struct {
	TelemetryDestGroupId        [63 + 1]byte
	TelemetryDestIp             uint32
	TelemetryDestPort           uint32
	TelemetryDestGroupRowStatus int32
}

type subGroup struct {
	TelemetrySubscriptionId    uint32
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
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 40000,
	})
	if err != nil {
		fmt.Println("listen failed. err:", err)
		return
	}
	defer listen.Close()
	serv := service.NewService(store.StoreClient)
	log.Info("start listen")
	for {
		var data [200]byte
		var sensorpathIndex int
		n, _, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read udp failed,err:", err)
			continue
		}

		//		fmt.Printf("data1:%v addr:%v count:%v\n", data[:n], addr, n)

		if data[0] == SensorGroupTableID {
			log.Info("SensorGroupTable")
			sersorGrp := &sensorGroup{}
			buf := bytes.NewReader(data[1:n])
			err = binary.Read(buf, binary.BigEndian, sersorGrp)
			if err != nil {
				log.Info("error")
				continue
			}
			//		buf = bytes.NewReader(data[85:n])
			//		binary.Read(buf, binary.BigEndian, &sersorGrp.TelemetrySensorRowStatus)
			log.Infof("%s %s \n", sersorGrp.TelemetrySensorGroupID[:], sersorGrp.TelemetrySensorPath[:])
			for index, ele := range sersorGrp.TelemetrySensorPath {
				//				log.Infof("index %d %c", ele, index)
				if ele == 0 {
					//					log.Infof("index %d\n", index)
					sensorpathIndex = index
					break
				}

			}

			sensorGrpEnt := &store.SensorGroupEntry{
				SensorGroupId: string(sersorGrp.TelemetrySensorGroupID[:]),
				SensorPath:    string(sersorGrp.TelemetrySensorPath[:sensorpathIndex])}
			log.Infof("%d\n", sersorGrp.TelemetrySensorRowStatus)
			if sersorGrp.TelemetrySensorRowStatus == 0 {
				err := serv.SensorGroup().Create(sensorGrpEnt)
				if err != nil {
					fmt.Printf("%+v", err)
				}
			} else if sersorGrp.TelemetrySensorRowStatus == 4 {
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
			} else if destGrp.TelemetryDestGroupRowStatus == 4 {
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
			log.Infof("subGrp:%d %s %s %d\n", subGrp.TelemetrySubscriptionId, subGrp.TelemetrySensorGroupId[:],
				subGrp.TelemetryDestGroupId[:], subGrp.TelemetryInterva)
			log.Infof("%d\n", subGrp.TelemetrySubGroupRowStatus)
			subGrpEnt := &store.SubscriptionGroupEntry{
				SubscriptionId: subGrp.TelemetrySubscriptionId,
				SensorGroupId:  string(subGrp.TelemetrySensorGroupId[:]),
				DestGroupId:    string(subGrp.TelemetryDestGroupId[:]),
				Interval:       subGrp.TelemetryInterva}
			if subGrp.TelemetrySubGroupRowStatus == 0 {
				err := serv.SubGroup().Create(subGrpEnt)
				if err != nil {
					fmt.Printf("%+v", err)
				}
			} else if subGrp.TelemetrySubGroupRowStatus == 4 {
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
func main() {
	//	log.Info("1111")
	/* 		go collect.ConnectToRedis()
	time.Sleep(time.Second)

	//接收上报的信息

	go collect.StartCollectDataLoop()
	go server.TelemetryDataTimer() */
	//receiveSnmpMsg()
	for {
		time.Sleep(time.Second)
	}
}
