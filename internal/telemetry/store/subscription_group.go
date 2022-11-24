package store

import (
	"strings"
	"sync"

	"github.com/marmotedu/errors"
)

type SubscriptionGroupEntry struct {
	SubscriptionId string
	SensorGroupId  string
	DestGroupId    string
	Interval       uint32
}

type SubscriptionGroupRecord struct {
	next  *SubscriptionGroupRecord
	value *SubscriptionGroupEntry
}

type SubscriptionGroupTable struct {
	head *SubscriptionGroupRecord
}

var (
	subgrponce  sync.Once
	subgrptable *SubscriptionGroupTable
)

func NewSubscriptionGroup() *SubscriptionGroupTable {
	subgrponce.Do(func() {
		subgrptable = &SubscriptionGroupTable{}
	})
	return subgrptable
}

func (t *SubscriptionGroupTable) Compare(leftCmpP *SubscriptionGroupEntry, rightCmpP *SubscriptionGroupEntry) int {
	if strings.Compare(leftCmpP.SubscriptionId, rightCmpP.SubscriptionId) > 0 {
		return 1
	}
	if strings.Compare(leftCmpP.SubscriptionId, rightCmpP.SubscriptionId) < 0 {
		return -1
	}
	/* 	if strings.Compare(leftCmpP.SensorGroupId, rightCmpP.SensorGroupId) > 0 {
	   		return 1
	   	}
	   	if strings.Compare(leftCmpP.SensorGroupId, rightCmpP.SensorGroupId) < 0 {
	   		return -1
	   	}

	   	if strings.Compare(leftCmpP.DestGroupId, rightCmpP.DestGroupId) > 0 {
	   		return 1
	   	}
	   	if strings.Compare(leftCmpP.DestGroupId, rightCmpP.DestGroupId) < 0 {
	   		return -1
	   	} */
	return 0
}
func (t *SubscriptionGroupTable) CreateRecord(val *SubscriptionGroupEntry) error {
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
	newRecord := &SubscriptionGroupRecord{value: val}
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
func (t *SubscriptionGroupTable) GetFirstRecord() (*SubscriptionGroupEntry, error) {
	if t.head != nil {
		return t.head.value, nil
	}
	return nil, errors.New("table is nil.")
}
func (t *SubscriptionGroupTable) GetRecord(val *SubscriptionGroupEntry) (*SubscriptionGroupEntry, error) {
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
	return nil, errors.New("SubscriptionGroupEntry GetRecord fail")
}

func (t *SubscriptionGroupTable) GetNextRecord(val *SubscriptionGroupEntry) (*SubscriptionGroupEntry, error) {
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
	return nil, errors.New("SubscriptionGroupEntry GetNextRecord fail")
}

func (t *SubscriptionGroupTable) DelRecord(val *SubscriptionGroupEntry) error {
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
	return errors.New("SubscriptionGroupEntry DelRecord fail")
}

func (t *SubscriptionGroupTable) SetRecord(val *SubscriptionGroupEntry) error {
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
	return errors.New("SubscriptionGroupEntry SetRecord fail")
}
