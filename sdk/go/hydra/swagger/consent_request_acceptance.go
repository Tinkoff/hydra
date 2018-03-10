/*
 * ORY Hydra - Cloud Native OAuth 2.0 and OpenID Connect Server
 *
 * Welcome to the ORY Hydra HTTP API documentation. You will find documentation for all HTTP APIs here. Keep in mind that this document reflects the latest branch, always. Support for versioned documentation is coming in the future.
 *
 * OpenAPI spec version: Latest
 * Contact: hi@ory.am
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type ConsentRequestAcceptance struct {

	// AccessTokenExtra represents arbitrary data that will be added to the access token and that will be returned on introspection and warden requests.
	AccessTokenExtra map[string]interface{} `json:"accessTokenExtra,omitempty"`

	// A list of scopes that the user agreed to grant. It should be a subset of requestedScopes from the consent request.
	GrantScopes []string `json:"grantScopes,omitempty"`

	// IDTokenExtra represents arbitrary data that will be added to the ID token. The ID token will only be issued if the user agrees to it and if the client requested an ID token.
	IdTokenExtra map[string]interface{} `json:"idTokenExtra,omitempty"`

	// Subject represents a unique identifier of the user (or service, or legal entity, ...) that accepted the OAuth2 request.
	Subject string `json:"subject,omitempty"`
}