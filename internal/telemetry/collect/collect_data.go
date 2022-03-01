package collect

import (
	"encoding/json"
	"errors"
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

	data := SensorpathData{}
	if err := json.Unmarshal([]byte(message.Payload), &data); err != nil {
		log.Errorf("Unmarshalling message body failed, malformed: %v", err)
		return
	}
	//	log.Infow("receive redis message", "command", data.Command, "payload", message.Payload)
	log.Infof("%s %d %d", data.SensorPath, data.RxByte, data.TxByte)
	load(&data)

}

//接收上报的信息
func load(data *SensorpathData) {
	//找到sensordata表
	log.Infof("%s\n", data.SensorPath)
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
	store.TelemetryDataMutex.Unlock()
}
