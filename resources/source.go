package resources

import (
	"encoding/json"
	"fmt"
	"github.com/louisevanderlith/droxolite/drx"
	"net/http"
	"strings"
)

type Source struct {
	client *http.Client
	r      *http.Request
}

func APIResource(clnt *http.Client, r *http.Request) *Source {
	return &Source{
		client: clnt,
		r:      r,
	}
}

func (src *Source) get(container interface{}, api, path string, params ...string) error {
	tkninfo := drx.GetIdentity(src.r)
	url, err := tkninfo.GetResourceURL(api)

	if err != nil {
		return err
	}

	fullURL := fmt.Sprintf("%s/%s/%s", url, path, strings.Trim(strings.Join(params, "/"), "/"))

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	req.Header.Set("Authorization", "Bearer "+drx.GetToken(src.r))

	if err != nil {
		return err
	}

	resp, err := src.client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	return dec.Decode(container)
}
