package store

import (
	"github.com/golang/glog"
	"labix.org/v2/mgo/bson"
)

const (
	OFFLINE = 11
	ONLINE  = 12
)

type Msg struct {
	Msg_id string // Msg ID
	Body   []byte
	Typ    int

	Owner string // Owner
}

func GetOfflineMsg(mID string, ch chan<- *Msg) {
	defer recover()
	// Find in the db
	sei := sei_msg.New()
	defer sei.Refresh()
	c := sei.DB(Config.MsgName).C(Config.OfflineName)
	iter := c.Find(bson.M{"Msg_id": mID}).Limit(Config.OfflineMsgs).Iter()
	defer iter.Close()
	msg := new(Msg)
	for iter.Next(msg) {
		ch <- msg
		msg = new(Msg)
	}
}

// Del the offile msg
func DelOfflineMsg(msg_id string, id string) {
	c := sei_msg.DB(Config.MsgName).C(Config.OfflineName)
	defer sei_msg.Refresh()
	err := c.Remove(bson.M{"Msg_id": msg_id, "Owner": id})
	if err != nil {
		glog.Errorf("Remove a offline msg(id:%v,Owner:%v) error:%v", msg_id, id, err)
	}
}