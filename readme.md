

## commands using while dev

### curl

Get Classes
    curl -v  localhost:9090/classes | jq

Classes request with body becomes post 
    curl -v  localhost:9090/classes -d '{"name":"jira101", "capacity":20, "startDate": "2021-03-25", "endDate": "2021-04-12"}' | jq

Updating a class using put request
    curl -v  localhost:9090/ff06487d-4f5a-41a8-a3d5-5319d9305110 -XPUT -d '{"name":"home old workout", "capacity":5, "startDate": "2021-03-25", "endDate": "2021-04-12"}' | jq