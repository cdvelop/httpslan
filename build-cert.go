package httpslan

import (
	"embed"
	"io"
	"log"
	"os"
	"os/exec"
)

//go:embed mkcert.exe
var embedFS embed.FS

func (h *httpsLan) BuildSSlCertificate() {
	const e = "BuildSSlCertificate error "
	// Crea un archivo temporal
	tempFile, er := os.CreateTemp("", "temp_exe_*.exe")
	if er != nil {
		log.Println(e+"creando archivo temporal:", er)
		return
	}
	defer os.Remove(tempFile.Name())

	// Abre el archivo exe embebido
	exeFile, er := embedFS.Open("mkcert.exe")
	if er != nil {
		log.Println(e+"abriendo el archivo embebido:", er)
		return
	}
	defer exeFile.Close()

	// Copia el contenido del exe embebido al archivo temporal
	_, er = io.Copy(tempFile, exeFile)
	if er != nil {
		log.Println(e+"escribiendo en el archivo temporal:", er)
		return
	}

	// Cierra el archivo temporal antes de ejecutarlo
	tempFile.Close()

	// Ejecuta el archivo exe desde la memoria
	cmd := exec.Command(tempFile.Name(), h.app_domain)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	er = cmd.Run()
	if er != nil {
		log.Println(e+"ejecutando el archivo exe:", er)
		return
	}

	err := h.moveCertFiles()
	if err != "" {
		log.Println(e + err)
	}
}
