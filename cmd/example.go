package main

import (
	"log"

	"github.com/aryannr97/pledge"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	plg, err := pledge.New(pledge.AuthMethodJWT, true, &pledge.Config{
		Issuer:         "Scratch Cred Securities",
		PublicKeyPath:  "/home/rnrupnar/scratch/creds/dev/public.pem",
		PrivateKeyPath: "/home/rnrupnar/scratch/creds/dev/private.pem",
	})
	if err != nil {
		log.Println(err)
		return
	}

	t, _ := plg.GenerateIdentitiy(map[string]interface{}{
		"userid": "guid",
		"name":   "peter",
		"email":  "p@example",
	})
	log.Printf("Generated token: %v", t)
}
