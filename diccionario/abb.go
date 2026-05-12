package diccionario

import TDAPila "tdas/pila"

const (
	_ERROR_ITERADOR            = "El iterador termino de iterar"
	_ERROR_CLAVE_NO_ENCONTRADA = "La clave no se encuentra al diccionario"
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

func (arbol *abb[K, V]) Borrar(clave K) V {
	nodoABorrar := arbol.buscarNodo(clave)

	if nodoABorrar == nil {
		panic(_ERROR_CLAVE_NO_ENCONTRADA)
	}
	nodoActual := *nodoABorrar
	datoElemento := nodoActual.dato

	if nodoActual.arbolIzq == nil && nodoActual.arbolDer == nil {
		arbol.borrarHoja(nodoABorrar)
	} else if nodoActual.arbolIzq != nil && nodoActual.arbolDer != nil {
		arbol.borrarNodoConDosHijos(nodoABorrar)
	} else {
		arbol.borrarNodoConUnHijo(nodoABorrar)
	}
	arbol.cantidad--
	return datoElemento
}

func (arbol *abb[K, V]) borrarHoja(nodoHoja **nodo[K, V]) {
	*nodoHoja = nil
}

func (arbol *abb[K, V]) borrarNodoConUnHijo(nodoABorrar **nodo[K, V]) {
	nodoActual := *nodoABorrar

	if nodoActual.arbolIzq != nil {
		*nodoABorrar = nodoActual.arbolIzq
	} else {
		*nodoABorrar = nodoActual.arbolDer
	}
}

func (arbol *abb[K, V]) borrarNodoConDosHijos(nodoABorrar **nodo[K, V]) {
	nodoActual := *nodoABorrar
	nodoAReemplazar := arbol.buscarMaximo(&(nodoActual.arbolIzq))

	if nodoAReemplazar == nil {
		panic("El nodo a reemplazar es invalido")
	}
	nodoReemplazo := *nodoAReemplazar

	nodoActual.clave = nodoReemplazo.clave
	nodoActual.dato = nodoReemplazo.dato

	if nodoReemplazo.arbolIzq != nil {
		arbol.borrarNodoConUnHijo(nodoAReemplazar)
	} else {
		arbol.borrarHoja(nodoAReemplazar)
	}
}

func (arbol *abb[K, V]) buscarMaximo(elemento **nodo[K, V]) **nodo[K, V] {
	if *elemento == nil {
		return nil
	}
	for (*elemento).arbolDer != nil {
		elemento = &(*elemento).arbolDer
	}
	return elemento
}

func (arbol *abb[K, V]) buscarNodo(clave K) **nodo[K, V] {
	enlace := &arbol.raiz

	for *enlace != nil {
		elementoActual := *enlace
		if arbol.cmp(clave, elementoActual.clave) == 0 {
			return enlace
		} else if arbol.cmp(clave, elementoActual.clave) > 0 {
			enlace = &elementoActual.arbolDer
		} else {
			enlace = &elementoActual.arbolIzq
		}
	}
	return nil
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
