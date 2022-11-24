package service

import (
	"fmt"
	"strings"
	"telemetry/internal/telemetry/collect"
	v1 "telemetry/internal/telemetry/gnmi/v1"
	"telemetry/internal/telemetry/store"
	"time"

	"github.com/marmotedu/iam/pkg/log"

	"github.com/pkg/errors"
	"github.com/thinkeridea/go-extend/exnet"
	"google.golang.org/grpc"
)

type SubGroupSrv interface {
	Create(subgrp *store.SubscriptionGroupEntry) error
	Delete(subgrp *store.SubscriptionGroupEntry) error
	GetALL() error
}

func NewsubscriptionGroupService(s *service) *subscriptionGroupService {
	return &subscriptionGroupService{s.store}
}

type subscriptionGroupService struct {
	db *store.Datastore
}

func (s *subscriptionGroupService) isExistSensorGroup(SensorGroupId string) error {
	if sensorGrp, err := s.db.SensorGroup().GetFirstRecord(); err == nil {
		if sensorGrp.SensorGroupId == SensorGroupId {
			return nil
		}
		for {
			if sensorGrp, err = s.db.SensorGroup().GetNextRecord(sensorGrp); err == nil {
				if sensorGrp.SensorGroupId == SensorGroupId {
					return nil
				}
			} else {
				return errors.New("SensorGroup is not exist")
			}
		}
	}
	return errors.New("SensorGroup is not exist")
}

