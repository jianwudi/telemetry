package server

import (
	"context"
	"fmt"
	"io"
	"strings"
	v1 "telemetry/internal/telemetry/gnmi/v1"
	"telemetry/internal/telemetry/store"
	"time"

	"github.com/marmotedu/iam/pkg/log"

	"google.golang.org/protobuf/proto"
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
	store.TelemetryDataMutex.Lock()
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
	store.TelemetryDataMutex.Unlock()

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

func clearGponPmOltTraffics(teledata *store.OltTrafficsRecordData) {
	teledata.Port_tx_bytes = 0
	teledata.Port_rx_bytes = 0
	teledata.Port_tx_pkt = 0
	teledata.Port_rx_pkt = 0
	teledata.Port_tx_discard_pkt = 0
	teledata.Port_rx_discard_pkt = 0
	teledata.Port_rx_crc_error_pkt = 0
	teledata.Port_rx_oversized_discard_pkt = 0
	teledata.Port_rx_undersized_discard_pkt = 0
	teledata.Port_rx_error_pkt = 0
	teledata.Port_tx_rate = 0
	teledata.Port_rx_rate = 0
	teledata.Port_tx_peak_rate = 0
	teledata.Port_rx_peak_rate = 0
}
func addGponPmOltTrafficsTeleData(teledata *v1.GponPmOltTraffics_GponPmOltTraffic, data *store.OltTrafficsRecordData) {
	teledata.Name = fmt.Sprintf("gpon.f.%d.%d", data.Slot, data.Link)
	teledata.PortTxBytes = data.Port_tx_bytes
	teledata.PortRxBytes = data.Port_rx_bytes
	teledata.PortTxPkt = data.Port_tx_pkt
	teledata.PortRxPkt = data.Port_rx_pkt
	teledata.PortTxDiscardPkt = data.Port_tx_discard_pkt
	teledata.PortRxDiscardPkt = data.Port_rx_discard_pkt
	teledata.PortRxCrcErrorPkt = data.Port_rx_crc_error_pkt
	teledata.PortRxOversizedDiscardPkt = data.Port_rx_oversized_discard_pkt
	teledata.PortRxUndersizedDiscardPkt = data.Port_rx_undersized_discard_pkt
	teledata.PortRxErrorPkt = data.Port_rx_error_pkt
	teledata.PortTxRate = data.Port_tx_rate
	teledata.PortRxRate = data.Port_rx_rate
	teledata.PortTxPeakRate = data.Port_tx_peak_rate
	teledata.PortRxPeakRate = data.Port_rx_peak_rate
}
func clearEthernetPortKpiRecordTeleData(teledata *store.EthernetPortKpiRecordData) {
	teledata.Port_tx_bytes = 0
	teledata.Port_rx_bytes = 0
	teledata.Port_tx_packets = 0
	teledata.Port_rx_packets = 0
	teledata.Port_tx_discard_packets = 0
	teledata.Port_rx_discard_packets = 0
	teledata.Port_rx_alignment_error_packets = 0
	teledata.Port_tx_crc_error_packets = 0
	teledata.Port_rx_crc_error_packets = 0
	teledata.Port_tx_oversized_packets = 0
	teledata.Port_rx_oversized_packets = 0
	teledata.Port_tx_undersized_packets = 0
	teledata.Port_rx_undersized_packets = 0
	teledata.Port_tx_fragment_packets = 0
	teledata.Port_rx_fragment_packets = 0
	teledata.Port_tx_jabber_packets = 0
	teledata.Port_rx_jabber_packets = 0
	teledata.Port_tx_error_packets = 0
	teledata.Port_rx_error_packets = 0
	teledata.Port_tx_rate = 0
	teledata.Port_rx_rate = 0
	teledata.Port_tx_peak_rate = 0
	teledata.Port_rx_peak_rate = 0
}
func addEthernetPortKpiRecordTeleData(teledata *v1.EthernetPortKpiRecords_EthernetPortKpiRecord, data *store.EthernetPortKpiRecordData) {
	teledata.Name = fmt.Sprintf("ethernetCsmacd.f.%d.%d", 1, data.Link)
	teledata.PortTxBytes = data.Port_tx_bytes
	teledata.PortRxBytes = data.Port_rx_bytes
	teledata.PortTxPackets = data.Port_tx_packets
	teledata.PortRxPackets = data.Port_rx_packets
	teledata.PortTxDiscardPackets = data.Port_tx_discard_packets
	teledata.PortRxDiscardPackets = data.Port_rx_discard_packets
	teledata.PortRxAlignmentErrorPackets = data.Port_rx_alignment_error_packets
	teledata.PortTxCrcErrorPackets = data.Port_tx_crc_error_packets
	teledata.PortRxCrcErrorPackets = data.Port_rx_crc_error_packets
	teledata.PortTxOversizedPackets = data.Port_tx_oversized_packets
	teledata.PortRxOversizedPackets = data.Port_rx_oversized_packets
	teledata.PortTxUndersizedPackets = data.Port_tx_undersized_packets
	teledata.PortRxUndersizedPackets = data.Port_rx_undersized_packets
	teledata.PortTxFragmentPackets = data.Port_tx_fragment_packets
	teledata.PortRxFragmentPackets = data.Port_rx_fragment_packets
	teledata.PortTxJabberPackets = data.Port_tx_jabber_packets
	teledata.PortRxJabberPackets = data.Port_rx_jabber_packets
	teledata.PortTxErrorPackets = data.Port_tx_error_packets
	teledata.PortRxErrorPackets = data.Port_rx_error_packets
	teledata.PortTxRate = data.Port_tx_rate
	teledata.PortRxRate = data.Port_rx_rate
	teledata.PortTxPeakRate = data.Port_tx_peak_rate
	teledata.PortRxPeakRate = data.Port_rx_peak_rate
}
func clearGponPmOltChannelTraffics(teledata *store.OltChannelTrafficsRecordData) {
	teledata.Port_tx_bytes = 0
	teledata.Port_rx_bytes = 0
	teledata.Port_tx_pkt = 0
	teledata.Port_rx_pkt = 0
	teledata.Port_tx_discard_pkt = 0
	teledata.Port_rx_discard_pkt = 0
	teledata.Port_rx_crc_error_pkt = 0
	teledata.Port_rx_oversized_discard_pkt = 0
	teledata.Port_rx_undersized_discard_pkt = 0
	teledata.Port_rx_error_pkt = 0
	teledata.Port_tx_rate = 0
	teledata.Port_rx_rate = 0
	teledata.Port_tx_peak_rate = 0
	teledata.Port_rx_peak_rate = 0
}
func addGponPmOltChannelTrafficsTeleData(teledata *v1.GponPmOltChannelTraffic_GponPmOltChannelTraffic, data *store.OltChannelTrafficsRecordData) {
	teledata.Name = fmt.Sprintf("gpon.f.%d.%d.%d", data.Slot, data.Link, data.Channel)
	teledata.PortTxBytes = data.Port_tx_bytes
	teledata.PortRxBytes = data.Port_rx_bytes
	teledata.PortTxPkt = data.Port_tx_pkt
	teledata.PortRxPkt = data.Port_rx_pkt
	teledata.PortTxDiscardPkt = data.Port_tx_discard_pkt
	teledata.PortRxDiscardPkt = data.Port_rx_discard_pkt
	teledata.PortRxCrcErrorPkt = data.Port_rx_crc_error_pkt
	teledata.PortRxOversizedDiscardPkt = data.Port_rx_oversized_discard_pkt
	teledata.PortRxUndersizedDiscardPkt = data.Port_rx_undersized_discard_pkt
	teledata.PortRxErrorPkt += data.Port_rx_error_pkt
	teledata.PortTxRate = data.Port_tx_rate
	teledata.PortRxRate = data.Port_rx_rate
	teledata.PortTxPeakRate = data.Port_tx_peak_rate
	teledata.PortRxPeakRate = data.Port_rx_peak_rate
}
func clearServiceFlowKpiRecordData(teledata *store.ServiceFlowKpiRecordData) {
	teledata.DownstreamQueueDropCnt = 0
	teledata.DownstreamQueuePassCnt = 0
	teledata.DownstreamQueueDropMax = 0
	teledata.DownstreamQueueDropMin = 0
	teledata.DownstreamQueueDropRateMax = 0
	teledata.DownstreamQueueDropRateMin = 0
	teledata.DownstreamQueueDropSecondsCnt = 0
	teledata.DownstreamQueuePassBytes = 0
	teledata.DownstreamMfrAvg = 0
	teledata.UpstreamPassBytes = 0
	teledata.UpstreamPassCnt = 0
	teledata.UpstreamDropCnt = 0
}
func addServiceFlowKpiRecordData(teledata *v1.ServiceFlowKpiRecords_ServiceFlowKpiRecord, data *store.ServiceFlowKpiRecordData) {
	teledata.Name = data.VlanId
	teledata.DownstreamQueueDropCnt = data.DownstreamQueueDropCnt
	teledata.DownstreamQueuePassCnt = data.DownstreamQueuePassCnt
	teledata.DownstreamQueueDropMax = data.DownstreamQueueDropMax
	teledata.DownstreamQueueDropMin = data.DownstreamQueueDropMin
	teledata.DownstreamQueueDropRateMax = data.DownstreamQueueDropRateMax
	teledata.DownstreamQueueDropRateMin = data.DownstreamQueueDropRateMin
	teledata.DownstreamQueueDropSecondsCnt = data.DownstreamQueueDropSecondsCnt
	teledata.DownstreamQueuePassBytes = data.DownstreamQueuePassBytes
	teledata.DownstreamMfrAvg = data.DownstreamMfrAvg
	teledata.UpstreamPassBytes = data.UpstreamPassBytes
	teledata.UpstreamPassCnt = data.UpstreamPassCnt
	teledata.UpstreamDropCnt = data.UpstreamDropCnt
}
func marshalTeleGrpcData(m proto.Message, curSensorPath string) *v1.ServiceArgs {
	buf2, err := proto.Marshal(m)
	if err != nil {
		log.Errorf("Marshal fail")
		return nil
	}
	var row v1.TelemetryRowGPB
	row.Content = buf2
	var dataGpb v1.TelemetryGPBTable
	dataGpb.Row = append(dataGpb.Row, &row)
	buf, err := proto.Marshal(
		&v1.Telemetry{
			SensorPath: curSensorPath,
			DataGpb:    &dataGpb,
		})
	if err != nil {
		log.Errorf("Marshal fail")
		return nil
	}
	return &v1.ServiceArgs{
		Data: buf, ReqId: 1,
	}
}
func runPublish(grpcDataEntry *store.GrpcDataEntry) {
	var args []*v1.ServiceArgs
	var curSensorPath string
	var pmOltTraffic v1.GponPmOltTraffics
	var pmOltChannelTraffic v1.GponPmOltChannelTraffic
	var ethernetPortKpiRecords v1.EthernetPortKpiRecords
	var serviceFlowKpiRecords v1.ServiceFlowKpiRecords
	//	store.TelemetryDataMutex.Lock()
	for i, sensorpath := range grpcDataEntry.SensorPath {
		p1 := strings.Split(sensorpath, ".")
		if strings.Compare("an_gpon_pm_olt_traffic:GponPmOltTraffics", p1[0]) == 0 {
			data, ok := grpcDataEntry.Data[i].(*store.OltTrafficsRecordData)
			if !ok {
				log.Errorf("type fail")
				return
			}
			if strings.Compare(curSensorPath, p1[0]) == 0 {
				var teledata v1.GponPmOltTraffics_GponPmOltTraffic
				addGponPmOltTrafficsTeleData(&teledata, data)
				clearGponPmOltTraffics(data)
				pmOltTraffic.PmOltTraffic = append(pmOltTraffic.PmOltTraffic, &teledata)
			} else {
				//是别的if跳进来的 所以如果不是第一次就得加入
				if i != 0 {
					//通过curSensorPath判断要将什么变量Marshal
					if curSensorPath == "an_gpon_pm_olt_traffic:GponPmOltChannelTraffics" {
						arg := marshalTeleGrpcData(&pmOltChannelTraffic, curSensorPath)
						if arg != nil {
							args = append(args, arg)
						}
					} else if curSensorPath == "an_ethernet_kpi:EthernetPortKpiRecords" {
						arg := marshalTeleGrpcData(&ethernetPortKpiRecords, curSensorPath)
						if arg != nil {
							args = append(args, arg)
						}
					} else if curSensorPath == "an_bb_service_flow_kpi:ServiceFlowKpiRecords" {
						arg := marshalTeleGrpcData(&serviceFlowKpiRecords, curSensorPath)
						if arg != nil {
							args = append(args, arg)
						}
					}
				}
				//将上一次的搞成buf 加入
				var teledata v1.GponPmOltTraffics_GponPmOltTraffic
				addGponPmOltTrafficsTeleData(&teledata, data)
				clearGponPmOltTraffics(data)
				pmOltTraffic.PmOltTraffic = append(pmOltTraffic.PmOltTraffic, &teledata)
				curSensorPath = "an_gpon_pm_olt_traffic:GponPmOltTraffics"
			}
		} else if strings.Compare("an_gpon_pm_olt_traffic:GponPmOltChannelTraffics", p1[0]) == 0 {
			data, ok := grpcDataEntry.Data[i].(*store.OltChannelTrafficsRecordData)
			if !ok {
				log.Errorf("type fail")
				return
			}
			if strings.Compare(curSensorPath, p1[0]) == 0 {
				var teledata v1.GponPmOltChannelTraffic_GponPmOltChannelTraffic
				addGponPmOltChannelTrafficsTeleData(&teledata, data)
				clearGponPmOltChannelTraffics(data)
				pmOltChannelTraffic.PmOltChannelTraffic = append(pmOltChannelTraffic.PmOltChannelTraffic, &teledata)
			} else {
				if i != 0 {
					//通过curSensorPath判断要将什么变量Marshal
					if curSensorPath == "an_gpon_pm_olt_traffic:GponPmOltTraffics" {
						arg := marshalTeleGrpcData(&pmOltTraffic, curSensorPath)
						if arg != nil {
							args = append(args, arg)
						}
					} else if curSensorPath == "an_ethernet_kpi:EthernetPortKpiRecords" {
						arg := marshalTeleGrpcData(&ethernetPortKpiRecords, curSensorPath)
						if arg != nil {
							args = append(args, arg)
						}
					} else if curSensorPath == "an_bb_service_flow_kpi:ServiceFlowKpiRecords" {
						arg := marshalTeleGrpcData(&serviceFlowKpiRecords, curSensorPath)
						if arg != nil {
							args = append(args, arg)
						}
					}

				}
				var teledata v1.GponPmOltChannelTraffic_GponPmOltChannelTraffic
				addGponPmOltChannelTrafficsTeleData(&teledata, data)
				clearGponPmOltChannelTraffics(data)
				pmOltChannelTraffic.PmOltChannelTraffic = append(pmOltChannelTraffic.PmOltChannelTraffic, &teledata)
				curSensorPath = "an_gpon_pm_olt_traffic:GponPmOltChannelTraffics"
			}
		} else if strings.Compare("an_ethernet_kpi:EthernetPortKpiRecords", p1[0]) == 0 {
			data, ok := grpcDataEntry.Data[i].(*store.EthernetPortKpiRecordData)
			if !ok {
				log.Errorf("type fail")
				return
			}
			fmt.Println(data)
			if strings.Compare(curSensorPath, p1[0]) == 0 {
				var teledata v1.EthernetPortKpiRecords_EthernetPortKpiRecord
				addEthernetPortKpiRecordTeleData(&teledata, data)
				clearEthernetPortKpiRecordTeleData(data)
				ethernetPortKpiRecords.EthernetPortKpiRecord = append(ethernetPortKpiRecords.EthernetPortKpiRecord, &teledata)
			} else {
				if i != 0 {
					if curSensorPath == "an_gpon_pm_olt_traffic:GponPmOltTraffics" {
						arg := marshalTeleGrpcData(&pmOltTraffic, curSensorPath)
						if arg != nil {
							args = append(args, arg)
						}
					} else if curSensorPath == "an_gpon_pm_olt_traffic:GponPmOltChannelTraffics" {
						arg := marshalTeleGrpcData(&pmOltChannelTraffic, curSensorPath)
						if arg != nil {
							args = append(args, arg)
						}
					} else if curSensorPath == "an_bb_service_flow_kpi:ServiceFlowKpiRecords" {
						arg := marshalTeleGrpcData(&serviceFlowKpiRecords, curSensorPath)
						if arg != nil {
							args = append(args, arg)
						}
					}
				}
				var teledata v1.EthernetPortKpiRecords_EthernetPortKpiRecord
				addEthernetPortKpiRecordTeleData(&teledata, data)
				clearEthernetPortKpiRecordTeleData(data)
				ethernetPortKpiRecords.EthernetPortKpiRecord = append(ethernetPortKpiRecords.EthernetPortKpiRecord, &teledata)
			}
		} else if strings.Compare("an_bb_service_flow_kpi:ServiceFlowKpiRecords", p1[0]) == 0 {
			data, ok := grpcDataEntry.Data[i].(*store.ServiceFlowKpiRecordData)
			if !ok {
				log.Errorf("type fail")
				return
			}
			fmt.Println(data)
			if strings.Compare(curSensorPath, p1[0]) == 0 {
				var teledata v1.ServiceFlowKpiRecords_ServiceFlowKpiRecord
				addServiceFlowKpiRecordData(&teledata, data)
				clearServiceFlowKpiRecordData(data)
				serviceFlowKpiRecords.ServiceFlowKpiRecord = append(serviceFlowKpiRecords.ServiceFlowKpiRecord, &teledata)
			} else {
				if i != 0 {
					//通过curSensorPath判断要将什么变量Marshal
					if curSensorPath == "an_gpon_pm_olt_traffic:GponPmOltTraffics" {
						arg := marshalTeleGrpcData(&pmOltTraffic, curSensorPath)
						if arg != nil {
							args = append(args, arg)
						}
					} else if curSensorPath == "an_gpon_pm_olt_traffic:GponPmOltChannelTraffics" {
						arg := marshalTeleGrpcData(&pmOltChannelTraffic, curSensorPath)
						if arg != nil {
							args = append(args, arg)
						}
					} else if curSensorPath == "an_ethernet_kpi:EthernetPortKpiRecords" {
						arg := marshalTeleGrpcData(&ethernetPortKpiRecords, curSensorPath)
						if arg != nil {
							args = append(args, arg)
						}
					}
				}
				var teledata v1.ServiceFlowKpiRecords_ServiceFlowKpiRecord
				addServiceFlowKpiRecordData(&teledata, data)
				clearServiceFlowKpiRecordData(data)
				serviceFlowKpiRecords.ServiceFlowKpiRecord = append(serviceFlowKpiRecords.ServiceFlowKpiRecord, &teledata)
			}
		}
	}
	//	store.TelemetryDataMutex.Unlock()
	//循环结束
	if curSensorPath == "an_gpon_pm_olt_traffic:GponPmOltTraffics" {
		arg := marshalTeleGrpcData(&pmOltTraffic, curSensorPath)
		if arg != nil {
			args = append(args, arg)
		}
	} else if curSensorPath == "an_ethernet_kpi:EthernetPortKpiRecords" {
		arg := marshalTeleGrpcData(&ethernetPortKpiRecords, curSensorPath)
		if arg != nil {
			args = append(args, arg)
		}
	} else if curSensorPath == "an_bb_service_flow_kpi:ServiceFlowKpiRecords" {
		arg := marshalTeleGrpcData(&serviceFlowKpiRecords, curSensorPath)
		if arg != nil {
			args = append(args, arg)
		}
	} else if curSensorPath == "an_gpon_pm_olt_traffic:GponPmOltChannelTraffics" {
		arg := marshalTeleGrpcData(&pmOltChannelTraffic, curSensorPath)
		if arg != nil {
			args = append(args, arg)
		}
	}

	//数据
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	for _, client := range grpcDataEntry.Client {
		stream, err := client.DataPublish(ctx)
		if err != nil {
			log.Fatalf("client.DataPublish failed: %v", err)
		}
		waitc := make(chan struct{})
		go func() {
			for {
				_, err := stream.Recv()
				if err == io.EOF {
					// read done.
					close(waitc)
					return
				}
				if err != nil {
					log.Fatalf("client.DataPublish failed: %v", err)
				}
				//		log.Printf("Got message %s at point(%d, %d)", in.Message, in.Location.Latitude, in.Location.Longitude)
			}
		}()
		for _, arg := range args {
			if err := stream.Send(arg); err != nil {
				log.Fatalf("client.DataPublish: stream.Send(%v) failed: %v", arg, err)
			}
		}
		stream.CloseSend()
		<-waitc
	}
}
