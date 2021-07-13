## Restful Documents
* https://tutorialedge.net/golang/creating-restful-api-with-golang/
* https://devqa.io/curl-sending-api-requests/
* https://blog.golang.org/json
* https://gobyexample.com/json
* https://developpaper.com/golang-uses-json-format-to-add-delete-query-and-modify/
* https://tutorialedge.net/golang/parsing-json-with-golang/
* https://tutorialedge.net/golang/creating-restful-api-with-golang/  [*****]


## List All
```
curl -s localhost:10000/all | jq .
```

## Add Record
```
curl -H "Content-Type: application/json" --data @body.json localhost:10000/article
```

## Delete Record:
```
curl -X DELETE http://localhost:10000/delete/2
```
