package mongo

import (
	"net"
	"testing"

	"github.com/normegil/docker"
	"github.com/normegil/zookeeper-rest/modules/test"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
)

const mongoInternalPort = 27017
const MONGO_PORTS string = "[50017;50037]"

// Test_NewMongoSession will create a session connected to a mongo instance inside a docker. For Test only.
func Test_NewMongoSession(t testing.TB) (Session, func()) {
	mongoInfo, closeFn := test.NewMongo(t)
	host := net.TCPAddr{mongoInfo.Address, mongoInfo.Port, ""}
	session, err := mgo.Dial(host.String())
	if nil != err {
		defer closeFn()
		t.Fatal(errors.Wrap(err, "Could not create a new mgo session {Host:"+host.String()+"}"))
	}
	return NewSession(session, ""), func() {
		session.Close()
		closeFn()
	}
}

// Test_NewMongo will create a new mongo instance inside a docker and return connection infos. For Test only.
func Test_NewMongo(t testing.TB) (MongoInfo, func()) {
	mainPortBinding := docker.PortBinding{"tcp", mongoInternalPort, MONGO_PORTS}
	info, close := NewDocker(t, docker.Options{
		Name:  "MongoDB",
		Image: "mongo:latest",
		Ports: []docker.PortBinding{mainPortBinding},
	})
	return MongoInfo{
		Address: info.Address,
		Port:    info.Ports[mainPortBinding],
	}, close
}

type MongoInfo struct {
	Address net.IP
	Port    int
}
