# Sprint 4

Video: 

## Work Completed

### Frontend
- Added footer with link to project Github.
- Displays error messages when user cannot sign up or log in correctly.
- Used Angular library to implement cookies to save user data.
- Added an 'Add to cart' column in the table to allow users to save items.
- Implemented separate cart page that shows all products a user has added to their cart using cookies.
- Added 'Sign out' button that only shows when user is logged in, and 'Sign up / Log in' button disappears when logged in.
- Title routes to home page upon click.

### Backend
- Created route, handler function, and unit test for user logout.
- Adapted User routes and unit tests to generate a JWT cookie that stores the username and keeps user logged in for 24 hours.
- Adapted UserProduct routes and unit tests to extract username from the given cookie in the HTTP request.

## Frontend Tests
### Angular Unit Tests:

### Cypress Tests:

## Backend Tests
### Store Unit Tests:
TestWalmart, TestTarget
### User Unit Tests:
TestSignUp_OK, TestSignUp_TakenUsername, TestSignIn_OK, TestSignIn_UsernameNotFound, TestSignIn_PasswordIncorrect, TestLogout_OK
### UserProduct Unit Tests:
TestSaveProduct_OK, TestSaveProduct_UnknownUsername, TestSaveProduct_AlreadySavedProduct, TestRemoveProduct_OK, TestRemoveProduct_UnknownUsername, TestRemoveProduct_ProductNotSaved, TestGetAllProducts_OK, TestGetAllProducts_UnknownUsername, TestGetAllProducts_NoProductsSaved

## Updated Documentation of Backend API
### Store Methods:
#### BothStores
Recieves a search request from the front end body and then passes that to the Target and Walmart APIs, returns standardized JSON list to frontend body
- Path: `/bothStores`
- HTTP Method: `GET`
- HTTP Status Responses:
    - 200 OK (success)

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
    - 500 Internal Server Error (JWT cookie creation failed)
- Example JSON body:
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
    - 500 Internal Server Error (JWT cookie creation failed)
- Example JSON body:
```json
{
	"username": "admin",
	"password": "123456"
}
```

#### Log User Out
Clears user cookie from browser
- Path: `/api/user/logout`
- HTTP Method: `POST`
- HTTP Status Responses:
    - 200 OK (success)

### UserProduct Methods
#### Save Product to User Account
Creates a UserProduct object in the local saved products database, given a username cookie (created on sign-in) and product information in JSON body
- Path: `/api/products`
- HTTP Method: `POST`
- HTTP Status Responses:
    - 200 OK (success)
    - 400 Bad Request (request was missing username or product info)
    - 401 Unauthorized (username does not belong to any created accounts OR if the cookie signature is invalid)
    - 409 Conflict (product has already been saved by user)
- Example JSON body:
```json
{
	"seller_name": "Target",
	"product_name": "North Face Backpack",
	"price": "$120.00",
	"rating": "4.6",
	"image_url": "https://example.com"
}
```

#### Remove Product From User Account
Removes a UserProduct object in the local saved products database, given a username cookie (created on sign-in) and product information in JSON body
- Path: `/api/products`
- HTTP Method: `DELETE`
- HTTP Status Responses:
    - 200 OK (success)
    - 400 Bad Request (request was missing username or product info)
    - 401 Unauthorized (username does not belong to any created accounts OR if the cookie signature is invalid)
    - 404 Not Found (product has not been saved by user)
- Example JSON body:
```json
{
	"seller_name": "Target",
	"product_name": "North Face Backpack",
	"price": "$120.00",
	"rating": "4.6",
	"image_url": "https://example.com"
}
```

#### Get User's Saved Products
Returns list of products saved to user's account, given a username cookie (created on sign-in)
- Path: `/api/products`
- HTTP Method: `GET`
- HTTP Status Responses:
    - 200 OK (success, will return empty list if user has zero saved products)
    - 404 Not Found (user does not exist in local account credentials database)
