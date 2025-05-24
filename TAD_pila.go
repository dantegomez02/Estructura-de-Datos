package main

import (
	"fmt"
)

type Nodo struct {
	dato      int
	siguiente *Nodo
}

type Pila struct {
	tope *Nodo
	tamanio int
}

func nuevaPila() *Pila {
	return &Pila{
		tope:    nil,
		tamanio: 0,
	}
}
func (p *Pila) apilar(valor int) {
	nuevoNodo := &Nodo{
		dato:     valor,
		siguiente: p.tope,
	}
	p.tope = nuevoNodo
	p.tamanio++
}

func (p *Pila) esVacia() bool {
	if p.tope == nil {
		return true
	}else{
		return false
	}
}
func (p *Pila) topePila() int {
	return p.tope.dato
}

func (p *Pila) desapilar() {
	if p.tope == nil {
		return
	}
	p.tope = p.tope.siguiente
	p.tamanio--
}

func (p *Pila) size() int {
	return p.tamanio
}

func main() {
	pila := nuevaPila()
	pila.apilar(30)
	pila.apilar(15)
	pila.apilar(45)

	fmt.Println(pila.topePila())
	fmt.Println(pila.size())
	fmt.Println(pila.esVacia())

}
