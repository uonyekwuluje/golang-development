## Restful Documents
* https://tutorialedge.net/golang/creating-restful-api-with-golang/

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
