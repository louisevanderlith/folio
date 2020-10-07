package resources

import (
	"github.com/louisevanderlith/folio/resources/models"
	"github.com/louisevanderlith/husk/records"
)

func (src *Source) FetchContent(key string) (interface{}, error) {
	var res interface{}
	err := src.get(&res, "cms", "content", key)
	return res, err
}

func (src *Source) FetchAllContent(pagesize string) (records.Page, error) {
	res := records.NewResultPage(models.Content{})
	err := src.get(res, "cms", "content", pagesize)

	if err != nil {
		return nil, err
	}

	return res, nil
}
