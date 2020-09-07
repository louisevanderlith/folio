package resources

import (
	"github.com/louisevanderlith/husk/records"
)

func (src *Source) FetchEntity(key string) (interface{}, error) {
	var res interface{}
	err := src.get(&res, "entity", "info", key)
	return res, err
}

func (src *Source) FetchEntities(pagesize string) (records.Page, error) {
	var res records.Page
	err := src.get(&res, "entity", "info", pagesize)

	if err != nil {
		return nil, err
	}

	return res, nil
}
