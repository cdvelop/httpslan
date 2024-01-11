package httpslan

import (
	"fmt"
	"os"
)

func exit(message ...any) {
	fmt.Printf("Error: %v\n", message)
	fmt.Println("")
	fmt.Println("Presiona enter para salir...")
	fmt.Scanln()
	os.Exit(1)
}