func (s *subscriptionGroupService) isExistDestrGroup(DestGroupId string) error {
	if destGrp, err := s.db.DestGroup().GetFirstRecord(); err == nil {
		if destGrp.DestGroupId == DestGroupId {
			return nil
		}
		for {
			if destGrp, err = s.db.DestGroup().GetNextRecord(destGrp); err == nil {
				if destGrp.DestGroupId == DestGroupId {
					return nil
				}
			} else {
				return errors.New("SensorGroup is not exist")
			}
		}
	}
	return errors.New("destGrp is not exist")
}
func (s *subscriptionGroupService) addSensorPath(SubscriptionId string, sensorgrp *store.SensorGroupEntry, grpcdata *store.GrpcDataEntry) error {
	if sensordata, err := s.db.SensorData().GetFirstRecord(); err == nil {
		if sensordata.SensorPath == sensorgrp.SensorPath {
			var oltTrafficsRecorddata store.OltTrafficsRecordData
			var oltChannelTrafficsRecorddata store.OltChannelTrafficsRecordData
			var ethernetPortKpiRecordData store.EthernetPortKpiRecordData
			var serviceFlowKpiRecordData store.ServiceFlowKpiRecordData
			//publish 频道
			if len(sensordata.SubscriptionIds) == 0 {
				log.Infof("Notify %s", sensorgrp.SensorPath)
				collect.Notify(sensorgrp.SensorPath, 1)
			}

			sensordata.SubscriptionIds = append(sensordata.SubscriptionIds, SubscriptionId)
			grpcdata.SensorPath = append(grpcdata.SensorPath, sensordata.SensorPath)
			store.TelemetryDataMutex.Lock()
			p1 := strings.Split(sensordata.SensorPath, ".")
			if p1[0] == "an_gpon_pm_olt_traffic:GponPmOltChannelTraffics" {
				sensordata.Data = append(sensordata.Data, &oltChannelTrafficsRecorddata)
				grpcdata.Data = append(grpcdata.Data, &oltChannelTrafficsRecorddata)
			} else if p1[0] == "an_gpon_pm_olt_traffic:GponPmOltTraffics" {
				sensordata.Data = append(sensordata.Data, &oltTrafficsRecorddata)
				grpcdata.Data = append(grpcdata.Data, &oltTrafficsRecorddata)
			} else if p1[0] == "an_ethernet_kpi:EthernetPortKpiRecords" {
				sensordata.Data = append(sensordata.Data, &ethernetPortKpiRecordData)
				grpcdata.Data = append(grpcdata.Data, &ethernetPortKpiRecordData)
			} else if p1[0] == "an_bb_service_flow_kpi:ServiceFlowKpiRecords" {
				sensordata.Data = append(sensordata.Data, &serviceFlowKpiRecordData)
				grpcdata.Data = append(grpcdata.Data, &serviceFlowKpiRecordData)
			}

			store.TelemetryDataMutex.Unlock()
		} else {
			for {
				if sensordata, err = s.db.SensorData().GetNextRecord(sensordata); err == nil {
					if sensordata.SensorPath == sensorgrp.SensorPath {
						var oltTrafficsRecorddata store.OltTrafficsRecordData
						var oltChannelTrafficsRecorddata store.OltChannelTrafficsRecordData
						var ethernetPortKpiRecordData store.EthernetPortKpiRecordData
						var serviceFlowKpiRecordData store.ServiceFlowKpiRecordData
						//publish 频道
						if len(sensordata.SubscriptionIds) == 0 {
							collect.Notify(sensorgrp.SensorPath, 1)
						}
						sensordata.SubscriptionIds = append(sensordata.SubscriptionIds, SubscriptionId)
						grpcdata.SensorPath = append(grpcdata.SensorPath, sensordata.SensorPath)
						store.TelemetryDataMutex.Lock()
						p1 := strings.Split(sensordata.SensorPath, ".")
						if p1[0] == "an_gpon_pm_olt_traffic:GponPmOltChannelTraffics" {
							sensordata.Data = append(sensordata.Data, &oltChannelTrafficsRecorddata)
							grpcdata.Data = append(grpcdata.Data, &oltChannelTrafficsRecorddata)
						} else if p1[0] == "an_gpon_pm_olt_traffic:GponPmOltTraffics" {
							sensordata.Data = append(sensordata.Data, &oltTrafficsRecorddata)
							grpcdata.Data = append(grpcdata.Data, &oltTrafficsRecorddata)
						} else if p1[0] == "an_ethernet_kpi:EthernetPortKpiRecords" {
							sensordata.Data = append(sensordata.Data, &ethernetPortKpiRecordData)
							grpcdata.Data = append(grpcdata.Data, &ethernetPortKpiRecordData)
						} else if p1[0] == "an_bb_service_flow_kpi:ServiceFlowKpiRecords" {
							sensordata.Data = append(sensordata.Data, &serviceFlowKpiRecordData)
							grpcdata.Data = append(grpcdata.Data, &serviceFlowKpiRecordData)
						}
						store.TelemetryDataMutex.Unlock()
						break
					}
				} else {
					break
				}
			}
		}

	} else {
		errors.New("SensorData is not exist")
	}
	return nil
}
func (s *subscriptionGroupService) addDataTable(subgrp *store.SubscriptionGroupEntry, grpcDataEntry *store.GrpcDataEntry) error {
	if sensorGrp, err := s.db.SensorGroup().GetFirstRecord(); err == nil {
		//先找个sensor group id相匹配的senor path
		if sensorGrp.SensorGroupId == subgrp.SensorGroupId {
			err = s.addSensorPath(subgrp.SubscriptionId, sensorGrp, grpcDataEntry)
			if err != nil {
				return err
			}
		}
		for {
			if sensorGrp, err = s.db.SensorGroup().GetNextRecord(sensorGrp); err == nil {
				if sensorGrp.SensorGroupId == subgrp.SensorGroupId {
					err = s.addSensorPath(subgrp.SubscriptionId, sensorGrp, grpcDataEntry)
					if err != nil {
						return err
					}
				}
			} else {
				break
			}
		}
	} else {
		return errors.New("sensorGrp is not exist")
	}

	return nil
}
func CreateGrpcClient(addr string) (v1.GRPCDataserviceClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithTimeout(time.Second), grpc.WithBlock())
	if err != nil {
		return nil, nil, errors.Wrapf(err, "grpc dial %s fail", addr)
	}
	//	log.Infof("CreateGrpcClient2 %v", conn)
	client := v1.NewGRPCDataserviceClient(conn)
	return client, conn, nil
}
func (s *subscriptionGroupService) addGrpcConn(subgrp *store.SubscriptionGroupEntry, grpcDataEntry *store.GrpcDataEntry) error {
	if destGrp, err := s.db.DestGroup().GetFirstRecord(); err == nil {
		if destGrp.DestGroupId == subgrp.DestGroupId {
			ip, _ := exnet.Long2IPString(uint(destGrp.DestIp))
			log.Infof("ip:%s,port:%d\n", ip, destGrp.DestPort)
			client, conn, err := CreateGrpcClient(fmt.Sprintf("%s:%d", ip, destGrp.DestPort))
			if err != nil {
				return err
			}
			grpcDataEntry.Conns = append(grpcDataEntry.Conns, conn)
			grpcDataEntry.Client = append(grpcDataEntry.Client, client)
		}
		for {
			if destGrp, err = s.db.DestGroup().GetNextRecord(destGrp); err == nil {
				if destGrp.DestGroupId == subgrp.DestGroupId {
					ip, _ := exnet.Long2IPString(uint(destGrp.DestIp))
					client, conn, err := CreateGrpcClient(fmt.Sprintf("%s:%d", ip, destGrp.DestPort))
					if err != nil {
						return err
					}
					grpcDataEntry.Conns = append(grpcDataEntry.Conns, conn)
					grpcDataEntry.Client = append(grpcDataEntry.Client, client)
				}
			} else {
				break
			}
		}
	} else {
		return errors.New("destGrp is not exist")
	}

	return nil
}

