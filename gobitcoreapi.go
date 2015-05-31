package gobitcoreapi

import (
    "crypto/tls"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "strconv"
    "bytes"
)

type API struct {
    endPoint      string
    version       string
    client        *http.Client
}

func NewAPI(endPoint string) *API {
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{Transport: tr}

    return &API{endPoint, "v1", client}
}

func (this *API) call(action, httpMethod string, params map[string]string) ([]byte, error) {

    var err error
    var res *http.Response

    if httpMethod=="POST" {

        jsonString, jsonError := json.Marshal(params)
        if jsonError!=nil {
            return nil, jsonError
        }

        req, _ := http.NewRequest("POST", this.endPoint+"/"+this.version+"/"+action, bytes.NewBuffer([]byte(jsonString)))
        res, err = this.client.Do(req)

    } else {
        valuesStr := ""
        for key, val := range params {
            valuesStr += "&"+key+"="+val
        }
        res, err = this.client.Get(this.endPoint+"/"+this.version+"/"+action+"?"+valuesStr)
    }
    body, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
    return body, err
}

func (this *API) SetVersion(version string) {
    this.version = version
}

func (this *API) Node(address string) (interface{}, error) {
    dataStream, err := this.call("node", "GET", nil)
    data := map[string]interface{}{}
    json.Unmarshal(dataStream, &data)
    return data, err
}

func (this *API) Blocks( from, to, offset, limit int) (interface{}, error) {
    if to==0 {
        to = 1000000
    }
    if limit==0 {
        limit = 10
    }
    var params = map[string]string{
        "from": strconv.Itoa(from),
        "to": strconv.Itoa(to),
        "offset": strconv.Itoa(offset),
        "limit": strconv.Itoa(limit),
    }
    dataStream, err := this.call("blocks", "GET", params)
    data := map[string]interface{}{}
    json.Unmarshal(dataStream, &data)
    return data, err
}

func (this *API) LatestBlock() (interface{}, error) {
    dataStream, err := this.call("blocks/latest", "GET", nil)
    data := map[string]interface{}{}
    json.Unmarshal(dataStream, &data)
    return data, err
}

func (this *API) Block(block string) (interface{}, error) {
    dataStream, err := this.call("blocks/"+block, "GET", nil)
    data := map[string]interface{}{}
    json.Unmarshal(dataStream, &data)
    return data, err
}

func (this *API) BlockByHeight(height int) (interface{}, error) {
    return this.Block(strconv.Itoa(height))
}

func (this *API) Transaction(hash string) (interface{}, error) {
    dataStream, err := this.call("transactions/"+hash, "GET", nil)
    data := map[string]interface{}{}
    json.Unmarshal(dataStream, &data)
    return data, err
}

// THIS NEED WORK, NEVER TESTED
func (this *API) SendTransaction(rawHex string) (interface{}, error) {
    var params = map[string]string{
        "raw": rawHex,
    }
    dataStream, err := this.call("transactions/send", "POST", params)
    data := map[string]interface{}{}
    json.Unmarshal(dataStream, &data)
    return data, err
}

func (this *API) TransactionAddresses(hash string) (interface{}, error) {
    dataStream, err := this.call("transactions/"+hash+"/addresses", "GET", nil)
    data := map[string]interface{}{}
    json.Unmarshal(dataStream, &data)
    return data, err
}

func (this *API) TransactionInputs(hash string, inputIndex int) (interface{}, error) {
    inputIndexStr := ""
    if inputIndex!=nil {
        inputIndexStr = "/"+inputIndex
    }
    dataStream, err := this.call("transactions/"+hash+"/inputs"+inputIndexStr, "GET", nil)
    data := map[string]interface{}{}
    json.Unmarshal(dataStream, &data)
    return data, err
}

func (this *API) TransactionOutputs(hash string, outputIndex int) (interface{}, error) {
    outputIndexStr := ""
    if outputIndex!=nil {
        outputIndexStr = "/"+outputIndex
    }
    dataStream, err := this.call("transactions/"+hash+"/inputs"+outputIndexStr, "GET", nil)
    data := map[string]interface{}{}
    json.Unmarshal(dataStream, &data)
    return data, err
}

func (this *API) Address(address string) (interface{}, error) {
    dataStream, err := this.call("addresses/"+address, "GET", nil)
    data := map[string]interface{}{}
    json.Unmarshal(dataStream, &data)
    return data, err
}

func (this *API) Transactions(address string) (map[string]interface{}, error) {
    dataStream, err := this.call("addresses/"+address+"/transactions", "GET", nil)
    data := map[string]interface{}{}
    json.Unmarshal(dataStream, &data)
    return data, err
}

// address or addresses separated by commas
func (this *API) UnspentOutputs(address string) (interface{}, error) {
    dataStream, err := this.call("addresses/"+address+"/utxos", "GET", nil)
    data := map[string]interface{}{}
    json.Unmarshal(dataStream, &data)
    return data, err
}

func (this *API) DoubleSpendsOutputs(address string) (interface{}, error) {
    dataStream, err := this.call("addresses/"+address+"/double_spends", "GET", nil)
    data := map[string]interface{}{}
    json.Unmarshal(dataStream, &data)
    return data, err
}


