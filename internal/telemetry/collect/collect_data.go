package collect

import (
	"encoding/json"
	"errors"
	"fmt"
	"telemetry/internal/telemetry/store"
	"time"

	redis "github.com/go-redis/redis/v7"
	"github.com/marmotedu/iam/pkg/log"
)

const (
	RedisSubChannel = "data.notifications"
)

type SensorpathData struct {
	SensorPath       string  `json:"sensorpath"`
	RxPackage        uint64  `json:"rx_package"`
	TxPackage        uint64  `json:"tx_package"`
	RxByte           uint64  `json:"rx_byte"`
	TxByte           uint64  `json:"tx_byte"`
	RxDropShort      uint64  `json:"rx_dropped_too_short"`
	RxDropLong       uint64  `json:"rx_dropped_too_long"`
	RxCrcError       uint64  `json:"rx_crc_error"`
	RxPackageDrop    uint64  `json:"rx_packets_dropped"`
	Voltage          float64 `json:"voltage"`
	Temperature      float64 `json:"temperature"`
	GponLaunchPower  float64 `json:"gponlLaunchPower"`
	XgponLaunchPower float64 `json:"xgponLaunchPower"`
	GponCurrent      float64 `json:"gponCurrent"`
	XgponCurrent     float64 `json:"xgponCurrent"`
	OnuLaunchPower   float64 `json:"onulLaunchPower"`
	OnuReceivePower  float64 `json:"onuReceivePower"`
	OnuCurrent       float64 `json:"onuCurrent"`
	SerialNumber     string  `json:"serialNumber"`
	VendorID         string  `json:"vendorID"`
	OnuUpTime        string  `json:"onuUpTime"`
}
type telemetryOltTrafficsData struct {
	Slot                           uint8  `json:"slot"`
	Link                           uint8  `json:"link"`
	Port_tx_bytes                  uint64 `json:"port_tx_bytes"`
	Port_rx_bytes                  uint64 `json:"port_rx_bytes"`
	Port_tx_pkt                    uint64 `json:"port_tx_pkt"`
	Port_rx_pkt                    uint64 `json:"port_rx_pkt"`
	Port_tx_discard_pkt            uint64 `json:"port_tx_discard_pkt"`
	Port_rx_discard_pkt            uint64 `json:"port_rx_discard_pkt"`
	Port_rx_crc_error_pkt          uint64 `json:"port_rx_crc_error_pkt"`
	Port_rx_oversized_discard_pkt  uint64 `json:"port_rx_oversized_discard_pkt"`
	Port_rx_undersized_discard_pkt uint64 `json:"port_rx_undersized_discard_pkt"`
	Port_rx_error_pkt              uint64 `json:"port_rx_error_pkt"`
	Port_tx_rate                   uint32 `json:"port_tx_rate"`
	Port_rx_rate                   uint32 `json:"port_rx_rate"`
	Port_tx_peak_rate              uint32 `json:"port_tx_peak_rate"`
	Port_rx_peak_rate              uint32 `json:"port_rx_peak_rate"`
}

