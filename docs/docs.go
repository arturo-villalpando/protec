// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/merchant/": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Create\nType \"Bearer\" followed by a space and JWT token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Merchants"
                ],
                "summary": "Create",
                "parameters": [
                    {
                        "description": "Create",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.MerchantRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Merchant"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    }
                }
            }
        },
        "/merchant/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Read\nType \"Bearer\" followed by a space and JWT token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Merchants"
                ],
                "summary": "Read",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "1",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Merchant"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Update\nType \"Bearer\" followed by a space and JWT token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Merchants"
                ],
                "summary": "Update",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "1",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.MerchantRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Merchant"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    }
                }
            }
        },
        "/report/transactions/": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Read\nType \"Bearer\" followed by a space and JWT token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reports Transactions"
                ],
                "summary": "Read",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Earns"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    }
                }
            }
        },
        "/report/transactions/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Read Transaction Merchants Earns\nType \"Bearer\" followed by a space and JWT token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reports Transactions"
                ],
                "summary": "Read Transaction Merchants Earns",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "1",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Earns"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    }
                }
            }
        },
        "/transaction/": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Create\nType \"Bearer\" followed by a space and JWT token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transactions"
                ],
                "summary": "Create",
                "parameters": [
                    {
                        "description": "Create",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.TransactionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Transaction"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Earns": {
            "type": "object",
            "required": [
                "earns"
            ],
            "properties": {
                "earns": {
                    "type": "number",
                    "example": 1
                }
            }
        },
        "models.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "status bad request"
                }
            }
        },
        "models.Merchant": {
            "type": "object",
            "required": [
                "commission",
                "createdAt",
                "merchant_id",
                "merchant_name",
                "updatedAt"
            ],
            "properties": {
                "commission": {
                    "type": "number",
                    "maximum": 100,
                    "example": 1
                },
                "createdAt": {
                    "type": "string",
                    "example": "2022-12-10T12:00:00Z"
                },
                "merchant_id": {
                    "type": "integer",
                    "example": 1
                },
                "merchant_name": {
                    "type": "string",
                    "maxLength": 128,
                    "example": "Qwerty"
                },
                "updatedAt": {
                    "type": "string",
                    "example": "2022-12-10T12:00:00Z"
                }
            }
        },
        "models.MerchantRequest": {
            "type": "object",
            "required": [
                "commission",
                "merchant_name"
            ],
            "properties": {
                "commission": {
                    "type": "number",
                    "maximum": 100,
                    "example": 1
                },
                "merchant_name": {
                    "type": "string",
                    "maxLength": 128,
                    "example": "Qwerty"
                }
            }
        },
        "models.Transaction": {
            "type": "object",
            "required": [
                "amount",
                "commission",
                "createdAt",
                "fee",
                "merchant_id",
                "transaction_id"
            ],
            "properties": {
                "amount": {
                    "type": "number",
                    "example": 100
                },
                "commission": {
                    "type": "number",
                    "maximum": 100,
                    "example": 1
                },
                "createdAt": {
                    "type": "string",
                    "example": "2022-12-10T12:00:00Z"
                },
                "fee": {
                    "type": "number",
                    "example": 1
                },
                "merchant_id": {
                    "type": "integer",
                    "example": 1
                },
                "transaction_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "models.TransactionRequest": {
            "type": "object",
            "required": [
                "amount",
                "merchant_id"
            ],
            "properties": {
                "amount": {
                    "type": "number",
                    "example": 100
                },
                "merchant_id": {
                    "type": "integer",
                    "example": 1
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