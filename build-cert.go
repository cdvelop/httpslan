package httpslan

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// mkcert.exe_folder ej: ./files_myapp/certs/
// cert_destine_folder ej: ./files_myapp/certs/
// app_domain ej: 192.168.0.1,miapp.cl,localhost
func (h httpsLan) BuildSSlCertificate() (err string) {

	_, er := exec.Command(filepath.Join(h.certs_dir, "mkcert.exe"), h.app_domain).Output()
	if er != nil {
		return er.Error() + " error no se logro generar certificado ssl para:" + h.app_domain
	}

	// mover los archivos a carpeta certs
	files, er := os.ReadDir("./")
	if er != nil {
		return "error al leer directorio raíz de certificados " + er.Error()
	}

	for _, archive := range files {

		if strings.Contains(archive.Name(), ".pem") {
			// fmt.Println(archive.Name())
			er := os.Rename(archive.Name(), filepath.Join(h.certs_dir, archive.Name()))
			if er != nil {
				return "error al mover certificado " + er.Error()
			}
		}
	}

	// fmt.Printf("SALIDA: %v", out)
	return ""
}
