package dao

type SysInfo struct{
	Ip 		    string
	CpuUsage    float64
	MemoryUsage float64
	DiskUsage   []float64
	NetworkIn   float64
	NetworkOut  float64
	DataTime    string
}