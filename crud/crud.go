package crud

import (
	"LWRworkshop/client"
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

//Insert will insert a document into the database
func Insert(collection string, data interface{}, w http.ResponseWriter) (*mongo.InsertOneResult, error) {

	col := client.GetCollection(collection)
	result, err := col.InsertOne(context.TODO(), data)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	return result, err

}

//GetOne will get one item from the database
func GetOne(collection, ID string, w http.ResponseWriter) ([]byte, error) {
	IDobj, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return []byte{}, err
	}

	filter := bson.M{
		"_id": IDobj,
	}

	dataRaw := bson.M{}

	col := client.GetCollection(collection)

	err = col.FindOne(context.TODO(), filter).Decode(&dataRaw)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	data, _ := json.Marshal(dataRaw)

	return data, err
}
