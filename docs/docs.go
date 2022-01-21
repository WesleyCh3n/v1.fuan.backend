// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Wesley",
            "email": "wesley.ch3n.0530@gmail.com"
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
        "/api/concat": {
            "put": {
                "description": "return concat selection csvs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Python"
                ],
                "summary": "concat 2 selection csv",
                "operationId": "concat_selection_file",
                "parameters": [
                    {
                        "description": "files need to be concated",
                        "name": "files",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.ResConcat"
                        }
                    }
                }
            }
        },
        "/api/export": {
            "put": {
                "description": "return processed selection csv",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Python"
                ],
                "summary": "export selection files",
                "operationId": "export_selection_file",
                "parameters": [
                    {
                        "description": "filtered files",
                        "name": "FltrFile",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.FltrFile"
                        }
                    },
                    {
                        "description": "selected ranges",
                        "name": "RangeIndex",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Range"
                            }
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.ResExport"
                        }
                    }
                }
            }
        },
        "/api/save": {
            "patch": {
                "description": "Save selected range in raw file",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Python"
                ],
                "summary": "Save selected range in raw file",
                "operationId": "save_selected_range",
                "parameters": [
                    {
                        "description": "Original file",
                        "name": "UploadFile",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "range(string)  to  write  in  csv  column",
                        "name": "Range",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/upload": {
            "post": {
                "description": "upload raw csv and return filtered csvs",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Python"
                ],
                "summary": "Create filtered files",
                "operationId": "upload_create_filtered_data",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Upload file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.ResUpload"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ConcatFile": {
            "type": "object",
            "properties": {
                "ConcatFile": {
                    "type": "string"
                }
            }
        },
        "models.ExportFile": {
            "type": "object",
            "properties": {
                "ExportFile": {
                    "type": "string",
                    "example": "exportfile.csv"
                }
            }
        },
        "models.Fltr": {
            "type": "object",
            "properties": {
                "FltrFile": {
                    "$ref": "#/definitions/models.FltrFile"
                },
                "Range": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Range"
                    }
                }
            }
        },
        "models.FltrFile": {
            "type": "object",
            "properties": {
                "cyDb": {
                    "type": "string",
                    "example": "cyDb.csv"
                },
                "cyGt": {
                    "type": "string",
                    "example": "cyGt.csv"
                },
                "cyLt": {
                    "type": "string",
                    "example": "cyLt.csv"
                },
                "cyRt": {
                    "type": "string",
                    "example": "cyRt.csv"
                },
                "rslt": {
                    "type": "string",
                    "example": "rslt.csv"
                }
            }
        },
        "models.Range": {
            "type": "object",
            "properties": {
                "End": {
                    "type": "number"
                },
                "Start": {
                    "type": "number"
                }
            }
        },
        "models.ResConcat": {
            "type": "object",
            "properties": {
                "python": {
                    "$ref": "#/definitions/models.ConcatFile"
                },
                "saveDir": {
                    "type": "string",
                    "example": "file/example"
                },
                "serverRoot": {
                    "type": "string",
                    "example": "http://example.com:3000"
                }
            }
        },
        "models.ResExport": {
            "type": "object",
            "properties": {
                "python": {
                    "$ref": "#/definitions/models.ExportFile"
                },
                "saveDir": {
                    "type": "string",
                    "example": "file/example"
                },
                "serverRoot": {
                    "type": "string",
                    "example": "http://example.com:3000"
                }
            }
        },
        "models.ResUpload": {
            "type": "object",
            "properties": {
                "python": {
                    "$ref": "#/definitions/models.Fltr"
                },
                "saveDir": {
                    "type": "string",
                    "example": "file/example"
                },
                "serverRoot": {
                    "type": "string",
                    "example": "http://example.com:3000"
                },
                "uploadFile": {
                    "type": "string",
                    "example": "example_raw.csv"
                }
            }
        }
    },
    "tags": [
        {
            "description": "running python process api",
            "name": "Python"
        }
    ]
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
	Host:        "localhost:3001",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "analyze API",
	Description: "analyze python backend",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
	swag.Register("swagger", &s{})
}
