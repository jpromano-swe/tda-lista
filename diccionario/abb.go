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
func (arbol *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
  arbol._iterarRango(arbol.raiz, desde, hasta, visitar)
}

func (arbol *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
  arbol._iterarRango(arbol.raiz, nil, nil, visitar)
}

func (arbol *abb[K, V]) _iterarRango(nodoActual *nodo[K, V], desde *K, hasta *K, visitar func(clave K, dato V) bool) bool {
  if nodoActual == nil {
    return true
  }

  claveActual := nodoActual.clave
  if desde == nil || arbol.cmp(claveActual, *desde) > 0 {
    if !arbol._iterarRango(nodoActual.arbolIzq, desde, hasta, visitar) {
      return false
    }
  }

  if (desde == nil || arbol.cmp(claveActual, *desde) >= 0) && (hasta == nil || arbol.cmp(claveActual, *hasta) <= 0) {
    if !visitar(claveActual, nodoActual.dato) {
      return false
    }

  }
  if hasta == nil || arbol.cmp(claveActual, *hasta) < 0 {
    if !arbol._iterarRango(nodoActual.arbolDer, desde, hasta, visitar) {
      return false
    }
  }
  return true
}
