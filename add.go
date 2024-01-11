package httpslan

import (
	"os"
	"path/filepath"
)

// os.Getenv("APP_DOMAIN") ej: 192.168.0.1,miapp.cl,localhost
func AddHttpsLan() (h *httpsLan) {
	const e = "AddHttpsLan error"

	APP_DOMAIN := os.Getenv("APP_DOMAIN")
	if APP_DOMAIN == "" {
		APP_DOMAIN = "localhost"
	}

	const certs_dir = "certs"

	h = &httpsLan{
		app_domain:         APP_DOMAIN,
		certs_dir:          certs_dir,
		publicCertFilePath: filepath.Join(certs_dir, APP_DOMAIN+".pem"),
		privateKeyFilePath: filepath.Join(certs_dir, APP_DOMAIN+"-key.pem"),
	}

	return h
}
