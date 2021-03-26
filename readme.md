# Instructions
I have developed this repo on windows10 with wsl2. Hopefully everthing should work didn't had time to do docker or some other fancy hexagonal structure stuff.



## Important Endpoints

## docs
put swagger and react docs together one for endping for now. time was the limitation.
http://localhost:9090/docs

## Create Class
http://localhost:9090/classes
POST {"name":"jira1887", "capacity":20, "startDate": "2021-05-10", "endDate": "2021-06-15"}

will check for startdate is not later then end date. Also will check date clahes with other classes in data store.

## Create Booking
http://localhost:9090/bookings
POST {"clientName":"tom", "bookingDate": "2021-03-25", "classId"="2aac3e20-2f2f-4afa-ab58-c08f2c175f4a"} 

check for classid is matching with any class in datastore.


## commands using while dev

### curl

Get Classes
    curl -v  localhost:9090/classes | jq

Classes request with body becomes post 
    curl -v  localhost:9090/classes -d '{"name":"jira101", "capacity":20, "startDate": "2021-03-25", "endDate": "2021-04-12"}' | jq

Updating a class using put request
    curl -v  localhost:9090/ff06487d-4f5a-41a8-a3d5-5319d9305110 -XPUT -d '{"name":"home old workout", "capacity":5, "startDate": "2021-03-25", "endDate": "2021-04-12"}' | jq

curl -v  localhost:9090/bookings -d '{"clientName":"tom", "bookingDate": "2021-03-25", "classId"="2aac3e20-2f2f-4afa-ab58-c08f2c175f4a"}' | jq

## Tasks

* load deata from envoirment like port and running mode (dev, prod)
* dockerize it
* some tests
* other model incluing booking POST
* class check for clashes with other class dates
* allow CORS