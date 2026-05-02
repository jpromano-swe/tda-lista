package diccionario


func (hash *tablaDeHash[K, V]) encontrarPrimerOcupado () int{
	i:=0
	contador:=0
	for i<hash.capacidad && hash.tablaHash[i].EstadoDeCelda !=_ocupado{
		i++
		contador++
	}
	return contador
}
