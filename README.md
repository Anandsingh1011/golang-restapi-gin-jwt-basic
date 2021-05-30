# Golang Gin Framework - Video API (

## Run

```bash
go run server.go
```

## Endpoint  

```Endpoint 
POST   /login
GET    /api/videos
POST   /api/videos 
```

### Login Details 

```Login Details 
Username : vuser
Password : newpassword
```

### Login  

```login 
POST   /login

Request
---


curl --location --request POST 'http://localhost:5000/login' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'username=vuser' \
--data-urlencode 'password=newpassword'



Response
---

{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidnVzZXIiLCJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjIyNjAzMzg5LCJpYXQiOjE2MjIzNDQxODksImlzcyI6Im5ldy5jb20ifQ.FEDoyhSduQPua23nyXfRTBE3z47u8W6ks_NSwtig1n4"
}

```


### Add New Video  

```addvideo 
POST   /api/videos

Request
---


curl --location --request POST 'http://localhost:5000/api/videos' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidnVzZXIiLCJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjIyNjAzMzg5LCJpYXQiOjE2MjIzNDQxODksImlzcyI6Im5ldy5jb20ifQ.FEDoyhSduQPua23nyXfRTBE3z47u8W6ks_NSwtig1n4' \
--data-raw '{
    "title": "Golang title1",
    "description": "Some description",
    "url": "http://name.com",
    "author": {
      "firstname": "Anand",
      "lastname": "Singh",
      "age": 20,
      "email": "anand@anand.com"
    }
}'



Response
---

{
    "message": "Video Input is Valid!!"
}

```


### Get All Videos

```getvideo 
GET   /api/videos

Request
---
curl --location --request GET 'http://localhost:5000/api/videos' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidnVzZXIiLCJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjIyNjAzMzg5LCJpYXQiOjE2MjIzNDQxODksImlzcyI6Im5ldy5jb20ifQ.FEDoyhSduQPua23nyXfRTBE3z47u8W6ks_NSwtig1n4'



Response
---

[
    {
        "title": "Golang title2",
        "description": "Some description",
        "url": "http://name.com",
        "author": {
            "firstname": "Anand",
            "lastname": "Singh",
            "age": 20,
            "email": "anand@anand.com"
        }
    },
    {
        "title": "Golang title2",
        "description": "Some description",
        "url": "http://name.com",
        "author": {
            "firstname": "Anand",
            "lastname": "Singh",
            "age": 20,
            "email": "anand@anand.com"
        }
    },
    {
        "title": "Golang title2",
        "description": "Some description",
        "url": "http://name.com",
        "author": {
            "firstname": "Anand",
            "lastname": "Singh",
            "age": 20,
            "email": "anand@anand.com"
        }
    },
    {
        "title": "Golang title1",
        "description": "Some description",
        "url": "http://name.com",
        "author": {
            "firstname": "Anand",
            "lastname": "Singh",
            "age": 20,
            "email": "anand@anand.com"
        }
    }
]

```
