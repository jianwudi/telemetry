package store

import (
	"strings"
	"sync"

	"github.com/marmotedu/errors"
)

type SensorDataEntry struct {
	SensorPath      string
	SubscriptionIds []string
	Data            []interface{}
	//	Interval        uint32
}
type SensorDataRecord struct {
	next  *SensorDataRecord
	value *SensorDataEntry
}

type SensorDataTable struct {
	head *SensorDataRecord
}

var (
	subdataonce  sync.Once
	subdatatable *SensorDataTable
)

func NewSensorData() *SensorDataTable {
	subdataonce.Do(func() {
		subdatatable = &SensorDataTable{}
	})
	return subdatatable
}

func (t *SensorDataTable) Compare(leftCmpP *SensorDataEntry, rightCmpP *SensorDataEntry) int {
	if strings.Compare(leftCmpP.SensorPath, rightCmpP.SensorPath) > 0 {
		return 1
	}
	if strings.Compare(leftCmpP.SensorPath, rightCmpP.SensorPath) < 0 {
		return -1
	}
	return 0
}
func (t *SensorDataTable) CreateRecord(val *SensorDataEntry) error {
	curRecP := t.head
	prevRecP := curRecP
	sensorDataMutex.Lock()
	for curRecP != nil {
		cmpRet := t.Compare(curRecP.value, val)
		if cmpRet < 0 {
			prevRecP = curRecP
			curRecP = curRecP.next
		} else if cmpRet == 0 {
			sensorDataMutex.Unlock()
			return errors.New("error : table already exist.")
		} else {
			break
		}
	}
	newRecord := &SensorDataRecord{value: val}
	if t.head == nil {
		t.head = newRecord
	} else if curRecP == prevRecP {
		t.head = newRecord
		newRecord.next = curRecP
	} else {
		newRecord.next = curRecP
		prevRecP.next = newRecord
	}
	sensorDataMutex.Unlock()
	return nil
}
func (t *SensorDataTable) GetFirstRecord() (*SensorDataEntry, error) {
	sensorDataMutex.Lock()
	if t.head != nil {
		sensorDataMutex.Unlock()
		return t.head.value, nil
	}
	sensorDataMutex.Unlock()
	return nil, errors.New("table is nil.")
}
func (t *SensorDataTable) GetRecord(val *SensorDataEntry) (*SensorDataEntry, error) {
	sensorDataMutex.Lock()
	curRecP := t.head
	for curRecP != nil {
		cmpRet := t.Compare(curRecP.value, val)
		if cmpRet < 0 {
			curRecP = curRecP.next
		} else if cmpRet == 0 {
			sensorDataMutex.Unlock()
			return curRecP.value, nil
		} else {
			break
		}
	}
	sensorDataMutex.Unlock()
	return nil, errors.New("SensorDataEntry GetRecord fail")
}

func (t *SensorDataTable) GetNextRecord(val *SensorDataEntry) (*SensorDataEntry, error) {
	sensorDataMutex.Lock()
	curRecP := t.head
	for curRecP != nil {
		cmpRet := t.Compare(curRecP.value, val)
		if cmpRet < 0 {
			curRecP = curRecP.next
		} else if cmpRet == 0 {
			curRecP = curRecP.next
			if curRecP != nil {
				sensorDataMutex.Unlock()
				return curRecP.value, nil
			} else {
				break
			}
		} else {
			sensorDataMutex.Unlock()
			return curRecP.value, nil
		}
	}
	sensorDataMutex.Unlock()
	return nil, errors.New("SensorDataEntry GetNextRecord fail")
}

func (t *SensorDataTable) DelRecord(val *SensorDataEntry) error {
	sensorDataMutex.Lock()
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
			sensorDataMutex.Unlock()
			return nil
		} else {
			break
		}
	}
	sensorDataMutex.Unlock()
	return errors.New("SensorDataEntry DelRecord fail")
}

func (t *SensorDataTable) SetRecord(val *SensorDataEntry) error {
	sensorDataMutex.Lock()
	curRecP := t.head
	for curRecP != nil {
		cmpRet := t.Compare(curRecP.value, val)
		if cmpRet < 0 {
			curRecP = curRecP.next
		} else if cmpRet == 0 {
			curRecP.value = val
			sensorDataMutex.Unlock()
			return nil
		} else {
			break
		}
	}
	sensorDataMutex.Unlock()
	return errors.New("SensorDataEntry SetRecord fail")
}
