package lista

//TODO Juan: faltan pruebas
//TODO Ramiro: faltan pruebas

type ListaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
	largo   int
}

type nodo[T any] struct {
	dato      T
	siguiente *nodo[T]
}

type Iterador[T any] struct {
	nodoAnterior *nodo[T]
	nodoActual   *nodo[T]
	lista        *ListaEnlazada[T]
}

func CrearListaEnlazada[T any]() *ListaEnlazada[T] {
	return &ListaEnlazada[T]{
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

func (lista ListaEnlazada[T]) EstaVacia() bool {
	return lista.primero == nil
}

func (lista *ListaEnlazada[T]) InsertarPrimero(elem T) {
	nuevoNodo := crearNodo(elem)
	if lista.EstaVacia() {
		lista.ultimo = nuevoNodo
	} else {
		nuevoNodo.siguiente = lista.primero
	}
	lista.primero = nuevoNodo
	lista.largo++
}

func (lista *ListaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	aux := lista.primero.dato

	if lista.primero == lista.ultimo {
		lista.primero = nil
		lista.ultimo = nil
	} else {
		lista.primero = lista.primero.siguiente
	}
	lista.largo--
	return aux
}

func (lista *ListaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.primero.dato
}

func (lista *ListaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista ListaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.ultimo.dato
}

func (lista *ListaEnlazada[T]) InsertarUltimo(elem T) {
	nodoNuevo := crearNodo(elem)
	if lista.EstaVacia() {
		lista.primero = nodoNuevo
		lista.ultimo = nodoNuevo
	} else {
		lista.ultimo.siguiente = nodoNuevo
		lista.ultimo = nodoNuevo
	}
	lista.largo++
}

func (lista *ListaEnlazada[T]) Iterador() IteradorLista[T] {
	return &Iterador[T]{nodoAnterior: nil, nodoActual: lista.primero, lista: lista}
}

func (lista *ListaEnlazada[T]) Iterar(visitar func(T) bool) {
	nodoActual := lista.primero

	for nodoActual != nil {
		if !visitar(nodoActual.dato) {
			return
		}
		nodoActual = nodoActual.siguiente
	}
}

func (iterador *Iterador[T]) HayAlgoMas() bool {
	return iterador.nodoActual != nil
}

func (iterador *Iterador[T]) Avanzar() {
	if !iterador.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	iterador.nodoAnterior = iterador.nodoActual
	iterador.nodoActual = iterador.nodoActual.siguiente
}

func (iterador *Iterador[T]) Insertar(elem T) {
	nodoNuevo := crearNodo(elem)

	if iterador.nodoAnterior == nil {
		iterador.insertarAlPrincipioDeLista(nodoNuevo)
	} else {
		iterador.insertarDespuesDeNodoAnterior(nodoNuevo)
	}
	iterador.nodoActual = nodoNuevo
	iterador.lista.largo++
}

func (iterador *Iterador[T]) VerActual() T {
	if !iterador.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	return iterador.nodoActual.dato
}

func (iterador *Iterador[T]) Borrar() T {
	if !iterador.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}

	aux := iterador.nodoActual.dato

	if iterador.lista.primero == iterador.lista.ultimo {
		iterador.lista.primero = nil
		iterador.lista.ultimo = nil
		iterador.nodoActual = nil
	} else if iterador.nodoAnterior == nil {
		iterador.lista.primero = iterador.nodoActual.siguiente
		iterador.nodoActual = iterador.lista.primero
	} else if iterador.nodoActual.siguiente == nil {
		iterador.lista.ultimo = iterador.nodoAnterior
		iterador.nodoAnterior.siguiente = nil
		iterador.nodoActual = nil
	} else {
		iterador.nodoAnterior.siguiente = iterador.nodoActual.siguiente
		iterador.nodoActual = iterador.nodoActual.siguiente
	}
	iterador.lista.largo -= 1
	return aux
}

func (iterador *Iterador[T]) insertarAlPrincipioDeLista(nodoNuevo *nodo[T]) {
	nodoNuevo.siguiente = iterador.lista.primero
	iterador.lista.primero = nodoNuevo

	if iterador.lista.ultimo != nil {
		return
	}
	iterador.lista.ultimo = nodoNuevo
}

func (iterador *Iterador[T]) insertarDespuesDeNodoAnterior(nodoNuevo *nodo[T]) {
	nodoNuevo.siguiente = iterador.nodoActual
	iterador.nodoAnterior.siguiente = nodoNuevo

	if iterador.nodoActual == nil {
		iterador.lista.ultimo = nodoNuevo
	}
}

/*type Iterador[T any] struct {
	nodoAnterior *nodo[T]
	nodoActual   *nodo[T]
	lista        *ListaEnlazada[T]
}*/