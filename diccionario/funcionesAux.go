package diccionario

func (hash *tablaDeHash[K, V]) encontrarPrimerOcupado() int {
  i := 0
  contador := 0
  for i < hash.capacidad && hash.tablaHash[i].estadoDeCelda != _OCUPADO {
    i++
    contador++
  }
  return contador
}

func esPrimo(numero int) bool {
  if numero < 2 {
    return false
  }
  return _esPrimo(numero, 2)
}

func _esPrimo(numero int, divisor int) bool {
  if divisor*divisor > numero {
    return true
  }
  if numero%divisor == 0 {
    return false
  }
  return _esPrimo(numero, divisor+1)
}

func siguientePrimo(numero int) int {
  if numero <= 2 {
    return 2
  }
  if numero%2 == 0 {
    numero++
  }
  for !esPrimo(numero) {
    numero += 2
  }
  return numero
}
