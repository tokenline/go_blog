// Code generated by swaggo/swag. DO NOT EDIT
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
        "/blog/auth/login": {
            "post": {
                "tags": [
                    "授权"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResult"
                        }
                    }
                }
            }
        },
        "/blog/user/add": {
            "post": {
                "tags": [
                    "用户"
                ],
                "summary": "增加用户",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.BlogUserAddDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResult"
                        }
                    }
                }
            }
        },
        "/blog/user/delete": {
            "post": {
                "tags": [
                    "用户"
                ],
                "summary": "软删除指定用户",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.IdInfoDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResult"
                        }
                    }
                }
            }
        },
        "/blog/user/list": {
            "get": {
                "tags": [
                    "用户"
                ],
                "summary": "返回1个全部状态为启用用户的数组",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.JsonResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/dto.BlogUserVo"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/blog/user/query": {
            "get": {
                "tags": [
                    "用户"
                ],
                "summary": "用户查询",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "pageNum",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页数量",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.JsonResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/response.PageResult"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "list": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/dto.BlogUserVo"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/blog/user/update": {
            "post": {
                "tags": [
                    "用户"
                ],
                "summary": "更新用户信息",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.BlogUserUpdateDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResult"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "constants.UserGender": {
            "type": "integer",
            "enum": [
                0,
                1,
                2
            ],
            "x-enum-comments": {
                "UserGenderFemale": "女",
                "UserGenderMale": "男",
                "UserGenderNone": "未知"
            },
            "x-enum-varnames": [
                "UserGenderNone",
                "UserGenderMale",
                "UserGenderFemale"
            ]
        },
        "constants.UserStatus": {
            "type": "integer",
            "enum": [
                0,
                1
            ],
            "x-enum-comments": {
                "UserStatusDisable": "禁用",
                "UserStatusEnable": "启用"
            },
            "x-enum-varnames": [
                "UserStatusEnable",
                "UserStatusDisable"
            ]
        },
        "dto.BlogUserAddDto": {
            "type": "object",
            "properties": {
                "gender": {
                    "description": "性别",
                    "allOf": [
                        {
                            "$ref": "#/definitions/constants.UserGender"
                        }
                    ]
                },
                "mail": {
                    "description": "邮箱地址",
                    "type": "string"
                },
                "nickname": {
                    "description": "昵称",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "phone": {
                    "description": "电话号码",
                    "type": "string"
                },
                "remark": {
                    "description": "备注",
                    "type": "string"
                },
                "status": {
                    "description": "状态",
                    "allOf": [
                        {
                            "$ref": "#/definitions/constants.UserStatus"
                        }
                    ]
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "dto.BlogUserUpdateDto": {
            "type": "object",
            "properties": {
                "gender": {
                    "$ref": "#/definitions/constants.UserGender"
                },
                "id": {
                    "description": "用户ID  想一想",
                    "type": "integer"
                },
                "mail": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "remark": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/constants.UserStatus"
                }
            }
        },
        "dto.BlogUserVo": {
            "type": "object",
            "properties": {
                "gender": {
                    "description": "性别",
                    "allOf": [
                        {
                            "$ref": "#/definitions/constants.UserGender"
                        }
                    ]
                },
                "id": {
                    "description": "用户ID",
                    "type": "integer"
                },
                "mail": {
                    "description": "邮箱地址",
                    "type": "string"
                },
                "nickname": {
                    "description": "昵称",
                    "type": "string"
                },
                "phone": {
                    "description": "电话号码",
                    "type": "string"
                },
                "remark": {
                    "description": "备注",
                    "type": "string"
                },
                "status": {
                    "description": "状态",
                    "allOf": [
                        {
                            "$ref": "#/definitions/constants.UserStatus"
                        }
                    ]
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "dto.LoginDto": {
            "type": "object"
        },
        "request.IdInfoDto": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "response.JsonResult": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "状态码（业务层状态）",
                    "type": "integer"
                },
                "data": {
                    "description": "数据"
                },
                "message": {
                    "description": "消息",
                    "type": "string"
                },
                "success": {
                    "description": "是否成功",
                    "type": "boolean"
                }
            }
        },
        "response.PageResult": {
            "type": "object",
            "properties": {
                "list": {
                    "description": "数据"
                },
                "totalCount": {
                    "description": "总条数",
                    "type": "integer"
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
