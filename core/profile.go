package core

import "github.com/louisevanderlith/husk"

type Profile struct {
	Title          string       `hsk:"size(128)" json:",omitempty"`
	Description    string       `hsk:"size(512)" json:",omitempty"`
	ContactEmail   string       `hsk:"size(128)" json:",omitempty"`
	ContactPhone   string       `hsk:"size(20)" json:",omitempty"`
	URL            string       `hsk:"size(128)" json:",omitempty"`
	ImageKey       husk.Key     `hsk:"null"`
	SocialLinks    []SocialLink `json:",omitempty"`
	PortfolioItems []Portfolio  `json:",omitempty"`
	Headers        []Header     `json:",omitempty"`
}

func (p Profile) Valid() (bool, error) {
	return husk.ValidateStruct(&p)
}

func getProfile(key husk.Key) (husk.Recorder, error) {
	return ctx.Profiles.FindByKey(key)
}

func getProfileByName(name string) (husk.Recorder, error) {
	return ctx.Profiles.FindFirst(byName(name))
}

func GetProfile(key husk.Key) (*Profile, error) {
	rec, err := getProfile(key)

	if err != nil {
		return nil, err
	}

	return rec.Data().(*Profile), nil
}

func GetProfileByName(name string) (*Profile, error) {
	rec, err := getProfileByName(name)

	if err != nil {
		return nil, err
	}

	return rec.Data().(*Profile), nil
}

func GetProfiles(page, size int) husk.Collection {
	return ctx.Profiles.Find(page, size, husk.Everything())
}

func (p Profile) Create() husk.CreateSet {
	return ctx.Profiles.Create(p)
}

func (p Profile) Update(key husk.Key) error {
	profile, err := getProfile(key)

	if err != nil {
		return err
	}

	profile.Set(p)

	return ctx.Profiles.Update(profile)
}
