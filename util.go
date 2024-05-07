package nesbconvertpin

import (
	"os"
)

func getRootDir() string {
	return os.Getenv("FAPWORKDIR")
}
