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

package ginhttp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"gopkg.in/raiqub/oauth.v2"
	"gopkg.in/raiqub/oauth.v2/oauthtest"
)

const (
	FormKeyCustom    = "my_key"
	FormValueCustom  = "any_custom_value"
	ListenerEndpoint = "/token"
)

func TestClientGrant(t *testing.T) {
	adapter := oauthtest.NewTokenAdapter()
	srv := NewTokenServer(adapter, oauth.GrantTypeClient)
	adapter.CustomValues[FormKeyCustom] = []string{FormValueCustom}

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.RedirectTrailingSlash = true

	router.POST(ListenerEndpoint, srv.AccessTokenRequest)
	ts := httptest.NewServer(router)
	defer ts.Close()

	client := &http.Client{}
	form := url.Values{
		oauth.FormKeyGrantType: []string{oauth.GrantTypeClient},
		oauth.FormKeyScope:     []string{adapter.Scope},
		FormKeyCustom:          []string{FormValueCustom},
	}
	req, _ := http.NewRequest("POST", ts.URL+ListenerEndpoint,
		strings.NewReader(form.Encode()))
	req.SetBasicAuth(adapter.ClientID, adapter.ClientSecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Error posting data to server: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Unexpected HTTP status: %s", resp.Status)
		strBody, _ := ioutil.ReadAll(resp.Body)
		t.Fatalf("Response body: %s", string(strBody))
	}

	var response oauth.TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Error parsing response: %v", err)
	}

	if response.AccessToken != adapter.AccessToken {
		t.Fatalf("Unexpected response: %#v", response)
	}
}
