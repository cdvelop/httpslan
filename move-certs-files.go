package httpslan

import (
	"os"
	"path/filepath"
	"strings"
)

// mover los archivos a carpeta certs
func (h httpsLan) moveCertFiles() (err string) {
	const e = "moveCertFiles "

	err = h.checkCertDir()
	if err != "" {
		return e + err
	}

	// mover los archivos a carpeta certs
	files, er := os.ReadDir("./")
	if er != nil {
		return e + "al leer directorio raíz de certificados " + er.Error()
	}

	for _, archive := range files {

		if strings.Contains(archive.Name(), ".pem") {
			// fmt.Println(archive.Name())
			er := os.Rename(archive.Name(), filepath.Join(h.certs_dir, archive.Name()))
			if er != nil {

				return e + "al mover certificado " + er.Error()
			}
		}
	}

	return ""
}
