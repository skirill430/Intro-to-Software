# Sprint 2

## Work Completed
### Frontend
### Backend

## Frontend Tests

## Backend Tests
### Store Tests
Calls the walmart and target API with a search for "shoes" in the body. Checks if the returned body has proper query result such as "shoes". If this is true, it passes the test

## Documentation of backend API

### Store Methods:
### Walmart
Recieves a search request from the front end body and then passes that to the Walmart API, returns JSON list to front end body
- Path: `/walmart`
- HTTP  METHOD: `GET`
- HTTP Status Responses:
    - 200 OK (success)


### Target
Recieves a search request from the front end body and then passes that to the Target API, returns JSON list to front end body
-Path: `/Target`
-HTTP  METHOD: `GET`
- HTTP Status Responses:
    - 200 OK (success)

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