# folio
Mango API: Folio

The Folio API controls profile information services.

## Run with Docker
* $ docker build -t avosa/folio:latest .
* $ docker rm FolioDEV
* $ docker run -d -e RUNMODE=DEV -p 8090:8090 -v db/:/db/ --network mango_net --name FolioDEV avosa/folio:latest 
* $ docker logs FolioDEV