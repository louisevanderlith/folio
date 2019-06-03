# folio
Mango API: Folio

The Folio API controls profile information services.

## Run with Docker
* $ docker build -t avosa/folio:dev .
* $ docker rm FolioDEV
* $ docker run -d -p 8090:8090 -v db/:/db/ --network mango_net --name FolioDEV avosa/folio:dev 
* $ docker logs FolioDEV

### Folio API Exposes the following Endpoints
**"/v1/profile" POST**
Posting to Profile will attempt to create a new database entry
```
url: /v1/profile
body: {
    Title:          "string"
	Description:    "string"
	ContactEmail    "string" 
	ContactPhone    "string"  
	URL:            "string"
	ImageKey:       husk.Key {0`0}
	SocialLinks:    []SocialLink{}
	PortfolioItems: []Portfolio{}
	Headers:        []Header{} 
}
response: {

}
```

**"/v1/profile" PUT**
Putting it to Profile will attempt to update an existing database entry
```
url: /v1/profile
body: {
	Key: '1554125525`0',
	Body: {
		Title:          "string"
		Description:    "string"
		ContactEmail    "string" 
		ContactPhone    "string"  
		URL:            "string"
		ImageKey:       husk.Key {0`0}
		SocialLinks:    []SocialLink{}
		PortfolioItems: []Portfolio{}
		Headers:        []Header{} 
	}
}
response: {

}
```

**"/v1/profile/:site" GET**
Getting a Profile by Title OR Key will attempt to find a matching database entry
```
url: /v1/profile/1554125525`0 [By Key]
url: /v1/profile/mysite [By Title]
body: {}
response: {

}
```

**"/v1/profile/all/:pagesize" GET**
Getting all Profiles will return the specified amount, if available, of entries in the database
```
url: /v1/profile/all/A10 [Page A, 10 Results]
body: {}
response: {
    
}
```