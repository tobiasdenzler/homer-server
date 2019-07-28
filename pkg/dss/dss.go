package dss

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

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

	var err error

	// We need to get a valid session token if it does not already exist
	if sessionToken != "" {
		log.Println("Session token already exists ->", sessionToken)
	} else {
		log.Println("Generating a new session token")
		loginRespose, err := client.Get(config.Config.Server.Address + "/json/system/loginApplication?loginToken=" + config.Config.Server.LoginToken)

		if err == nil {
			data, _ := ioutil.ReadAll(loginRespose.Body)
			log.Println(string(data))

			// {"result":{"token":"9267..."},"ok":true}
			// Extract the session token from the response
			var jsonData map[string]interface{}
			err := json.Unmarshal(data, &jsonData)
			if err == nil {
				result := jsonData["result"].(map[string]interface{})
				sessionToken = result["token"].(string)
				log.Println("current session token ->", sessionToken)
			}
		}
		defer loginRespose.Body.Close()
	}

	return err
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
// The function will return a string with the Json result from the request.
func Call(path string, params map[string]string) (string, error) {

	var result string

	// Login to the server
	err := login()

	if err == nil {

		// Create new http request and add the session token in the header
		req, err := http.NewRequest("GET", config.Config.Server.Address+path+"?token="+sessionToken+"&"+createParamString(params), nil)
		req.Header.Add("Accept", "application/json")

		response, err := client.Do(req)
		if err == nil {
			read, _ := ioutil.ReadAll(response.Body)
			result = string(read)
		}
		defer response.Body.Close()
	}

	return result, err
}
