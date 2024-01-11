package httpslan

import (
	"log"
	"os"
	"path/filepath"
)

// os.Getenv("APP_DOMAIN") ej: 192.168.0.1,miapp.cl,localhost
// schedule (cronograma) interface para ejecutar la renovación de certificados cada 6 meses
func AddHttpsLan(s schedule) (h *httpsLan) {
	const e = "AddHttpsLan "

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
		schedule:           s,
	}

	h.BuildSSlCertificate()

	if h.schedule != nil {
		err := h.AddFuncToSchedule("0 0 1 */6 *", "generación de certificado ssl local cada 6 meses", h.BuildSSlCertificate)
		if err != "" {
			log.Println(e, err)
		}
	}

	return h
}
