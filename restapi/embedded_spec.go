// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "description": "dimple api\n",
    "title": "dimple",
    "version": "0.1.0"
  },
  "paths": {
    "/list": {
      "get": {
        "description": "return past sent crypto",
        "summary": "list past sent crypto",
        "operationId": "list",
        "responses": {
          "200": {
            "description": "Crypto sent",
            "schema": {
              "$ref": "#/definitions/listing"
            }
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/send": {
      "post": {
        "description": "send some crypto",
        "summary": "send some crypto",
        "operationId": "send",
        "parameters": [
          {
            "type": "string",
            "name": "to",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "return txid",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "fields": {
          "type": "string"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "listing": {
      "type": "object",
      "properties": {
        "addr": {
          "description": "Some string",
          "type": "string"
        },
        "amount": {
          "description": "amount sent",
          "type": "integer"
        },
        "txid": {
          "description": "txid",
          "type": "string"
        }
      }
    }
  }
}`))
}
