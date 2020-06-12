# FarmSale API

## Setup
Use go version 1.10 due to a known issue with [mongo-driver](https://github.com/golang/go/issues/37362)

## Routes

### Credentials
1. Admin
 ```
 {
 "email": "admin@user.com",
  "password": "admin12345"
 }
 ```
 2. User
 ```
 {
  "email": "user@user.com",
  "password": "user12345"
  }
 ```
 3. Agent
 ```
 {
 "email": "agent@user.com",
 "password": "agent12345"
 }
 ```
 
 4. Manager
 ```
 {
 "email": "manager@user.com",
 "password": "manager12345"
 }
 ```
 
 ### Public Routes
`GET https://farmsaledev.herokuapp.com`

### User authenticated route
`GET https://farmsaledev.herokuapp.com/auth/products`

### Admin Routes
`GET https://farmsaledev.herokuapp.com/admin`

### Manager Routes
`GET https://farmsaledev.herokuapp.com/manager`

### AgentRoutes
`GET https://farmsaledev.herokuapp.com/agent`

## WORKING DEMO

https://farmsaledev.herokuapp.com
