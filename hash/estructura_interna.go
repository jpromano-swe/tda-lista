package hash

type Diccionario[K comparable, V any] struct {
}

type IterDiccionario[K comparable, V any] struct {
}

func CrearHash[K comparable, V any]() Diccionario[K, V]

func (hash *Diccionario[K, V]) Pertenece(clave K) bool {
  return true
}

func (hash *Diccionario[K, V]) Obtener(clave K) V {
}

func (hash *Diccionario[K, V]) Borrar(clave K) V {}

func (hash *Diccionario[K, V]) Guardar(clave K, valor V) {}

func (hash *Diccionario[K, V]) Cantidad() int {
  return 1
}

func (hash *Diccionario[K, V]) Iterar(visitar func(clave K, dato V) bool) {
  return true
}

func (iter *IterDiccionario[K, V]) HayAlgoMas() bool {}

func (iter *IterDiccionario[K, V]) VerActual() (K, V) {}

func (iter *IterDiccionario[K, V]) Avanzar() {}
