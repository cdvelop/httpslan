package httpslan

import "os"

func (h httpsLan) checkCertDir() (err string) {

	// Utiliza os.Stat para obtener información sobre el archivo o directorio
	_, er := os.Stat(h.certs_dir)

	// Verifica si hay un error y si el error indica que el archivo o directorio no existe
	if os.IsNotExist(er) {
		// Intenta crear el directorio y maneja posibles errores
		er := os.MkdirAll(h.certs_dir, os.ModePerm)
		if er != nil {
			return "checkCertDir al crear el directorio: " + er.Error()
		}
	}

	return ""
}
