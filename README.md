## go-hdwallet

A multi-cryptocurrency HD wallet implementated by golang.

## supported coins

- BTC
- LTC
- DOGE
- DASH
- ETH
- ETC
- BCH
- QTUM
- USDT
- IOST
- USDC

## install

```sh
go get -v -u github.com/foxnut/go-hdwallet
```

## example

```go
package main

import (
    "fmt"

    "github.com/foxnut/go-hdwallet"
)

var (
    mnemonic = "range sheriff try enroll deer over ten level bring display stamp recycle"
)

func main() {
    master, err :