# CREATE
curl -d '{"name":"Jhon", "description":"test event", "location": "test location", "datetime": "2024-05-05T05:05:05.000Z"}' -H "Content-Type: application/json" -X POST http://localhost:8080/events

# UPDATE
curl -d '{"name":"Jhon Updated", "description":"test event updated", "location": "location", "datetime": "2024-05-05T05:05:05.000Z"}' -H "Content-Type: application/json" -X PUT http://localhost:8080/events/1

# DELETE
curl -X DELETE http://localhost:8080/events/1