package service

import (
	"telemetry/internal/telemetry/store"
)

type Service interface {
	SensorGroup() SensorGroupSrv
	DestGroup() DestGroupSrv
	SubGroup() SubGroupSrv
	SensorData() SensorDataSrv
	GrpcData() GrpcDataSrv
}

type service struct {
	store *store.Datastore
}

func NewService(store *store.Datastore) Service {
	return &service{
		store: store,
	}
}

func (s *service) SensorGroup() SensorGroupSrv {
	return NewSensorGroupService(s)
}

func (s *service) DestGroup() DestGroupSrv {
	return NewDestGroupService(s)
}

func (s *service) SubGroup() SubGroupSrv {
	return NewsubscriptionGroupService(s)
}

func (s *service) SensorData() SensorDataSrv {
	return NewSensorDataService(s)
}

func (s *service) GrpcData() GrpcDataSrv {
	return NewGrpcDataService(s)
}
