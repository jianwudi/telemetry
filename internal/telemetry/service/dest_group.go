package service

import (
	"telemetry/internal/telemetry/store"

	"github.com/marmotedu/iam/pkg/log"
)

type DestGroupSrv interface {
	Create(destgrp *store.DestGroupEntry) error
	Delete(destgrp *store.DestGroupEntry) error
	GetALL() error
}
type destGroupService struct {
	db *store.Datastore
}

var _ DestGroupSrv = (*destGroupService)(nil)

func NewDestGroupService(s *service) *destGroupService {
	return &destGroupService{s.store}
}

func (s *destGroupService) Create(destgrp *store.DestGroupEntry) error {

	err := s.db.DestGroup().CreateRecord(destgrp)
	if err != nil {
		return err
	}

	return nil
}

func (s *destGroupService) Delete(destgrp *store.DestGroupEntry) error {
	if destgrp.DestIp == 0 && destgrp.DestPort == 0 {
		for {
			dstgrptmp, err := s.db.DestGroup().GetNextRecord(&store.DestGroupEntry{
				DestGroupId: destgrp.DestGroupId,
			})
			if err == nil {
				err = s.db.DestGroup().DelRecord(dstgrptmp)
				if err != nil {
					return err
				}
			} else {
				break
			}
		}
	} else {
		err := s.db.DestGroup().DelRecord(destgrp)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *destGroupService) GetALL() error {
	destgrp, err := s.db.DestGroup().GetFirstRecord()
	if err != nil {
		return nil
	}
	log.Infof("destgrp:%v", destgrp)
	for {
		destgrp, err := s.db.DestGroup().GetNextRecord(destgrp)
		if err != nil {
			return nil
		}
		log.Infof("destgrp:%v", destgrp)
	}
}
