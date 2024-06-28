package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)
	if total != 30 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado: 30", total)
	}
}

func TestDivisao(t *testing.T) {
	total := Divisao(15, 15)
	if total != 1 {
		t.Errorf("Resultado da divisão é inválido: Resultado %d. Esperado: 1", total)
	}
}