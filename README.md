# GoTask (Golang)
A simple REST API for studies made in *Golang* simulating a control of tasks to do.

- 📚 gorilla/mux
- 📚 joho/godotenv
- 📚 lib/pq

## Structure

```
Task {
	UID : uint32
	Description : string
	UserID : string
	CreatedAt : string
	Closed : boolean
}
```

## Endpoints

- (GET) tasks
  - Headers:
    - UserId
___

- (POST) tasks
  - Headers:
    - UserId
  
  - Body:
    - description: string
___

- (DELETE) tasks/:id
  - Headers:
    - UserId

  - Params:
    - id
___

- (PATCH) tasks/:id
  - Headers:
    - UserId
    
  - Params:
    - id

  - Body:
    - closed: boolean
___
