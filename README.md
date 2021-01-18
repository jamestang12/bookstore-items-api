# Book Store Item API

Book store API is a microservice architecture API that group 3 individual API and 2 custom build library into a docker container which is highly maintainable and testable. This API is implemented with Golang, MySQL, Cassandra, Docker, and Elasticsearch as persistence layers, which allow user to make http request and interact with the database

# Tools

- Golang
- MySQL
- Cassandra
- Docker
- Elasticsearch
- Gin
- GoCQL
- Golang restclient
- Mux
- Elastic

## Functionality

- [x] Store user data
- [x] Dynamic item searching
- [x] Role based users
- [x] User authentication with OAuth token
- [x] Display items and user info
- [x] Register/login to track inquiries
- [x] Logging System
- [x] Custome error handler library
- [x] OAuth token library

## [Book Store User API](https://github.com/jamestang12/bookstore-user-api)
- This is the user api which handle the user data and uses MySQL, gin , and Golang

## [Book Store OAuth API](https://github.com/jamestang12/bookstore-oauth-api)
- This is the oauth api which handle authentication and generate user token and uses Cassandra, gin , and Golang

## [Book Store Utils Library](https://github.com/jamestang12/bookstore_utils_go)
- This is the custome build error handing library for this application

## [Book Store Ouath Library](https://github.com/jamestang12/bookstore_oauth_go)
- This is the custome build OAuth Library for this application which handle authentication and communicat with the Book Store OAuth API




