# BestMixer.io API

[![Build Status](https://travis-ci.org/MarinX/bestmixer.svg)](https://travis-ci.org/MarinX/bestmixer)
[![Go Report Card](https://goreportcard.com/badge/github.com/MarinX/bestmixer)](https://goreportcard.com/report/github.com/MarinX/bestmixer)
[![GoDoc](https://godoc.org/github.com/MarinX/bestmixer?status.svg)](https://godoc.org/github.com/MarinX/bestmixer)
[![License MIT](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat)](LICENSE)

bestmixer is a Go client library for the [BestMixer.io API](https://bestmixer.io/en/api).


## Start using it

1. Download and install it:

    ```sh
    $ go get github.com/MarinX/bestmixer
    ```

2. Import it in your code:

    ```go
    import "github.com/MarinX/bestmixer"
    ```

3. (Optional) Run test 

    ```sh 
    $ go test -v
    ```

## API

### creating a client
```go
	// Create a new BestMixer.io client with http.DefaultClient
	// replace api_key with your own key
	client := bestmixer.New("api_key", nil)

	// if you have your own http.Client you can supply it
	// as a second param
	client := bestmixer.New("api_key", myHttpClient)
```

### Code
#### Info
```go
	codeInfo, err := client.Code.Info("bm_code")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("BTC:", codeInfo.BTC)
	fmt.Println("BCH:", codeInfo.BCH)
	fmt.Println("LTC:", codeInfo.LTC)
```

### Fee
#### Info

```go
	feeInfo, err := client.Fee.Info()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("fee btc:", feeInfo.BTC)
	fmt.Println("fee bch:", feeInfo.BCH)
	fmt.Println("fee ltc:", feeInfo.LTC)
```

### Order
#### creating order

```go
	order, err := client.Order.Create(
		"bm_code",
		bestmixer.CoinBTC,
		1.001,
		[]bestmixer.Output{
			bestmixer.Output{
				Address: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				Percent: 5.5,
				Delay:   33,
			},
		})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Order created:", order.OrderID)
```

#### creating fixed order

```go
	orderFixed, err := client.Order.CreateFixed(
		"bm_code",
		bestmixer.CoinBTC,
		1.001,
		[]bestmixer.Output{
			bestmixer.Output{
				Address: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				Percent: 0.1,
				Delay:   33,
			},
		})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Order fixed created:", orderFixed.OrderID)
```

#### get order info

```go
	orderStatus, err := client.Order.Info("my_order_id")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Order status:", orderStatus.Status)
```

## License

This library is under the MIT License