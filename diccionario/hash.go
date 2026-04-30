package diccionario

const (
  _libre EstadoDeCelda = iota
  _ocupado
  _borrado

  _capacidadInicial    = 7
  _factorDeCrecimiento = 2
  _factorDeCarga       = 0.7
)

type EstadoDeCelda = int

type tablaDeHash[K comparable, V any] struct {
  tablaHash []celdaDiccionario[K, V]
  cantidad  int
}

type celdaDiccionario[K comparable, V any] struct {
  clave K
  valor V
  EstadoDeCelda
}
type iteradorHash[K comparable, V any] struct{}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
  diccionario := new(tablaDeHash[K, V])
  diccionario.tablaHash = make([]celdaDiccionario[K, V], _capacidadInicial)
  diccionario.cantidad = 0
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
  panic("No se implemento Borrar")
}

func (hash *tablaDeHash[K, V]) Guardar(clave K, valor V) { //REFACTOR, QUEDO MUY LARGO

  indiceActual := hash.buscarIndiceParaInsertar(clave)
  capacidadActual := len(hash.tablaHash)
  factorCargado := float64(hash.cantidad+1) / float64(capacidadActual)

  if hash.tablaHash[indiceActual].EstadoDeCelda == _ocupado {
    hash.tablaHash[indiceActual].valor = valor
    return
  }

  if factorCargado > _factorDeCarga {
    nuevaCapacidad := siguientePrimo(_factorDeCrecimiento * capacidadActual)
    hash.redimensionarDiccionario(nuevaCapacidad)
    indiceActual = hash.buscarIndiceParaInsertar(clave)
  }

  hash.tablaHash[indiceActual].clave = clave
  hash.tablaHash[indiceActual].valor = valor
  hash.tablaHash[indiceActual].EstadoDeCelda = _ocupado
  hash.cantidad++
}

func (hash *tablaDeHash[K, V]) Cantidad() int {
  return hash.cantidad
}

func (hash *tablaDeHash[K, V]) Iterar(visitar func(clave K, dato V) bool) {
  return
}

func (hash *tablaDeHash[K, V]) Iterador() IterDiccionario[K, V] {
  return &iteradorHash[K, V]{}
}

func (iter *iteradorHash[K, V]) HayAlgoMas() bool {
  return false
}

func (iter *iteradorHash[K, V]) VerActual() (K, V) {
  panic("El iterador termino de iterar")
}

func (iter *iteradorHash[K, V]) Avanzar() {
  panic("El iterador termino de iterar")
}

func (hash *tablaDeHash[K, V]) redimensionarDiccionario(nuevaCapacidad int) {
  tablaActual := hash.tablaHash
  hash.tablaHash = make([]celdaDiccionario[K, V], nuevaCapacidad)
  hash.cantidad = 0

  for i := 0; i < len(tablaActual); i++ {
    if tablaActual[i].EstadoDeCelda == _ocupado {
      indiceActual := hash.buscarIndiceParaInsertar(tablaActual[i].clave)
      hash.tablaHash[indiceActual].clave = tablaActual[i].clave
      hash.tablaHash[indiceActual].valor = tablaActual[i].valor
      hash.tablaHash[indiceActual].EstadoDeCelda = _ocupado
      hash.cantidad++
    }
  }
}
