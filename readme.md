

## commands using while dev

### curl

Get Classes
    curl -v  localhost:9090/classes | jq

Classes request with body becomes post 
    curl -v  localhost:9090/classes -d '{"name":"jira101", "capacity":20, "startDate": "2021-03-25", "endDate": "2021-04-12"}' | jq