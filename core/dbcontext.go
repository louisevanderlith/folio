package core

import (
	"encoding/json"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/collections"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/op"
	"github.com/louisevanderlith/husk/records"
	"os"
	"reflect"
)

type FolioContext interface {
	GetContent(key hsk.Key) (Content, error)
	GetAllContent(page, size int) (records.Page, error)
	GetDisplay(realm, client string) (hsk.Record, error)
	CreateContent(o Content) (hsk.Key, error)
	UpdateContent(k hsk.Key, o Content) error
	Shutdown() error
}

type context struct {
	Content husk.Table
}

var ctx FolioContext

func CreateContext() FolioContext {
	c := context{
		Content: husk.NewTable(Content{}),
	}

	c.Seed()

	ctx = c
	return ctx
}

func Context() FolioContext {
	return ctx
}

func (c context) Seed() {
	contents, err := contentSeeds()

	if err != nil {
		panic(err)
	}

	err = c.Content.Seed(contents)

	if err != nil {
		panic(err)
	}
}

func contentSeeds() (collections.Enumerable, error) {
	f, err := os.Open("db/contents.seed.json")

	if err != nil {
		return nil, err
	}

	var items []Content
	dec := json.NewDecoder(f)
	err = dec.Decode(&items)

	if err != nil {
		return nil, err
	}

	return collections.ReadOnlyList(reflect.ValueOf(items)), nil
}

func (c context) Shutdown() error {
	return c.Content.Save()
}

func (c context) GetContent(key hsk.Key) (Content, error) {
	rec, err := c.Content.FindByKey(key)

	if err != nil {
		return Content{}, err
	}

	return rec.GetValue().(Content), nil
}

func (c context) GetAllContent(page, size int) (records.Page, error) {
	return c.Content.Find(page, size, op.Everything())
}

func (c context) GetDisplay(realm, client string) (hsk.Record, error) {
	return c.Content.FindFirst(byRealmClient(realm, client))
}

func (c context) CreateContent(o Content) (hsk.Key, error) {
	return c.Content.Create(o)
}
func (c context) UpdateContent(k hsk.Key, o Content) error {
	return c.Content.Update(k, o)
}
