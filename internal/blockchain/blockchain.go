
package blockchain

import (
    "os/exec"
	"os"
	"encoding/hex"
	"github.com/limserhane/sgcsfuse/internal/crypto"
	"github.com/limserhane/sgcsfuse/internal/logger"
)

var path = os.Getenv("SGCSFUSEPATH") + "/blockchain"

func Add(filename string, data []byte) {
	hashBytes := crypto.Hash(data)
	hashString := hex.EncodeToString(hashBytes)
    err := exec.Command("/bin/bash", path+"/add.sh", filename, hashString).Run()
	logger.Debugf("(BlockchainAdd) %v", err)
}

func Close() {
    err := exec.Command("/bin/bash", path+"/close.sh").Run()
	logger.Debugf("(BlockchainClose) %v", err)
}