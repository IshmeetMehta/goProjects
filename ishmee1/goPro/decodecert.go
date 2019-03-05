package main

import(
    "fmt"
    "log"
    "io/ioutil"
    "encoding/pem"
    "crypto/rsa"
    "crypto/x509"
     b64 "encoding/base64"
//     "strconv"
//     "encoding/binary"
     "math/big"
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
    exponent := big.NewInt(int64(rsaPublicKey.E)) 
//    n := b64.URLEncoding.EncodeToString([]byte((rsaPublicKey.N).Bytes()))
    n := b64.RawURLEncoding.EncodeToString([]byte((rsaPublicKey.N).Bytes()))
    e := b64.RawURLEncoding.EncodeToString([]byte((exponent).Bytes()))
    fmt.Println("n:  ", n)
    fmt.Println("e:  ", e)
    }
