package libs

import (
	"fmt"
	"time"
	mydao "github.com/wayne011872/systemMonitor/dao"
	mclt "github.com/wayne011872/systemMonitor/database/mongo"
	mcrud "github.com/wayne011872/systemMonitor/model/sysInfo"
	mecrud "github.com/wayne011872/systemMonitor/model/errorSysInfo"
)

func SaveSysInfoInMgo() error {
	networkName := GetNetworkName()
	if networkName == "" {
		panic("取不到NETWORK_NAME")
	}
	mgoClient,err := mclt.GetMgoDBClient()
	if err != nil {
		return err
	}
	mgoCRUD := mcrud.NewCRUD(mgoClient.GetCtx(), mgoClient.GetDB())
	netIn, netOut := GetNetPerSecond(GetNetInfo, networkName)
	err = mgoCRUD.Save(
		&mydao.SysInfo{
			Ip:			 GetLocalIP(),
			CpuUsage:    GetCpuPercent(),
			MemoryUsage: GetMemoryPercent(),
			DiskUsage:   GetDiskPercent(GetDiskPartitions, GetDiskUsageState),
			NetworkIn:   netIn,
			NetworkOut:  netOut,
			DataTime:    time.Now().Format("2006-01-02 15:04:05"),
		})
	if err != nil {
		return err
	}
	return nil
}

func SaveErrorSysInfoInMgo() error {
	fmt.Printf("[%s] Save Error System Data In Mongo\n",time.Now().Format("2006-01-02 15:04:05"))
	networkName := GetNetworkName()
	if networkName == "" {
		panic("取不到NETWORK_NAME")
	}
	mgoClient,err := mclt.GetMgoDBClient()
	if err != nil {
		return err
	}
	mgoCRUD := mecrud.NewCRUD(mgoClient.GetCtx(), mgoClient.GetErrDB())
	netIn, netOut := GetNetPerSecond(GetNetInfo, networkName)
	err = mgoCRUD.Save(
		&mydao.ErrorSysInfo{
			Ip:			 GetLocalIP(),
			CpuUsage:    GetCpuPercent(),
			CpuProcess: GetProcessesCPU(),
			MemoryUsage: GetMemoryPercent(),
			MemoryProcess: GetProcessesMemory(),
			DiskUsage:   GetDiskPercent(GetDiskPartitions, GetDiskUsageState),
			NetworkIn:   netIn,
			NetworkOut:  netOut,
			DataTime:    time.Now().Format("2006-01-02 15:04:05"),
		})
	if err != nil {
		return err
	}
	return nil
}