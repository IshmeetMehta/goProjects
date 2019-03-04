package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
)

func main() {
	var pubPEMData = []byte(`
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvRsYyGcGDi53cKftxHes
maN7ASeBe4x14x0tiAEnxTlkLHX8WgEVO1c+Pis2SWrt01gmDVOkU7BjY5TRt2E8
1BmC4C8CA4Eci/Cr0R59+GYftZwvkplwodqOoENwcl5xukjr9mwZsbgSReMu1ES6
fU/iFmO2biQotltKD09lE0xjREK0xtxLOf6GNSRGWO/DHECG0RfCACxX19sZMUV6
njlsMQR/IEp6Yrmxm/BMilqTT5h3PuozI+rhnIsE6A9F36A9B9tbalsKpBLi62Qn
tqKGITuqFOypghJTxi9LnlQIQbxsSaEjC5o3Jo+kpzNQmCgZJs7u7icfhHPoo6A8
twIDAQAB
-----END PUBLIC KEY-----`)

	block, rest := pem.Decode(pubPEMData)
	if block == nil || block.Type != "PUBLIC KEY" {
		log.Fatal("failed to decode PEM block containing public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Got a %T, with remaining data: %q", pub, rest)
}
