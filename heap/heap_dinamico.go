package heap

const (
  _CAPACIDAD_INICIAL     = 2
  _FACTOR_DE_CRECIMIENTO = 2
  _FACTOR_DE_REDUCCION   = 4
  _ERROR_HEAP_VACIO      = "La cola esta vacia"
)

type heapDinamico[T any] struct {
  datos    []T
  cantidad int
  cmp      func(T, T) int
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
  heap := new(heapDinamico[T])
  heap.datos = make([]T, 0)
  heap.cantidad = _CAPACIDAD_INICIAL
  heap.cmp = funcion_cmp
  return heap

}
func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
  heap := new(heapDinamico[T])
  heap.datos = make([]T, len(arreglo))
  copy(heap.datos, arreglo)
  heap.cantidad = len(arreglo)
  heap.cmp = funcion_cmp

  //APLICAR DOWNHEAP a heap

  return heap
}

func (heap *heapDinamico[T]) EstaVacia() bool {
  return len(heap.datos) == 0
}

func (heap *heapDinamico[T]) Encolar(elemento T) {
  if (len(heap.datos)+1) == heap.cantidad{
    heap.cantidad = heap.cantidad * _FACTOR_DE_CRECIMIENTO
    heap.redimensionar(heap.cantidad)
  }
  heap.datos[len(heap.datos)-1] = elemento
  heap.upheap(len(heap.datos)-1)
}

func (heap *heapDinamico[T]) VerMax() T {
  if heap.EstaVacia() {
    panic(_ERROR_HEAP_VACIO)
  }
  return heap.datos[0]
}
func (heap *heapDinamico[T]) Cantidad() int {
  return len(heap.datos)
}

func (heap *heapDinamico[T]) Desencolar() T {
  if heap.EstaVacia() {
    panic(_ERROR_HEAP_VACIO)
  }
  return heap._desencolar(0, heap.cantidad-1)
}

func (heap *heapDinamico[T]) _desencolar(ini, fin int) T {
  maximoDelHeap := heap.datos[ini]
  if heap.cantidad == 1 {
    heap.cantidad--
    heap.datos = heap.datos[:heap.cantidad]
    return maximoDelHeap
  }
  heap.datos[ini] = heap.datos[fin]
  heap.cantidad--
  heap.datos = heap.datos[:heap.cantidad]
  //downheap(ini)
  return maximoDelHeap
}


func (heap *heapDinamico[T]) upheap(pos int) {
  if heap.cmp(heap.datos[pos], heap.datos[(pos-1)/2]) <= 0{
    return
  }
  if heap.cmp(heap.datos[pos], heap.datos[(pos-1)/2]) > 0{
   heap.datos[pos], heap.datos[(pos-1)/2] =  heap.datos[(pos-1)/2], heap.datos[pos]
  }
  heap.upheap((pos-1)/2)
} 

func (heap *heapDinamico[T]) downheap(pos int) {

}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
  
}

func (heap *heapDinamico[T]) redimensionar(nuevaCapacidad int) {
  nuevosDatos := make([]T, nuevaCapacidad)
  copy(nuevosDatos, heap.datos[:heap.cantidad])
  heap.datos = nuevosDatos
  
}
