// Code generated by paramgen. DO NOT EDIT.
// Source: github.com/ConduitIO/conduit-connector-sdk/tree/main/cmd/paramgen

package activemq

import (
	sdk "github.com/conduitio/conduit-connector-sdk"
)

func (SourceConfig) Parameters() map[string]sdk.Parameter {
	return map[string]sdk.Parameter{
		"contentType": {
			Default:     "text/plain",
			Description: "contentType is the content type of the message.",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{},
		},
		"password": {
			Default:     "",
			Description: "password is the password to use when connecting to the broker.",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{
				sdk.ValidationRequired{},
			},
		},
		"queue": {
			Default:     "",
			Description: "queue is the name of the queue to read from.",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{
				sdk.ValidationRequired{},
			},
		},
		"tlsConfig.caCertPath": {
			Default:     "",
			Description: "caCertPath is the path to the CA certificate file.",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{},
		},
		"tlsConfig.clientCertPath": {
			Default:     "",
			Description: "clientCertPath is the path to the client certificate file.",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{},
		},
		"tlsConfig.clientKeyPath": {
			Default:     "",
			Description: "clientKeyPath is the path to the client key file.",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{},
		},
		"tlsConfig.useTLS": {
			Default:     "false",
			Description: "useTLS is a flag to enable or disable TLS.",
			Type:        sdk.ParameterTypeBool,
			Validations: []sdk.Validation{},
		},
		"url": {
			Default:     "",
			Description: "url is the url of the ActiveMQ classic broker.",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{
				sdk.ValidationRequired{},
			},
		},
		"user": {
			Default:     "",
			Description: "user is the username to use when connecting to the broker.",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{
				sdk.ValidationRequired{},
			},
		},
	}
}
