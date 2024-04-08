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
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/cars/add": {
            "post": {
                "description": "Add one or more cars",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cars"
                ],
                "summary": "AddCarHandler",
                "parameters": [
                    {
                        "description": "Registration numbers of cars (comma-separated)",
                        "name": "regNums",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Successful response with an array of added cars",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Car"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/main.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/main.APIError"
                        }
                    }
                }
            }
        },
        "/cars/delete/{id}": {
            "delete": {
                "description": "Delete a car by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cars"
                ],
                "summary": "DeleteCarHandler",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Car ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ID of the deleted car",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/main.APIError"
                        }
                    },
                    "404": {
                        "description": "Resource not found",
                        "schema": {
                            "$ref": "#/definitions/main.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/main.APIError"
                        }
                    }
                }
            }
        },
        "/cars/get": {
            "get": {
                "description": "Get a list of cars with pagination support",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cars"
                ],
                "summary": "GetCarsHandler",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Number of items per page",
                        "name": "page_size",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Car make",
                        "name": "make",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Car model",
                        "name": "model",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Car year",
                        "name": "year",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful response with an array of cars",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Car"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/main.APIError"
                        }
                    },
                    "404": {
                        "description": "Resource not found",
                        "schema": {
                            "$ref": "#/definitions/main.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/main.APIError"
                        }
                    }
                }
            }
        },
        "/cars/update/{id}": {
            "put": {
                "description": "Update a car by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cars"
                ],
                "summary": "UpdateCarHandler",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Car ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ID of the updated car",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/main.APIError"
                        }
                    },
                    "404": {
                        "description": "Resource not found",
                        "schema": {
                            "$ref": "#/definitions/main.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/main.APIError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.APIError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "main.Car": {
            "type": "object",
            "properties": {
                "mark": {
                    "type": "string"
                },
                "model": {
                    "type": "string"
                },
                "owner": {
                    "$ref": "#/definitions/main.People"
                },
                "regNum": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "main.People": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "patronymic": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
