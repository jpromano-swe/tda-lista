package diccionario

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func hasheoDeBytes(arreglo []byte) uint64 {
	hash := sha256.Sum256(arreglo)
	return binary.BigEndian.Uint64(hash[0:8])
}

func asignarIndice(hash uint64, capacidad int) int {
	return int(hash % uint64(capacidad))
}

func hashingDeClaves[K comparable](clave K, capacidad int) int {
	bytes := convertirABytes(clave)
	hash := hasheoDeBytes(bytes)
	return asignarIndice(hash, capacidad)
}
