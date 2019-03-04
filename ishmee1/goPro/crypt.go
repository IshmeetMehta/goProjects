package main

import (
    "crypto/x509"
    "encoding/pem"
    "io/ioutil"
    "fmt"
)


func main() {

    // Read the	 pem file from OS
    pemString, err := ioutil.ReadFile("service-account.pem")
	if err != nil {

		fmt.Println("file reading error", err)
		return
	}
 pubkey_bytes, err := x509.MarshalPKIXPublicKey(pemString)
    if err != nil {
            return "", err
    }
    pubkey_pem := pem.EncodeToMemory(
            &pem.Block{
                    Type:  "RSA PUBLIC KEY",
                    Bytes: pubkey_bytes,
            },
    )

    return string(pubkey_pem), nil

}	
