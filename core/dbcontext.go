package core

import (
	"github.com/louisevanderlith/husk"
)

type context struct {
	Profiles husk.Tabler
}

var ctx context

func init() {
	defer createDefaultWebsite()

	ctx = context{
		Profiles: husk.NewTable(new(Profile)),
	}
}

func createDefaultWebsite() {
	if ctx.Profiles.Exists(husk.Everything()) {
		return
	}

	vosa := Profile{
		Title:        "avosa",
		ContactEmail: "info@avosa.co.za",
		Description:  "We're a software development company specialising in the Automotive and Transport industry. ",
		ContactPhone: "0893523423",
		URL:          "https://www.localhost",
		/*PortfolioItems: []Portfolio{
			Portfolio{
				Name: "Cars",
				URL:  "https://cars.localhost",
			},
			Portfolio{
				Name: "Admin",
				URL:  "https://admin.localhost",
			},
			Portfolio{
				Name: "Logbook",
				URL:  "https://logbook.localhost",
			},
			Portfolio{
				Name: "Shop",
				URL:  "https://shop.localhost",
			},
		},*/
		SocialLinks: []SocialLink{
			SocialLink{
				Icon: "fa-facebook",
				URL:  "https://www.facebook.com/avosasoftware",
			},
			SocialLink{
				Icon: "fa-twitter",
				URL:  "https://twitter.com/avosasoftware",
			},
		},
	}

	rec := vosa.Create()

	if rec.Error != nil {
		panic(rec.Error)
	}

	defer ctx.Profiles.Save()
}
