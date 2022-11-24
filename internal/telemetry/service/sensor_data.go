package service

import (
	"telemetry/internal/telemetry/store"

	"github.com/marmotedu/iam/pkg/log"
)

type SensorDataSrv interface {
	GetALL() error
}

type sensorDataService struct {
	db *store.Datastore
}

var _ SensorDataSrv = (*sensorDataService)(nil)

func NewSensorDataService(s *service) *sensorDataService {
	return &sensorDataService{s.store}
}

func (s *sensorDataService) GetALL() error {
	sensordata, err := s.db.SensorData().GetFirstRecord()
	if err != nil {
		log.Infof("sensorDataService get first fail")
		return nil
	}
	log.Infof("sensorgrp:%+v", sensordata)

	for {
		sensordata, err = s.db.SensorData().GetNextRecord(sensordata)
		if err != nil {
			return nil
		}
		log.Infof("sensorgrp1:%+v", sensordata)
	}
}
