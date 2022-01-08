package measuretype

import "github.com/sofisoft-tech/ms-measureunit/seedwork"

type MeasureType struct {
	Name              string `bson:"name"`
	seedwork.Document `bson:"inline"`
}

func (_ MeasureType) GetCollectionName() string {
	return "measureType"
}
