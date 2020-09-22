# go_api_example

Files for course from https://github.com/bradtraversy/go_restapi
Modified since then!

 Simple GO Lang REST API

> Simple RESTful API to create, read, update and delete books. No database implementation yet

## Quick Start


``` bash
# Install mods
go get -u github.com/gorilla/mux
go get -u github.com/go-resty/resty
go get -u github.com/stretchr/testify/assert
go get -u github.com/stretchr/testify/suite
```

``` bash
go build
./go_restapi
```

## Endpoints

### Get All Books
``` bash
GET api/books
```
### Get Single Book
``` bash
GET api/books/{id}
```

### Delete Book
``` bash
DELETE api/books/{id}
```

### Delete all Books
``` bash
DELETE api/books
```

### Create Book
``` bash
POST api/books

# Request sample
# {
#   "isbn":"4545454",
#   "title":"Book Three",
#   "author":{"firstname":"Harry",  "lastname":"White"}
# }
```

### Update Book
``` bash
PUT api/books/{id}

# Request sample
# {
#   "isbn":"4545454",
#   "title":"Updated Title",
#   "author":{"firstname":"Harry",  "lastname":"White"}
# }

```
### Build steps
``` bash
Set Paths:
export GOPATH=/Users/$USER/go
export PATH=$PATH:$GOPATH/bin;

From:
$GOPATH/src/github.com/$USER/restapi
docker build -t my-go-app .

Enter container using Almquist shell after starting:
docker run -it --rm my-go-app  /bin/ash
Run container in detached mode
docker run -p 8080:8001 -d my-go-app
Run Unit tests only
go test -v -run Unit
Run tests with coverage report
go test -coverprofile=c.out
See report in browser (generates html file)
go tool cover -html=c.out -o coverage.html

```
### Minikube setup
```
minikube delete
minikube start
eval $(minikube docker-env)
```
### Minikube steps for error:   
https://registry-1.docker.io/v2 connection error while pulling image   
`minikube ssh`   
`sudo vi /etc/systemd/network/10-eth1.network` add `DNS=8.8.8.8` under [Network]   
`sudo vi /etc/systemd/network/20-dhcp.network` add `DNS=8.8.8.8` under [Network]   
`sudo systemctl restart systemd-networkd`    
### Local registry steps:   
**Start registry**   
`docker run -d -p 5000:5000 --restart=always --name registry registry:2`   
**Build image**   
`docker build -t my-go-app .`   
**Tag image**   
`docker tag my-go-app:latest localhost:5000/my-go-app`   
**Push to local registry**   
`docker push localhost:5000/my-go-app`   
**Remove image local - does not remove from registry**   
`docker image remove localhost:5000/my-go-app`   
### Create deployment in minikube
`minikube kubectl create deployment testdev -- --image=localhost:5000/my-go-app`   

### Alt steps for runnning minikube and to avoid going down a rabbit hole of nonsense   
**Build image**   
`docker build -t my-go-app .`   
**Tag image**   
`docker tag my-go-app:latest localhost:5000/my-go-app`   
**Point your terminal to use the docker daemon inside minikube**   
`eval $(minikube docker-env)`   
**Check to see all images**      
`docker images`   
Notice that the image we just created and tagged is not here? We're using minikube's docker daemon now   
**Push to cache in minikube**   
`minikube cache add localhost:5000/my-go-app`   
This will fail as the image doesn't exist in this daemon   
So in another terminal OR after undoing `eval $(minikube docker-env)`   
Re-run `minikube cache add localhost:5000/my-go-app`   
**Check image is in minikube cache**   
`minikube cache list`   
We will see an image with no repo & tag, check the imageId, it's the same as the local image in non minikube docker   
**Tag image in minikube docker daemon**   
`docker tag c22dbba37091 localhost:5000/my-go-app`   
### Create deployment in minikube   
`minikube kubectl create deployment testdev -- --image=localhost:5000/my-go-app`  
### Alter deployment yaml to pull from minikube local   
Get deployment yaml   
`minikube kubectl get deploy testdev -- -o yaml --export >> testdev.yaml`   
In testdev.yaml   
`imagePullPolicy: Always`   
change to   
`imagePullPolicy: Never`   
Apply yaml   
`minikube kubectl apply -- -f testdev.yaml`   
### Expose container port using service   
Set deployment to expose deployment of type node port   
`minikube kubectl expose deployment testdev -- --type=NodePort`   
Get port that has been exposed externally   
` minikube kubectl get svc testdev`   
Get minikube ip   
`minikube ip`   
Test interaction with cluster   
`curl 192.168.50:32145/api/books`     
