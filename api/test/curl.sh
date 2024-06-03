## Users Login
    curl -d '{"email":"jhon@example.com", "password": "12345"}' -H "Content-Type: application/json" -X POST http://localhost:8080/login
    curl -d '{"email":"jhon.doe@example.com", "password": "12345"}' -H "Content-Type: application/json" -X POST http://localhost:8080/login

## Users

    # CREATE
    curl -d '{"email":"jhon@example.com", "password": "12345"}' -H "Content-Type: application/json" -X POST http://localhost:8080/signup
    curl -d '{"email":"jhon.doe@example.com", "password": "12345"}' -H "Content-Type: application/json" -X POST http://localhost:8080/signup

## Events

    # CREATE
    curl -d '{"name":"Jhon", "description":"test event", "location": "test location", "datetime": "2024-05-05T05:05:05.000Z"}' \
        -H "Content-Type: application/json" \
        -H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Impob25AZXhhbXBsZS5jb20iLCJleHAiOjE3MTc0NTg5OTIsInVzZXJJZCI6MH0.siYg2H78FehK0kcpwsiOXQPn01ZmIP9A2TN385WR6pY" \
        -X POST http://localhost:8080/events

    curl -d '{"name":"Jhon", "description":"test event", "location": "test location", "datetime": "2024-05-05T05:05:05.000Z"}' \
    -H "Content-Type: application/json" \
    -H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Impob24uZG9lQGV4YW1wbGUuY29tIiwiZXhwIjoxNzE3NDU5MzA5LCJ1c2VySWQiOjB9.D7_8FgYVbf5hkv2M_FG0Fgyxp08W3ZI2jS5Eo_ztgi0" \
    -X POST http://localhost:8080/events

    # UPDATE
    curl -d '{"name":"Jhon Updated", "description":"test event updated", "location": "location", "datetime": "2024-05-05T05:05:05.000Z"}' -H "Content-Type: application/json" -X PUT http://localhost:8080/events/1

    # DELETE
    curl -X DELETE http://localhost:8080/events/1