// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-04-20 01:48:10.734157 +0500 +05 m=+0.053801876

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/calc-delivery-cost": {
            "post": {
                "description": "API for getting total delivery cost",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "geo"
                ],
                "summary": "Calculate Delivery Price For Clients",
                "parameters": [
                    {
                        "description": "calc-delivery-cost",
                        "name": "calc",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CalcDeliveryCostRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.CalcDeliveryCostResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            }
        },
        "/v1/couriers/{courier_id}": {
            "get": {
                "description": "API for getting courier",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "courier"
                ],
                "summary": "Get Courier",
                "parameters": [
                    {
                        "type": "string",
                        "description": "courier_id",
                        "name": "courier_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GetCourierModel"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            }
        },
        "/v1/fares": {
            "get": {
                "description": "API for getting fares",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "fare"
                ],
                "summary": "Get Fares",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GetAllFaresModel"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            },
            "put": {
                "description": "API for updating fare",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "fare"
                ],
                "summary": "Update Fare",
                "parameters": [
                    {
                        "description": "fare",
                        "name": "fare",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateFareModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GetFareModel"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            },
            "post": {
                "description": "API for creating fare",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "fare"
                ],
                "summary": "Create Fare",
                "parameters": [
                    {
                        "description": "fare",
                        "name": "fare",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateFareModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GetFareModel"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            }
        },
        "/v1/fares/{fare_id}": {
            "get": {
                "description": "API for getting fare",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "fare"
                ],
                "summary": "Get Fare",
                "parameters": [
                    {
                        "type": "string",
                        "description": "fare_id",
                        "name": "fare_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GetFareModel"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            },
            "delete": {
                "description": "API for deleting fare",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "fare"
                ],
                "summary": "Delete Fare",
                "parameters": [
                    {
                        "type": "string",
                        "description": "fare_id",
                        "name": "fare_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseOK"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            }
        },
        "/v1/geozones/": {
            "get": {
                "description": "Get Geozones",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "geo"
                ],
                "summary": "Get Geozones",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Parent ID",
                        "name": "parent_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "PAGE",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "PAGE",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GeozoneModel"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            }
        },
        "/v1/getdistance/from/{from_lat}/{from_long}/to/{to_lat}/{to_long}": {
            "get": {
                "description": "API for getting distance",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "geo"
                ],
                "summary": "Get Distance",
                "parameters": [
                    {
                        "type": "string",
                        "description": "from_lat",
                        "name": "from_lat",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "from_long",
                        "name": "from_long",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "to_lat",
                        "name": "to_lat",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "to_long",
                        "name": "to_long",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GetDistanceResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            }
        },
        "/v1/order": {
            "get": {
                "description": "API for getting orders",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order"
                ],
                "summary": "Get Orders",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GetOrders"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            }
        },
        "/v1/order/": {
            "post": {
                "description": "API for creating order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order"
                ],
                "summary": "Create Order",
                "parameters": [
                    {
                        "description": "order",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateOrder"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseOK"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            }
        },
        "/v1/order/{order_id}": {
            "get": {
                "description": "API for getting order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order"
                ],
                "summary": "Get Order",
                "parameters": [
                    {
                        "type": "string",
                        "description": "order_id",
                        "name": "order_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GetOrder"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.CalcDeliveryCostRequest": {
            "type": "object",
            "properties": {
                "from_location": {
                    "type": "object",
                    "$ref": "#/definitions/models.Location"
                },
                "min_distance": {
                    "type": "number"
                },
                "min_price": {
                    "type": "number"
                },
                "per_km_price": {
                    "type": "number"
                },
                "to_location": {
                    "type": "object",
                    "$ref": "#/definitions/models.Location"
                }
            }
        },
        "models.CalcDeliveryCostResponse": {
            "type": "object",
            "properties": {
                "distance": {
                    "type": "number"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "models.CreateFareModel": {
            "type": "object",
            "properties": {
                "delivery_time": {
                    "type": "integer"
                },
                "min_distance": {
                    "type": "number"
                },
                "min_price": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                },
                "price_per_km": {
                    "type": "number"
                }
            }
        },
        "models.CreateOrder": {
            "type": "object",
            "properties": {
                "branch_id": {
                    "type": "string",
                    "example": "a010f178-da52-4373-aacd-e477d871e27a"
                },
                "co_id": {
                    "type": "string",
                    "example": "a010f178-da52-4373-aacd-e477d871e27a"
                },
                "creator_type_id": {
                    "type": "string",
                    "example": "a010f178-da52-4373-aacd-e477d871e27a"
                },
                "deliver_price": {
                    "type": "number",
                    "example": 10000
                },
                "fare_id": {
                    "type": "string",
                    "example": "a010f178-da52-4373-aacd-e477d871e27a"
                },
                "from_address": {
                    "type": "string",
                    "example": "Hamid Olimjon maydoni 10A dom 40-kvartira"
                },
                "from_location": {
                    "type": "object",
                    "$ref": "#/definitions/models.Location"
                },
                "phone_number": {
                    "type": "string",
                    "example": "998998765432"
                },
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.productModel"
                    }
                },
                "to_address": {
                    "type": "string",
                    "example": "Hamid Olimjon maydoni 10A dom 40-kvartira"
                },
                "to_location": {
                    "type": "object",
                    "$ref": "#/definitions/models.Location"
                },
                "user_id": {
                    "type": "string",
                    "example": "a010f178-da52-4373-aacd-e477d871e27a"
                }
            }
        },
        "models.GeozoneModel": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "geozones": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.geozone"
                    }
                }
            }
        },
        "models.GetAllFaresModel": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "fares": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.GetFareModel"
                    }
                }
            }
        },
        "models.GetCourierModel": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "models.GetDistanceResponse": {
            "type": "object",
            "properties": {
                "distance": {
                    "type": "number"
                }
            }
        },
        "models.GetFareModel": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "delivery_time": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "min_distance": {
                    "type": "number"
                },
                "min_price": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                },
                "price_per_km": {
                    "type": "number"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.GetOrder": {
            "type": "object",
            "properties": {
                "branch_id": {
                    "type": "string",
                    "example": "a010f178-da52-4373-aacd-e477d871e27a"
                },
                "co_id": {
                    "type": "string",
                    "example": "a010f178-da52-4373-aacd-e477d871e27a"
                },
                "creator_type_id": {
                    "type": "string",
                    "example": "a010f178-da52-4373-aacd-e477d871e27a"
                },
                "deliver_price": {
                    "type": "number",
                    "example": 10000
                },
                "fare_id": {
                    "type": "string",
                    "example": "a010f178-da52-4373-aacd-e477d871e27a"
                },
                "from_address": {
                    "type": "string",
                    "example": "Hamid Olimjon maydoni 10A dom 40-kvartira"
                },
                "from_location": {
                    "type": "object",
                    "$ref": "#/definitions/models.Location"
                },
                "id": {
                    "type": "string",
                    "example": "701dc270-0adc-4d00-ae78-4f2f78d794cc"
                },
                "phone_number": {
                    "type": "string",
                    "example": "998998765432"
                },
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.getOrderProductModel"
                    }
                },
                "status_id": {
                    "type": "string",
                    "example": "52f248b4-23a0-4350-80b7-1704eaff6c8c"
                },
                "to_address": {
                    "type": "string",
                    "example": "Hamid Olimjon maydoni 10A dom 40-kvartira"
                },
                "to_location": {
                    "type": "object",
                    "$ref": "#/definitions/models.Location"
                },
                "user_id": {
                    "type": "string",
                    "example": "a010f178-da52-4373-aacd-e477d871e27a"
                }
            }
        },
        "models.GetOrders": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "orders": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.getOrderModel"
                    }
                }
            }
        },
        "models.Location": {
            "type": "object",
            "properties": {
                "lat": {
                    "type": "number",
                    "example": 40.123
                },
                "long": {
                    "type": "number",
                    "example": 60.123
                }
            }
        },
        "models.ResponseError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "object"
                }
            }
        },
        "models.ResponseOK": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "object"
                }
            }
        },
        "models.UpdateFareModel": {
            "type": "object",
            "properties": {
                "delivery_time": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "min_distance": {
                    "type": "number"
                },
                "min_price": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                },
                "price_per_km": {
                    "type": "number"
                }
            }
        },
        "models.geozone": {
            "type": "object",
            "properties": {
                "geometry": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "parent_id": {
                    "type": "string"
                }
            }
        },
        "models.getOrderModel": {
            "type": "object",
            "properties": {
                "branch_id": {
                    "type": "string",
                    "example": "a010f178-da52-4373-aacd-e477d871e27a"
                },
                "co_id": {
                    "type": "string",
                    "example": "a010f178-da52-4373-aacd-e477d871e27a"
                },
                "creator_type_id": {
                    "type": "string",
                    "example": "a010f178-da52-4373-aacd-e477d871e27a"
                },
                "deliver_price": {
                    "type": "number",
                    "example": 10000
                },
                "fare_id": {
                    "type": "string",
                    "example": "a010f178-da52-4373-aacd-e477d871e27a"
                },
                "from_address": {
                    "type": "string",
                    "example": "Hamid Olimjon maydoni 10A dom 40-kvartira"
                },
                "from_location": {
                    "type": "object",
                    "$ref": "#/definitions/models.Location"
                },
                "id": {
                    "type": "string",
                    "example": "701dc270-0adc-4d00-ae78-4f2f78d794cc"
                },
                "phone_number": {
                    "type": "string",
                    "example": "998998765432"
                },
                "status_id": {
                    "type": "string",
                    "example": "52f248b4-23a0-4350-80b7-1704eaff6c8c"
                },
                "to_address": {
                    "type": "string",
                    "example": "Hamid Olimjon maydoni 10A dom 40-kvartira"
                },
                "to_location": {
                    "type": "object",
                    "$ref": "#/definitions/models.Location"
                },
                "user_id": {
                    "type": "string",
                    "example": "a010f178-da52-4373-aacd-e477d871e27a"
                }
            }
        },
        "models.getOrderProductModel": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "example": "Choyxona Osh"
                },
                "price": {
                    "type": "number",
                    "example": 25000
                },
                "quantity": {
                    "type": "number",
                    "example": 2
                },
                "total_amount": {
                    "type": "number"
                }
            }
        },
        "models.productModel": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "Choyxona Osh"
                },
                "price": {
                    "type": "number",
                    "example": 25000
                },
                "quantity": {
                    "type": "number",
                    "example": 2
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

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
