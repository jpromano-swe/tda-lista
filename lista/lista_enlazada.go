package lista

const (
	_errorListaVacia        = "La lista esta vacia"
	_errorIteradorTerminado = "El iterador termino de iterar"
)

type listaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
	largo   int
}

type nodo[T any] struct {
	dato      T
	siguiente *nodo[T]
}

type iterador[T any] struct {
	nodoAnterior *nodo[T]
	nodoActual   *nodo[T]
	lista        *listaEnlazada[T]
}

func CrearListaEnlazada[T any]() *listaEnlazada[T] {
	return &listaEnlazada[T]{
		primero: nil,
		ultimo:  nil,
		largo:   0,
	}
}

func crearNodo[T any](elem T) *nodo[T] {
	nuevoNodo := nodo[T]{
		dato:      elem,
		siguiente: nil,
	}
	return &nuevoNodo
}

func (lista listaEnlazada[T]) EstaVacia() bool {
	return lista.primero == nil
}

func (lista *listaEnlazada[T]) InsertarPrimero(elem T) {
	nuevoNodo := crearNodo(elem)
	if lista.EstaVacia() {
		lista.ultimo = nuevoNodo
	} else {
		nuevoNodo.siguiente = lista.primero
	}
	lista.primero = nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic(_errorListaVacia)
	}
	aux := lista.primero.dato

	lista.primero = lista.primero.siguiente
	if lista.primero == nil {
		lista.ultimo = nil
	}
	lista.largo--
	return aux
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic(_errorListaVacia)
	}
	return lista.primero.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic(_errorListaVacia)
	}
	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) InsertarUltimo(elem T) {
	nodoNuevo := crearNodo(elem)
	if lista.EstaVacia() {
		lista.primero = nodoNuevo
	} else {
		lista.ultimo.siguiente = nodoNuevo
	}
	lista.ultimo = nodoNuevo
	lista.largo++
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterador[T]{nodoAnterior: nil, nodoActual: lista.primero, lista: lista}
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	nodoActual := lista.primero

	for nodoActual != nil {
		if !visitar(nodoActual.dato) {
			return
		}
		nodoActual = nodoActual.siguiente
	}
}

func (iterador *iterador[T]) HayAlgoMas() bool {
	return iterador.nodoActual != nil
}

func (iterador *iterador[T]) Avanzar() {
	if !iterador.HayAlgoMas() {
		panic(_errorIteradorTerminado)
	}
	iterador.nodoAnterior = iterador.nodoActual
	iterador.nodoActual = iterador.nodoActual.siguiente
}

func (iterador *iterador[T]) Insertar(elem T) {
	nodoNuevo := crearNodo(elem)
	nodoNuevo.siguiente = iterador.nodoActual

	if iterador.nodoAnterior == nil {
		iterador.lista.primero = nodoNuevo
	} else {
		iterador.nodoAnterior.siguiente = nodoNuevo
	}

	if iterador.nodoActual == nil {
		iterador.lista.ultimo = nodoNuevo
	}

	iterador.nodoActual = nodoNuevo
	iterador.lista.largo++
}

func (iterador *iterador[T]) VerActual() T {
	if !iterador.HayAlgoMas() {
		panic(_errorIteradorTerminado)
	}
	return iterador.nodoActual.dato
}

func (iterador *iterador[T]) Borrar() T {
	if !iterador.HayAlgoMas() {
		panic(_errorIteradorTerminado)
	}

	aux := iterador.nodoActual.dato

	if iterador.nodoAnterior == nil {
		iterador.lista.primero = iterador.nodoActual.siguiente
	} else {
		iterador.nodoAnterior.siguiente = iterador.nodoActual.siguiente
	}

	if iterador.nodoActual.siguiente == nil {
		iterador.lista.ultimo = iterador.nodoAnterior
	}

	iterador.nodoActual = iterador.nodoActual.siguiente

	if iterador.lista.primero == nil {
		iterador.lista.ultimo = nil
	}

	iterador.lista.largo--
	return aux
}
