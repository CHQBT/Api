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
        "/v1/auth/login": {
            "post": {
                "description": "it generates new tokens",
                "tags": [
                    "auth"
                ],
                "summary": "login user",
                "parameters": [
                    {
                        "description": "login and password",
                        "name": "userinfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Token",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid date",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error while reading from server",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/twit": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "it will Create Twit",
                "tags": [
                    "TWIT API"
                ],
                "summary": "Create Twit",
                "parameters": [
                    {
                        "description": "info",
                        "name": "info",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateTwitRequestApi"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "twit_id",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Invalid token",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/twit/all": {
            "get": {
                "description": "it will get all Twits",
                "tags": [
                    "TWIT API"
                ],
                "summary": "get all Twits",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/twit/latest-uploaded": {
            "get": {
                "description": "it will get latest twits",
                "tags": [
                    "TWIT API"
                ],
                "summary": "get latest twits",
                "parameters": [
                    {
                        "type": "string",
                        "description": "limit",
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
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/twit/location": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "it will Create Twit's Location",
                "tags": [
                    "TWIT API"
                ],
                "summary": "Create Twit's Location",
                "parameters": [
                    {
                        "description": "info",
                        "name": "info",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateLocationRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "location_id",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/twit/most-viewed": {
            "get": {
                "description": "it will get most view twits",
                "tags": [
                    "TWIT API"
                ],
                "summary": "get most view twits",
                "parameters": [
                    {
                        "type": "string",
                        "description": "limit",
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
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/twit/music/{twit_id}": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Upload Twit Music",
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "TWIT API"
                ],
                "summary": "CreateMusic",
                "parameters": [
                    {
                        "type": "string",
                        "description": "twit_id",
                        "name": "twit_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "UploadMediaForm",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/twit/photo/{twit_id}": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Upload Twit Photo",
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "TWIT API"
                ],
                "summary": "CreatePhoto",
                "parameters": [
                    {
                        "type": "string",
                        "description": "twit_id",
                        "name": "twit_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "UploadMediaForm",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/twit/search": {
            "get": {
                "description": "it will search twit by keyword from twit text and twit title and publisher-name",
                "tags": [
                    "TWIT API"
                ],
                "summary": "search twit by keyword from twit text and twit title and publisher-name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "keywoard",
                        "name": "keywoard",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/twit/type/{type}": {
            "get": {
                "description": "it will get twits by type",
                "tags": [
                    "TWIT API"
                ],
                "summary": "get twits by type",
                "parameters": [
                    {
                        "type": "string",
                        "description": "type",
                        "name": "type",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/twit/url": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "it will Create Twit's urls like youtube url",
                "tags": [
                    "TWIT API"
                ],
                "summary": "Create Twit's urls like youtube url",
                "parameters": [
                    {
                        "description": "info",
                        "name": "info",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateURLRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "url_id",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/twit/video/{twit_id}": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Upload Twit Video",
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "TWIT API"
                ],
                "summary": "CreateVideo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "twit_id",
                        "name": "twit_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "UploadMediaForm",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/twit/{id}": {
            "get": {
                "description": "it will Get Twit By id",
                "tags": [
                    "TWIT API"
                ],
                "summary": "Get Twit",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.TwitResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "it will add views Twit",
                "tags": [
                    "TWIT API"
                ],
                "summary": "Add view to Twit",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "message",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
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
                "description": "it will Delete Twit",
                "tags": [
                    "TWIT API"
                ],
                "summary": "Delete Twit",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "message",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.CreateLocationRequest": {
            "type": "object",
            "properties": {
                "lat": {
                    "type": "string"
                },
                "lon": {
                    "type": "string"
                },
                "twit_id": {
                    "type": "string"
                }
            }
        },
        "model.CreateTwitRequestApi": {
            "type": "object",
            "properties": {
                "publisher_fio": {
                    "type": "string"
                },
                "texts": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "model.CreateURLRequest": {
            "type": "object",
            "properties": {
                "twit_id": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "model.LocationInfo": {
            "type": "object",
            "properties": {
                "latitude": {
                    "type": "string"
                },
                "location_id": {
                    "type": "string"
                },
                "longitude": {
                    "type": "string"
                }
            }
        },
        "model.MusicInfo": {
            "type": "object",
            "properties": {
                "music_id": {
                    "type": "string"
                },
                "music_url": {
                    "type": "string"
                }
            }
        },
        "model.PhotoInfo": {
            "type": "object",
            "properties": {
                "photo_id": {
                    "type": "string"
                },
                "photo_url": {
                    "type": "string"
                }
            }
        },
        "model.TwitResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "locations": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.LocationInfo"
                    }
                },
                "musics": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.MusicInfo"
                    }
                },
                "photos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.PhotoInfo"
                    }
                },
                "publisher_fio": {
                    "type": "string"
                },
                "readers_count": {
                    "type": "integer"
                },
                "texts": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "urls": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.UrlInfo"
                    }
                },
                "user_id": {
                    "type": "string"
                },
                "videos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.VideoInfo"
                    }
                }
            }
        },
        "model.UrlInfo": {
            "type": "object",
            "properties": {
                "url": {
                    "type": "string"
                },
                "url_id": {
                    "type": "string"
                }
            }
        },
        "model.UserLogin": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.VideoInfo": {
            "type": "object",
            "properties": {
                "video_id": {
                    "type": "string"
                },
                "video_url": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "description": "API Gateway",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
