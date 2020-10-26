package core

import (
	"github.com/louisevanderlith/husk/hsk"
)

type contentFilter func(obj Content) bool

func (f contentFilter) Filter(obj hsk.Record) bool {
	return f(obj.GetValue().(Content))
}

func byRealmClient(realm, client string) contentFilter {
	return func(obj Content) bool {
		return obj.Realm == realm && obj.Client == client
	}
}
