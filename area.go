package area

import "math"

//Pi é uma proporção numéria definida pela relação entre
//o perimetro de uma circunferencia e seu diâmetro

const Pi = 3.1415

//Circ é responsável por calcular a área da circunferência
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Rect é responsável por calcular a área de um quadrado
func Rect(base, altura float64) float64 {
	return base * altura
}

//Não é visível
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
