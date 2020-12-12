// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "test",
    "title": "test",
    "version": "1.0.0"
  },
  "paths": {
    "/login": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "login",
        "parameters": [
          {
            "name": "user info",
            "in": "body",
            "required": true,
            "schema": {
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
            }
          }
        ],
        "responses": {
          "200": {
            "description": "returns user",
            "schema": {
              "$ref": "#/definitions/user"
            },
            "headers": {
              "Authenthication": {
                "type": "string"
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/product/list": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "operationId": "productList",
        "parameters": [
          {
            "type": "string",
            "format": "date",
            "name": "date",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "returns list of product",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/product"
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/product/{productID}": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "operationId": "getProduct",
        "parameters": [
          {
            "type": "string",
            "name": "productID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "get product",
            "schema": {
              "$ref": "#/definitions/product"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "put": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "updateProduct",
        "parameters": [
          {
            "type": "string",
            "name": "productID",
            "in": "path",
            "required": true
          },
          {
            "name": "product",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/product"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "update product",
            "schema": {
              "$ref": "#/definitions/product"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "createProduct",
        "parameters": [
          {
            "name": "product",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/product"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "creates product",
            "schema": {
              "$ref": "#/definitions/product"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "productID",
          "in": "path",
          "required": true
        }
      ]
    },
    "/scanCheck": {
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "consumes": [
          "multipart/form-data"
        ],
        "summary": "Uploads a file.",
        "operationId": "scanCheck",
        "parameters": [
          {
            "type": "file",
            "description": "The file to upload.",
            "name": "upfile",
            "in": "formData"
          },
          {
            "type": "string",
            "format": "date-time",
            "description": "Date when scan was done",
            "name": "scanDate",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/productCount"
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/scanProducts": {
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Uploads a file.",
        "operationId": "scanProducts",
        "parameters": [
          {
            "name": "image",
            "in": "body",
            "required": true,
            "schema": {
              "required": [
                "body"
              ],
              "properties": {
                "body": {
                  "type": "string"
                }
              }
            }
          },
          {
            "type": "string",
            "format": "date-time",
            "description": "Date when scan was done",
            "name": "scanDate",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "schema": {
              "$ref": "#/definitions/scanResponse"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/signup": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "signup",
        "parameters": [
          {
            "name": "user info",
            "in": "body",
            "required": true,
            "schema": {
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
            }
          }
        ],
        "responses": {
          "200": {
            "description": "returns user",
            "schema": {
              "$ref": "#/definitions/user"
            },
            "headers": {
              "Authenthication": {
                "type": "string"
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/status": {
      "get": {
        "produces": [
          "application/json"
        ],
        "operationId": "getStatus",
        "responses": {
          "200": {
            "description": "returns status of server",
            "schema": {
              "properties": {
                "status": {
                  "type": "string",
                  "example": "OK"
                }
              }
            }
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "product": {
      "type": "object",
      "properties": {
        "boughtAt": {
          "type": "string",
          "format": "date-time",
          "x-nullable": true
        },
        "id": {
          "type": "string"
        },
        "inStock": {
          "type": "boolean"
        },
        "name": {
          "type": "string"
        },
        "type": {
          "type": "string"
        }
      }
    },
    "productCount": {
      "type": "object",
      "properties": {
        "count": {
          "type": "integer"
        },
        "product": {
          "type": "object",
          "$ref": "#/definitions/product"
        }
      }
    },
    "scanResponse": {
      "type": "object",
      "properties": {
        "productCounts": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/productCount"
          }
        },
        "products": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/product"
          }
        }
      }
    },
    "user": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "id": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "Bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "test",
    "title": "test",
    "version": "1.0.0"
  },
  "paths": {
    "/login": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "login",
        "parameters": [
          {
            "name": "user info",
            "in": "body",
            "required": true,
            "schema": {
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
            }
          }
        ],
        "responses": {
          "200": {
            "description": "returns user",
            "schema": {
              "$ref": "#/definitions/user"
            },
            "headers": {
              "Authenthication": {
                "type": "string"
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/product/list": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "operationId": "productList",
        "parameters": [
          {
            "type": "string",
            "format": "date",
            "name": "date",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "returns list of product",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/product"
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/product/{productID}": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "operationId": "getProduct",
        "parameters": [
          {
            "type": "string",
            "name": "productID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "get product",
            "schema": {
              "$ref": "#/definitions/product"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "put": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "updateProduct",
        "parameters": [
          {
            "type": "string",
            "name": "productID",
            "in": "path",
            "required": true
          },
          {
            "name": "product",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/product"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "update product",
            "schema": {
              "$ref": "#/definitions/product"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "createProduct",
        "parameters": [
          {
            "name": "product",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/product"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "creates product",
            "schema": {
              "$ref": "#/definitions/product"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "productID",
          "in": "path",
          "required": true
        }
      ]
    },
    "/scanCheck": {
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "consumes": [
          "multipart/form-data"
        ],
        "summary": "Uploads a file.",
        "operationId": "scanCheck",
        "parameters": [
          {
            "type": "file",
            "description": "The file to upload.",
            "name": "upfile",
            "in": "formData"
          },
          {
            "type": "string",
            "format": "date-time",
            "description": "Date when scan was done",
            "name": "scanDate",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/productCount"
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/scanProducts": {
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Uploads a file.",
        "operationId": "scanProducts",
        "parameters": [
          {
            "name": "image",
            "in": "body",
            "required": true,
            "schema": {
              "required": [
                "body"
              ],
              "properties": {
                "body": {
                  "type": "string"
                }
              }
            }
          },
          {
            "type": "string",
            "format": "date-time",
            "description": "Date when scan was done",
            "name": "scanDate",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "schema": {
              "$ref": "#/definitions/scanResponse"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/signup": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "signup",
        "parameters": [
          {
            "name": "user info",
            "in": "body",
            "required": true,
            "schema": {
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
            }
          }
        ],
        "responses": {
          "200": {
            "description": "returns user",
            "schema": {
              "$ref": "#/definitions/user"
            },
            "headers": {
              "Authenthication": {
                "type": "string"
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/status": {
      "get": {
        "produces": [
          "application/json"
        ],
        "operationId": "getStatus",
        "responses": {
          "200": {
            "description": "returns status of server",
            "schema": {
              "properties": {
                "status": {
                  "type": "string",
                  "example": "OK"
                }
              }
            }
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "product": {
      "type": "object",
      "properties": {
        "boughtAt": {
          "type": "string",
          "format": "date-time",
          "x-nullable": true
        },
        "id": {
          "type": "string"
        },
        "inStock": {
          "type": "boolean"
        },
        "name": {
          "type": "string"
        },
        "type": {
          "type": "string"
        }
      }
    },
    "productCount": {
      "type": "object",
      "properties": {
        "count": {
          "type": "integer"
        },
        "product": {
          "type": "object",
          "$ref": "#/definitions/product"
        }
      }
    },
    "scanResponse": {
      "type": "object",
      "properties": {
        "productCounts": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/productCount"
          }
        },
        "products": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/product"
          }
        }
      }
    },
    "user": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "id": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "Bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  }
}`))
}