func delGrpcConn(grpcDataEntry *store.GrpcDataEntry) {
	for _, conn := range grpcDataEntry.Conns {
		conn.Close()
	}
}
func (s *subscriptionGroupService) Create(subgrp *store.SubscriptionGroupEntry) error {
	var grpcDataEntry store.GrpcDataEntry
	//应该在snmp上做
	if err := s.isExistSensorGroup(subgrp.SensorGroupId); err != nil {
		return err
	}
	//应该在snmp上做
	if err := s.isExistDestrGroup(subgrp.DestGroupId); err != nil {
		return err
	}
	if err := s.db.SubscriptionGroup().CreateRecord(subgrp); err != nil {
		return err
	}
	if err := s.addDataTable(subgrp, &grpcDataEntry); err != nil {
		return err
	}
	//注释掉grpc相关
	grpcDataEntry.SubscriptionId = subgrp.SubscriptionId
	grpcDataEntry.Interval = subgrp.Interval
	if err := s.addGrpcConn(subgrp, &grpcDataEntry); err != nil {
		delGrpcConn(&grpcDataEntry) //断开grpc连接
		s.delDataTable(subgrp)      //去除sensor data中Data
		return err
	}
	if err := s.db.GrpcDataGroup().CreateRecord(&grpcDataEntry); err != nil {
		return err
	}
	return nil
}

