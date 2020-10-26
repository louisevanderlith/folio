package api

import (
	"encoding/json"
	"fmt"
	"github.com/louisevanderlith/folio/core"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/records"
	"net/http"
)

func FetchContent(web *http.Client, host string, k hsk.Key) (core.Content, error) {
	url := fmt.Sprintf("%s/content/%s", host, k.String())
	resp, err := web.Get(url)

	if err != nil {
		return core.Content{}, err
	}

	defer resp.Body.Close()

	result := core.Content{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}

func FetchAllContent(web *http.Client, host, pagesize string) (records.Page, error) {
	url := fmt.Sprintf("%s/content/%s", host, pagesize)
	resp, err := web.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	result := records.NewResultPage(core.Content{})
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}
