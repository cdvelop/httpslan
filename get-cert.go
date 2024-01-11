package httpslan

import (
	"crypto/tls"
	"fmt"
)

func (h *httpsLan) GetCertificate(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {

	fmt.Println("hello tls:", hello)

	if !h.built_cert {

		h.BuildSSlCertificate()

		h.built_cert = true
	}

	cert, er := tls.LoadX509KeyPair(h.publicCertFilePath, h.privateKeyFilePath)
	if er != nil {
		return nil, er
	}
	return &cert, nil
}
