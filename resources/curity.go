package resources

func (src *Source) FetchSecurityReport(key string) (interface{}, error) {
	return src.get("secure", "report", key)
}

func (src *Source) FetchSecurityReports(pagesize string) (map[string]interface{}, error) {
	res, err := src.get("secure", "report", pagesize)

	if err != nil {
		return nil, err
	}

	return res.(map[string]interface{}), nil
}

func (src *Source) FetchProfile(key string) (interface{}, error) {
	return src.get("secure", "profiles", key)
}

func (src *Source) FetchProfiles(pagesize string) (map[string]interface{}, error) {
	res, err := src.get("secure", "profiles", pagesize)

	if err != nil {
		return nil, err
	}

	return res.(map[string]interface{}), nil
}

func (src *Source) FetchUser(key string) (interface{}, error) {
	return src.get("secure", "users", key)
}

func (src *Source) FetchUsers(pagesize string) (map[string]interface{}, error) {
	res, err := src.get("secure", "users", pagesize)

	if err != nil {
		return nil, err
	}

	return res.(map[string]interface{}), nil
}

func (src *Source) FetchResource(key string) (interface{}, error) {
	return src.get("secure", "resources", key)
}

func (src *Source) FetchResources(pagesize string) (map[string]interface{}, error) {
	res, err := src.get("secure", "resources", pagesize)

	if err != nil {
		return nil, err
	}

	return res.(map[string]interface{}), nil
}