type telemetryOltChannelTrafficsData struct {
	Slot                           uint8  `json:"slot"`
	Link                           uint8  `json:"link"`
	Channel                        uint8  `json:"channel"`
	Port_tx_bytes                  uint64 `json:"port_tx_bytes"`
	Port_rx_bytes                  uint64 `json:"port_rx_bytes"`
	Port_tx_pkt                    uint64 `json:"port_tx_pkt"`
	Port_rx_pkt                    uint64 `json:"port_rx_pkt"`
	Port_tx_discard_pkt            uint64 `json:"port_tx_discard_pkt"`
	Port_rx_discard_pkt            uint64 `json:"port_rx_discard_pkt"`
	Port_rx_crc_error_pkt          uint64 `json:"port_rx_crc_error_pkt"`
	Port_rx_oversized_discard_pkt  uint64 `json:"port_rx_oversized_discard_pkt"`
	Port_rx_undersized_discard_pkt uint64 `json:"port_rx_undersized_discard_pkt"`
	Port_rx_error_pkt              uint64 `json:"port_rx_error_pkt"`
	Port_tx_rate                   uint32 `json:"port_tx_rate"`
	Port_rx_rate                   uint32 `json:"port_rx_rate"`
	Port_tx_peak_rate              uint32 `json:"port_tx_peak_rate"`
	Port_rx_peak_rate              uint32 `json:"port_rx_peak_rate"`
}
type telemetryEthernetPortKpiRecordData struct {
	Link                            uint8  `json:"link"`
	Port_tx_bytes                   uint64 `json:"port_tx_bytes"`
	Port_rx_bytes                   uint64 `json:"port_rx_bytes"`
	Port_tx_packets                 uint64 `json:"port_tx_packets"`
	Port_rx_packets                 uint64 `json:"port_rx_packets"`
	Port_tx_discard_packets         uint64 `json:"port_tx_discard_packets"`
	Port_rx_discard_packets         uint64 `json:"port_rx_discard_packets"`
	Port_rx_alignment_error_packets uint64 `json:"port_rx_alignment_error_packets"`
	Port_tx_crc_error_packets       uint64 `json:"port_tx_crc_error_packets"`
	Port_rx_crc_error_packets       uint64 `json:"port_rx_crc_error_packets"`
	Port_tx_oversized_packets       uint64 `json:"port_tx_oversized_packets"`
	Port_rx_oversized_packets       uint64 `json:"port_rx_oversized_packets"`
	Port_tx_undersized_packets      uint64 `json:"port_tx_undersized_packets"`
	Port_rx_undersized_packets      uint64 `json:"port_rx_undersized_packets"`
	Port_tx_fragment_packets        uint64 `json:"port_tx_fragment_packets"`
	Port_rx_fragment_packets        uint64 `json:"port_rx_fragment_packets"`
	Port_tx_jabber_packets          uint64 `json:"port_tx_jabber_packets"`
	Port_rx_jabber_packets          uint64 `json:"port_rx_jabber_packets"`
	Port_tx_error_packets           uint64 `json:"port_tx_error_packets"`
	Port_rx_error_packets           uint64 `json:"port_rx_error_packets"`
	Port_tx_rate                    uint64 `json:"port_tx_rate"`
	Port_rx_rate                    uint64 `json:"port_rx_rate"`
	Port_tx_peak_rate               uint64 `json:"port_tx_peak_rate"`
	Port_rx_peak_rate               uint64 `json:"port_rx_peak_rate"`
}

type telemetryServiceFlowKpiRecordData struct {
	VlanId                        uint32 `json:"vlanid"`
	DownstreamQueueDropCnt        uint64 `json:"downstream_queue_drop_cnt"`
	DownstreamQueuePassCnt        uint64 `json:"downstream_queue_pass_cnt"`
	DownstreamQueueDropMax        uint32 `json:"downstream_queue_drop_max"`
	DownstreamQueueDropMin        uint32 `json:"downstream_queue_drop_min"`
	DownstreamQueueDropRateMax    uint32 `json:"downstream_queue_drop_rate_max"`
	DownstreamQueueDropRateMin    uint32 `json:"downstream_queue_drop_rate_min"`
	DownstreamQueueDropSecondsCnt uint32 `json:"downstream_queue_drop_seconds_cnt"`
	DownstreamQueuePassBytes      uint64 `json:"downstream_queue_pass_bytes"`
	DownstreamMfrAvg              uint32 `json:"downstream_mfr_avg"`
	UpstreamPassBytes             uint64 `json:"upstream_pass_bytes"`
	UpstreamPassCnt               uint64 `json:"upstream_pass_cnt"`
	UpstreamDropCnt               uint64 `json:"upstream_drop_cnt"`
}
type SensorPathType struct {
	SensorPath int `json:"sensorpath"`
}

func StartCollectDataLoop() {
	client := RedisCluster{}
	// On message, synchronize
	for {
		err := client.StartSubHandler(RedisSubChannel, func(v interface{}) {
			handleRedisEvent(v)
		})
		if err != nil {
			if !errors.Is(err, ErrRedisIsDown) {
				log.Errorf("Connection to Redis failed, reconnect in 10s: %s", err.Error())
			}

			time.Sleep(10 * time.Second)
			log.Warnf("Reconnecting: %s", err.Error())
		}
	}
}

