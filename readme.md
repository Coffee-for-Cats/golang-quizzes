# Quizes app
This was intended to be a "Would you rather" web app, but I turned it into a "quiz" app.
- Every player eeeds to register.
- You can play without being registered, but your score will not be tracked.
- Status 201 for correct answer, 202 for wrong.

Database schemas in ./script/set-database/set.go

## Routes
See main.go for more details.
- GET /quiz/{id}
- GET /quiz/random
- POST /register  | body: *username*, *password*.
- POST /quiz      | body: *title* | auth needed
- POST /questions | body: *question*, *correct*, *wrong*, *quiz_id* | auth needed
- POST /answer    | body: *question_id*, *answer_text* | auth optional

## Commands I am using:
setup:
- sudo docker compose up -d
- go run .
manage db:
- psql -h 172.19.0.2 -p 5432 -d heavycake -U fluffycat -W 