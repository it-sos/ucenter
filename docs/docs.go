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
                "security": [
                    {
                        "token": [
                            "read"
                        ]
                    }
                ],
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
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/datamodels.User"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/errors.Errors"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "token": [
                            "read"
                        ]
                    }
                ],
                "description": "更新用户信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "更新用户信息",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/parameter.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/datamodels.User"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/errors.Errors"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "token": [
                            "read"
                        ]
                    }
                ],
                "description": "创建用户信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "创建用户信息",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/parameter.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/datamodels.User"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/errors.Errors"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "token": [
                            "read"
                        ]
                    }
                ],
                "description": "删除用户信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "删除用户信息",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/parameter.UserDeleted"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/errors.Errors"
                        }
                    }
                }
            }
        },
        "/users/byuuid": {
            "get": {
                "security": [
                    {
                        "token": [
                            "read"
                        ]
                    }
                ],
                "description": "通过用户uuid获取用户信息",
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
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/datamodels.User"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/errors.Errors"
                        }
                    }
                }
            }
        },
        "/users/disabled": {
            "put": {
                "security": [
                    {
                        "token": [
                            "read"
                        ]
                    }
                ],
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
                        "description": "request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/parameter.UserDisabled"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/errors.Errors"
                        }
                    }
                }
            }
        },
        "/users/list": {
            "get": {
                "security": [
                    {
                        "token": [
                            "read"
                        ]
                    }
                ],
                "description": "获取全部用户分页列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "获取用户分页列表",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/parameter.Page"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/vo.UserVO"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/errors.Errors"
                        }
                    }
                }
            }
        },
        "/users/password": {
            "put": {
                "security": [
                    {
                        "token": [
                            "read"
                        ]
                    }
                ],
                "description": "修改登陆用户密码",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "修改登陆用户密码",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/parameter.Password"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error",
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
                    "type": "integer"
                },
                "disabled": {
                    "description": "禁用状态1=是，0=否",
                    "type": "integer"
                },
                "expired": {
                    "description": "有效期0=永久，unix时间戳",
                    "type": "integer"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "nickname": {
                    "description": "昵称",
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
                    "example": "5bbc-4ala-3dja-1djs-0aja"
                }
            }
        },
        "errors.Errors": {
            "type": "object",
            "properties": {
                "errCode": {
                    "type": "integer",
                    "example": 4000000
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "parameter.Page": {
            "type": "object",
            "properties": {
                "page": {
                    "description": "页码",
                    "type": "integer"
                },
                "size": {
                    "description": "每页条数",
                    "type": "integer"
                }
            }
        },
        "parameter.Password": {
            "type": "object",
            "required": [
                "id",
                "password"
            ],
            "properties": {
                "id": {
                    "description": "用户id",
                    "type": "integer"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                }
            }
        },
        "parameter.User": {
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
                "nickname": {
                    "description": "昵称",
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "parameter.UserDeleted": {
            "type": "object",
            "required": [
                "deleted",
                "id"
            ],
            "properties": {
                "deleted": {
                    "description": "删除标志0=否，1=是",
                    "type": "string"
                },
                "id": {
                    "description": "用户id",
                    "type": "integer"
                }
            }
        },
        "parameter.UserDisabled": {
            "type": "object",
            "required": [
                "disabled",
                "id"
            ],
            "properties": {
                "disabled": {
                    "description": "禁用标志0=否，1=是",
                    "type": "string"
                },
                "id": {
                    "description": "用户id",
                    "type": "integer"
                }
            }
        },
        "vo.User": {
            "type": "object",
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
                    "type": "integer"
                },
                "disabled": {
                    "description": "禁用状态1=是，0=否",
                    "type": "integer"
                },
                "disabledName": {
                    "type": "string",
                    "example": "已禁用"
                },
                "expired": {
                    "description": "有效期0=永久，unix时间戳",
                    "type": "integer"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "nickname": {
                    "description": "昵称",
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
                    "example": "5bbc-4ala-3dja-1djs-0aja"
                }
            }
        },
        "vo.UserVO": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "每页数据",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/vo.User"
                    }
                },
                "page": {
                    "description": "当前页数",
                    "type": "integer"
                },
                "size": {
                    "description": "每页条数",
                    "type": "integer"
                },
                "totalPage": {
                    "description": "总页数",
                    "type": "integer"
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
