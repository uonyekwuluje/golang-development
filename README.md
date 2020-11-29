# golang-development
Go Development

## Install Go
Install GO In linux
```
GOVERSION="1.15.5"
sudo wget https://dl.google.com/go/go${GOVERSION}.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go${GOVERSION}.linux-amd64.tar.gz
sudo rm go${GOVERSION}.linux-amd64.tar.gz
```
Test and verify installation:
```
/usr/local/go/bin/go version
```
You should see
```
go version go1.15.5 linux/amd64
```

## Run Go Code
Test Sample Code
```
go run hello.go
```

## Build Go Code
Test Sample Code
```
go build hello.go
```
after the build, execute
```
./build
```
