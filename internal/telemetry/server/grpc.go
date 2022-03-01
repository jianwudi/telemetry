package server

import (
	"context"
	"encoding/json"
	"telemetry/internal/telemetry/gnmi"
	"telemetry/internal/telemetry/store"
	"time"

	"github.com/marmotedu/iam/pkg/log"
)

func TelemetryDataTimer() {
	ticker := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			TelemetryDataPublish()
		}
	}
}
func TelemetryDataPublish() {
	//	log.Info("publish grpc data")
	if grpcdata, err := store.StoreClient.GrpcDataGroup().GetFirstRecord(); err == nil {
		grpcdata.Count++
		//		log.Infof("%d", grpcdata.Count)
		if 500*grpcdata.Count >= grpcdata.Interval {
			runPublish(grpcdata)
			grpcdata.Count = 0
		}
		if grpcdata, err = store.StoreClient.GrpcDataGroup().GetNextRecord(grpcdata); err == nil {
			grpcdata.Count++
			if 500*grpcdata.Count >= grpcdata.Interval {
				runPublish(grpcdata)
				grpcdata.Count = 0
			}
		}
	}
}
func clearData(data *store.TelemetryData) {
	data.RxPackage = 0
	data.TxPackage = 0
	data.RxByte = 0
	data.TxByte = 0
	data.RxDropShort = 0
	data.RxDropLong = 0
	data.RxCrcError = 0
	data.RxPackageDrop = 0
}

type TelemetryOltTransceiverData struct {
	Voltage          float64
	Temperature      float64
	GponLaunchPower  float64
	XgponLaunchPower float64
	GponCurrent      float64
	XgponCurrent     float64
}

type TelemetryOnuTransceiverData struct {
	Voltage         float64
	Temperature     float64
	OnuLaunchPower  float64
	OnuReceivePower float64
	OnuCurrent      float64
}
type TelemetryOnuLocalInfo struct {
	SerialNumber string
	VendorID     string
}
type TelemetryOnuRemoteInfo struct {
	OnuUpTime string
	VendorID  string
}
type TelemetryTrafficData struct {
	RxPackage     uint64
	TxPackage     uint64
	RxByte        uint64
	TxByte        uint64
	RxDropShort   uint64
	RxDropLong    uint64
	RxCrcError    uint64
	RxPackageDrop uint64
}

func runPublish(grpcDataEntry *store.GrpcDataEntry) {
	oltTransceiverData := TelemetryOltTransceiverData{}
	oltTrafficData := TelemetryTrafficData{}
	onuTrafficData := TelemetryOnuTransceiverData{}
	onuLocalInfo := TelemetryOnuLocalInfo{}
	onuRemoteInfo := TelemetryOnuRemoteInfo{}
	for _, client := range grpcDataEntry.Client {
		store.TelemetryDataMutex.Lock()
		for index, telData := range grpcDataEntry.Data {
			var responses []*gnmi.SubscribeResponse
			var notif gnmi.Notification
			notif.Timestamp = time.Now().Unix()

			update := new(gnmi.Update)
			path := new(gnmi.Path)
			path.Target = grpcDataEntry.SensorPath[index]
			update.Path = path
			log.Infof("path:%s", grpcDataEntry.SensorPath[index][:len("an-gpon-pm-olt-transceivers")])
			if grpcDataEntry.SensorPath[index][:len("an-gpon-pm-olt-transceivers")] == "an-gpon-pm-olt-transceivers" {
				oltTransceiverData.Voltage = telData.Voltage
				oltTransceiverData.Temperature = telData.Temperature
				oltTransceiverData.GponLaunchPower = telData.GponLaunchPower
				oltTransceiverData.XgponLaunchPower = telData.XgponLaunchPower
				oltTransceiverData.GponCurrent = telData.GponCurrent
				oltTransceiverData.XgponCurrent = telData.XgponCurrent
				telDataJ, _ := json.Marshal(oltTransceiverData)
				val := &gnmi.TypedValue{
					Value: &gnmi.TypedValue_AsciiVal{AsciiVal: string(telDataJ)},
				}
				update.Val = val
			} else if grpcDataEntry.SensorPath[index][:len("an-gpon-onu-transceivers")] == "an-gpon-onu-transceivers" {
				onuTrafficData.Voltage = telData.Voltage
				onuTrafficData.Temperature = telData.Temperature
				onuTrafficData.OnuLaunchPower = telData.OnuLaunchPower
				onuTrafficData.OnuReceivePower = telData.OnuReceivePower
				onuTrafficData.OnuCurrent = telData.OnuCurrent
				telDataJ, _ := json.Marshal(onuTrafficData)
				val := &gnmi.TypedValue{
					Value: &gnmi.TypedValue_AsciiVal{AsciiVal: string(telDataJ)},
				}
				update.Val = val

			} else if grpcDataEntry.SensorPath[index][:len("an-gpon-pm-onu-local-info")] == "an-gpon-pm-onu-local-info" {
				onuLocalInfo.SerialNumber = telData.SerialNumber
				telDataJ, _ := json.Marshal(onuLocalInfo)
				val := &gnmi.TypedValue{
					Value: &gnmi.TypedValue_AsciiVal{AsciiVal: string(telDataJ)},
				}
				update.Val = val
			} else if grpcDataEntry.SensorPath[index][:len("an-gpon-pm-onu-remote-info")] == "an-gpon-pm-onu-remote-info" {
				onuRemoteInfo.OnuUpTime = telData.OnuUpTime
				onuRemoteInfo.VendorID = telData.VendorID
				telDataJ, _ := json.Marshal(onuRemoteInfo)
				val := &gnmi.TypedValue{
					Value: &gnmi.TypedValue_AsciiVal{AsciiVal: string(telDataJ)},
				}
				update.Val = val
			} else {
				oltTrafficData.RxPackage = telData.RxPackage
				oltTrafficData.TxPackage = telData.TxPackage
				oltTrafficData.RxByte = telData.RxByte
				oltTrafficData.TxByte = telData.TxByte
				oltTrafficData.RxDropShort = telData.RxDropShort
				oltTrafficData.RxDropLong = telData.RxDropLong
				oltTrafficData.RxCrcError = telData.RxCrcError
				oltTrafficData.RxPackageDrop = telData.RxPackageDrop
				clearData(telData)
				telDataJ, _ := json.Marshal(oltTrafficData)
				val := &gnmi.TypedValue{
					Value: &gnmi.TypedValue_AsciiVal{AsciiVal: string(telDataJ)},
				}
				update.Val = val
			}

			//			log.Infof("%p\n", telData)

			notif.Update = append(notif.Update, update)
			res := &gnmi.SubscribeResponse{
				Response: &gnmi.SubscribeResponse_Update{Update: &notif},
			}

			responses = append(responses, res)
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()
			stream, err := client.Publish(ctx)
			if err != nil {
				log.Fatalf("%v.RecordRoute(_) = _, %v", client, err)
			}

			for _, res := range responses {
				if err := stream.Send(res); err != nil {
					log.Fatalf("%v.Send(%v) = %v", stream, res, err)
				}
			}
			_, err = stream.CloseAndRecv()
			if err != nil {
				log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
			}
			//			log.Infof("Route summary: %v", reply)
		}
		store.TelemetryDataMutex.Unlock()

	}
}
