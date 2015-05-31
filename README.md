GoBitcoreApi
============
Simple wrapper in Go for Bitcore's API


## Installation


```bash
$ go get github.com/conejoninja/gobitcoreapi
```

## Documentation
See [Go Doc](http://godoc.org/github.com/conejoninja/gobitcoreapi) or [Go Walker](http://gowalker.org/github.com/conejoninja/gobitcoreapi) for usage and details.

## Example of use

```go
package main

import (
    "github.com/conejoninja/gobitcoreapi"
    "fmt"
)

func main() {

    var api *gobitcoreapi.API

    api = gobitcoreapi.NewAPI("YOUR_END_POINT")

    res, err := api.Address("1NcXPMRaanz43b1kokpPuYDdk6GGDvxT2T")
    fmt.Println("Address", res, err)

    res, err = api.Transactions("1NcXPMRaanz43b1kokpPuYDdk6GGDvxT2T", 1, 200, "asc")
    fmt.Println("Transactions", res, err)

    res, err = api.UnspentOutputs("1NcXPMRaanz43b1kokpPuYDdk6GGDvxT2T", 1, 200, "asc")
    fmt.Println("UnspentOutputs", res, err)

    res, err = api.Block("290000")
    fmt.Println("Block", res, err)

    res, err = api.BlockByHeight(290000)
    fmt.Println("BlockByHeight", res, err)

    res, err = api.Block("0000000000000000fa0b2badd05db0178623ebf8dd081fe7eb874c26e27d0b3b")
    fmt.Println("Block", res, err)

    res, err = api.Blocks(1, 20, "asc")
    fmt.Println("Blocks", res, err)

    res, err = api.LatestBlock()
    fmt.Println("LatestBlock", res, err)

    res, err = api.Transaction("c326105f7fbfa4e8fe971569ef8858f47ee7e4aa5e8e7c458be8002be3d86aad")
    fmt.Println("Transaction", res, err)


}
```

## Noted
I wouldn't use it for anything serious or important.

## Contributing to GoBitcoreAPI:

If you find any improvement or issue you want to fix, feel free to send me a pull request with testing.

Feel free to donate some bits : 1FYWFgvXKo5q1LPCYxo6NDCDzFd8Um72hH


## License

This is distributed under the Apache License v2.0

Copyright 2015 Daniel Esteban  -  conejo@conejo.me

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

