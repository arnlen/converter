# Go Converter

Training project to teach myself how to use CLI + API call within a go app.

## Purpose

A simple USD to EUR converter written in Go.

## Setup

It uses [exchangeratesapi.io](https://exchangeratesapi.io/) behind the scene, so be sure:

1. Create a [free account](https://manage.exchangeratesapi.io/signup/free)
2. Store your API Access Key inside the environment variable "`EXCHANGERATESAPI_ACCESS_KEY`".

```zsh
export EXCHANGERATESAPI_ACCESS_KEY=<your_exchange_rates_api_access_key>
```

3. Restart your shell

## Example

To convert **$5 USD** to EUR:

```zsh
$ go run main.go -usd=5

# 📡 Calling Exchangerateapi to get latest rates...
# ✅ Latest rates received: 1 EUR = 1.209715 USD
# 5 USD =  6.0485750000000005 EUR
```