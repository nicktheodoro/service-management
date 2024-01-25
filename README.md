# Service Management API

## Overview

This API, built with the Gin framework and Gorm, provides a foundation for managing services within Golang applications. It's currently in its initial stage!

## Resources

### User

- **Get User by ID:** `GET /users/:id` - Retrieve information about a specific user.
- **Get All Users:** `GET /users` - Retrieve information about all users.
- **Create User:** `POST /users` - Create a new user.
- **Update User:** `PUT /users/:id` - Update information for a user.
- **Delete User:** `DELETE /users/:id` - Delete a user.

### User Role

- **Get User Role by ID:** `GET /users/roles/:id` - Retrieve information about a specific user role.
- **Get All User Roles:** `GET /users/roles` - Retrieve information about all user roles.
- **Create User Role:** `POST /users/roles` - Create a new user role.
- **Update User Role:** `PUT /user/roles/:id` - Update information for a user role.
- **Delete User Role:** `DELETE /user/roles/:id` - Delete a user role.

### Authentication

- **User Login:** `POST /login` - Authenticate a user and return a token.

## Getting Started

1. **Clone the repository:**
   ```bash
   https://github.com/nicktheodoro/service-management.git
   ```

2. **Install dependencies:**
   ```bash
   go get -u github.com/gin-gonic
   go get -u github.com/dgrijalva/jwt-go
   go get -u github.com/spf13/viper
   go get -u golang.org/x/crypto
   go get -u gorm.io/gorm
   ```

3. **Run the application:**
   ```bash
   go run .\cmd\server\main.go
   ```

## Testing the API

1. **Create a User:**
   ```bash
   curl -X POST -H "Content-Type: application/json" -d '{"first_name":"John","last_name":"Doe","email":"john.doe@example.com","password":"password123","role_id":1}' http://localhost:3007/users
   ```
   Creates a new user with the provided information.

2. **Authenticate the User:**
   ```bash
   curl -X POST -H "Content-Type: application/json" -d '{"email":"john.doe@example.com","password":"password123"}' http://localhost:3007/login
   ```
   Authenticates the user and retrieves an authentication token.

3. **Use Resources with the Token in the Authorization Header:**
   ```bash
   # Replace '[your_token_here]' with the actual token obtained from the authentication step and replace '[resource]' with the endpoint
   curl -H "Authorization: Bearer [your_token_here]" http://localhost:3007/[resource]
   ```
   Accesses resources using the obtained authentication token in the Authorization header.
   
## Contributing

Contributions are welcome! Please feel free to open issues or pull requests to suggest improvements or add new features.

## License

This project is licensed under the [MIT License](./LICENSE).
