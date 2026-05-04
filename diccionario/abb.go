package diccionario

type abb[K comparable, V any] struct {
  raiz     *nodo[K, V]
  cantidad int
  cmp      func(K, K) int
}

type nodo[K comparable, V any] struct {
  clave    K
  dato     V
  arbolIzq *nodo[K, V]
  arbolDer *nodo[K, V]
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
  nuevoArbol := new(abb[K, V])
  return &nuevoArbol
}

func crearNodo[K comparable, V any](elem any) *nodo[K, V] {
  nuevoNodo := nodo[K, V]{
    dato:     elem,
    arbolIzq: new(nodo[K, V]),
    arbolDer: new(nodo[K, V]),
  }
  return &nuevoNodo
}
