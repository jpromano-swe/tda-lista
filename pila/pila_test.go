package pila_test

import (
	"github.com/stretchr/testify/require"
	TDAPila "tdas/pila"
	"testing"
)

func TestCreacionPila(t *testing.T) {
	p := TDAPila.CrearPilaDinamica[int]()
	require.True(t, p.EstaVacia())
	require.Panics(t, func() { p.VerTope() })
	require.Panics(t, func() { p.Desapilar() })
}

func TestApilar(t *testing.T) {
	p := TDAPila.CrearPilaDinamica[int]()

	p.Apilar(10)
	p.Apilar(20)
	p.Apilar(30)
}

func TestDesapilar(t *testing.T) {
	p := TDAPila.CrearPilaDinamica[int]()

	p.Apilar(10)
	p.Apilar(20)
	p.Apilar(30)
	require.Equal(t, 30, p.Desapilar())
	require.Equal(t, 20, p.Desapilar())
	p.Apilar(40)
	require.Equal(t, 40, p.Desapilar())
	require.Equal(t, 10, p.Desapilar())

	require.True(t, p.EstaVacia())
}

func TestVolumen(t *testing.T) {
	p := TDAPila.CrearPilaDinamica[int]()

	for i := 0; i < 5000; i++ {
		p.Apilar(i)
		require.Equal(t, i, p.VerTope())

	}

	for i := 4999; i >= 0; i-- {
		require.Equal(t, i, p.VerTope())
		require.Equal(t, i, p.Desapilar())

	}

	require.True(t, p.EstaVacia())
}

func TestConStrings(t *testing.T) {
	p := TDAPila.CrearPilaDinamica[string]()

	p.Apilar("algoritmos")
	p.Apilar("2")
	p.Apilar("fiuba")

	require.Equal(t, "fiuba", p.Desapilar())
	require.Equal(t, "2", p.Desapilar())
	require.Equal(t, "algoritmos", p.Desapilar())

}
