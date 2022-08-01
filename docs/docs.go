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
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/api/download": {
            "get": {
                "description": "Download video",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Download video",
                "operationId": "download",
                "parameters": [
                    {
                        "type": "string",
                        "description": "URL of youtube video",
                        "name": "url",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Itag of youtube video",
                        "name": "itag",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/v1/api/info": {
            "post": {
                "description": "Get information of video",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get video info",
                "operationId": "get-info",
                "parameters": [
                    {
                        "type": "string",
                        "description": "URL of youtube video",
                        "name": "url",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Youtube Download REST",
	Description:      "This is a sample server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
