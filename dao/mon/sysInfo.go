package mon

import(
	myDao "github.com/wayne011872/systemMonitor/dao"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"strconv"
)

func GetSysInfoCollection()(string,string){
	year := strconv.Itoa(time.Now().Year())
	month := time.Now().Format("01")
	return year,month
}
type SysInfo struct {
	ID primitive.ObjectID `bson:"_id"`
	*myDao.SysInfo `bson:",inline"`
}

func (si *SysInfo) GetID() interface{} {
	return si.ID
}

func (si *SysInfo) GetC() string {
	year,month := GetSysInfoCollection()
	sysInfoC := year+month
	return sysInfoC
}

func (si *SysInfo) GetDoc() interface{} {
	return si	
}

func (si *SysInfo)GetIndexes()[]mongo.IndexModel {
	return []mongo.IndexModel{}
}