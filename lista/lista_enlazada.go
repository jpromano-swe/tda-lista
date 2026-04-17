package lista

//TODO Juan: func Insertar() lista, pruebas
//TODO Ramiro: funciones de lista, func VerActual() del iterador, func Borrar() lista, pruebas

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

func crearNodo[T any](elem T) *nodo[T] {
  nuevoNodo := nodo[T]{
    dato:      elem,
    siguiente: nil,
  }
  return &nuevoNodo
}

/* TODO IMPLEMENTAR FUNCIONES DE LISTA

func EstaVacia()
func InsertarPrimero()
func BorrarPrimero()
func VerPrimero()
func Largo()
*/

func (lista ListaEnlazada[T]) VerUltimo() T {
  if lista.EstaVacia() {
    panic("La lista esta vacia")
  }
  return lista.ultimo.dato
}

func (lista ListaEnlazada[T]) InsertarUltimo(elem T) {
  nodoNuevo := crearNodo(elem)
  if lista.EstaVacia() {
    lista.primero = nodoNuevo
  }
  lista.ultimo.siguiente = nodoNuevo
  lista.ultimo = nodoNuevo
  lista.largo++
}

/* TODO IMPLEMENTAR FUNCIONES DE ITERADOR

func VerActual() T
func Borrar() T
*/

func (lista *ListaEnlazada[T]) CrearIteradorExterno() *Iterador[T] {
  return &Iterador[T]{nodoAnterior: nil, nodoActual: lista.primero, lista: lista}
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
    iterador.insertarAlPrincipiodeLista(nodoNuevo)
  } else {
    iterador.insertarDespuesDeNodoAnterior(nodoNuevo)
  }
  iterador.nodoActual = nodoNuevo
  iterador.lista.largo++

}

func (iterador *Iterador[T]) insertarAlPrincipiodeLista(nodoNuevo *nodo[T]) {
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
