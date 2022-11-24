package service

import (
	"strings"
	"telemetry/internal/telemetry/store"

	"github.com/marmotedu/iam/pkg/log"
	"github.com/pkg/errors"
)

type SensorGroupSrv interface {
	Create(sensorgrp *store.SensorGroupEntry) error
	Delete(sensorgrp *store.SensorGroupEntry) error
	GetALL() error
}
type sensorGroupService struct {
	db *store.Datastore
}

var _ SensorGroupSrv = (*sensorGroupService)(nil)

func NewSensorGroupService(s *service) *sensorGroupService {
	return &sensorGroupService{s.store}
}

func (s *sensorGroupService) Create(sensorgrp *store.SensorGroupEntry) error {
	sensorpath := strings.Split(sensorgrp.SensorPath, ".")
	if sensorpath[0] != "an_gpon_pm_olt_traffic:GponPmOltChannelTraffics" &&
		sensorpath[0] != "an_gpon_pm_olt_traffic:GponPmOltTraffics" &&
		sensorpath[0] != "an_ethernet_kpi:EthernetPortKpiRecords" &&
		sensorpath[0] != "an_bb_service_flow_kpi:ServiceFlowKpiRecords" {
		return errors.New("error : sensorpath error.")
	}

	err := s.db.SensorGroup().CreateRecord(sensorgrp)
	if err != nil {
		return err
	}
	var sensordata store.SensorDataEntry
	sensordata.SensorPath = sensorgrp.SensorPath
	if _, err = s.db.SensorData().GetRecord(&sensordata); err != nil {
		err = s.db.SensorData().CreateRecord(&sensordata)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *sensorGroupService) Delete(sensorgrp *store.SensorGroupEntry) error {
	if sensorgrp.SensorPath == "" {
		for {
			sensorgrptmp, err := s.db.SensorGroup().GetNextRecord(&store.SensorGroupEntry{
				SensorGroupId: sensorgrp.SensorGroupId,
			})
			if err == nil {
				err = s.db.SensorGroup().DelRecord(sensorgrptmp)
				if err != nil {
					return err
				}
			} else {
				break
			}
		}
	} else {
		err := s.db.SensorGroup().DelRecord(sensorgrp)
		if err != nil {
			return err
		}
	}
	return nil
	/* 	var sensordata store.SensorDataEntry
	   	sensordata.SensorPath = sensorgrp.SensorPath
	   	err = s.db.SensorData().DelRecord(&sensordata)
	   	if err != nil {
	   		return err
	   	}
	   	return nil */
}

func (s *sensorGroupService) GetALL() error {
	sensorgrp, err := s.db.SensorGroup().GetFirstRecord()
	if err != nil {
		return nil
	}
	log.Infof("sensorgrp:%+v", sensorgrp)
	for {
		sensorgrp, err = s.db.SensorGroup().GetNextRecord(sensorgrp)
		if err != nil {
			return nil
		}
		log.Infof("sensorgrp:%+v", sensorgrp)
	}
}