func handleRedisEvent(v interface{}) {
	message, ok := v.(*redis.Message)
	if !ok {
		return
	}
	//	log.Info(message.Payload)
	sensorpath := SensorPathType{}
	if err := json.Unmarshal([]byte(message.Payload), &sensorpath); err != nil {
		log.Errorf("Unmarshalling message body failed, malformed: %v", err)
		return
	}
	if sensorpath.SensorPath == 1 {
		teleData := telemetryOltTrafficsData{}
		if err := json.Unmarshal([]byte(message.Payload), &teleData); err != nil {
			log.Errorf("Unmarshalling message body failed, malformed: %v", err)
			return
		}
		//		log.Infof("%d %d", teleData.Port_tx_bytes, teleData.Port_rx_bytes)
		load(teleData)
	} else if sensorpath.SensorPath == 2 {
		teleData := telemetryOltChannelTrafficsData{}
		if err := json.Unmarshal([]byte(message.Payload), &teleData); err != nil {
			log.Errorf("Unmarshalling message body failed, malformed: %v", err)
			return
		}
		//		log.Infof("%d %d", teleData.Port_tx_bytes, teleData.Port_rx_bytes)
		load(teleData)
	} else if sensorpath.SensorPath == 3 {
		teleData := telemetryEthernetPortKpiRecordData{}
		if err := json.Unmarshal([]byte(message.Payload), &teleData); err != nil {
			log.Errorf("Unmarshalling message body failed, malformed: %v", err)
			return
		}
		//		log.Infof("%d %d", teleData.Port_tx_bytes, teleData.Port_rx_bytes)
		load(teleData)
	} else if sensorpath.SensorPath == 4 {
		teleData := telemetryServiceFlowKpiRecordData{}
		if err := json.Unmarshal([]byte(message.Payload), &teleData); err != nil {
			log.Errorf("Unmarshalling message body failed, malformed: %v", err)
			return
		}
		//		log.Infof("%d %d", teleData.Port_tx_bytes, teleData.Port_rx_bytes)
		load(teleData)
	}
}

