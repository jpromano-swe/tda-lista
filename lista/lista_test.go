package lista_test

import (
  TDALista "tdas/lista"
  "testing"

  "github.com/stretchr/testify/require"
)

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
  const volumen = 100000

  for i := 0; i < volumen; i++ {
    lista.InsertarUltimo(i)
  }
  require.False(t, lista.EstaVacia())

  for i := 0; i < volumen; i++ {
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

// 1. Al insertar un elemento en la posición en la que se crea el iterador, efectivamente se inserta al principio.
func TestIteradorInsertarAlPrincipio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	// Se crea el iterador (nace apuntando al principio, o sea al 2)
	iter := lista.Iterador()

	// Insertamos en esa posición
	iter.Insertar(1)

	// CHEQUEOS:
	// El nuevo primero de la lista debería ser 1
	require.Equal(t, 1, lista.VerPrimero())
	// El iterador ahora debería quedar apuntando al elemento que acabamos de insertar
	require.Equal(t, 1, iter.VerActual())
	// El largo ahora tiene que ser 3
	require.Equal(t, 3, lista.Largo())
}

// 2. Insertar un elemento cuando el iterador está al final efectivamente es equivalente a insertar al final.
func TestIteradorInsertarAlFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)

	iter := lista.Iterador()
	// Avanzamos hasta que el iterador se caiga de la lista (llegue al final)
	for iter.HayAlgoMas() {
		iter.Avanzar()
	}

	// Insertamos el 3 estando al final
	iter.Insertar(3)

	// TU TURNO DE CHEQUEAR:
	// - require: El último elemento de la lista (lista.VerUltimo()) tiene que ser 3.
	// - require: El iterador tiene que tener a 3 como VerActual().
	// - require: El largo tiene que ser 3.
}

// 3. Insertar un elemento en el medio se hace en la posición correcta.
func TestIteradorInsertarEnElMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(3) // Queremos meter un 2 entre el 1 y el 3

	iter := lista.Iterador()
	
	// Avanzamos una vez para estar parados sobre el 3
	iter.Avanzar()

	// Insertamos el 2. Esto debería empujar el 3 para atrás.
	iter.Insertar(2)

	// TU TURNO DE CHEQUEAR:
	// - require: iter.VerActual() tiene que ser 2.
	// - require: Si avanzás el iterador (iter.Avanzar()), el nuevo VerActual() tiene que ser 3.
	// - require: El primero sigue siendo 1 y el último sigue siendo 3.
	// - require: Largo es 3.
}

// 4. Al remover el elemento cuando se crea el iterador, cambia el primer elemento de la lista.
func TestIteradorRemoverAlPrincipio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(10)
	lista.InsertarUltimo(20)
	lista.InsertarUltimo(30)

	// Creamos el iterador (apunta al 10)
	iter := lista.Iterador()

	// Borramos el elemento actual
	borrado := iter.Borrar()

	// TU TURNO DE CHEQUEAR:
	// - require: 'borrado' tiene que ser 10.
	// - require: lista.VerPrimero() ahora tiene que ser 20.
	// - require: iter.VerActual() ahora también tiene que ser 20 (porque el 10 desapareció y el 20 tomó su lugar).
	// - require: Largo tiene que ser 2.
}