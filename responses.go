package rangomorg

type License struct {
	LicenseType string `json:"type"`
	Text        string `json:"text"`
	InfoUrl     string `json:"infoUrl"`
}

type SignedIntegersResult struct {
	Random        SignedIntegersRandom `json:"random"`
	Signature     string               `json:"signature"`
	BitsUsed      int                  `json:"bitsUsed"`
	BitsLeft      int                  `json:"bitsLeft"`
	RequestsLeft  int                  `json:"requestsLeft"`
	AdvisoryDelay int                  `json:"advisoryDelay"`
}
type SignedIntegersRandom struct {
	Method         string  `json:"method"`
	HashedApiKey   string  `json:"hashedApiKey"`
	N              int     `json:"n"`
	Min            int     `json:"min"`
	Max            int     `json:"max"`
	Replacement    bool    `json:"replacement"`
	Base           int     `json:"base"`
	Data           []int   `json:"data"`
	License        License `json:"license"`
	UserData       string  `json:"userData"`
	SerialNumber   int     `json:"serialNumber"`
	CompletionTime string  `json:"completionTime"`
}

type SignedIntegerSequencesResult struct {
	Random        SignedIntegerSequencesRandom `json:"random"`
	Signature     string                       `json:"signature"`
	BitsUsed      int                          `json:"bitsUsed"`
	BitsLeft      int                          `json:"bitsLeft"`
	RequestsLeft  int                          `json:"requestsLeft"`
	AdvisoryDelay int                          `json:"advisoryDelay"`
}
type SignedIntegerSequencesRandom struct {
	Method         string  `json:"method"`
	HashedApiKey   string  `json:"hashedApiKey"`
	N              int     `json:"n"`
	Min            int     `json:"min"`
	Max            int     `json:"max"`
	Length         int     `json:"length"`
	Replacement    bool    `json:"replacement"`
	Base           int     `json:"base"`
	Data           [][]int `json:"data"`
	License        License `json:"license"`
	UserData       string  `json:"userData"`
	SerialNumber   int     `json:"serialNumber"`
	CompletionTime string  `json:"completionTime"`
}

type SignedDecimalFractionsResult struct {
	Random        SignedDecimalFractionsRandom `json:"random"`
	Signature     string                       `json:"signature"`
	BitsUsed      int                          `json:"bitsUsed"`
	BitsLeft      int                          `json:"bitsLeft"`
	RequestsLeft  int                          `json:"requestsLeft"`
	AdvisoryDelay int                          `json:"advisoryDelay"`
}
type SignedDecimalFractionsRandom struct {
	Method         string    `json:"method"`
	HashedApiKey   string    `json:"hashedApiKey"`
	N              int       `json:"n"`
	DecimalPlaces  int       `json:"decimalPlaces"`
	Replacement    bool      `json:"replacement"`
	Data           []float32 `json:"data"`
	License        License   `json:"license"`
	UserData       string    `json:"userData"`
	SerialNumber   int       `json:"serialNumber"`
	CompletionTime string    `json:"completionTime"`
}

type SignedGaussiansResult struct {
	Random        SignedGaussiansRandom `json:"random"`
	Signature     string                `json:"signature"`
	BitsUsed      int                   `json:"bitsUsed"`
	BitsLeft      int                   `json:"bitsLeft"`
	RequestsLeft  int                   `json:"requestsLeft"`
	AdvisoryDelay int                   `json:"advisoryDelay"`
}
type SignedGaussiansRandom struct {
	Method            string    `json:"method"`
	HashedApiKey      string    `json:"hashedApiKey"`
	N                 int       `json:"n"`
	Mean              int       `json:"mean"`
	StandardDeviation int       `json:"standardDeviation"`
	SignificantDigits int       `json:"significantDigits"`
	Replacement       bool      `json:"replacement"`
	Data              []float32 `json:"data"`
	License           License   `json:"license"`
	UserData          string    `json:"userData"`
	SerialNumber      int       `json:"serialNumber"`
	CompletionTime    string    `json:"completionTime"`
}

type SignedStringsResult struct {
	Random        SignedStringsRandom `json:"random"`
	Signature     string              `json:"signature"`
	BitsUsed      int                 `json:"bitsUsed"`
	BitsLeft      int                 `json:"bitsLeft"`
	RequestsLeft  int                 `json:"requestsLeft"`
	AdvisoryDelay int                 `json:"advisoryDelay"`
}
type SignedStringsRandom struct {
	Method         string   `json:"method"`
	HashedApiKey   string   `json:"hashedApiKey"`
	N              int      `json:"n"`
	Length         int      `json:"length"`
	Characters     string   `json:"characters"`
	Replacement    bool     `json:"replacement"`
	Base           int      `json:"base"`
	Data           []string `json:"data"`
	License        License  `json:"license"`
	UserData       string   `json:"userData"`
	SerialNumber   int      `json:"serialNumber"`
	CompletionTime string   `json:"completionTime"`
}
