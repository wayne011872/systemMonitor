package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
	"github.com/joho/godotenv"
	
	ss "github.com/wayne011872/systemMonitor/libs"
	myMail "github.com/wayne011872/systemMonitor/mail"
)

var (
	output = flag.String("o", "mongo", "output(mongo or print)")
)

func main() {
	var sendTime time.Time
	isSend := false
	flag.Parse()
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("No .env file")
	}
	ipAddress := ss.GetLocalIP()
	if ipAddress == ""{
		panic("找不到ip位置")
	}
	switch *output {
	case "print":
		for {
			mailContent,isSysError := ss.DetectError()
			isSend = myMail.IsSendMail(sendTime,isSend)
			if isSysError && !isSend{
				err := ss.SaveErrorSysInfoInMgo()
				if err != nil{
					myMail.SendMgoErrorMail(ss.GetLocalIP() ,err)
				}
				sendTime,isSend = myMail.SendSystemErrorMail(ipAddress,mailContent)
			}
			RunPrint()
		}
	case "mongo":
		for {
			mailContent,isSysError := ss.DetectError()
			isSend = myMail.IsSendMail(sendTime,isSend)
			if isSysError && !isSend{
				err := ss.SaveErrorSysInfoInMgo()
				if err != nil{
					myMail.SendMgoErrorMail(ss.GetLocalIP() ,err)
				}
				sendTime,isSend = myMail.SendSystemErrorMail(ipAddress,mailContent)
			}
			RunMongo()
		}
	default:
		panic("invalid output")
	}
}

func RunPrint() {
	networkName := ss.GetNetworkName()
	sleepTime,_ := strconv.Atoi(os.Getenv(("PRINT_INTERVAL_TIME")))
	if networkName == ""{
		panic("取不到NETWORK_NAME")
	}
	fmt.Printf("IP: %s\n",ss.GetLocalIP())
	fmt.Printf("CPU使用率: %f%%\n", ss.GetCpuPercent())
	cpuMessage := ss.GetProcessesCPU()
	for _,p := range cpuMessage{
		fmt.Printf("Pid:%-10s 程序名稱: %-30s CPU使用率:%.2f%%\n", strconv.FormatInt(int64(p.Pid), 10), p.Name, p.Cpu)
	}
	fmt.Printf("記憶體使用率: %f%%\n", ss.GetMemoryPercent())
	memoryMessage := ss.GetProcessesMemory()
	for _,p := range memoryMessage{
		fmt.Printf("Pid:%-10s 程序名稱: %-30s 記憶體使用率:%.2f%%\n", strconv.FormatInt(int64(p.Pid), 10), p.Name, p.MemRate)
	}
	diskPercents := ss.GetDiskPercent(ss.GetDiskPartitions, ss.GetDiskUsageState)
	for k, d := range diskPercents {
		fmt.Printf("硬碟%d使用率: %f%%\n", k, d)
	}
	netIn, netOut := ss.GetNetPerSecond(ss.GetNetInfo, networkName)
	netInTrans,netInUnit := ss.TransferNetworkUnit(float64(netIn),0)
	netOutTrans,netOutUnit := ss.TransferNetworkUnit(float64(netOut),0)
	fmt.Printf("網路接收: %f %s\n", netInTrans,netInUnit)
	fmt.Printf("網路傳送: %f %s\n", netOutTrans,netOutUnit)
	fmt.Printf("--------------------------------------per %d seconds------------------------------------------\n",sleepTime)
	time.Sleep(time.Duration(sleepTime) * time.Second)
}

func RunMongo() {
	fmt.Printf("[%s] Save System Resource Data In Mongo\n",time.Now().Format("2006-01-02 15:04:05"))
	sleepTime,_ := strconv.Atoi(os.Getenv(("MONGO_INTERVAL_TIME")))
	err := ss.SaveSysInfoInMgo()
	if err != nil {
		myMail.SendMgoErrorMail(ss.GetLocalIP(),err)
	}
	fmt.Printf("--------------------------------------per %d seconds------------------------------------------\n",sleepTime)
	time.Sleep(time.Duration(sleepTime) * time.Second)
}