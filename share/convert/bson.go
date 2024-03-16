package convert

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
)

func BsonArray(data any) bson.D {
	bd := bson.D{}
	typeOf := reflect.TypeOf(data)
	valueOf := reflect.ValueOf(data)

	for i := 0; i < typeOf.Elem().NumField(); i++ {
		key := typeOf.Elem().Field(i).Name
		value := valueOf.Elem().Field(0)
		e := bson.E{Key: key, Value: value}
		bd = append(bd, e)
	}

	return bd
}
