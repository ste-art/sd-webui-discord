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
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API",
    "title": "ApiService",
    "version": "1.0.0"
  },
  "host": "127.0.0.1:18890",
  "basePath": "/api",
  "paths": {
    "/auth": {
      "get": {
        "description": "Auth With Discord",
        "tags": [
          "user"
        ],
        "operationId": "auth",
        "parameters": [
          {
            "type": "string",
            "description": "Discord Code",
            "name": "code",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "Discord State",
            "name": "state",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "302": {
            "description": "Redirect",
            "headers": {
              "Location": {
                "type": "string",
                "format": "url"
              }
            }
          }
        }
      }
    },
    "/cluster": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "Get Cluster Info",
        "tags": [
          "system"
        ],
        "operationId": "cluster",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/ClusterInfo"
            }
          }
        }
      }
    },
    "/community_history": {
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "Get Community History",
        "tags": [
          "user"
        ],
        "operationId": "community_history",
        "parameters": [
          {
            "description": "Page Info",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "properties": {
                "page_info": {
                  "$ref": "#/definitions/PageInfoRequest"
                },
                "query": {
                  "$ref": "#/definitions/HistoryQuery"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/HistoryList"
            }
          }
        }
      }
    },
    "/discord_server": {
      "get": {
        "description": "Get Discord Server",
        "tags": [
          "system"
        ],
        "operationId": "discord_server",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/DiscordServer"
            }
          }
        }
      }
    },
    "/login": {
      "get": {
        "description": "Login With Discord",
        "tags": [
          "user"
        ],
        "operationId": "login",
        "responses": {
          "302": {
            "description": "Redirect",
            "headers": {
              "Location": {
                "type": "string",
                "format": "url"
              }
            }
          }
        }
      }
    },
    "/open_discord_server": {
      "get": {
        "description": "Open Discord Server",
        "tags": [
          "system"
        ],
        "operationId": "open_discord_server",
        "responses": {
          "302": {
            "description": "Open Discord Servers",
            "headers": {
              "Location": {
                "type": "string",
                "format": "url"
              }
            }
          }
        }
      }
    },
    "/user_history": {
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "Get User History",
        "tags": [
          "user"
        ],
        "operationId": "user_history",
        "parameters": [
          {
            "description": "Page Info",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "properties": {
                "page_info": {
                  "$ref": "#/definitions/PageInfoRequest"
                },
                "query": {
                  "$ref": "#/definitions/HistoryQuery"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/HistoryList"
            }
          }
        }
      }
    },
    "/user_info": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "Get User Info",
        "tags": [
          "user"
        ],
        "operationId": "user_info",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/UserInfo"
            }
          }
        }
      }
    },
    "/user_list": {
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "Get User List",
        "tags": [
          "admin"
        ],
        "operationId": "user_list",
        "parameters": [
          {
            "description": "Page Info",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "properties": {
                "page_info": {
                  "$ref": "#/definitions/PageInfoRequest"
                },
                "query": {
                  "$ref": "#/definitions/UserQuery"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/UserList"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ClusterInfo": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32",
          "x-omitempty": false
        },
        "data": {
          "type": "object",
          "properties": {
            "cluster": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/NodeItem"
              }
            }
          },
          "x-omitempty": false
        },
        "message": {
          "type": "string"
        }
      }
    },
    "DiscordServer": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32",
          "x-omitempty": false
        },
        "data": {
          "type": "object",
          "properties": {
            "url": {
              "type": "string"
            }
          },
          "x-omitempty": false
        },
        "message": {
          "type": "string"
        }
      }
    },
    "HistoryItem": {
      "type": "object",
      "properties": {
        "command": {
          "type": "string"
        },
        "created": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "images": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "options": {
          "type": "object"
        },
        "user_avatar": {
          "type": "string"
        },
        "user_id": {
          "type": "string"
        },
        "user_name": {
          "type": "string"
        }
      }
    },
    "HistoryList": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32",
          "x-omitempty": false
        },
        "data": {
          "type": "object",
          "properties": {
            "history": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/HistoryItem"
              }
            },
            "page_info": {
              "$ref": "#/definitions/PageInfoResponse"
            }
          },
          "x-omitempty": false
        },
        "message": {
          "type": "string"
        }
      }
    },
    "HistoryQuery": {
      "type": "object",
      "properties": {
        "command": {
          "type": "string"
        }
      }
    },
    "NodeItem": {
      "type": "object",
      "properties": {
        "host": {
          "type": "string"
        },
        "max_concurrent": {
          "type": "integer",
          "format": "int32",
          "x-omitempty": false
        },
        "name": {
          "type": "string"
        },
        "pending": {
          "type": "integer",
          "format": "int32",
          "x-omitempty": false
        },
        "running": {
          "type": "integer",
          "format": "int32",
          "x-omitempty": false
        }
      }
    },
    "PageInfoRequest": {
      "type": "object",
      "properties": {
        "page": {
          "type": "integer",
          "format": "int32"
        },
        "page_size": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "PageInfoResponse": {
      "type": "object",
      "properties": {
        "page": {
          "type": "integer",
          "format": "int32"
        },
        "page_size": {
          "type": "integer",
          "format": "int32"
        },
        "total": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "UserInfo": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32",
          "x-omitempty": false
        },
        "data": {
          "type": "object",
          "properties": {
            "user": {
              "$ref": "#/definitions/UserItem"
            }
          },
          "x-omitempty": false
        },
        "message": {
          "type": "string"
        }
      }
    },
    "UserItem": {
      "type": "object",
      "properties": {
        "avatar": {
          "type": "string"
        },
        "created": {
          "type": "string"
        },
        "enable": {
          "type": "boolean"
        },
        "id": {
          "type": "string"
        },
        "roles": {
          "type": "string"
        },
        "stable_config": {
          "type": "object"
        },
        "username": {
          "type": "string"
        }
      }
    },
    "UserList": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32",
          "x-omitempty": false
        },
        "data": {
          "type": "object",
          "properties": {
            "page_info": {
              "$ref": "#/definitions/PageInfoResponse"
            },
            "users": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/UserItem"
              }
            }
          },
          "x-omitempty": false
        },
        "message": {
          "type": "string"
        }
      }
    },
    "UserQuery": {
      "type": "object",
      "properties": {
        "enable": {
          "type": "boolean"
        },
        "id": {
          "type": "string"
        },
        "username": {
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
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API",
    "title": "ApiService",
    "version": "1.0.0"
  },
  "host": "127.0.0.1:18890",
  "basePath": "/api",
  "paths": {
    "/auth": {
      "get": {
        "description": "Auth With Discord",
        "tags": [
          "user"
        ],
        "operationId": "auth",
        "parameters": [
          {
            "type": "string",
            "description": "Discord Code",
            "name": "code",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "Discord State",
            "name": "state",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "302": {
            "description": "Redirect",
            "headers": {
              "Location": {
                "type": "string",
                "format": "url"
              }
            }
          }
        }
      }
    },
    "/cluster": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "Get Cluster Info",
        "tags": [
          "system"
        ],
        "operationId": "cluster",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/ClusterInfo"
            }
          }
        }
      }
    },
    "/community_history": {
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "Get Community History",
        "tags": [
          "user"
        ],
        "operationId": "community_history",
        "parameters": [
          {
            "description": "Page Info",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "properties": {
                "page_info": {
                  "$ref": "#/definitions/PageInfoRequest"
                },
                "query": {
                  "$ref": "#/definitions/HistoryQuery"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/HistoryList"
            }
          }
        }
      }
    },
    "/discord_server": {
      "get": {
        "description": "Get Discord Server",
        "tags": [
          "system"
        ],
        "operationId": "discord_server",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/DiscordServer"
            }
          }
        }
      }
    },
    "/login": {
      "get": {
        "description": "Login With Discord",
        "tags": [
          "user"
        ],
        "operationId": "login",
        "responses": {
          "302": {
            "description": "Redirect",
            "headers": {
              "Location": {
                "type": "string",
                "format": "url"
              }
            }
          }
        }
      }
    },
    "/open_discord_server": {
      "get": {
        "description": "Open Discord Server",
        "tags": [
          "system"
        ],
        "operationId": "open_discord_server",
        "responses": {
          "302": {
            "description": "Open Discord Servers",
            "headers": {
              "Location": {
                "type": "string",
                "format": "url"
              }
            }
          }
        }
      }
    },
    "/user_history": {
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "Get User History",
        "tags": [
          "user"
        ],
        "operationId": "user_history",
        "parameters": [
          {
            "description": "Page Info",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "properties": {
                "page_info": {
                  "$ref": "#/definitions/PageInfoRequest"
                },
                "query": {
                  "$ref": "#/definitions/HistoryQuery"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/HistoryList"
            }
          }
        }
      }
    },
    "/user_info": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "Get User Info",
        "tags": [
          "user"
        ],
        "operationId": "user_info",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/UserInfo"
            }
          }
        }
      }
    },
    "/user_list": {
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "Get User List",
        "tags": [
          "admin"
        ],
        "operationId": "user_list",
        "parameters": [
          {
            "description": "Page Info",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "properties": {
                "page_info": {
                  "$ref": "#/definitions/PageInfoRequest"
                },
                "query": {
                  "$ref": "#/definitions/UserQuery"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/UserList"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ClusterInfo": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32",
          "x-omitempty": false
        },
        "data": {
          "type": "object",
          "properties": {
            "cluster": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/NodeItem"
              }
            }
          },
          "x-omitempty": false
        },
        "message": {
          "type": "string"
        }
      }
    },
    "ClusterInfoData": {
      "type": "object",
      "properties": {
        "cluster": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/NodeItem"
          }
        }
      },
      "x-omitempty": false
    },
    "DiscordServer": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32",
          "x-omitempty": false
        },
        "data": {
          "type": "object",
          "properties": {
            "url": {
              "type": "string"
            }
          },
          "x-omitempty": false
        },
        "message": {
          "type": "string"
        }
      }
    },
    "DiscordServerData": {
      "type": "object",
      "properties": {
        "url": {
          "type": "string"
        }
      },
      "x-omitempty": false
    },
    "HistoryItem": {
      "type": "object",
      "properties": {
        "command": {
          "type": "string"
        },
        "created": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "images": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "options": {
          "type": "object"
        },
        "user_avatar": {
          "type": "string"
        },
        "user_id": {
          "type": "string"
        },
        "user_name": {
          "type": "string"
        }
      }
    },
    "HistoryList": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32",
          "x-omitempty": false
        },
        "data": {
          "type": "object",
          "properties": {
            "history": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/HistoryItem"
              }
            },
            "page_info": {
              "$ref": "#/definitions/PageInfoResponse"
            }
          },
          "x-omitempty": false
        },
        "message": {
          "type": "string"
        }
      }
    },
    "HistoryListData": {
      "type": "object",
      "properties": {
        "history": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/HistoryItem"
          }
        },
        "page_info": {
          "$ref": "#/definitions/PageInfoResponse"
        }
      },
      "x-omitempty": false
    },
    "HistoryQuery": {
      "type": "object",
      "properties": {
        "command": {
          "type": "string"
        }
      }
    },
    "NodeItem": {
      "type": "object",
      "properties": {
        "host": {
          "type": "string"
        },
        "max_concurrent": {
          "type": "integer",
          "format": "int32",
          "x-omitempty": false
        },
        "name": {
          "type": "string"
        },
        "pending": {
          "type": "integer",
          "format": "int32",
          "x-omitempty": false
        },
        "running": {
          "type": "integer",
          "format": "int32",
          "x-omitempty": false
        }
      }
    },
    "PageInfoRequest": {
      "type": "object",
      "properties": {
        "page": {
          "type": "integer",
          "format": "int32"
        },
        "page_size": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "PageInfoResponse": {
      "type": "object",
      "properties": {
        "page": {
          "type": "integer",
          "format": "int32"
        },
        "page_size": {
          "type": "integer",
          "format": "int32"
        },
        "total": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "UserInfo": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32",
          "x-omitempty": false
        },
        "data": {
          "type": "object",
          "properties": {
            "user": {
              "$ref": "#/definitions/UserItem"
            }
          },
          "x-omitempty": false
        },
        "message": {
          "type": "string"
        }
      }
    },
    "UserInfoData": {
      "type": "object",
      "properties": {
        "user": {
          "$ref": "#/definitions/UserItem"
        }
      },
      "x-omitempty": false
    },
    "UserItem": {
      "type": "object",
      "properties": {
        "avatar": {
          "type": "string"
        },
        "created": {
          "type": "string"
        },
        "enable": {
          "type": "boolean"
        },
        "id": {
          "type": "string"
        },
        "roles": {
          "type": "string"
        },
        "stable_config": {
          "type": "object"
        },
        "username": {
          "type": "string"
        }
      }
    },
    "UserList": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32",
          "x-omitempty": false
        },
        "data": {
          "type": "object",
          "properties": {
            "page_info": {
              "$ref": "#/definitions/PageInfoResponse"
            },
            "users": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/UserItem"
              }
            }
          },
          "x-omitempty": false
        },
        "message": {
          "type": "string"
        }
      }
    },
    "UserListData": {
      "type": "object",
      "properties": {
        "page_info": {
          "$ref": "#/definitions/PageInfoResponse"
        },
        "users": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/UserItem"
          }
        }
      },
      "x-omitempty": false
    },
    "UserQuery": {
      "type": "object",
      "properties": {
        "enable": {
          "type": "boolean"
        },
        "id": {
          "type": "string"
        },
        "username": {
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
