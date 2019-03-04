package main

import (
	
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
    privKey, err := ioutil.ReadFile("service-account-key.pem")

    block, _ := pem.Decode([]byte(privKey))
    if block == nil || block.Type != "RSA PRIVATE KEY" {
        log.Fatal("failed to decode PEM block containing public key")
    }
    key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
        log.Fatal(err)
    }

    // publicKeyDer := x509.MarshalPKCS1PublicKey(&key.PublicKey)
    publicKeyDer, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
    if err != nil {
        log.Fatal(err)
    }
    pubKeyBlock := pem.Block{
        Type:    "PUBLIC KEY",
        Headers: nil,
        Bytes:   publicKeyDer,
    }
    pubKeyPem := string(pem.EncodeToMemory(&pubKeyBlock))
    fmt.Println(pubKeyPem)
}
