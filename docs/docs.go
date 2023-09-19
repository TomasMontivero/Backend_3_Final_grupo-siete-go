// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://developers.ctd.com.ar/es_ar/terminos-y-condiciones",
        "contact": {
            "name": "Rocio Belen Ghillino, Tomás Montivero, Agustin Damelio and Nicolás Gambino",
            "url": "https://github.com/rrrrho/grupo-siete-go"
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
        "/odontologos": {
            "put": {
                "description": "Replaces an existing odontologo from the repository",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "odontologos"
                ],
                "summary": "Replaces an odontologo",
                "parameters": [
                    {
                        "description": "New odontologo used to replace",
                        "name": "odontologo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/odontologo.Odontologo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/odontologo.Odontologo"
                        }
                    }
                }
            },
            "post": {
                "description": "Saves a odontologo into the repository",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "odontologos"
                ],
                "summary": "Saves a odontologo",
                "parameters": [
                    {
                        "description": "Odontologo to save",
                        "name": "odontologo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/odontologo.Odontologo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/odontologo.Odontologo"
                        }
                    }
                }
            },
            "patch": {
                "description": "Updates an existing odontologo from the repository with one o more features",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "odontologos"
                ],
                "summary": "Udpates an odontologo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Odontologo's ID in order to update from it",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/odontologo.Odontologo"
                        }
                    }
                }
            }
        },
        "/odontologos/{id}": {
            "get": {
                "description": "Gets a odontologo by id from the repository",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "odontologos"
                ],
                "summary": "Gets a odontologo by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Odontologo's ID in order to get from it",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/odontologo.Odontologo"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes an existing odontologo from the repository",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "odontologos"
                ],
                "summary": "Deletes an odontologo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Odontologo's ID in order to delete from it",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Odontologo eliminado exitosamente"
                    }
                }
            }
        },
        "/pacientes": {
            "post": {
                "description": "Saves a paciente into the repository",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "paciente"
                ],
                "summary": "Saves a paciente",
                "parameters": [
                    {
                        "description": "Paciente to save",
                        "name": "paciente",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/paciente.Paciente"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/paciente.Paciente"
                        }
                    }
                }
            },
            "patch": {
                "description": "Updates an existing paciente from the repository with one o more features",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "paciente"
                ],
                "summary": "Udpates a paciente",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Paciente's ID in order to update from it",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/paciente.Paciente"
                        }
                    }
                }
            }
        },
        "/pacientes/{id}": {
            "get": {
                "description": "Gets a paciente by id from the repository",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "paciente"
                ],
                "summary": "Gets a paciente by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Paciente's ID in order to get from it",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/paciente.Paciente"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes an existing paciente from the repository",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "paciente"
                ],
                "summary": "Deletes a paciente",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Paciente's ID in order to delete from it",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Paciente eliminado exitosamente"
                    }
                }
            }
        },
        "/turnos": {
            "get": {
                "description": "Gets a turno by paciente dni from the repository",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "turno"
                ],
                "summary": "Gets a turno by paciente DNI",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Paciente's DNI in order to get Turnos from it",
                        "name": "dni",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/turno.Turno"
                        }
                    }
                }
            },
            "put": {
                "description": "Replaces an existing turno from the repository",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "turno"
                ],
                "summary": "Replaces a turno",
                "parameters": [
                    {
                        "description": "New turno used to replace",
                        "name": "turno",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/turno.Turno"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/turno.Turno"
                        }
                    }
                }
            },
            "post": {
                "description": "Saves a turno into the repository",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "turno"
                ],
                "summary": "Saves a turno",
                "parameters": [
                    {
                        "description": "Turno to save",
                        "name": "turno",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/turno.Turno"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/turno.Turno"
                        }
                    }
                }
            },
            "patch": {
                "description": "Updates an existing turno from the repository with one o more features",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "turno"
                ],
                "summary": "Udpates a turno",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Turno's ID in order to update from it",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/turno.Turno"
                        }
                    }
                }
            }
        },
        "/turnos/{id}": {
            "get": {
                "description": "Gets a turno by id from the repository",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "turno"
                ],
                "summary": "Gets a turno by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Turno's ID in order to get from it",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/turno.Turno"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes an existing turno from the repository",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "turno"
                ],
                "summary": "Deletes a turno",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Turno's ID in order to delete from it",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Turnos eliminado exitosamente"
                    }
                }
            }
        }
    },
    "definitions": {
        "odontologo.Odontologo": {
            "type": "object",
            "properties": {
                "apellido": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "matricula": {
                    "type": "string"
                },
                "nombre": {
                    "type": "string"
                }
            }
        },
        "paciente.Paciente": {
            "type": "object",
            "properties": {
                "apellido": {
                    "type": "string"
                },
                "dni": {
                    "type": "string"
                },
                "domicilio": {
                    "type": "string"
                },
                "fecha_de_alta": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "nombre": {
                    "type": "string"
                }
            }
        },
        "turno.Turno": {
            "type": "object",
            "properties": {
                "descripcion": {
                    "type": "string"
                },
                "fecha_y_hora": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "odontologo": {
                    "$ref": "#/definitions/odontologo.Odontologo"
                },
                "paciente": {
                    "$ref": "#/definitions/paciente.Paciente"
                }
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
	Title:            "Clinica Odontologia Back End 3 - Grupo 7 - Certified Tech Developer - Digital House",
	Description:      "This API Handle Pacients, Dentist and Appointments.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
