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
TestWalmart, TestTarget
### User Unit Tests:
TestSignUp_OK, TestSignUp_TakenUsername, TestSignIn_OK, TestSignIn_UsernameNotFound, TestSignIn_PasswordIncorrect
### UserProduct Unit Tests:
TestSaveProduct_OK, TestSaveProduct_UnknownUsername, TestSaveProduct_AlreadySavedProduct, TestRemoveProduct_OK, TestRemoveProduct_UnknownUsername, TestRemoveProduct_ProductNotSaved, TestGetAllProducts_OK, TestGetAllProducts_UnknownUsername, TestGetAllProducts_NoProductsSaved

## Updated Documentation of Backend API

### Store Methods:
#### Walmart
Recieves a search request from the front end body and then passes that to the Walmart API, returns JSON list to front end body
- Path: `/walmart`
- HTTP Method: `GET`
- HTTP Status Responses:
    - 200 OK (success)


#### Target
Recieves a search request from the front end body and then passes that to the Target API, returns JSON list to front end body
- Path: `/Target`
- HTTP Method: `GET`
- HTTP Status Responses:
    - 200 OK (success)

### User Methods:
#### Create User
Creates a user object in the local account credentials database
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
    - 404 Not Found (user does not exist in local account credentials database)
- Example body:
```json
{
	"username": "admin",
	"password": "123456"
}
```
### UserProduct Methods
#### Save Product to User Account
Creates a UserProduct object in the local saved products database
- Path: `/api/products`
- HTTP Method: `POST`
- HTTP Status Responses:
    - 200 OK (success)
    - 400 Bad Request (request was missing username or product info)
    - 401 Unauthorized (username does not belong to any created accounts)
    - 409 Conflict (product has already been saved by user)
- Example body:
```json
{
	"username": "example_user",
	"seller_name": "Target",
	"product_name": "North Face Backpack",
	"price": "$120.00",
	"rating": "4.6",
	"image_url": "https://example.com"
}
```

#### Remove Product From User Account
Removes a UserProduct object in the local saved products database
- Path: `/api/products`
- HTTP Method: `DELETE`
- HTTP Status Responses:
    - 200 OK (success)
    - 400 Bad Request (request was missing username or product info)
    - 401 Unauthorized (username does not belong to any created accounts)
    - 404 Not Found (product has not been saved by user)
- Example body:
```json
{
	"username": "example_user",
	"seller_name": "Target",
	"product_name": "North Face Backpack",
	"price": "$120.00",
	"rating": "4.6",
	"image_url": "https://example.com"
}
```

#### Get User's Saved Products
Returns list of products saved to user's account
- Path: `/api/products`
- HTTP Method: `GET`
- HTTP Status Responses:
    - 200 OK (success, will return empty list if user has zero saved products)
    - 404 Not Found (user does not exist in local account credentials database)
- Example body:
```json
{
	"username": "example_user"
}
```