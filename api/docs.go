// Package api Code generated by swaggo/swag. DO NOT EDIT
package api

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
        "/project": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Project"
                ],
                "summary": "Get all projects",
                "operationId": "getAllProjects",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Project"
                            }
                        }
                    }
                }
            }
        },
        "/project/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "New Project Creating",
                "operationId": "createProject",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Dataset file",
                        "name": "Dataset",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "project name",
                        "name": "ProjectName",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "enum": [
                            "Логистическая регрессия",
                            "Случайный лес"
                        ],
                        "type": "string",
                        "description": "project type",
                        "name": "ProjectType",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "project info",
                        "name": "ProjectTarget",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/project/edit": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Project"
                ],
                "summary": "Edit Project",
                "operationId": "editProject",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Dataset file",
                        "name": "Dataset",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "project name",
                        "name": "ProjectName",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "enum": [
                            "Логистическая регрессия",
                            "Случайный лес",
                            "Линейная регрессия",
                            "Метод опорных векторов"
                        ],
                        "type": "string",
                        "description": "project type",
                        "name": "ProjectType",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "project info",
                        "name": "ProjectTarget",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/project/type/{type}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Project"
                ],
                "summary": "Get Projects by type",
                "operationId": "getProjByType",
                "parameters": [
                    {
                        "enum": [
                            "logreg",
                            "randf",
                            "linreg",
                            "vect"
                        ],
                        "type": "string",
                        "description": "Project type",
                        "name": "type",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Project"
                            }
                        }
                    }
                }
            }
        },
        "/project/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Project"
                ],
                "summary": "Get project by id",
                "operationId": "getProjectById",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Project ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/models.Project"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Project"
                ],
                "summary": "Delete Project",
                "operationId": "deleteProject",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Project ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get all users",
                "operationId": "getAllUsers",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.User"
                            }
                        }
                    }
                }
            }
        },
        "/user/edit": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "User"
                ],
                "summary": "Edit Profile",
                "operationId": "editProfile",
                "parameters": [
                    {
                        "description": "UserData",
                        "name": "UserData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "User"
                ],
                "summary": "User login",
                "operationId": "login",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "Username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Password",
                        "name": "Password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "New User registration",
                "operationId": "registerUser",
                "parameters": [
                    {
                        "description": "UserData",
                        "name": "UserData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get user profile by id",
                "operationId": "getUserById",
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
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "User"
                ],
                "summary": "Delete profile",
                "operationId": "deleteProfile",
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
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Project": {
            "type": "object",
            "properties": {
                "project_name": {
                    "type": "string"
                },
                "project_type": {
                    "type": "string"
                },
                "result_link": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "target": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "projects": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Project"
                    }
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
	Version:          "0.0.1",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "LinAutoML API",
	Description:      "API specs for LinAutoML",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
