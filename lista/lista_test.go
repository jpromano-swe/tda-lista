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

func TestInsertarPrimeroIterador(t *testing.T){
  lista:=TDALista.CrearListaEnlazada[int]()

}
