package inputs

import (
    "os"
    "path/filepath"
)

func ReadInput(day string) ([]byte, error) {
    return os.ReadFile(filepath.Join("inputs", day + ".txt"))
}
