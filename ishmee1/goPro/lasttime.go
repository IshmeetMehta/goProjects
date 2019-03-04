package main

import(
    "fmt"
    "log"
    "encoding/pem"
    "crypto/x509"
    "crypto/rsa"
    "io/ioutil"
)

func main(){
    rootPEM, err := ioutil.ReadFile("service-account.pem")
     if err != nil {
        log.Fatal(err)
    }

    block, _ := pem.Decode([]byte(rootPEM))
    var cert* x509.Certificate
    cert, _ = x509.ParseCertificate(block.Bytes)
    fmt.Println(cert.subject)
    rsaPublicKey := cert.PublicKey.(*rsa.PublicKey)
    //fmt.Println(rsaPublicKey.N)
   // fmt.Println(rsaPublicKey.E)
   // fmt.Println(cert.PublicKey.Subject)

 publicKeyDer, err := x509.MarshalPKIXPublicKey(rsaPublicKey)
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
