package store

import (
	"strings"
	"sync"

	"github.com/marmotedu/errors"
)

type SensorGroupEntry struct {
	SensorGroupId string
	SensorPath    string
}

type SensorGroupRecord struct {
	next  *SensorGroupRecord
	value *SensorGroupEntry
}

type SensorGroupTable struct {
	head *SensorGroupRecord
}

var (
	sensorgrponce  sync.Once
	sensorgrptable *SensorGroupTable
)

func NewSensorGroup() *SensorGroupTable {
	sensorgrponce.Do(func() {
		sensorgrptable = &SensorGroupTable{}
	})
	return sensorgrptable
}

func (t *SensorGroupTable) Compare(leftCmpP *SensorGroupEntry, rightCmpP *SensorGroupEntry) int {
	if strings.Compare(leftCmpP.SensorGroupId, rightCmpP.SensorGroupId) > 0 {
		return 1
	}
	if strings.Compare(leftCmpP.SensorGroupId, rightCmpP.SensorGroupId) < 0 {
		return -1
	}

	if strings.Compare(leftCmpP.SensorPath, rightCmpP.SensorPath) > 0 {
		return 1
	}
	if strings.Compare(leftCmpP.SensorPath, rightCmpP.SensorPath) < 0 {
		return -1
	}
	return 0
}
func (t *SensorGroupTable) CreateRecord(val *SensorGroupEntry) error {
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
	newRecord := &SensorGroupRecord{value: val}
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
func (t *SensorGroupTable) GetFirstRecord() (*SensorGroupEntry, error) {
	if t.head != nil {
		return t.head.value, nil
	}
	return nil, errors.New("table is nil.")
}
func (t *SensorGroupTable) GetRecord(val *SensorGroupEntry) (*SensorGroupEntry, error) {
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
	return nil, errors.New("SensorGroupTable GetRecord fail")
}

func (t *SensorGroupTable) GetNextRecord(val *SensorGroupEntry) (*SensorGroupEntry, error) {
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
			break
		}
	}
	return nil, errors.New("SensorGroupTable GetNextRecord fail")
}

func (t *SensorGroupTable) DelRecord(val *SensorGroupEntry) error {
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
	return errors.New("SensorGroupTable DelRecord fail")
}

func (t *SensorGroupTable) SetRecord(val *SensorGroupEntry) error {
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
	return errors.New("SensorGroupTable SetRecord fail")
}
