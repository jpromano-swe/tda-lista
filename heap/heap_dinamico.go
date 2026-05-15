package heap

const (
	_ERROR_HEAP_VACIO = "La cola esta vacia"
)

type heapDinamico[T any] struct {
	datos    []T
	cantidad int
	cmp      func(T, T) int
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := new(heapDinamico[T])
	heap.datos = make([]T, 0)
	heap.cantidad = 0
	heap.cmp = funcion_cmp
	return heap

}
func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := new(heapDinamico[T])
	heap.datos = make([]T, len(arreglo))
	copy(heap.datos, arreglo)
	heap.cantidad = len(arreglo)
	heap.cmp = funcion_cmp

	//APLICAR DOWHEAP a heap

	return heap
}

func (heap *heapDinamico[T]) EstaVacia() bool {
	return heap.cantidad == 0
}

func (heap *heapDinamico[T]) Encolar(T) {

}

func (heap *heapDinamico[T]) VerMax() T {
	if heap.EstaVacia() {
		panic(_ERROR_HEAP_VACIO)
	}
	return heap.datos[0]
}

func (heap *heapDinamico[T]) Desencolar() T {
	if heap.EstaVacia() {
		panic(_ERROR_HEAP_VACIO)
	}
	return
}

func (heap *heapDinamico[T]) Cantidad() int {
	return heap.cantidad
}

func (heap *heapDinamico[T]) upheap(pos int) {

}

func (heap *heapDinamico[T]) downheap(pos int) {
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {

}
