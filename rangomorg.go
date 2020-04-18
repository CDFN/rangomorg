package rangomorg

import (
	"github.com/ybbus/jsonrpc"
)

const ENDPOINT = "https://api.random.org/json-rpc/2/invoke"

type Rangom struct {
	ApiKey    string
	RPCClient jsonrpc.RPCClient
}

func New(apiKey string) Rangom {
	return Rangom{ApiKey: apiKey, RPCClient: jsonrpc.NewClient(ENDPOINT)}
}

func (r *Rangom) GenerateSignedIntegers(n int, min int, max int, optional map[string]interface{}) (SignedIntegersResult, error) {

	out := SignedIntegersResult{}

	params := map[string]interface{}{
		"apiKey": r.ApiKey,
		"n":      n,
		"min":    min,
		"max":    max,
	}

	if optional != nil {
		for k, v := range optional {
			params[k] = v
		}
	}

	err := r.RPCClient.CallFor(&out, "generateSignedIntegers", params)

	if err != nil {
		return out, err
	}

	return out, nil
}
func (r *Rangom) GenerateSignedIntegerSequences(n int, length int, min int, max int, optional map[string]interface{}) (SignedIntegerSequencesResult, error) {

	out := SignedIntegerSequencesResult{}

	params := map[string]interface{}{
		"apiKey": r.ApiKey,
		"n":      n,
		"length": length,
		"min":    min,
		"max":    max,
	}

	if optional != nil {
		for k, v := range optional {
			params[k] = v
		}
	}

	err := r.RPCClient.CallFor(&out, "generateSignedIntegerSequences", params)

	if err != nil {
		return out, err
	}

	return out, nil
}
func (r *Rangom) GenerateSignedDecimalFractions(n int, decimalPlaces int, optional map[string]interface{}) (SignedDecimalFractionsResult, error) {

	out := SignedDecimalFractionsResult{}

	params := map[string]interface{}{
		"apiKey":        r.ApiKey,
		"n":             n,
		"decimalPlaces": decimalPlaces,
	}

	if optional != nil {
		for k, v := range optional {
			params[k] = v
		}
	}

	err := r.RPCClient.CallFor(&out, "generateSignedDecimalFractions", params)

	if err != nil {
		return out, err
	}

	return out, nil
}
func (r *Rangom) GenerateSignedGaussians(n int, mean int, standardDeviation int, significantDigits int, optional map[string]interface{}) (SignedGaussiansResult, error) {

	out := SignedGaussiansResult{}

	params := map[string]interface{}{
		"apiKey":            r.ApiKey,
		"n":                 n,
		"mean":              mean,
		"standardDeviation": standardDeviation,
		"significantDigits": significantDigits,
	}

	if optional != nil {
		for k, v := range optional {
			params[k] = v
		}
	}

	err := r.RPCClient.CallFor(&out, "generateSignedGaussians", params)

	if err != nil {
		return out, err
	}

	return out, nil
}
func (r *Rangom) GenerateSignedStrings(n int, length int, characters string, optional map[string]interface{}) (SignedStringsResult, error) {
	out := SignedStringsResult{}

	params := map[string]interface{}{
		"apiKey":     r.ApiKey,
		"n":          n,
		"length":     length,
		"characters": characters,
	}

	if optional != nil {
		for k, v := range optional {
			params[k] = v
		}
	}

	err := r.RPCClient.CallFor(&out, "generateSignedStrings", params)

	if err != nil {
		return out, err
	}

	return out, nil
}
func (r *Rangom) GenerateSignedUUIDs(n int, optional map[string]interface{}) (SignedUUIDsResult, error) {
	out := SignedUUIDsResult{}

	params := map[string]interface{}{
		"apiKey": r.ApiKey,
		"n":      n,
	}

	if optional != nil {
		for k, v := range optional {
			params[k] = v
		}
	}

	err := r.RPCClient.CallFor(&out, "generateSignedUUIDs", params)

	if err != nil {
		return out, err
	}

	return out, nil
}
func (r *Rangom) GenerateSignedBlobs(n int, size int, optional map[string]interface{}) (SignedBlobsResult, error) {
	out := SignedBlobsResult{}

	params := map[string]interface{}{
		"apiKey": r.ApiKey,
		"n":      n,
		"size":   size,
	}

	if optional != nil {
		for k, v := range optional {
			params[k] = v
		}
	}

	err := r.RPCClient.CallFor(&out, "generateSignedBlobs", params)

	if err != nil {
		return out, err
	}

	return out, nil
}
