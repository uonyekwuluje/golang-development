## Docker Image and Container Cleanup
```
docker stop $(docker ps -aq)
docker rmi --force $(docker images -q)
docker rm $(docker ps -aq)
```

## Docker Login (in-progress)
```
docker login
```


## Intro Golang Web Dev
```
docker build -t introweb .
docker run --name introweb -d -p 8080:8080 introweb
```

## Test App
```
http://localhost:8080/
```

## Tag and Upload
```
docker tag introweb cicdpoc/introweb:poc
docker push cicdpoc/introweb:poc
```
