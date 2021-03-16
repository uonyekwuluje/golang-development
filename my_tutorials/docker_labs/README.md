## Docker Application Resource
* [Golang Docker Image](https://hub.docker.com/_/golang?tab=tags&page=1&ordering=last_updated)

## Docker Image and Container Cleanup
```
docker stop $(docker ps -aq)
docker rmi --force $(docker images -q)
docker rm $(docker ps -aq)
```

## Intro Golang Web Dev
```
docker build -t introweb .
docker run --name introweb -d -p 8080:8080 introweb
```
