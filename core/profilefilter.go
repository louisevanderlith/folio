package core

import (
	"encoding/base64"
	"encoding/json"
	"github.com/louisevanderlith/husk"
	"log"
)

type profileFilter func(obj Profile) bool

func (f profileFilter) Filter(obj husk.Dataer) bool {
	return f(obj.(Profile))
}

func byName(name string) profileFilter {
	return func(obj Profile) bool {
		return obj.Title == name
	}
}

//base64 obj
func byHash(hsh string) profileFilter {
	if len(hsh) == 0 {
		return nil
	}

	profParm := Profile{}
	parm, err := base64.StdEncoding.DecodeString(hsh)

	if err != nil {
		log.Println(err)
		return nil
	}

	err = json.Unmarshal(parm, &profParm)

	if err != nil {
		log.Println(err)
		return nil
	}

	return func(obj Profile) bool {
		return (len(profParm.ClientID) == 0 || obj.ClientID == profParm.ClientID) && (len(profParm.Title) == 0 || obj.Title == profParm.Title)
	}

}
