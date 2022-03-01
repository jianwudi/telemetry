package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"telemetry/internal/telemetry/collect"
	"telemetry/internal/telemetry/service"
	"telemetry/internal/telemetry/store"
	"time"
)

type SensorGroup struct {
	TelemetrySensorGroupID   [40 + 1]byte /* <attr=1,key,persist> */
	TelemetrySensorPath      [40 + 1]byte /* <attr=2,key,persist> */
	TelemetrySensorRowStatus int32
}

const (
	SensorGroupTable = iota + 1
	DestGroupTable
	SubscriptionGroupTable
)

/* func startPubSubLoop() {
	rediscli := storage.RedisCluster{} //订阅所有频道
	for {
		err := cacheStore.StartPubSubHandler(RedisPubSubChannel, func(v interface{}) {
			handleRedisEvent(v, nil, nil)
		})
		if err != nil {
			if !errors.Is(err, storage.ErrRedisIsDown) {
				log.Errorf("Connection to Redis failed, reconnect in 10s: %s", err.Error())
			}

			time.Sleep(10 * time.Second)
			log.Warnf("Reconnecting: %s", err.Error())
		}
	}
} */

/* func main() {
	go channel.ConnectToRedis()
	time.Sleep(time.Second)
	go channel.StartPubSubLoop()
	sensorgrp := &store.SensorGroupEntry{"sensorgroupid1", "traffic/nni/slot=1/port=1"}
	sensorgrpsvc := service.NewSensorGroupService(store.StoreClient)
	err := sensorgrpsvc.Create(sensorgrp)
	if err != nil {
		fmt.Printf("%+v", err)
	}

	destgrp := &store.DestGroupEntry{DestGroupId: "destgroupid1", DestIp: 0x7f000001, DestPort: 50051}
	err = store.StoreClient.DestGroup().CreateRecord(destgrp)
	if err != nil {
		fmt.Printf("%+v", err)
	}

	subgrp := &store.SubscriptionGroupEntry{
		SubscriptionId: 1,
		SensorGroupId:  "sensorgroupid1",
		DestGroupId:    "destgroupid1",
	}
	s := service.NewsubscriptionGroupService(store.StoreClient)

	//	err = store.StoreClient.SubscriptionGroup().CreateRecord(subgrp)
	err = s.Create(subgrp)
	if err != nil {
		fmt.Printf("%+v", err)
	}

	subgrp, err = store.StoreClient.SubscriptionGroup().GetFirstRecord()
	if err == nil {
		fmt.Printf("%s %s \n", subgrp.SensorGroupId, sensorgrp.SensorGroupId)
	}
	server.TelemetryDataTimer()
} */
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
	fmt.Println("start listen")
	for {
		var data [100]byte
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read udp failed,err:", err)
			continue
		}

		fmt.Printf("data:%v addr:%v count:%v\n", data[:n], addr, n)

		if data[0] == SensorGroupTable {
			fmt.Println("SensorGroupTable")
			sersorGrp := &SensorGroup{}
			//			fmt.Println(unsafe.Sizeof(p))
			buf := bytes.NewReader(data[1:n])
			err = binary.Read(buf, binary.BigEndian, sersorGrp)
			if err != nil {
				continue
			}
			buf = bytes.NewReader(data[85:n])
			binary.Read(buf, binary.BigEndian, &sersorGrp.TelemetrySensorRowStatus)
			serv := service.NewService(store.StoreClient)
			fmt.Println(sersorGrp.TelemetrySensorGroupID, sersorGrp.TelemetrySensorPath)
			sensorgrp := &store.SensorGroupEntry{string(sersorGrp.TelemetrySensorGroupID[:]), string(sersorGrp.TelemetrySensorPath[:])}
			fmt.Printf("%d\n", sersorGrp.TelemetrySensorRowStatus)
			if sersorGrp.TelemetrySensorRowStatus == 1 {
				err := serv.SensorGroup().Create(sensorgrp)
				if err != nil {
					fmt.Printf("%+v", err)
				}
			} else {
				err := serv.SensorGroup().Delete(sensorgrp)
				if err != nil {
					fmt.Printf("%+v", err)
				}
			}
		}
	}
}
func main() {
	//连接redis
	go collect.ConnectToRedis()
	time.Sleep(time.Second)

	//接收上报的信息
	go collect.StartCollectDataLoop()
	receiveSnmpMsg()
	//	service.NewService(store.StoreClient)
	//上报grpc server数据
	//	server.TelemetryDataTimer()
}
