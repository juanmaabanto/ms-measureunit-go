package measureunit

import "github.com/sofisoft-tech/ms-measureunit/seedwork"

type MeasureUnit struct {
	CompanyId         string  `bson:"companyId"`
	MeasureTypeId     string  `bson:"measureTypeId"`
	Name              string  `bson:"name"`
	Symbol            string  `bson:"symbol"`
	Reference         bool    `bson:"reference"`
	Ratio             float64 `bson:"ratio"`
	Code              string  `bson:"code"`
	Active            bool    `bson:"active"`
	Canceled          bool    `bson:"canceled"`
	seedwork.Document `bson:"inline"`
}

func (_ MeasureUnit) GetCollectionName() string {
	return "measureUnit"
}
