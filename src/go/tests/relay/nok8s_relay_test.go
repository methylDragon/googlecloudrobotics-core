// Copyright 2022 The Cloud Robotics Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"testing"
	"time"
)

const (
	RelayClientPath = "src/go/cmd/http-relay-client/http-relay-client-bin_/http-relay-client-bin"
	RelayServerPath = "src/go/cmd/http-relay-server/http-relay-server-bin_/http-relay-server-bin"
)

var (
	RelayClientArgs = []string{
		"--backend_scheme=http",
		"--relay_scheme=http",
		"--server_name=remote1",
		"--disable_auth_for_remote",
	}
	RelayServerArgs = []string{
		"--port=0",
	}
	rsPortMatcher = regexp.MustCompile(`Relay server listening on: 127.0.0.1:(\d\d*)\n$`)
)

// TestHttpRelay launches a local http relay (client + server) and connects a
// test-hhtp-server as a backend. The test is then interacting with the backend
// through the local relay.
func TestHttpRelay(t *testing.T) {
	tests := []struct {
		desc       string
		urlPath    string
		statusCode int
		body       string
	}{
		{
			desc:       "simple get",
			urlPath:    "/client/remote1/",
			statusCode: http.StatusOK,
			body:       "Hello",
		},
		{
			desc:       "backend status is preserved",
			urlPath:    "/client/remote1/bad-path",
			statusCode: http.StatusNotFound,
			body:       "",
		},
		{
			desc:       "invalid client",
			urlPath:    "/client/wrong/",
			statusCode: http.StatusInternalServerError,
			body:       "doesn't appear to be running the relay client",
		},
	}

	// setup http test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			fmt.Fprintln(w, "Hello")
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer ts.Close()

	// run relay server exposing the relay client
	var rsOut bytes.Buffer
	rs := exec.Command(RelayServerPath, RelayServerArgs...)
	rs.Stdout = os.Stdout
	rs.Stderr = io.MultiWriter(os.Stderr, &rsOut)
	if err := rs.Start(); err != nil {
		t.Fatal("failed to start relay-server: ", err)
	}
	rsPort := ""
	for i := 0; i < 10; i++ {
		if m := rsPortMatcher.FindStringSubmatch(rsOut.String()); m != nil {
			rsPort = m[1]
			log.Printf("Server port: %s", rsPort)
			break
		}
		log.Print("Waiting for relay to be up-and-running ...")
		time.Sleep(1 * time.Second)
	}
	if rsPort == "" {
		t.Fatal("timeout waiting for relay-server to launch")
	}

	// run relay client exposing the test-backend
	rcArgs := append(RelayClientArgs, []string{
		fmt.Sprintf("--backend_address=%s", strings.TrimPrefix(ts.URL, "http://")),
		"--relay_address=127.0.0.1:" + rsPort,
	}...)

	rc := exec.Command(RelayClientPath, rcArgs...)
	rc.Stdout = os.Stdout
	rc.Stderr = os.Stderr
	if err := rc.Start(); err != nil {
		t.Fatal("failed to start relay-client: ", err)
	}

	connected := false
	for i := 0; i < 10; i++ {
		if strings.Contains(rsOut.String(), "Relay client connected") {
			connected = true
			break
		}
		log.Print("Waiting for relay to be up-and-running ...")
		time.Sleep(1 * time.Second)
	}
	if !connected {
		t.Fatal("timeout waiting for relay-cleint to connect to relay-server")
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			res, err := http.Get("http://127.0.0.1:" + rsPort + tc.urlPath)
			if err != nil {
				t.Fatal(err)
			}
			body, err := io.ReadAll(res.Body)
			res.Body.Close()
			if err != nil {
				t.Fatal(err)
			}

			if res.StatusCode != tc.statusCode {
				t.Errorf("Wrong status code - got %d, expected %d", res.StatusCode, tc.statusCode)
			}
			if !strings.Contains(string(body), tc.body) {
				t.Errorf("Wrong body - got %q, expected it to contain %q, ", body, tc.body)
			}
		})
	}

	// tear down relay
	if err := rs.Process.Kill(); err != nil {
		t.Fatal("failed to kill relay-server: ", err)
	}
	if err := rc.Process.Kill(); err != nil {
		t.Fatal("failed to kill relay-client: ", err)
	}
}