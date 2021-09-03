
package blockchain

import (
    "os/exec"
	"encoding/hex"
	"github.com/limserhane/sgcsfuse/internal/crypto"
	// "github.com/limserhane/sgcsfuse/internal/logger"
)

func Add(filename string, data []byte) {
	hashBytes := crypto.Hash(data)
	hashString := hex.EncodeToString(hashBytes)
    exec.Command("/bin/bash", "add.sh", filename, hashString).Run()
}

func Close() {
    exec.Command("/bin/bash", "close.sh").Run()
}