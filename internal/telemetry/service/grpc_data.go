package service

import (
	"telemetry/internal/telemetry/store"

	"github.com/marmotedu/iam/pkg/log"
)

type GrpcDataSrv interface {
	GetALL() error
}

type grpcDataService struct {
	db *store.Datastore
}

var _ GrpcDataSrv = (*grpcDataService)(nil)

func NewGrpcDataService(s *service) *grpcDataService {
	return &grpcDataService{s.store}
}

func (s *grpcDataService) GetALL() error {
	grpcdata, err := s.db.GrpcDataGroup().GetFirstRecord()
	if err != nil {
		return nil
	}
	log.Infof("grpcdata:%v", grpcdata)
	for _, path := range grpcdata.SensorPath {
		log.Infof("path:%s", path)
	}
	for _, data := range grpcdata.Data {
		log.Infof("data:%+v", data)
	}
	for {
		grpcdata, err = s.db.GrpcDataGroup().GetNextRecord(grpcdata)
		if err != nil {
			return nil
		}
		log.Infof("grpcdata:%v", grpcdata)
		for _, path := range grpcdata.SensorPath {
			log.Infof("path:%s", path)
		}
		for _, data := range grpcdata.Data {
			log.Infof("data:%+v", data)
		}
	}
}
