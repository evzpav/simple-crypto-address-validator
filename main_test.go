package main

import (
	"testing"
)

type walletAddressInput struct {
	Crypto  string
	Address string
}

var walletAddresses = []walletAddressInput{
	{Crypto: "btc", Address: "1CFNjwLjZdSKB8nZopxhLaR8vvqaQKD3Bi"},                             //old btc type
	{Crypto: "BTC", Address: "bc1qar0srrr7xfkvy5l643lydnw9re59gtzzwf5mdq"},                     // bech32
	// {Crypto: "bch", Address: "qq7ujnfl6tqx7xcdsdsrsqlqgqz8rm5stsvgx2kcvu"},                     // cash address
	// {Crypto: "bch", Address: "bitcoincash:qq7ujnfl6tqx7xcdsdsrsqlqgqz8rm5stsvgx2kcvu"},         // cash address
	// {Crypto: "bch", Address: "16dhNPnPp346wzrRTkArKhqPM1ELeJDvRr"},
	{Crypto: "BTG", Address: "GakMJVF7Du16VK9dpN6nhJyLUPLXkTfqSY"},
	{Crypto: "DGB", Address: "D59P8MiMXkjs7HPn31zAnUSvRNwvNZUBYa"},
	{Crypto: "DASH", Address: "XiHMBEic8q8wX5aKqVv6zRFec7cAuYGjBV"},
	{Crypto: "ETH", Address: "0x15cc4bf4fe84fea178d2b10f89f1a6c914dfc8c2"},
	{Crypto: "SMART", Address: "SbsLb8eM583oraW89qhbkcqZmuR4aYKkea"},
	{Crypto: "XRP", Address: "rMkfgicNKuCfXojDhcX4W2LnGoHFqhFrr6"},
	{Crypto: "ZEC", Address: "t1SBt3V8MfG4ZJ2ZDTuWfDshn4PuyvqjJV3"},
	{Crypto: "ZCR", Address: "ZXvpr2M6wvKoFcTJ57WCjT9Wkd38xkL8Fo"},
}

func TestValidateAddress(t *testing.T) {
	for _, w := range walletAddresses {
		isValid, err := validateAddress(w.Crypto, w.Address)
		if err != nil {
			t.Errorf("%v", err)
			return
		}

		if !isValid {
			t.Errorf("Crypto: %s, Address: %s", w.Crypto, w.Address)
		}
	}

}
