// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

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
        "termsOfService": "http://www.yibuyiyin.com/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.yibuyiyin.com/support",
            "email": "peng.yu@yibuyiyin.com"
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
        "/users": {
            "get": {
                "description": "通过用户ID获取用户信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "获取用户信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/datamodels.User"
                        }
                    }
                }
            },
            "put": {
                "description": "更新用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "更新用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "昵称",
                        "name": "nickname",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "手机号",
                        "name": "phone",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "有效期，0=永久/指定过期时间，默认:0",
                        "name": "expired",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "是否禁用状态1=是;0=否，默认:否",
                        "name": "disabled",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/datamodels.User"
                        }
                    }
                }
            },
            "post": {
                "description": "创建用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "创建用户",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/datamodels.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/datamodels.User"
                        }
                    }
                }
            },
            "delete": {
                "description": "删除用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "删除用户",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/disabled": {
            "put": {
                "description": "设置用户禁用状态",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "设置用户禁用状态",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "0=启用；1=禁用",
                        "name": "disabled",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/password": {
            "put": {
                "description": "修改当前用户密码",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "修改当前用户密码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "query",
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
        "/users/uuid": {
            "get": {
                "description": "通过用户ID获取用户信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "获取用户信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户uuid",
                        "name": "uuid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error message",
                        "schema": {
                            "$ref": "#/definitions/errors.Errors"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "datamodels.User": {
            "type": "object",
            "required": [
                "account",
                "nickname",
                "password"
            ],
            "properties": {
                "account": {
                    "description": "帐号",
                    "type": "string"
                },
                "create_time": {
                    "type": "string",
                    "readOnly": true
                },
                "deleted": {
                    "description": "删除标志1=是，0=否",
                    "type": "integer",
                    "default": 0
                },
                "disabled": {
                    "description": "禁用状态1=是，0=否",
                    "type": "integer",
                    "default": 0
                },
                "expired": {
                    "description": "有效期0=永久，unix时间戳",
                    "type": "integer",
                    "default": 0
                },
                "id": {
                    "type": "integer",
                    "readOnly": true,
                    "example": 1
                },
                "nickname": {
                    "description": "昵称",
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "update_time": {
                    "type": "string",
                    "readOnly": true
                },
                "uuid": {
                    "type": "string",
                    "readOnly": true,
                    "example": "5bbc-4ala-3dja-1djs-0aja"
                }
            }
        },
        "errors.Errors": {
            "type": "object",
            "properties": {
                "errorCode": {
                    "type": "integer"
                },
                "errorMessage": {
                    "type": "string"
                }
            }
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
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "ucenter api",
	Description: "用户中心接口",
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
