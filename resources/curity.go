package resources

import (
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/husk/records"
	"github.com/louisevanderlith/kong/prime"
)

func (src *Source) FetchSecurityReport(key string) (interface{}, error) {
	var res interface{}
	err := src.get(&res, "secure", "report", key)
	return res, err
}

func (src *Source) FetchSecurityReports(pagesize string) (records.Page, error) {
	res := records.NewResultPage(nil)
	err := src.get(res, "secure", "report", pagesize)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (src *Source) FetchProfile(key string) (prime.Profile, error) {
	res := prime.Profile{ImageKey: &keys.TimeKey{}}
	err := src.get(&res, "secure", "profiles", key)
	return res, err
}

func (src *Source) FetchProfiles(pagesize string) (records.Page, error) {
	res := records.NewResultPage(prime.Profile{ImageKey: &keys.TimeKey{}})
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
	res := records.NewResultPage(prime.User{})
	err := src.get(res, "secure", "users", pagesize)

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
	res := records.NewResultPage(prime.Resource{})
	err := src.get(res, "secure", "resources", pagesize)

	if err != nil {
		return nil, err
	}

	return res, nil
}
