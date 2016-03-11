/*
 * Copyright 2016 Fabrício Godoy
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package oauth

// A TokenAdapter provides an adapter for token management.
type TokenAdapter interface {
	// AccessToken creates and returns a new access token.
	AccessToken(c *TokenContext) *TokenResponse

	// Client gets the client information if valid.
	Client(c *TokenContext) *ClientEntry

	// Refresh validate provided refresh token.
	Refresh(c *TokenContext) bool

	// SupportedGrantTypes gets a list of supported grant types.
	SupportedGrantTypes() []string

	// User validate resource owner credentials for password grant type.
	User(c *TokenContext) bool
}
