package utils

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MakeFilter(filters []string) (bsonFilters bson.M) {
	bsonFilters = bson.M{}
	for _, v := range filters {
		slFilter := strings.Split(v, ":")
		key := slFilter[0]
		operations := slFilter[1]

		switch operations {
		case "ne":
			bsonFilters[key] = bson.M{"$ne": slFilter[2]}
			break
		case "like":
			bsonFilters[key] = bson.M{
				"$regex":   slFilter[2],
				"$options": "i",
			}
			break
		case "eq":
			bsonFilters[key] = slFilter[2]
			break
		case "eqInt":
			bsonFilters[key] = slFilter[2]
			break
		case "isNull":
			bsonFilters[key] = nil
			break
		case "isNotNull":
			bsonFilters[key] = bson.M{"$ne": nil}
			break
		case "id":
			oid, _ := primitive.ObjectIDFromHex(slFilter[2])
			bsonFilters[key] = oid
			break
		default:
			bsonFilters[key] = slFilter[2]
			break
		}
	}

	return bsonFilters
}
