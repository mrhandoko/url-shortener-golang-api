# url-shortener-golang-api

## Functional Requirements

1. Given a long URL, the service should generate a shorter and **unique** alias of it with **exactly 6 alphanumeric** characters and return it to the client so the user can get the shortened URL
2. Given a short link, the service should be able to redirect the request to the original link from server-side.
3. Given a short link, the service should be able to return these stats to the client so the user can view them:
- Redirect count
- Created at

## Non-Functional Requirements

- Since we’re using Golang as our main BE language, it’s recommended to use it for this exercise. In case you’re not familiar with it, you may use any language you’re most comfortable with.
- The service should expose **API endpoints** that satisfy the above requirements. You don’t need to implement the view.
- For the purpose of the coding exercise, you may store the data in-memory without any external dependencies to ease implementation. Extra point if you can make an interface such that it’s easy to swap the implementation.
