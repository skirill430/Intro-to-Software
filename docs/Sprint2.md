# Sprint 2

Frontend video: https://youtu.be/5rg8xnS_-LQ

Backend video: https://youtu.be/i9pVfyWH468

## Work Completed

### General / Quality of Life
 - Implemented a continuous integration system using GitHub actions to run unit tests, cypress tests, and build production files that are then deployed to a server.


### Frontend
 - Implemented an HTTP service to retrieve data from the back end and then display it to a table.
 - Implemented routing in Angular in order to support switching between multiple pages. 
 - Reorganized the file structure for the components to make for more coherent code. 
 - Created a separate login page from the home page where the main functionality is.
 - Implemented HTTP post requests to allow users to sign up or log in.
 - Login page has two input forms for users to fill out and buttons for new users to sign up and returning users to log in.

### Backend
- Added a second API that gets search results from Target store. 
- Implemented unit tests for the two APIS.
- Set up a SQLite database to store user information.
- Created routes and handler functions for user sign-up and sign-in that interact with local database to verify the given credentials, returning the appropriate HTTP status response.
- Implemented hashing and salting of user passwords.
- Created unit tests for user sign-up and sign-in routes.

## Frontend Tests
### Angular Unit Tests:
Render Toolbar, Render Sidebar, Render Search Text Box, Create HomeComponent, Create AppComponent, Title 'Quick Shop', Create HttpService, Create LoginComponent
### Cypress Test:
Home (Tests home page by typing a query in the search box and testing for results)
## Backend Tests
### Store Unit Tests:
TestWalmart, TestTarget
### User Unit Tests:
TestSignUp_OK, TestSignUp_TakenUsername, TestSignIn_OK, TestSignIn_UsernameNotFound, TestSignIn_PasswordIncorrect

## Documentation of backend API

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
