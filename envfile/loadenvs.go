// Package envfile loads environment variables from env files.
package envfile

import (
	"os"
	"path"
	"path/filepath"

	"github.com/joho/godotenv"
)

const (
	envFileName        = ".env"
	envPrivateFileName = ".envprivate"
)

// Load looks for the default .env and .envprivate files in the current directory
// and the path of the binary. It sets all environment variables from it for the
// current process. It will overwrite existing environment variables.
func Load() {
	LoadFiles(envFileName, envPrivateFileName)
}

// LoadFiles loads the given files instead of the default ones.
func LoadFiles(files ...string) {
	currentDirectory, err := os.Getwd()
	if err == nil {
		loadEnvsFromPath(currentDirectory, files...)
	}

	executable, err := os.Executable()
	if err == nil {
		executableDirectory := filepath.Dir(executable)
		loadEnvsFromPath(executableDirectory, files...)
	}
}

func loadEnvsFromPath(directoryPath string, files ...string) {
	paths := make([]string, 0, len(files))

	for _, file := range files {
		filePath := path.Join(directoryPath, file)
		paths = append(paths, filePath)
	}

	_ = godotenv.Overload(paths...)
}
