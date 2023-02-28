# Sprint 2

## Work Completed
### Frontend
### Backend

## Frontend Tests

## Backend Tests

## Documentation of backend API

### Store Methods:
### Walmart
### Target

### User Methods:
#### Create User
Creates a user object in the local database
- Path: `/api/user/signup`
- HTTP Method: `POST`
- HTTP Status Responses:
    - 200 OK (success)
    - 400 Bad Request (request was missing username/password)
    - 409 Conflict (user already exists)
- Example body:
```json
{
	"username": "Daniel",
	"password": "password"
}
```

#### Sign User In
Returns user info
- Path: `/api/user/signin`
- HTTP Method: `POST`
- HTTP Status Responses:
    - 200 OK (success)
    - 400 Bad Request (request was missing username/password)
    - 401 Unauthorized (incorrect password)
    - 404 Not Found (user does not exist in local database)
- Example body:
```json
{
	"username": "admin",
	"password": "123456"
}
```