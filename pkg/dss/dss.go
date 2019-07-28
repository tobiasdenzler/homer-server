package dss

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/tobiasdenzler/homer-server/config"
)

var client *http.Client
var sessionToken string

func init() {
	// Create http client
	// Ignore certificate of Digitalstrom server
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: tr}
}

// Login connects to the DSS and login to the server using the login token.
func login() error {

	log.Debug("Login to DSS and retrieve session token")

	// We need to get a valid session token if it does not already exist
	if sessionToken == "" {
		loginResponse, err := client.Get(config.Config.Server.Address + "/json/system/loginApplication?loginToken=" + config.Config.Server.LoginToken)

		if err != nil {
			return errors.New("Failed to login to DSS -> " + err.Error())
		}
		data, _ := ioutil.ReadAll(loginResponse.Body)
		log.Tracef("Result from %s -> %s", "/json/system/loginApplication", string(data))

		// {"result":{"token":"9267..."},"ok":true}
		// Extract the session token from the response
		var jsonData map[string]interface{}
		err = json.Unmarshal(data, &jsonData)
		if err != nil {
			return errors.New("Failed to unmarshal the result from DSS -> " + err.Error())
		}
		result := jsonData["result"].(map[string]interface{})

		// DSS might also return ok = false
		if !jsonData["ok"].(bool) {
			return errors.New("DSS returns false from loginApplication")
		}
		sessionToken = result["token"].(string)
		log.Tracef("Current session token is -> %s", sessionToken)

		defer loginResponse.Body.Close()
	}
	return nil
}

// This will take a map of query parameters and format them.
// Format: key1=value1&key2=value2
func createParamString(params map[string]string) string {
	b := new(bytes.Buffer)
	for key, value := range params {
		fmt.Fprintf(b, "%s=%s&", key, value)
	}
	return strings.TrimSuffix(b.String(), "&")
}

// Call will request the DSS with a path and a map of request parameters.
// The function will return a []byte with the JSON result from the request.
func Call(path string, params map[string]string) []byte {

	log.Debugf("Call to DSS for path %s with parameters %s", path, params)

	var result []byte

	// Login to the server
	err := login()
	if err != nil {
		log.Panic("Not able to login to DSS -> ", err)
	}

	// Create new http request using the session token
	req, err := http.NewRequest("GET", config.Config.Server.Address+path+"?token="+sessionToken+"&"+createParamString(params), nil)
	req.Header.Add("Accept", "application/json")

	response, err := client.Do(req)
	if err != nil {
		log.Panicf("Failed to request %s from DSS -> %s", path, err)
	} else {
		result, _ := ioutil.ReadAll(response.Body)
		log.Tracef("Result from %s -> %s", path, string(result))
	}
	defer response.Body.Close()

	return result
}
