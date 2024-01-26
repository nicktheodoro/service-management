// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login": {
            "post": {
                "description": "authenticate user and generate access token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "parameters": [
                    {
                        "description": "Login Input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.LoginInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "JWT Token",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "security": [
                    {
                        "Authorization Token": []
                    }
                ],
                "description": "get all users",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/users.User"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Authorization Token": []
                    }
                ],
                "description": "create a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "parameters": [
                    {
                        "description": "User Input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.UserInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/users.User"
                        }
                    }
                }
            }
        },
        "/users/roles": {
            "get": {
                "security": [
                    {
                        "Authorization Token": []
                    }
                ],
                "description": "get all user roles",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user_roles"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/users.UserRole"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Authorization Token": []
                    }
                ],
                "description": "create a new user role",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user_roles"
                ],
                "parameters": [
                    {
                        "description": "User Role Input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.UserRoleInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/users.UserRole"
                        }
                    }
                }
            }
        },
        "/users/roles/{id}": {
            "get": {
                "security": [
                    {
                        "Authorization Token": []
                    }
                ],
                "description": "get user role by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user_roles"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User Role ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/users.UserRole"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Authorization Token": []
                    }
                ],
                "description": "update an existing user role by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user_roles"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User Role ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User Role Input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.UserRoleInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/users.UserRole"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Authorization Token": []
                    }
                ],
                "description": "delete an existing user role by ID",
                "tags": [
                    "user_roles"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User Role ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "security": [
                    {
                        "Authorization Token": []
                    }
                ],
                "description": "get User by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/users.User"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Authorization Token": []
                    }
                ],
                "description": "update an existing user by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User Input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.UserInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/users.User"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Authorization Token": []
                    }
                ],
                "description": "delete an existing user by ID",
                "tags": [
                    "users"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.LoginInput": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "controllers.UserInput": {
            "type": "object",
            "required": [
                "email",
                "first_name",
                "last_name",
                "password",
                "role_id"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role_id": {
                    "type": "integer"
                }
            }
        },
        "controllers.UserRoleInput": {
            "type": "object",
            "required": [
                "role_name"
            ],
            "properties": {
                "role_name": {
                    "type": "string"
                }
            }
        },
        "users.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "firstname": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastname": {
                    "type": "string"
                },
                "role": {
                    "$ref": "#/definitions/users.UserRole"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "users.UserRole": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "role_name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "",
	Description:      "API REST in Golang with Gin Framework",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
