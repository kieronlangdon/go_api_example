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
#   "author":{"firstname":"Harry",  "lastnae":"White"}
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


```
