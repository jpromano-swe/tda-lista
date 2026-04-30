package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const _volumenPruebas = 100000

func TestNuevaListaComienzaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.Equal(t, lista.Largo(), 0)
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerPrimero()
	})
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerUltimo()
	})
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.BorrarPrimero()
	})
}

func TestAgregarQuitarElementoIncrementaDecrementaLargo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.Equal(t, lista.Largo(), 0)
	lista.InsertarPrimero(5)
	require.Equal(t, lista.Largo(), 1)
	lista.InsertarUltimo(8)
	require.Equal(t, lista.Largo(), 2)
	lista.BorrarPrimero()
	require.Equal(t, lista.Largo(), 1)
	lista.BorrarPrimero()
	require.Equal(t, lista.Largo(), 0)
}

func TestListaVaciadaSeComportaComoVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	lista.InsertarPrimero(5)
	require.False(t, lista.EstaVacia())
	require.Equal(t, lista.VerPrimero(), 5)
	require.Equal(t, lista.VerUltimo(), 5)
	lista.BorrarPrimero()
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerPrimero()
	})
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerUltimo()
	})
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.BorrarPrimero()
	})
}

func TestListaDeStringsSeComportaComoCola(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	require.True(t, lista.EstaVacia())
	lista.InsertarPrimero("Fundamentos")
	require.False(t, lista.EstaVacia())
	require.Equal(t, lista.VerPrimero(), "Fundamentos")
	require.Equal(t, lista.VerUltimo(), "Fundamentos")
	lista.BorrarPrimero()
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerPrimero()
	})
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.BorrarPrimero()
	})
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerUltimo()
	})
}

func TestInsertarPrimeroMantieneElUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	lista.InsertarPrimero(5)
	require.Equal(t, lista.VerPrimero(), 5)
	require.Equal(t, lista.VerUltimo(), 5)
	lista.InsertarPrimero(3)
	require.Equal(t, lista.VerPrimero(), 3)
	require.Equal(t, lista.VerUltimo(), 5)
	lista.InsertarPrimero(1)
	require.Equal(t, lista.VerPrimero(), 1)
	require.Equal(t, lista.VerUltimo(), 5)
}

func TestInsertarUltimoMantieneElPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	lista.InsertarUltimo(5)
	require.Equal(t, lista.VerPrimero(), 5)
	require.Equal(t, lista.VerUltimo(), 5)
	lista.InsertarUltimo(3)
	require.Equal(t, lista.VerPrimero(), 5)
	require.Equal(t, lista.VerUltimo(), 3)
	lista.InsertarUltimo(1)
	require.Equal(t, lista.VerPrimero(), 5)
	require.Equal(t, lista.VerUltimo(), 1)
}

func TestInsertarYBorrarVolumenNoRompeElPrograma(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	for i := 0; i < _volumenPruebas; i++ {
		lista.InsertarUltimo(i)
	}
	require.False(t, lista.EstaVacia())

	for i := 0; i < _volumenPruebas; i++ {
		lista.BorrarPrimero()
	}
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.BorrarPrimero()
	})
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerPrimero()
	})
}

func TestIteradorListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()

	require.False(t, iter.HayAlgoMas())

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter.VerActual()
	})

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter.Avanzar()
	})

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter.Borrar()
	})
}

func TestIteradorInsertarAlInicio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(5)
	lista.InsertarUltimo(3)
	require.Equal(t, lista.Largo(), 2)

	iter := lista.Iterador()
	iter.Insertar(1)

	require.Equal(t, lista.VerPrimero(), 1)
	require.Equal(t, lista.VerUltimo(), 3)
	require.Equal(t, lista.Largo(), 3)
	require.Equal(t, iter.VerActual(), 1)

	datosCopiados := []int{}

	lista.Iterar(func(dato int) bool {
		datosCopiados = append(datosCopiados, dato)
		return true
	})
	require.Equal(t, []int{1, 5, 3}, datosCopiados)
}

func TestIteradorInsertarEnElMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(5)
	lista.InsertarUltimo(3)
	require.Equal(t, lista.Largo(), 2)

	iter := lista.Iterador()
	iter.Avanzar()

	require.Equal(t, iter.VerActual(), 3)

	iter.Insertar(4)

	require.Equal(t, lista.VerPrimero(), 5)
	require.Equal(t, lista.VerUltimo(), 3)
	require.Equal(t, lista.Largo(), 3)
	require.Equal(t, iter.VerActual(), 4)

	datosCopiados := []int{}

	lista.Iterar(func(dato int) bool {
		datosCopiados = append(datosCopiados, dato)
		return true
	})
	require.Equal(t, []int{5, 4, 3}, datosCopiados)
}

func TestIteradorInsertarAlFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(5)
	lista.InsertarUltimo(3)
	require.Equal(t, lista.Largo(), 2)

	iter := lista.Iterador()
	iter.Avanzar()
	iter.Avanzar()

	require.False(t, iter.HayAlgoMas())

	iter.Insertar(1)

	require.Equal(t, lista.VerPrimero(), 5)
	require.Equal(t, lista.VerUltimo(), 1)
	require.Equal(t, lista.Largo(), 3)
	require.Equal(t, iter.VerActual(), 1)

	datosCopiados := []int{}

	lista.Iterar(func(dato int) bool {
		datosCopiados = append(datosCopiados, dato)
		return true
	})
	require.Equal(t, []int{5, 3, 1}, datosCopiados)
}

func TestIteradorAlInicioCambiaPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)

	iter := lista.Iterador()

	require.Equal(t, iter.VerActual(), 3)
	require.Equal(t, iter.Borrar(), 3)

	require.Equal(t, lista.VerPrimero(), 4)
	require.Equal(t, lista.VerUltimo(), 5)
	require.Equal(t, lista.Largo(), 2)
	require.Equal(t, iter.VerActual(), 4)

	datosCopiados := []int{}

	lista.Iterar(func(dato int) bool {
		datosCopiados = append(datosCopiados, dato)
		return true
	})
	require.Equal(t, []int{4, 5}, datosCopiados)
}

func TestIteradorBorrarUltimoCambiaUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)

	iter := lista.Iterador()
	iter.Avanzar()
	iter.Avanzar()

	require.Equal(t, iter.VerActual(), 5)
	require.Equal(t, iter.Borrar(), 5)

	require.Equal(t, lista.VerPrimero(), 3)
	require.Equal(t, lista.VerUltimo(), 4)
	require.Equal(t, lista.Largo(), 2)
	require.False(t, iter.HayAlgoMas())
}

func TestIteradorBorraElementoEnMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)

	iter := lista.Iterador()
	iter.Avanzar()

	require.Equal(t, iter.VerActual(), 4)
	require.Equal(t, iter.Borrar(), 4)

	require.Equal(t, lista.Largo(), 2)
	require.Equal(t, lista.VerPrimero(), 3)
	require.Equal(t, lista.VerUltimo(), 5)
	require.Equal(t, iter.VerActual(), 5)

	datosCopiados := []int{}

	lista.Iterar(func(dato int) bool {
		datosCopiados = append(datosCopiados, dato)
		return true
	})
	require.Equal(t, []int{3, 5}, datosCopiados)
}

func TestIteradorBorrarUnicoElementoDejaListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(3)

	iter := lista.Iterador()

	require.Equal(t, iter.VerActual(), 3)
	require.Equal(t, iter.Borrar(), 3)

	require.True(t, lista.EstaVacia())
	require.Equal(t, lista.Largo(), 0)
	require.False(t, iter.HayAlgoMas())

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerPrimero()
	})
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerUltimo()
	})
}

func TestIteradorInsertarElementoEnListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()

	require.False(t, iter.HayAlgoMas())

	iter.Insertar(2)
	require.False(t, lista.EstaVacia())
	require.Equal(t, lista.Largo(), 1)
	require.Equal(t, lista.VerPrimero(), 2)
	require.Equal(t, lista.VerUltimo(), 2)
	require.True(t, iter.HayAlgoMas())
	require.Equal(t, iter.VerActual(), 2)
}

func TestIteradorInternoEnListaVaciaNoRecorreElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	seUsoIterador := false

	lista.Iterar(func(elem int) bool {
		seUsoIterador = true
		return true
	})

	require.False(t, seUsoIterador)
}

func TestIteradorInternoRecorreTodosLosElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)

	datosCopiados := []int{}

	lista.Iterar(func(dato int) bool {
		datosCopiados = append(datosCopiados, dato)
		return true
	})
	require.Equal(t, []int{3, 4, 5}, datosCopiados)
}

func TestIteradorInternoCortaConCondicion(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)

	datosCopiados := []int{}

	lista.Iterar(func(dato int) bool {
		datosCopiados = append(datosCopiados, dato)
		return dato != 3
	})
	require.Equal(t, []int{1, 2, 3}, datosCopiados)
}
