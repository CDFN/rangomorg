#RangomOrg
RangomOrg is wrapper for random.org API. It's very simple to use!<br>
Dependencies:
Name    |Command
--------|-------
jsonrpc | github.com/ybbus/jsonrpc

Usage:
```Go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/CDFN/rangomorg"
	"log"
)

const (
	randomApiKey = "your-api-key"
)

func main() {
	random := rangomorg.New(randomApiKey)
	result, err := random.GenerateSignedStrings(5, 10, "rangom", map[string]interface{}{
		"userData": "YourUserData", // These options are optional
		"replacement": true,        // see https://api.random.org/json-rpc/2 for more
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	jsonBytes, _ := json.MarshalIndent(result, "", "    ")
	fmt.Println("Random: ")
	fmt.Println(string(jsonBytes))                      // Display result in json form
	fmt.Println("Requested data: ", result.Random.Data) // Display requested data
	fmt.Println("Signature: ", result.Signature)        // In case of signed api, display signature
}
```