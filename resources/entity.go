package resources

import (
	"github.com/louisevanderlith/folio/resources/models"
	"github.com/louisevanderlith/husk/records"
)

func (src *Source) FetchEntity(key string) (interface{}, error) {
	var res interface{}
	err := src.get(&res, "entity", "", key)
	return res, err
}

func (src *Source) FetchEntities(pagesize string) (records.Page, error) {
	res := records.NewResultPage(models.Entity{})
	err := src.get(res, "entity", "", pagesize)

	if err != nil {
		return nil, err
	}

	return res, nil
}
