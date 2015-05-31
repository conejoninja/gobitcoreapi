package gobitcoreapi

import (
    "crypto/tls"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "strconv"
    "bytes"
    "fmt"
    "reflect"
    "time"
)

type API struct {
    endPoint      string
    version       string
    client        *http.Client
}

type NodeStatus struct {
    Sync        float64
    PeerCount   int
    Version     string
    Network     string
    Height      int
}

type Address struct {
    Address         string
    Transactions    string
    Unconfirmed     string
    Confirmed       string
    LastActivity    string
}

type BlockHeader struct {
    Version         int
    PrevHash        string
    MerkleRoot      string
    Time            int
    Bits            int
    Nonce           int
}

type BlockTransaction struct {
    Version         int
    Inputs          []TransactionInput
    Outputs         []TransactionOutput
    NLockTime       int
}

type TransactionInput struct {
    PrevTxId        string
    OutputIndex     int
    SequenceNumber  int
    Script          string
    ScriptString    string
}

type TransactionOutput struct {
    Satoshis        string
    Script          string
}

type Block struct {
    Header          BlockHeader
    Transactions    []BlockTransaction
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
        fmt.Println("GET => ", this.endPoint+"/"+this.version+"/"+action+"?"+valuesStr)
        res, err = this.client.Get(this.endPoint+"/"+this.version+"/"+action+"?"+valuesStr)
    }
    body, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
    return body, err
}

func (this *API) SetVersion(version string) {
    this.version = version
}

func (this *API) Node() (NodeStatus, error) {
    dataStream, err := this.call("node", "GET", nil)
    var data NodeStatus
    json.Unmarshal(dataStream, &data)

    return data, err
}

func (this *API) Blocks( from, limit, offset, to int) ([]Block, error) {
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
    var data []Block
    fmt.Println("DATASTREAM", dataStream)
    json.Unmarshal(dataStream, &data)

    fmt.Println("BLOCKS", data)

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
    if inputIndex>-1 {
        inputIndexStr = "/"+strconv.Itoa(inputIndex)
    }
    dataStream, err := this.call("transactions/"+hash+"/inputs"+inputIndexStr, "GET", nil)
    data := map[string]interface{}{}
    json.Unmarshal(dataStream, &data)
    return data, err
}

func (this *API) TransactionOutputs(hash string, outputIndex int) (interface{}, error) {
    outputIndexStr := ""
    if outputIndex>-1 {
        outputIndexStr = "/"+strconv.Itoa(outputIndex)
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


func toString(x interface{}) string {
    switch y := x.(type) {

        // Handle dates with special logic
        // This needs to come above the fmt.Stringer
        // test since time.Time's have a .String()
        // method
        case time.Time:
        return y.Format("A Monday")

        // Handle type string
        case string:
        return y

        // Handle type with .String() method
        case fmt.Stringer:
        return y.String()

        // Handle type with .Error() method
        case error:
        return y.Error()

    }

    // Handle named string type
    if v := reflect.ValueOf(x); v.Kind() == reflect.String {
        return v.String()
    }

    // Fallback to fmt package for anything else like numeric types
    return fmt.Sprint(x)
}
