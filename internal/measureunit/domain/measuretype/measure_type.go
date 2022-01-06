package measuretype

import (
	"github.com/sofisoft-tech/ms-measureunit/seedwork"
)

type MeasureType struct {
	Name              string `bson:"name"`
	seedwork.Document `bson:"inline"`
}

func (mt *MeasureType) GetCollectionName() string {
	return "measureType"
}