//接收上报的信息
func load(data interface{}) {
	//找到sensordata表
	switch data := data.(type) {
	case telemetryOltTrafficsData:
		sensordata := &store.SensorDataEntry{}
		sensordata.SensorPath = fmt.Sprintf("an_gpon_pm_olt_traffic:GponPmOltTraffics.%d.%d", data.Slot, data.Link)
		store.TelemetryDataMutex.Lock()
		sensordatastore, err := store.StoreClient.SensorData().GetRecord(sensordata)
		if err != nil {
			log.Errorf("load fail: %s", err.Error())
			return
		}
		for _, storedata := range sensordatastore.Data {
			//		log.Infof("data: %p", storedata)
			teledata, ok := storedata.(*store.OltTrafficsRecordData)
			if !ok {
				log.Errorf("type fail")
				store.TelemetryDataMutex.Unlock()
				return
			}
			teledata.Slot = data.Slot
			teledata.Link = data.Link
			teledata.Port_tx_bytes += data.Port_tx_bytes
			teledata.Port_rx_bytes += data.Port_rx_bytes
			teledata.Port_tx_pkt += data.Port_tx_pkt
			teledata.Port_rx_pkt += data.Port_rx_pkt
			teledata.Port_tx_discard_pkt += data.Port_tx_discard_pkt
			teledata.Port_rx_discard_pkt += data.Port_rx_discard_pkt
			teledata.Port_rx_crc_error_pkt += data.Port_rx_crc_error_pkt
			teledata.Port_rx_oversized_discard_pkt += data.Port_rx_oversized_discard_pkt
			teledata.Port_rx_undersized_discard_pkt += data.Port_rx_undersized_discard_pkt
			teledata.Port_rx_error_pkt += data.Port_rx_error_pkt
			teledata.Port_tx_rate = data.Port_tx_rate
			teledata.Port_rx_rate = data.Port_rx_rate
			teledata.Port_tx_peak_rate = data.Port_tx_peak_rate
			teledata.Port_rx_peak_rate = data.Port_rx_peak_rate
		}
		store.TelemetryDataMutex.Unlock()
	case telemetryOltChannelTrafficsData:
		sensordata := &store.SensorDataEntry{}
		sensordata.SensorPath = fmt.Sprintf("an_gpon_pm_olt_traffic:GponPmOltChannelTraffics.%d.%d.%d", data.Slot, data.Link, data.Channel)
		store.TelemetryDataMutex.Lock()
		sensordatastore, err := store.StoreClient.SensorData().GetRecord(sensordata)
		if err != nil {
			log.Errorf("load fail: %s", err.Error())
			return
		}
		for _, storedata := range sensordatastore.Data {
			teledata, ok := storedata.(*store.OltChannelTrafficsRecordData)
			if !ok {
				log.Errorf("type fail")
				store.TelemetryDataMutex.Unlock()
				return
			}
			teledata.Slot = data.Slot
			teledata.Link = data.Link
			teledata.Channel = data.Channel
			teledata.Port_tx_bytes += data.Port_tx_bytes
			teledata.Port_rx_bytes += data.Port_rx_bytes
			teledata.Port_tx_pkt += data.Port_tx_pkt
			teledata.Port_rx_pkt += data.Port_rx_pkt
			teledata.Port_tx_discard_pkt += data.Port_tx_discard_pkt
			teledata.Port_rx_discard_pkt += data.Port_rx_discard_pkt
			teledata.Port_rx_crc_error_pkt += data.Port_rx_crc_error_pkt
			teledata.Port_rx_oversized_discard_pkt += data.Port_rx_oversized_discard_pkt
			teledata.Port_rx_undersized_discard_pkt += data.Port_rx_undersized_discard_pkt
			teledata.Port_rx_error_pkt += data.Port_rx_error_pkt
			teledata.Port_tx_rate = data.Port_tx_rate
			teledata.Port_rx_rate = data.Port_rx_rate
			teledata.Port_tx_peak_rate = data.Port_tx_peak_rate
			teledata.Port_rx_peak_rate = data.Port_rx_peak_rate
		}
		store.TelemetryDataMutex.Unlock()
	case telemetryEthernetPortKpiRecordData:
		sensordata := &store.SensorDataEntry{}
		sensordata.SensorPath = fmt.Sprintf("an_ethernet_kpi:EthernetPortKpiRecords.%d.%d", 0, data.Link)
		store.TelemetryDataMutex.Lock()
		sensordatastore, err := store.StoreClient.SensorData().GetRecord(sensordata)
		if err != nil {
			log.Errorf("load fail: %s", err.Error())
			return
		}
		for _, storedata := range sensordatastore.Data {
			teledata, ok := storedata.(*store.EthernetPortKpiRecordData)
			if !ok {
				log.Errorf("type fail")
				store.TelemetryDataMutex.Unlock()
				return
			}
			teledata.Link = data.Link
			teledata.Port_tx_bytes = data.Port_tx_bytes
			teledata.Port_rx_bytes += data.Port_rx_bytes
			teledata.Port_tx_packets += data.Port_tx_packets
			teledata.Port_rx_packets += data.Port_rx_packets
			teledata.Port_tx_discard_packets += data.Port_tx_discard_packets
			teledata.Port_rx_discard_packets += data.Port_rx_discard_packets
			teledata.Port_rx_alignment_error_packets += data.Port_rx_alignment_error_packets
			teledata.Port_tx_crc_error_packets += data.Port_tx_crc_error_packets
			teledata.Port_rx_crc_error_packets += data.Port_rx_crc_error_packets
			teledata.Port_tx_oversized_packets += data.Port_tx_oversized_packets
			teledata.Port_rx_oversized_packets += data.Port_rx_oversized_packets
			teledata.Port_tx_undersized_packets += data.Port_tx_undersized_packets
			teledata.Port_rx_undersized_packets += data.Port_rx_undersized_packets
			teledata.Port_tx_fragment_packets += data.Port_tx_fragment_packets
			teledata.Port_rx_fragment_packets += data.Port_rx_fragment_packets
			teledata.Port_tx_jabber_packets += data.Port_tx_jabber_packets
			teledata.Port_rx_jabber_packets += data.Port_rx_jabber_packets
			teledata.Port_tx_error_packets += data.Port_tx_error_packets
			teledata.Port_rx_error_packets += data.Port_rx_error_packets
			teledata.Port_tx_rate = data.Port_tx_rate
			teledata.Port_rx_rate = data.Port_rx_rate
			teledata.Port_tx_peak_rate = data.Port_tx_peak_rate
			teledata.Port_rx_peak_rate = data.Port_rx_peak_rate
		}
		store.TelemetryDataMutex.Unlock()
	case telemetryServiceFlowKpiRecordData:
		sensordata := &store.SensorDataEntry{}
		sensordata.SensorPath = fmt.Sprintf("an_bb_service_flow_kpi:ServiceFlowKpiRecords.%d", data.VlanId)
		store.TelemetryDataMutex.Lock()
		sensordatastore, err := store.StoreClient.SensorData().GetRecord(sensordata)
		if err != nil {
			log.Errorf("load fail: %s", err.Error())
			return
		}
		for _, storedata := range sensordatastore.Data {
			teledata, ok := storedata.(*store.ServiceFlowKpiRecordData)
			if !ok {
				log.Errorf("type fail")
				store.TelemetryDataMutex.Unlock()
				return
			}
			teledata.VlanId = data.VlanId
			teledata.DownstreamQueueDropCnt += data.DownstreamQueueDropCnt
			teledata.DownstreamQueuePassCnt += data.DownstreamQueuePassCnt
			teledata.DownstreamQueueDropMax = data.DownstreamQueueDropMax
			teledata.DownstreamQueueDropMin = data.DownstreamQueueDropMin
			teledata.DownstreamQueueDropRateMax = data.DownstreamQueueDropRateMax
			teledata.DownstreamQueueDropRateMin = data.DownstreamQueueDropRateMin
			teledata.DownstreamQueueDropSecondsCnt += data.DownstreamQueueDropSecondsCnt
			teledata.DownstreamQueuePassBytes += data.DownstreamQueuePassBytes
			teledata.DownstreamMfrAvg += data.DownstreamMfrAvg
			teledata.UpstreamPassBytes += data.UpstreamPassBytes
			teledata.UpstreamPassCnt = data.UpstreamPassCnt
			teledata.UpstreamDropCnt = data.UpstreamDropCnt
		}
		store.TelemetryDataMutex.Unlock()
	}
	/* 	log.Infof("%s\n", data.SensorPath)
	   	sensordata := &store.SensorDataEntry{}
	   	sensordata.SensorPath = data.SensorPath
	   	sensordatastore, err := store.StoreClient.SensorData().GetRecord(sensordata)
	   	if err != nil {
	   		log.Errorf("load fail: %s", err.Error())
	   		return
	   	}
	   	store.TelemetryDataMutex.Lock()
	   	for _, storedata := range sensordatastore.Data {
	   		storedata.RxPackage += data.RxPackage
	   		storedata.TxPackage += data.TxPackage
	   		storedata.RxByte += data.RxByte
	   		storedata.TxByte += data.TxByte
	   		storedata.RxDropShort += data.RxDropShort
	   		storedata.RxDropLong += data.RxDropLong
	   		storedata.RxCrcError += data.RxCrcError
	   		storedata.RxPackageDrop += data.RxPackageDrop
	   		storedata.Voltage = data.Voltage
	   		storedata.Temperature = data.Temperature
	   		storedata.GponLaunchPower = data.GponLaunchPower
	   		storedata.XgponLaunchPower = data.XgponLaunchPower
	   		storedata.GponCurrent = data.GponCurrent
	   		storedata.XgponCurrent = data.XgponCurrent
	   		storedata.OnuLaunchPower = data.OnuLaunchPower
	   		storedata.OnuReceivePower = data.OnuReceivePower
	   		storedata.OnuCurrent = data.OnuCurrent
	   		storedata.SerialNumber = data.SerialNumber
	   		storedata.VendorID = data.VendorID
	   		storedata.OnuUpTime = data.OnuUpTime
	   		//		log.Infof("%d\n", storedata.RxByte)
	   	}
	   	store.TelemetryDataMutex.Unlock() */
}
