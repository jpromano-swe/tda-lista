package pila

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := &pilaDinamica[T]{
		datos:    make([]T, 2),
		cantidad: 0,
	}
	return pila
}

func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	tope := p.datos[p.cantidad-1]
	return tope
}

func (p *pilaDinamica[T]) Apilar(dato T) {
	if p.cantidad == len(p.datos) {
		p.redimensionar(len(p.datos) * 2)
	}
	p.datos[p.cantidad] = dato
	p.cantidad++
}

func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	aux := p.datos[p.cantidad-1]
	p.cantidad--
	if p.cantidad*4 <= len(p.datos) && len(p.datos) > 1 {
		p.redimensionar(len(p.datos) / 2)
	}
	return aux
}
func (p *pilaDinamica[T]) redimensionar(tam int) {
	aux := make([]T, tam)
	copy(aux, p.datos[:p.cantidad])
	p.datos = aux
}
