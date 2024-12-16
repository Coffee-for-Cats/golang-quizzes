# Quizzes app
This was intended to be a "Would you rather" web app, but I turned it into a "quiz" app.
- Every player needs to register.
- You can play without being registered, but *your* score will not be tracked.
- Status 201 for correct answer, 202 for wrong.

Database schemas in ./script/set-database/set.go

## Routes
See routes.http for more details and tests.
See main.go for the implementation.

## Commands:
setup:
- Start postgres and update the connection string:
`sudo docker compose up -d` To start postgres on 
`docker ps -a` > `go inspect <id>` > Update IpAddres in database/connect.go.
- Run the application:
```bash
go mod tidy
go run scripts/set-database/set.go
go run .
```

manage db:
- psql "postgresql://fluffycat@172.19.0.2:5432/heavycake?sslmode=disable"
- password is *s3cret* in my case.
- remember to change the ip ;)