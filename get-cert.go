package httpslan

import (
	"crypto/tls"
)

func (h *httpsLan) GetCertificate(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {

	// fmt.Println("hello tls:", hello)
	cert, er := tls.LoadX509KeyPair(h.publicCertFilePath, h.privateKeyFilePath)
	if er != nil {
		return nil, er
	}
	return &cert, nil
}
