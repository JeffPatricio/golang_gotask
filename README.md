# GoTask (Golang)
A simple REST API for studies made in *Golang* simulating a control of tasks to do.

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
  - Headers
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
    - id: string
___

- (PATCH) tasks/:id
  - Headers:
    - UserId
    
  - Params:
    - id: string

  - Body:
    - closed: boolean
___
