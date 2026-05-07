package diccionario

import TDAPila "tdas/pila"

const (
  _ERROR_ITERADOR = "El iterador termino de iterar"
)

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

type iteradorABB[K comparable, V any] struct {
  pilaNodos  TDAPila.Pila[any]
  inicio     *K
  fin        *K
  comparador func(K, K) int
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
  arbol.IterarRango(nil, nil, visitar)
}

func (arbol *abb[K, V]) Iterador() IterDiccionario[K, V] {
  return arbol.IteradorRango(nil, nil)
}

func (iter *iteradorABB[K, V]) HayAlgoMas() bool {
  if iter.pilaNodos.EstaVacia() {
    return false
  }
  if iter.fin == nil {
    return true
  }
  claveActual := iter.pilaNodos.VerTope().(*nodo[K, V]).clave
  return iter.comparador(claveActual, *iter.fin) <= 0
}

func (iter *iteradorABB[K, V]) Avanzar() {
  if !iter.HayAlgoMas() {
    panic(_ERROR_ITERADOR)
  }
  nodoActual := iter.pilaNodos.VerTope().(*nodo[K, V])
  iter.pilaNodos.Desapilar()
  if nodoActual.arbolDer != nil {
    iter.apilarRamaIzq(nodoActual.arbolDer)
  }
}

func (iter *iteradorABB[K, V]) VerActual() (K, V) {
  if !iter.HayAlgoMas() {
    panic(_ERROR_ITERADOR)
  }
  nodoActual := iter.pilaNodos.VerTope().(*nodo[K, V])
  return nodoActual.clave, nodoActual.dato
}

func (arbol *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
  iter := iteradorABB[K, V]{
    pilaNodos:  TDAPila.CrearPilaDinamica[any](),
    inicio:     desde,
    fin:        hasta,
    comparador: arbol.cmp,
  }
  iter.apilarRango(arbol.raiz)
  return &iter
}

func (iter *iteradorABB[K, V]) apilarRamaIzq(nodoActual *nodo[K, V]) {
  for nodoActual != nil {
    iter.pilaNodos.Apilar(nodoActual)
    nodoActual = nodoActual.arbolIzq
  }
}

func (iter *iteradorABB[K, V]) apilarRango(nodoActual *nodo[K, V]) {
  if iter.inicio == nil {
    iter.apilarRamaIzq(nodoActual)
    return
  }
  for nodoActual != nil {
    claveActual := nodoActual.clave
    if iter.comparador(claveActual, *iter.inicio) < 0 {
      nodoActual = nodoActual.arbolDer
    } else {
      iter.pilaNodos.Apilar(nodoActual)
      nodoActual = nodoActual.arbolIzq
    }
  }
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
