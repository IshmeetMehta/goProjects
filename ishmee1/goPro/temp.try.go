package main

import(
    "fmt"
    "log"
    "io/ioutil"
//    "encoding/pem"
//    "crypto/rsa"
//    "crypto/x509"
     b64 "encoding/base64"
//     "math/big"
)

func main(){
    rootPEM, err := ioutil.ReadFile("service-account.pem")
     if err != nil {
        log.Fatal(err)
    }


    cert := b64.URLEncoding.EncodeToString([]byte(rootPEM))
//    fmt.Println(uEnc)
//    fmt.Println(cert)
//    fmt.Println(rsaPublicKey.N)
//    fmt.Println(rsaPublicKey.E)
//	uEnc := b64.URLEncoding.EncodeToString([]byte(str))
	fmt.Println(cert)
    }
