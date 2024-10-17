# Quizzes app
This was intended to be a "Would you rather" web app, but I turned it into a "quiz" app.
- Every player needs to register.
- You can play without being registered, but *your* score will not be tracked.
- Status 201 for correct answer, 202 for wrong.

Database schemas in ./script/set-database/set.go

## Routes
See main.go for more details.
- GET /quiz/{id}
- GET /quiz/random
- POST /register  | body: *username*, *password*.
- POST /quiz      | body: *title*                                   | auth needed
- POST /questions | body: *question*, *correct*, *wrong*, *quiz_id* | auth needed
- POST /answer    | body: *question_id*, *answer_text*              | auth optional

## Commands:
setup:
- `sudo docker compose up -d`
- Get the container ip and change it in the database string in /database/connect.go
- `go run .`

manage db:
- psql "postgresql://fluffycat@172.19.0.2:5432/heavycake?sslmode=disable"
- password is *s3cret* in my case.
- remember to change the ip ;)