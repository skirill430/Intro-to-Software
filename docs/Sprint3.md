# Sprint 2

Frontend video:

Backend video:

## Work Completed

### Frontend

### Backend
- Created a secondary database to store the data of the products saved by users that connects to the initial database of user account details.
- Created routes and handler functions for a user saving or removing a product that interact with both databases to verify given username and attach the product to that user account, returning the appropriate HTTP status response.
- Created a route and handler function to get the products saved by a given user, returning the appropriate HTTP status response and the list of saved products in JSON format if the user was valid.
- Wrote unit tests for the following routes: user saving a product, user removing a product, and getting a user's saved products.
- Simplified unit test output.

## Frontend Tests
### Angular Unit Tests:
### Cypress Test:

## Backend Tests
### Store Unit Tests:
### User Unit Tests:
### UserProduct Unit Tests:

## Updated Documentation of Backend API

### Store Methods:
#### Walmart
Recieves a search request from the front end body and then passes that to the Walmart API, returns JSON list to front end body
- Path: `/walmart`
- HTTP  METHOD: `GET`
- HTTP Status Responses:
    - 200 OK (success)


#### Target
Recieves a search request from the front end body and then passes that to the Target API, returns JSON list to front end body
- Path: `/Target`
- HTTP  METHOD: `GET`
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
