package store

import "sync"

type TelemetryData struct {
	SensorPathType   uint64
	RxPackage        uint64
	TxPackage        uint64
	RxByte           uint64
	TxByte           uint64
	RxDropShort      uint64
	RxDropLong       uint64
	RxCrcError       uint64
	RxPackageDrop    uint64
	Voltage          float64
	Temperature      float64
	GponLaunchPower  float64
	XgponLaunchPower float64
	GponCurrent      float64
	XgponCurrent     float64
	OnuLaunchPower   float64
	OnuReceivePower  float64
	OnuCurrent       float64
	SerialNumber     string
	VendorID         string
	OnuUpTime        string
}

type OltTrafficsRecordData struct {
	Slot                           uint8
	Link                           uint8
	Port_tx_bytes                  uint64
	Port_rx_bytes                  uint64
	Port_tx_pkt                    uint64
	Port_rx_pkt                    uint64
	Port_tx_discard_pkt            uint64
	Port_rx_discard_pkt            uint64
	Port_rx_crc_error_pkt          uint64
	Port_rx_oversized_discard_pkt  uint64
	Port_rx_undersized_discard_pkt uint64
	Port_rx_error_pkt              uint64
	Port_tx_rate                   uint32
	Port_rx_rate                   uint32
	Port_tx_peak_rate              uint32
	Port_rx_peak_rate              uint32
}

type OltChannelTrafficsRecordData struct {
	Slot                           uint8
	Link                           uint8
	Channel                        uint8
	Port_tx_bytes                  uint64
	Port_rx_bytes                  uint64
	Port_tx_pkt                    uint64
	Port_rx_pkt                    uint64
	Port_tx_discard_pkt            uint64
	Port_rx_discard_pkt            uint64
	Port_rx_crc_error_pkt          uint64
	Port_rx_oversized_discard_pkt  uint64
	Port_rx_undersized_discard_pkt uint64
	Port_rx_error_pkt              uint64
	Port_tx_rate                   uint32
	Port_rx_rate                   uint32
	Port_tx_peak_rate              uint32
	Port_rx_peak_rate              uint32
}

type EthernetPortKpiRecordData struct {
	//ethernetCsmacd.f.s.p
	Link                            uint8
	Port_tx_bytes                   uint64
	Port_rx_bytes                   uint64
	Port_tx_packets                 uint64
	Port_rx_packets                 uint64
	Port_tx_discard_packets         uint64
	Port_rx_discard_packets         uint64
	Port_rx_alignment_error_packets uint64
	Port_tx_crc_error_packets       uint64
	Port_rx_crc_error_packets       uint64
	Port_tx_oversized_packets       uint64
	Port_rx_oversized_packets       uint64
	Port_tx_undersized_packets      uint64
	Port_rx_undersized_packets      uint64
	Port_tx_fragment_packets        uint64
	Port_rx_fragment_packets        uint64
	Port_tx_jabber_packets          uint64
	Port_rx_jabber_packets          uint64
	Port_tx_error_packets           uint64
	Port_rx_error_packets           uint64
	Port_tx_rate                    uint64
	Port_rx_rate                    uint64
	Port_tx_peak_rate               uint64
	Port_rx_peak_rate               uint64
}

type ServiceFlowKpiRecordData struct {
	VlanId                        uint32
	DownstreamQueueDropCnt        uint64
	DownstreamQueuePassCnt        uint64
	DownstreamQueueDropMax        uint32
	DownstreamQueueDropMin        uint32
	DownstreamQueueDropRateMax    uint32
	DownstreamQueueDropRateMin    uint32
	DownstreamQueueDropSecondsCnt uint32
	DownstreamQueuePassBytes uint64
	DownstreamMfrAvg         uint32
	UpstreamPassBytes        uint64
	UpstreamPassCnt          uint64
	UpstreamDropCnt          uint64
}

//var sensorDataMutex sync.Mutex

var (
	TelemetryDataMutex   sync.Mutex
	sensorDataMutex      sync.Mutex
	TelemetryClientMutex sync.Mutex
	StoreClient          = &Datastore{}
)

type Datastore struct {
}

func (ds *Datastore) DestGroup() *DestGroupTable {
	return NewDestGroup()
}
func (ds *Datastore) SubscriptionGroup() *SubscriptionGroupTable {
	return NewSubscriptionGroup()
}
func (ds *Datastore) SensorGroup() *SensorGroupTable {
	return NewSensorGroup()
}
func (ds *Datastore) SensorData() *SensorDataTable {
	return NewSensorData()
}
func (ds *Datastore) GrpcDataGroup() *GrpcDataTable {
	return NewGrpcData()
}
