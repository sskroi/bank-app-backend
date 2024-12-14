// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/account": {
            "post": {
                "security": [
                    {
                        "UserBearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create bank account",
                "parameters": [
                    {
                        "description": "Account info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/apihandler.createAccountInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/apihandler.createAccountResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apihandler.response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/apihandler.response"
                        }
                    },
                    "403": {
                        "description": "User deleted or banned",
                        "schema": {
                            "$ref": "#/definitions/apihandler.response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/apihandler.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apihandler.response"
                        }
                    }
                }
            }
        },
        "/accounts": {
            "get": {
                "security": [
                    {
                        "UserBearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get all user's accounts",
                "parameters": [
                    {
                        "minimum": 0,
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "maximum": 100,
                        "minimum": 0,
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Account"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apihandler.response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/apihandler.response"
                        }
                    },
                    "403": {
                        "description": "User deleted or banned",
                        "schema": {
                            "$ref": "#/definitions/apihandler.response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/apihandler.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apihandler.response"
                        }
                    }
                }
            }
        },
        "/auth/sign-in": {
            "post": {
                "description": "Authorizes the user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Sign in",
                "parameters": [
                    {
                        "description": "Sign in info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/apihandler.signInInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User successfully authorized",
                        "schema": {
                            "$ref": "#/definitions/apihandler.tokenResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apihandler.response"
                        }
                    },
                    "401": {
                        "description": "Invalid login credentials",
                        "schema": {
                            "$ref": "#/definitions/apihandler.response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/apihandler.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apihandler.response"
                        }
                    }
                }
            }
        },
        "/auth/sign-up": {
            "post": {
                "description": "Register new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Sign up",
                "parameters": [
                    {
                        "description": "Sign up info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/apihandler.signUpInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User successfully created",
                        "schema": {
                            "$ref": "#/definitions/apihandler.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apihandler.response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/apihandler.response"
                        }
                    },
                    "409": {
                        "description": "User with such email already exists",
                        "schema": {
                            "$ref": "#/definitions/apihandler.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apihandler.response"
                        }
                    }
                }
            }
        },
        "/transaction": {
            "post": {
                "security": [
                    {
                        "UserBearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create transaction",
                "parameters": [
                    {
                        "description": "Transaction info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/apihandler.createTransactionInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/apihandler.createTransactionResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/apihandler.response"
                        }
                    },
                    "403": {
                        "description": "User deleted or banned",
                        "schema": {
                            "$ref": "#/definitions/apihandler.response"
                        }
                    },
                    "404": {
                        "description": "Receiver or sender account not foundr",
                        "schema": {
                            "$ref": "#/definitions/apihandler.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apihandler.response"
                        }
                    }
                }
            }
        },
        "/user/update-profile": {
            "post": {
                "security": [
                    {
                        "UserBearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update user profile",
                "parameters": [
                    {
                        "description": "New profile data and current password",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully updated",
                        "schema": {
                            "$ref": "#/definitions/apihandler.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apihandler.response"
                        }
                    },
                    "401": {
                        "description": "Incorrect current password",
                        "schema": {
                            "$ref": "#/definitions/apihandler.response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/apihandler.response"
                        }
                    },
                    "409": {
                        "description": "User with such email already exists",
                        "schema": {
                            "$ref": "#/definitions/apihandler.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apihandler.response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "apihandler.createAccountInput": {
            "type": "object",
            "required": [
                "currency"
            ],
            "properties": {
                "currency": {
                    "type": "string",
                    "enum": [
                        "rub"
                    ]
                }
            }
        },
        "apihandler.createAccountResponse": {
            "type": "object",
            "properties": {
                "number": {
                    "type": "string"
                }
            }
        },
        "apihandler.createTransactionInput": {
            "type": "object",
            "required": [
                "amount",
                "receiver_account_number",
                "sender_account_number"
            ],
            "properties": {
                "amount": {
                    "type": "number"
                },
                "receiver_account_number": {
                    "type": "string"
                },
                "sender_account_number": {
                    "type": "string"
                }
            }
        },
        "apihandler.createTransactionResponse": {
            "type": "object",
            "properties": {
                "conversion_rate": {
                    "type": "number"
                },
                "is_conversion": {
                    "type": "boolean"
                },
                "public_id": {
                    "type": "string"
                },
                "received": {
                    "type": "number"
                },
                "receiver_account_number": {
                    "type": "string"
                },
                "sender_account_number": {
                    "type": "string"
                },
                "sent": {
                    "type": "number"
                }
            }
        },
        "apihandler.response": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "apihandler.signInInput": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 64
                },
                "password": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 8
                }
            }
        },
        "apihandler.signUpInput": {
            "type": "object",
            "required": [
                "email",
                "name",
                "passport",
                "password",
                "surname"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 64
                },
                "name": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 1
                },
                "passport": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 8
                },
                "password": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 8
                },
                "patronymic": {
                    "type": "string",
                    "maxLength": 64
                },
                "surname": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 1
                }
            }
        },
        "apihandler.tokenResponse": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                }
            }
        },
        "domain.Account": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "number"
                },
                "currency": {
                    "type": "string"
                },
                "is_close": {
                    "type": "boolean"
                },
                "number": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "UserBearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "bankapi.iorkss.ru",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Backend part of educational banking application",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
