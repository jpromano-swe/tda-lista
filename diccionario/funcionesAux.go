package diccionario

func (hash *tablaDeHash[K, V]) buscarIndicePorClave(clave K) (int, bool) {
  indiceInicial := hashingDeClaves(clave, len(hash.tablaHash))
  for i := 0; i < len(hash.tablaHash); i++ {
    indiceActual := (indiceInicial + i) % len(hash.tablaHash)
    celdaActual := hash.tablaHash[indiceActual]

    if celdaActual.EstadoDeCelda == _LIBRE {
      return -1, false
    }

    if celdaActual.EstadoDeCelda == _OCUPADO && celdaActual.clave == clave {
      return indiceActual, true
    }
  }

  return -1, false
}
func (hash *tablaDeHash[K, V]) buscarIndiceParaInsertar(clave K) int {
  indiceInicial := hashingDeClaves(clave, len(hash.tablaHash))
  hayPrimerBorrado := false
  indicePrimerBorrado := 0

  for i := 0; i < len(hash.tablaHash); i++ {
    indiceActual := (indiceInicial + i) % len(hash.tablaHash)
    celdaActual := hash.tablaHash[indiceActual]

    switch celdaActual.EstadoDeCelda {
    case _BORRADO:
      if !hayPrimerBorrado {
        hayPrimerBorrado = true
        indicePrimerBorrado = indiceActual
      }
    case _LIBRE:
      if hayPrimerBorrado {
        return indicePrimerBorrado
      }
      return indiceActual

    case _OCUPADO:
      if celdaActual.clave == clave {
        return indiceActual
      }
    }
  }
  if hayPrimerBorrado {
    return indicePrimerBorrado
  }
  return -1
}

func (hash *tablaDeHash[K, V]) encontrarPrimerOcupado() int {
  i := 0
  contador := 0
  for i < hash.capacidad && hash.tablaHash[i].EstadoDeCelda != _OCUPADO {
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
