package diccionario


func (hash *tablaDeHash[K, V]) buscarIndicePorClave(clave K) (int, bool) {
	indiceInicial := hashingDeClaves(clave, len(hash.tablaHash))
	for i := 0; i < len(hash.tablaHash); i++ {
		indiceActual := (indiceInicial + i) % len(hash.tablaHash)
		celdaActual := hash.tablaHash[indiceActual]

		if celdaActual.EstadoDeCelda == _libre || celdaActual.EstadoDeCelda == _borrado {
			return -1, false
		}

		if celdaActual.EstadoDeCelda == _ocupado && celdaActual.clave == clave {
			return indiceActual, true
		}
	}

	return -1, false
}
func (hash *tablaDeHash[K, V]) buscarIndiceParaInsertar(clave K) int {
	indiceInicial := hashingDeClaves(clave, len(hash.tablaHash))
	for i := 0; i < len(hash.tablaHash); i++ {
		indiceActual := (indiceInicial + i) % len(hash.tablaHash)
		celdaActual := hash.tablaHash[indiceActual]

		if celdaActual.EstadoDeCelda == _libre || celdaActual.EstadoDeCelda == _borrado{
			return indiceActual
		}

		if celdaActual.EstadoDeCelda == _ocupado && celdaActual.clave == clave {
			return indiceActual
		}
	}
	return -1
}

func (iter *iteradorHash[K, V]) convertirIndiceEnClave() {

}

func esPrimo(n int) bool {
	if n < 2 {
		return false
	}
	return _esPrimo(n, 2)
}

func _esPrimo(n int, divisor int) bool {
	if divisor*divisor > n {
		return true
	}
	if n%divisor == 0 {
		return false
	}
	return _esPrimo(n, divisor+1)
}

func siguientePrimo(n int) int {
	if n <= 2 {
		return 2
	}
	if n%2 == 0 {
		n++
	}
	for !esPrimo(n) {
		n += 2
	}
	return n
}