func (s *subscriptionGroupService) delSensorPath(SubscriptionId string, sensorgrp *store.SensorGroupEntry) error {
	if sensordata, err := s.db.SensorData().GetFirstRecord(); err == nil {
		if sensordata.SensorPath == sensorgrp.SensorPath {
			for index, subId := range sensordata.SubscriptionIds {
				if subId == SubscriptionId {
					sensordata.SubscriptionIds = append(sensordata.SubscriptionIds[:index], sensordata.SubscriptionIds[index+1:]...)
					sensordata.Data = append(sensordata.Data[:index], sensordata.Data[index+1:]...)
					if len(sensordata.SubscriptionIds) == 0 {
						collect.Notify(sensorgrp.SensorPath, 0)
					}
					break
				}
			}
		} else {
			for {
				if sensordata, err = s.db.SensorData().GetNextRecord(sensordata); err == nil {
					if sensordata.SensorPath == sensorgrp.SensorPath {
						for index, subId := range sensordata.SubscriptionIds {
							if subId == SubscriptionId {
								sensordata.SubscriptionIds = append(sensordata.SubscriptionIds[:index], sensordata.SubscriptionIds[index+1:]...)
								sensordata.Data = append(sensordata.Data[:index], sensordata.Data[index+1:]...)
								if len(sensordata.SubscriptionIds) == 0 {
									collect.Notify(sensorgrp.SensorPath, 0)
								}
								break
							}
						}
					}
				} else {
					break
				}
			}
		}

	} else {
		errors.New("SensorData is not exist")
	}
	return nil
}
func (s *subscriptionGroupService) delDataTable(subgrp *store.SubscriptionGroupEntry) error {
	if sensorGrp, err := s.db.SensorGroup().GetFirstRecord(); err == nil {
		//先找个sensor group id相匹配的senor path
		if sensorGrp.SensorGroupId == subgrp.SensorGroupId {
			err = s.delSensorPath(subgrp.SubscriptionId, sensorGrp)
			if err != nil {
				return err
			}
		}
		for {
			if sensorGrp, err = s.db.SensorGroup().GetNextRecord(sensorGrp); err == nil {
				if sensorGrp.SensorGroupId == subgrp.SensorGroupId {
					err = s.delSensorPath(subgrp.SubscriptionId, sensorGrp)
					if err != nil {
						return err
					}
				}
			} else {
				break
			}
		}
	} else {
		return errors.New("destGrp is not exist")
	}
	return nil
}
func (s *subscriptionGroupService) deleteSubData(subgrp *store.SubscriptionGroupEntry) error {
	if err := s.delDataTable(subgrp); err != nil {
		return err
	}
	var grpcDataEnt store.GrpcDataEntry
	var grpcDataEntp *store.GrpcDataEntry
	grpcDataEnt.SubscriptionId = subgrp.SubscriptionId
	grpcDataEntp, err := s.db.GrpcDataGroup().GetRecord(&grpcDataEnt)
	if err != nil {
		return err
	}
	delGrpcConn(grpcDataEntp)
	if err := s.db.GrpcDataGroup().DelRecord(grpcDataEntp); err != nil {
		return err
	}
	if err := s.db.SubscriptionGroup().DelRecord(subgrp); err != nil {
		return err
	}
	return nil
}
func (s *subscriptionGroupService) Delete(subgrp *store.SubscriptionGroupEntry) error {
	if subgrp.SensorGroupId == "" && subgrp.DestGroupId == "" {
		//切割字符串
		subId := strings.Split(subgrp.SubscriptionId, "-")
		subgrptmp := &store.SubscriptionGroupEntry{
			SubscriptionId: subId[0],
		}
		var err error
		for {
			subgrptmp, err = s.db.SubscriptionGroup().GetNextRecord(subgrptmp)
			if err == nil {
				tmpsubId := strings.Split(subgrptmp.SubscriptionId, "-")
				if tmpsubId[0] != subId[0] {
					break
				}
				store.TelemetryDataMutex.Lock()
				err = s.deleteSubData(subgrptmp)
				store.TelemetryDataMutex.Unlock()
				if err != nil {
					return err
				}
			} else {
				break
			}
		}
	} else if subgrp.SensorGroupId == "" {
		subId := strings.Split(subgrp.SubscriptionId, "-")
		subgrptmp := &store.SubscriptionGroupEntry{
			SubscriptionId: subId[0],
		}
		var err error
		for {
			subgrptmp, err = s.db.SubscriptionGroup().GetNextRecord(subgrptmp)
			if err == nil {
				tmpsubId := strings.Split(subgrptmp.SubscriptionId, "-")
				if tmpsubId[0] != subId[0] {
					break
				}
				if tmpsubId[2] == subId[2] {
					store.TelemetryDataMutex.Lock()
					err = s.deleteSubData(subgrptmp)
					store.TelemetryDataMutex.Unlock()
					if err != nil {
						return err
					}
				}
			} else {
				break
			}
		}
	} else if subgrp.DestGroupId == "" {
		subId := strings.Split(subgrp.SubscriptionId, "-")
		subgrptmp := &store.SubscriptionGroupEntry{
			SubscriptionId: subId[0],
		}
		var err error
		for {
			subgrptmp, err = s.db.SubscriptionGroup().GetNextRecord(subgrptmp)
			if err == nil {
				tmpsubId := strings.Split(subgrptmp.SubscriptionId, "-")
				if tmpsubId[0] != subId[0] {
					break
				}
				if tmpsubId[1] == subId[1] {
					store.TelemetryDataMutex.Lock()
					err = s.deleteSubData(subgrptmp)
					store.TelemetryDataMutex.Unlock()
					if err != nil {
						return err
					}
				}
			} else {
				break
			}
		}
	} else {
		store.TelemetryDataMutex.Lock()
		err := s.deleteSubData(subgrp)
		store.TelemetryDataMutex.Unlock()
		return err
	}
	return nil
}

func (s *subscriptionGroupService) GetALL() error {
	subgrp, err := s.db.SubscriptionGroup().GetFirstRecord()
	if err != nil {
		return nil
	}
	log.Infof("subgrp:%v", subgrp)
	for {
		subgrp, err := s.db.SubscriptionGroup().GetNextRecord(subgrp)
		if err != nil {
			return nil
		}
		log.Infof("subgrp:%v", subgrp)
	}
}

func (s *subscriptionGroupService) Set(subgrp *store.SubscriptionGroupEntry) error {
	var grpcDataEnt store.GrpcDataEntry
	var grpcDataEntp *store.GrpcDataEntry
	var subGrpEnt store.SubscriptionGroupEntry
	var subGrpEntp *store.SubscriptionGroupEntry
	grpcDataEnt.SubscriptionId = subgrp.SubscriptionId
	subGrpEnt.SubscriptionId = subgrp.SubscriptionId
	grpcDataEntp, err := s.db.GrpcDataGroup().GetRecord(&grpcDataEnt)
	if err != nil {
		return err
	}
	grpcDataEntp.Interval = subgrp.Interval
	if err := s.db.GrpcDataGroup().SetRecord(grpcDataEntp); err != nil {
		return err
	}

	subGrpEntp, err = s.db.SubscriptionGroup().GetRecord(&subGrpEnt)
	if err != nil {
		return err
	}
	subGrpEntp.Interval = subgrp.Interval
	if err := s.db.SubscriptionGroup().SetRecord(subGrpEntp); err != nil {
		return err
	}

	return nil
}
