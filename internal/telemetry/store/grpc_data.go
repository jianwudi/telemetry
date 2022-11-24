package store

import (
	"strings"
	"sync"

	v1 "telemetry/internal/telemetry/gnmi/v1"

	"github.com/marmotedu/errors"
	"google.golang.org/grpc"
)

type GrpcDataEntry struct {
	SubscriptionId string
	Conns          []*grpc.ClientConn
	Client         []v1.GRPCDataserviceClient
	SensorPath     []string
	Data           []interface{}
	Interval       uint32
	Count          uint32
}

type GrpcDataRecord struct {
	next  *GrpcDataRecord
	value *GrpcDataEntry
}

type GrpcDataTable struct {
	head *GrpcDataRecord
}

var (
	grpcdataonce  sync.Once
	grpcdatatable *GrpcDataTable
)

func NewGrpcData() *GrpcDataTable {
	grpcdataonce.Do(func() {
		grpcdatatable = &GrpcDataTable{}
	})
	return grpcdatatable
}

func (t *GrpcDataTable) Compare(leftCmpP *GrpcDataEntry, rightCmpP *GrpcDataEntry) int {
	if strings.Compare(leftCmpP.SubscriptionId, rightCmpP.SubscriptionId) > 0 {
		return 1
	}
	if strings.Compare(leftCmpP.SubscriptionId, rightCmpP.SubscriptionId) < 0 {
		return -1
	}
	return 0
}
func (t *GrpcDataTable) CreateRecord(val *GrpcDataEntry) error {
	curRecP := t.head
	prevRecP := curRecP
	for curRecP != nil {
		cmpRet := t.Compare(curRecP.value, val)
		if cmpRet < 0 {
			prevRecP = curRecP
			curRecP = curRecP.next
		} else if cmpRet == 0 {
			return errors.New("error : table already exist.")
		} else {
			break
		}
	}
	newRecord := &GrpcDataRecord{value: val}
	if t.head == nil {
		t.head = newRecord
	} else if curRecP == prevRecP {
		t.head = newRecord
		newRecord.next = curRecP
	} else {
		newRecord.next = curRecP
		prevRecP.next = newRecord
	}
	return nil
}
func (t *GrpcDataTable) GetFirstRecord() (*GrpcDataEntry, error) {
	if t.head != nil {
		return t.head.value, nil
	}
	return nil, errors.New("table is nil.")
}
func (t *GrpcDataTable) GetRecord(val *GrpcDataEntry) (*GrpcDataEntry, error) {
	curRecP := t.head
	for curRecP != nil {
		cmpRet := t.Compare(curRecP.value, val)

		if cmpRet < 0 {
			curRecP = curRecP.next
		} else if cmpRet == 0 {
			return curRecP.value, nil
		} else {
			break
		}
	}
	return nil, errors.New("GrpcDataEntry GetRecord fail")
}

func (t *GrpcDataTable) GetNextRecord(val *GrpcDataEntry) (*GrpcDataEntry, error) {
	curRecP := t.head
	for curRecP != nil {
		cmpRet := t.Compare(curRecP.value, val)
		if cmpRet < 0 {
			curRecP = curRecP.next
		} else if cmpRet == 0 {
			curRecP = curRecP.next
			if curRecP != nil {
				return curRecP.value, nil
			} else {
				break
			}
		} else {
			return curRecP.value, nil
		}
	}
	return nil, errors.New("GrpcDataEntry GetNextRecord fail")
}

func (t *GrpcDataTable) DelRecord(val *GrpcDataEntry) error {
	curRecP := t.head
	prevRecP := curRecP
	for curRecP != nil {
		cmpRet := t.Compare(curRecP.value, val)
		if cmpRet < 0 {
			prevRecP = curRecP
			curRecP = curRecP.next
		} else if cmpRet == 0 {
			if curRecP == t.head {
				t.head = curRecP.next
			} else {
				prevRecP.next = curRecP.next
			}
			return nil
		} else {
			break
		}
	}
	return errors.New("GrpcDataEntry DelRecord fail")
}

func (t *GrpcDataTable) SetRecord(val *GrpcDataEntry) error {
	curRecP := t.head
	for curRecP != nil {
		cmpRet := t.Compare(curRecP.value, val)
		if cmpRet < 0 {
			curRecP = curRecP.next
		} else if cmpRet == 0 {
			curRecP.value = val
			return nil
		} else {
			break
		}
	}
	return errors.New("GrpcDataEntry SetRecord fail")
}
