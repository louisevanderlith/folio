package core

import "github.com/louisevanderlith/husk"

type Profile struct {
	ClientID       string       `hsk:"size(48)"`
	Title          string       `hsk:"size(128)" json:",omitempty"`
	Description    string       `hsk:"size(512)" json:",omitempty"`
	ContactEmail   string       `hsk:"size(128)" json:",omitempty"`
	ContactPhone   string       `hsk:"size(20)" json:",omitempty"`
	URL            string       `hsk:"size(128)" json:",omitempty"`
	ImageKey       husk.Key     `hsk:"null"`
	GTag           string       `hsk:"size(14)"`
	SocialLinks    []SocialLink `json:",omitempty"`
	PortfolioItems []Portfolio  `json:",omitempty"`
	Headers        []Header     `json:",omitempty"`
}

func (p Profile) Valid() (bool, error) {
	return husk.ValidateStruct(&p)
}

func GetProfile(key husk.Key) (Profile, error) {
	rec, err := ctx.Profiles.FindByKey(key)

	if err != nil {
		return Profile{}, err
	}

	return rec.Data().(Profile), nil
}

func GetProfileByName(name string) (Profile, error) {
	rec, err := ctx.Profiles.FindFirst(byName(name))

	if err != nil {
		return Profile{}, err
	}

	return rec.Data().(Profile), nil
}

func GetAllProfiles(page, size int) husk.Collection {
	return ctx.Profiles.Find(page, size, husk.Everything())
}

func GetProfiles(page, size int, hsh string) husk.Collection {
	if len(hsh) == 0 {
		return GetAllProfiles(page, size)
	}

	return ctx.Profiles.Find(page, size, byHash(hsh))
}

func (p Profile) Create() husk.CreateSet {
	return ctx.Profiles.Create(p)
}

func (p Profile) Update(key husk.Key) error {
	profile, err := ctx.Profiles.FindByKey(key)

	if err != nil {
		return err
	}

	err = profile.Set(p)

	if err != nil {
		return err
	}

	err = ctx.Profiles.Update(profile)

	if err != nil {
		return err
	}

	return ctx.Profiles.Save()
}
