package lista

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

func (lista ListaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.ultimo.dato
}

func (lista *ListaEnlazada[T]) CrearIteradorExterno() *Iterador[T] {
	return &Iterador[T]{nodoAnterior: lista.primero, nodoActual: lista.primero, lista: lista}
}

func (iterador *Iterador[T]) HayAlgoMas() bool {
	return iterador.nodoActual.siguiente != nil
}

func (iterador *Iterador[T]) Avanzar() {
	if iterador.nodoActual.siguiente == nil {
		panic("El iterador termino de iterar")
	}
	iterador.nodoActual = iterador.nodoActual.siguiente
	iterador.nodoAnterior = iterador.nodoActual
}
