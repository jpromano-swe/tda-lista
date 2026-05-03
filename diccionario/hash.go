package diccionario

//key: id
//Valor: valor
//indice: ubicacion en memoria

const (
	_LIBRE estadoDeCelda = iota
	_OCUPADO
	_BORRADO

	_CAPACIDAD_INICIAL     = 7
	_FACTOR_DE_CRECIMIENTO = 2
	_FACTOR_DE_CARGA       = 0.7
)

type estadoDeCelda = int

type tablaDeHash[K comparable, V any] struct {
	tablaHash []celdaDiccionario[K, V]
	cantidad  int
	capacidad int
}

type celdaDiccionario[K comparable, V any] struct {
	clave         K
	valor         V
	estadoDeCelda int
}
type iteradorHash[K comparable, V any] struct {
	hashPasado   *tablaDeHash[K, V]
	indiceActual int
	posicion     int
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	diccionario := new(tablaDeHash[K, V])
	diccionario.tablaHash = make([]celdaDiccionario[K, V], _CAPACIDAD_INICIAL)
	diccionario.cantidad = 0
	diccionario.capacidad = _CAPACIDAD_INICIAL
	return diccionario
}

func (hash *tablaDeHash[K, V]) Pertenece(clave K) bool {
	_, clavePertenece := hash.buscarIndicePorClave(clave)
	return clavePertenece
}

func (hash *tablaDeHash[K, V]) Obtener(clave K) V {
	indiceActual, clavePertenece := hash.buscarIndicePorClave(clave)
	if !clavePertenece {
		panic("La clave no pertenece al diccionario")
	}
	return hash.tablaHash[indiceActual].valor
}

func (hash *tablaDeHash[K, V]) Borrar(clave K) V {
	indiceActual, clavePetenece := hash.buscarIndicePorClave(clave)
	if !clavePetenece {
		panic("La clave no pertenece al diccionario")
	}
	hash.cantidad--
	hash.tablaHash[indiceActual].estadoDeCelda = _BORRADO
	if hash.cantidad*4 <= hash.capacidad && hash.capacidad > _CAPACIDAD_INICIAL {
		hash.redimensionarDiccionario(hash.capacidad / 2)
	}
	return hash.tablaHash[indiceActual].valor
}

func (hash *tablaDeHash[K, V]) Guardar(clave K, valor V) { //REFACTOR, QUEDO MUY LARGO

	indiceActual, clavePertenece := hash.buscarIndice(clave)
	capacidadActual := len(hash.tablaHash)
	factorCargado := float64(hash.cantidad+1) / float64(capacidadActual)

	if clavePertenece {
		hash.tablaHash[indiceActual].valor = valor
		return
	}

	if factorCargado > _FACTOR_DE_CARGA {
		nuevaCapacidad := siguientePrimo(_FACTOR_DE_CRECIMIENTO * capacidadActual)
		hash.redimensionarDiccionario(nuevaCapacidad)
		indiceActual, _ = hash.buscarIndice(clave)
	}

	hash.tablaHash[indiceActual].clave = clave
	hash.tablaHash[indiceActual].valor = valor
	hash.tablaHash[indiceActual].estadoDeCelda = _OCUPADO
	hash.cantidad++
}

func (hash *tablaDeHash[K, V]) Cantidad() int {
	return hash.cantidad
}

// PRIMITIVAS PARA EL ITERADOR

func (hash *tablaDeHash[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	indice := hash.encontrarPrimerOcupado()
	cantidad := 0
	for indice < hash.capacidad && cantidad < hash.cantidad {
		if hash.tablaHash[indice].estadoDeCelda == _OCUPADO {
			cantidad++
			if !visitar(hash.tablaHash[indice].clave, hash.tablaHash[indice].valor) {
				return
			}
		}
		indice++
	}
}

func (hash *tablaDeHash[K, V]) Iterador() IterDiccionario[K, V] {
	return &iteradorHash[K, V]{hashPasado: hash, indiceActual: hash.encontrarPrimerOcupado(), posicion: 0}
}

func (iter *iteradorHash[K, V]) HayAlgoMas() bool {
	return iter.posicion < iter.hashPasado.cantidad
}

func (iter *iteradorHash[K, V]) VerActual() (K, V) {
	if !iter.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	return iter.hashPasado.tablaHash[iter.indiceActual].clave, iter.hashPasado.tablaHash[iter.indiceActual].valor
}

func (iter *iteradorHash[K, V]) Avanzar() {
	if !iter.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	iter.indiceActual++
	for iter.indiceActual < iter.hashPasado.capacidad && iter.hashPasado.tablaHash[iter.indiceActual].estadoDeCelda != _OCUPADO {
		iter.indiceActual++
	}
	iter.posicion++
}

func (hash *tablaDeHash[K, V]) buscarIndice(clave K) (int, bool) {
	indiceInicial := hashingDeClaves(clave, len(hash.tablaHash))
	hayPrimerBorrado := false
	indicePrimerBorrado := 0

	for i := 0; i < len(hash.tablaHash); i++ {
		indiceActual := (indiceInicial + i) % len(hash.tablaHash)
		celdaActual := hash.tablaHash[indiceActual]

		switch celdaActual.estadoDeCelda {
		case _BORRADO:
			if !hayPrimerBorrado {
				hayPrimerBorrado = true
				indicePrimerBorrado = indiceActual
			}
		case _LIBRE:
			if hayPrimerBorrado {
				return indicePrimerBorrado, false
			}
			return indiceActual, false

		case _OCUPADO:
			if celdaActual.clave == clave {
				return indiceActual, true
			}
		}
	}
	if hayPrimerBorrado {
		return indicePrimerBorrado, false
	}
	return -1, false
}

func (hash *tablaDeHash[K, V]) redimensionarDiccionario(nuevaCapacidad int) {
	tablaActual := hash.tablaHash
	hash.tablaHash = make([]celdaDiccionario[K, V], nuevaCapacidad)
	hash.cantidad = 0
	hash.capacidad = nuevaCapacidad
	for i := 0; i < len(tablaActual); i++ {
		if tablaActual[i].estadoDeCelda == _OCUPADO {
			indiceActual := hash.buscarIndice(tablaActual[i].clave)
			hash.tablaHash[indiceActual].clave = tablaActual[i].clave
			hash.tablaHash[indiceActual].valor = tablaActual[i].valor
			hash.tablaHash[indiceActual].estadoDeCelda = _OCUPADO
			hash.cantidad++
		}
	}
}
