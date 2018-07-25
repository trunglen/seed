package category

import (
	"rocky-springs-86767/g/x/web"
	"rocky-springs-86767/x/logger"
	"rocky-springs-86767/x/mongodb"
	"rocky-springs-86767/x/validator"
)

var categoryLog = logger.NewLogger("tbl_category")
var categoryTable = mongodb.NewTable("category", "category")

type Category struct {
	mongodb.Model `bson:",inline"`
	Name          string `bson:"name" json:"name" validate:"string"`
}

func (cat *Category) Create() error {
	err := validator.Struct(cat)
	if err != nil {
		categoryLog.Error(err)
		return web.WrapBadRequest(err, "")
	}
	return categoryTable.Create(cat)
}

func DeleteCategoryByID(id string) error {
	return categoryTable.DeleteByID(id)
}
