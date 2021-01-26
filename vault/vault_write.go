package main

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"os"
)

var token = os.Getenv("VAULT_TOKEN")
var vault_addr = os.Getenv("VAULT_ADDR")

func main() {
	config := &api.Config{
		Address: vault_addr,
	}
	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	client.SetToken(token)
	c := client.Logical()
	
	dataMap := make(map[string]interface{})
	dataMap["storeKey"] = "data"
	data := map[string]interface{}{
		"hello": dataMap,
		"bla":"adfasdf",
	}

	secret, err := c.Write("secret/foo/subfoo", data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(secret)
}
