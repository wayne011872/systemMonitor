package mon

import(
	myDao "github.com/wayne011872/systemMonitor/dao"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"strconv"
)

func GetErrorSysInfoCollection()(string,string){
	year := strconv.Itoa(time.Now().Year())
	month := time.Now().Format("01")
	return year,month
}

type ErrorSysInfo struct {
	ID primitive.ObjectID `bson:"_id"`
	*myDao.ErrorSysInfo `bson:"inline,omitempty"`
}

func (esi *ErrorSysInfo) GetID() interface{} {
	return esi.ID
}

func (esi *ErrorSysInfo) GetC() string {
	year,month := GetSysInfoCollection()
	sysInfoC := year+month
	return sysInfoC
}

func (esi *ErrorSysInfo) GetDoc() interface{} {
	return esi	
}

func (eesi *ErrorSysInfo)GetIndexes()[]mongo.IndexModel {
	return []mongo.IndexModel{}
}