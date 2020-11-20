## go-shorterer-webapp
simple url shorterer writen in golang using gin, gorm (postgres) and deployed at heroku.

## Motivation
I build this webapp to learn more about golang, dockerize and doploying to cloud service. I will keep improving this repo and adding more features which interest me to learn about it.  

## Build With
- [golang](https://golang.org/)
  - [gorm](https://gorm.io/index.html/)
  - [gin-gionic](https://github.com/gin-gonic/gin/)
- [postgresql](https://www.postgresql.org/)
- [heroku](https://www.heroku.com/)

## Try my code
```
$ git clone https://github.com/Henri1024/go-shorterer-webapp/
$ cd go-shorterer-webapp

$ docker-compose up --build
```

```
// create new api key
// this will send your api key in response, without verification
$ curl -d '{"email":"example@mail.com"}' -H "Content-Type: application/json" -X POST http://localhost:8888/api/newkey

// create new shortererlink
$ curl -d '{"destination_value":"https://google.com/"}' -H "Content-Type: application/json" -X POST http://localhost:8888/api/new?apikey=YOUR_API_KEY
```

## ScreenShot
### Create API key
![Create-API-key](https://github.com/Henri1024/go-shorterer-webapp/blob/master/screenshot/create_api_key.png)
### Create Shorterer url
![Create-Shorterer-url](https://github.com/Henri1024/go-shorterer-webapp/blob/master/screenshot/create_shorterer_url.png)
