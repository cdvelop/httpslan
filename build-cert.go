package httpslan

import (
	"embed"
	"fmt"
	"io"
	"os"
	"os/exec"
)

//go:embed mkcert.exe
var embedFS embed.FS

func (h httpsLan) BuildSSlCertificate() {
	// Crea un archivo temporal
	tempFile, err := os.CreateTemp("", "temp_exe_*.exe")
	if err != nil {
		fmt.Println("Error creando archivo temporal:", err)
		return
	}
	defer os.Remove(tempFile.Name())

	// Abre el archivo exe embebido
	exeFile, err := embedFS.Open("mkcert.exe")
	if err != nil {
		fmt.Println("Error abriendo el archivo embebido:", err)
		return
	}
	defer exeFile.Close()

	// Copia el contenido del exe embebido al archivo temporal
	_, err = io.Copy(tempFile, exeFile)
	if err != nil {
		fmt.Println("Error escribiendo en el archivo temporal:", err)
		return
	}

	// Cierra el archivo temporal antes de ejecutarlo
	tempFile.Close()

	// Ejecuta el archivo exe desde la memoria
	cmd := exec.Command(tempFile.Name(), h.app_domain)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println("Error ejecutando el archivo exe:", err)
		return
	}

	h.moveCertFiles()
}
