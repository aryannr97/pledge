package main

import (
	"log"

	"github.com/aryannr97/pledge"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	plg, err := pledge.New(pledge.AuthMethodJWT, true, &pledge.Config{
		Issuer:         "<Issuer Org>",
		PublicKeyPath:  "<path_to_key>",
		PrivateKeyPath: "<path_to_key>",
	})
	if err != nil {
		log.Println(err)
		return
	}

	t, _ := plg.GenerateIdentitiy(map[string]interface{}{
		"name":   "peter",
		"email":  "peter@example",
	})
	log.Printf("Generated token: %v", t)
}
