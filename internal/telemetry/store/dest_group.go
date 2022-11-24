package store

import (
	"strings"
	"sync"

	"github.com/marmotedu/errors"
)

type DestGroupEntry struct {
	DestGroupId string
	DestIp      uint32
	DestPort    uint32
}

type DestGroupRecord struct {
	next  *DestGroupRecord
	value *DestGroupEntry
}

type DestGroupTable struct {
	head *DestGroupRecord
}

var (
	destgrponce  sync.Once
	destgrptable *DestGroupTable
)

func NewDestGroup() *DestGroupTable {
	destgrponce.Do(func() {
		destgrptable = &DestGroupTable{}
	})
	return destgrptable
}

func (t *DestGroupTable) Compare(leftCmpP *DestGroupEntry, rightCmpP *DestGroupEntry) int {
	if strings.Compare(leftCmpP.DestGroupId, rightCmpP.DestGroupId) > 0 {
		return 1
	}
	if strings.Compare(leftCmpP.DestGroupId, rightCmpP.DestGroupId) < 0 {
		return -1
	}

	if leftCmpP.DestIp > rightCmpP.DestIp {
		return 1
	}
	if leftCmpP.DestIp < rightCmpP.DestIp {
		return -1
	}
	if leftCmpP.DestPort > rightCmpP.DestPort {
		return 1
	}
	if leftCmpP.DestPort < rightCmpP.DestPort {
		return -1
	}
	return 0
}
func (t *DestGroupTable) CreateRecord(val *DestGroupEntry) error {
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
	newRecord := &DestGroupRecord{value: val}
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
func (t *DestGroupTable) GetFirstRecord() (*DestGroupEntry, error) {
	if t.head != nil {
		return t.head.value, nil
	}
	return nil, errors.New("table is nil.")
}
func (t *DestGroupTable) GetRecord(val *DestGroupEntry) (*DestGroupEntry, error) {
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
	return nil, errors.New("DestGroupTable GetRecord fail")
}

func (t *DestGroupTable) GetNextRecord(val *DestGroupEntry) (*DestGroupEntry, error) {
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
	return nil, errors.New("DestGroupTable GetNextRecord fail")
}

func (t *DestGroupTable) DelRecord(val *DestGroupEntry) error {
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
	return errors.New("DestGroupTable DelRecord fail")
}

func (t *DestGroupTable) SetRecord(val *DestGroupEntry) error {
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
	return errors.New("DestGroupTable SetRecord fail")
}
