package resources

import (
	"github.com/louisevanderlith/husk/records"
	"github.com/louisevanderlith/kong/prime"
)

func (src *Source) FetchSecurityReport(key string) (interface{}, error) {
	var res interface{}
	err := src.get(&res, "secure", "report", key)
	return res, err
}

func (src *Source) FetchSecurityReports(pagesize string) (records.Page, error) {
	var res records.Page
	err := src.get(&res, "secure", "report", pagesize)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (src *Source) FetchProfile(key string) (interface{}, error) {
	var res interface{}
	err := src.get(&res, "secure", "profiles", key)
	return res, err
}

func (src *Source) FetchProfiles(pagesize string) (records.Page, error) {
	res := records.NewResultPage(prime.Profile{})
	err := src.get(res, "secure", "profiles", pagesize)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (src *Source) FetchUser(key string) (interface{}, error) {
	var res interface{}
	err := src.get(&res, "secure", "users", key)
	return res, err
}

func (src *Source) FetchUsers(pagesize string) (records.Page, error) {
	var res records.Page
	err := src.get(&res, "secure", "users", pagesize)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (src *Source) FetchResource(key string) (interface{}, error) {
	var res interface{}
	err := src.get(&res, "secure", "resources", key)
	return res, err
}

func (src *Source) FetchResources(pagesize string) (records.Page, error) {
	var res records.Page
	err := src.get(&res, "secure", "resources", pagesize)

	if err != nil {
		return nil, err
	}

	return res, nil
}
