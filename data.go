package mongo

import (
	"testing"

	"github.com/pkg/errors"
)

// MongoData is a structure used to insert massively data in a MongoDB
type MongoData struct {
	Database   string
	Collection string
	Datas      []interface{}
}

// Insert will massively nsert MongoData in a MongoDatabase
func Insert(mongo Session, datas []Test_MongoData) error {
	for _, data := range datas {
		t.Log("Inserting data: " + data.Database + " " + data.Collection)
		collection := mongo.DB(data.Database).C(data.Collection)
		for _, toInsert := range data.Datas {
			t.Logf("Inserting data - %+v", toInsert)
			err := collection.Insert(toInsert)
			if nil != err {
				return errors.Wrapf(err, "Could not insert %+v")
			}
		}
	}
}

// Clean will clean a Mongo Instance of all it's databases
func Clean(t testing.TB, mongo Session) error {
	names, err := mongo.DatabaseNames()
	if err != nil {
		return errors.Wrapf(err, "Loading databases names")
	}
	for _, name := range names {
		if err = mongo.DB(name).DropDatabase(); nil != err {
			return errors.Wrapf(err, "Droping DB %s", name)
		}
	}
}
