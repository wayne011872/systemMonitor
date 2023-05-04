package sysInfo

import (
	"context"
	"github.com/wayne011872/systemMonitor/dao"
	"github.com/wayne011872/systemMonitor/dao/mon"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewCRUD(ctx context.Context, db *mongo.Database) CRUD {
	return &mongoCRUD{
		db: db,
		ctx:ctx,
	}
}

type CRUD interface {
	Save(*dao.SysInfo) error
}

type mongoCRUD struct {
	db *mongo.Database
	ctx context.Context
}

func (m *mongoCRUD) Save(s *dao.SysInfo) error{
	o := &mon.SysInfo{
		ID: primitive.NewObjectID(),
		SysInfo: s,
	}
	collection := m.db.Collection(o.GetC())
	_,err := collection.InsertOne(m.ctx,o.GetDoc())
	if err != nil {
		return err
	}
	return err
}