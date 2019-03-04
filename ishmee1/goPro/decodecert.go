package main

import(
    "fmt"
    "log"
    "io/ioutil"
    "encoding/pem"
    "crypto/rsa"
    "crypto/x509"
     b64 "encoding/base64"
//     "math/big"
)

func main(){
    rootPEM, err := ioutil.ReadFile("service-account.pem")
     if err != nil {
        log.Fatal(err)
    }


    block, _ := pem.Decode([]byte(rootPEM))
    var cert* x509.Certificate
    cert, _ = x509.ParseCertificate(block.Bytes)
    rsaPublicKey := cert.PublicKey.(*rsa.PublicKey)
    uEnc := b64.URLEncoding.EncodeToString([]byte((rsaPublicKey.N).Bytes()))
    uEnce := b64.URLEncoding.EncodeToString([]byte([]byte(rsaPublicKey.E)))
//    fmt.Println(uEnc)
//    fmt.Println(cert)
//    fmt.Println(rsaPublicKey.N)
//    fmt.Println(rsaPublicKey.E)
//	uEnc := b64.URLEncoding.EncodeToString([]byte(str))
	fmt.Println(uEnc)
	fmt.Println("Hello")
	fmt.Println(uEnce)
    }
