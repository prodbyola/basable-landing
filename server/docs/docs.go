// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/collaborators": {
            "get": {
                "description": "get all collaborators",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Collaborator"
                ],
                "summary": "returns all collaborators",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "page size",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "fullName",
                        "name": "fullName",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "role",
                        "name": "role",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "all Collaborators returned",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/SuccessResponseDto"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/Collaborator"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "unexpected internal server error",
                        "schema": {
                            "$ref": "#/definitions/FailedResponseDto"
                        }
                    }
                }
            },
            "post": {
                "description": "create new collaborator",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Collaborator"
                ],
                "summary": "registers a new collaborator",
                "parameters": [
                    {
                        "description": "New Collaborator Details JSON",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/CreateCollaboratorRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Collaborator created successfully",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/SuccessResponseDto"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/Collaborator"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "request body validation error",
                        "schema": {
                            "$ref": "#/definitions/FailedResponseDto"
                        }
                    },
                    "409": {
                        "description": "another Collaborator with supplied email exists",
                        "schema": {
                            "$ref": "#/definitions/FailedResponseDto"
                        }
                    },
                    "500": {
                        "description": "unexpected internal server error",
                        "schema": {
                            "$ref": "#/definitions/FailedResponseDto"
                        }
                    }
                }
            }
        },
        "/v1/collaborators/{id}": {
            "get": {
                "description": "get collaborator by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Collaborator"
                ],
                "summary": "returns a collaborator by its 16 chaarcter uuid",
                "parameters": [
                    {
                        "type": "string",
                        "description": "collaborator ID(UUID)",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/SuccessResponseDto"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/Collaborator"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "request param validation error",
                        "schema": {
                            "$ref": "#/definitions/FailedResponseDto"
                        }
                    },
                    "404": {
                        "description": "collaborator with the specified ID not found",
                        "schema": {
                            "$ref": "#/definitions/FailedResponseDto"
                        }
                    },
                    "500": {
                        "description": "unexpected internal server error",
                        "schema": {
                            "$ref": "#/definitions/FailedResponseDto"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete collaborator",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Collaborator"
                ],
                "summary": "deletes a collaborator",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Collaborator ID(UUID)",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Collaborator deleted suuceesfully",
                        "schema": {
                            "$ref": "#/definitions/SuccessResponseDto"
                        }
                    },
                    "400": {
                        "description": "request param validation error",
                        "schema": {
                            "$ref": "#/definitions/FailedResponseDto"
                        }
                    },
                    "500": {
                        "description": "unexpected internal server error",
                        "schema": {
                            "$ref": "#/definitions/FailedResponseDto"
                        }
                    }
                }
            },
            "patch": {
                "description": "update collaborator",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Collaborator"
                ],
                "summary": "updates a collaborator",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Collaborator ID(UUID)",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Collaborator Details JSON",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/UpdateCollaborator"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Collaborator updated successfully",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/SuccessResponseDto"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/Collaborator"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "request body/param validation error",
                        "schema": {
                            "$ref": "#/definitions/FailedResponseDto"
                        }
                    },
                    "404": {
                        "description": "Collaborator with specified ID not found",
                        "schema": {
                            "$ref": "#/definitions/FailedResponseDto"
                        }
                    },
                    "500": {
                        "description": "unexpected internal server error",
                        "schema": {
                            "$ref": "#/definitions/FailedResponseDto"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "Collaborator": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "email": {
                    "type": "string"
                },
                "fullName": {
                    "type": "string"
                },
                "github": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "imageUrl": {
                    "type": "string"
                },
                "linkedin": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "CreateCollaboratorRequest": {
            "type": "object",
            "required": [
                "email",
                "fullName"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "fullName": {
                    "type": "string"
                },
                "github": {
                    "type": "string"
                },
                "imageUrl": {
                    "type": "string"
                },
                "linkedin": {
                    "type": "string"
                },
                "role": {
                    "type": "string",
                    "default": "user"
                }
            }
        },
        "FailedResponseDto": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "errorType": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "integer"
                },
                "statusText": {
                    "type": "string"
                }
            }
        },
        "SuccessResponseDto": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "integer"
                },
                "statusText": {
                    "type": "string"
                }
            }
        },
        "UpdateCollaborator": {
            "type": "object",
            "properties": {
                "fullName": {
                    "type": "string"
                },
                "github": {
                    "type": "string"
                },
                "imageUrl": {
                    "type": "string"
                },
                "linkedin": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:5000",
	BasePath:         "/api",
	Schemes:          []string{"http"},
	Title:            "Basable Landing REST API",
	Description:      "REST API for Basable Landing",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
