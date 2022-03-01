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

//var sensorDataMutex sync.Mutex

var (
	TelemetryDataMutex sync.Mutex
	sensorDataMutex    sync.Mutex
	StoreClient        = &Datastore{}
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
