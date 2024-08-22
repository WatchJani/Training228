package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := CreateCommunicationChanel()
	receive := NewReceiver(ch)
	sender := NewSender(ch)

	go receive.DoSomething()

	var (
		counter int
		buffer  []byte = make([]byte, 4096)
	)

	for {
		counter = 0
		numIteration := rand.Intn(30)

		for range numIteration {
			time.Sleep(time.Duration(100) * time.Millisecond)

			counter++
			buffer[0] = byte(counter)
		}

		sender.Send(buffer, counter)
	}
}

type Linker struct {
	store   []byte //reference
	counter int
}

func NewLinker(store []byte, counter int) Linker {
	return Linker{
		store:   store,
		counter: counter,
	}
}

type Receiver struct {
	linker chan Linker
}

func NewReceiver(linker chan Linker) Receiver {
	return Receiver{
		linker: linker,
	}
}

func (r *Receiver) DoSomething() {
	for {
		data := r.Receive()
		fmt.Println("receive")
		fmt.Println(data.counter)
	}
}

func (r *Receiver) Receive() Linker {
	return <-r.linker
}

type Sender struct {
	linker chan Linker
}

func NewSender(linker chan Linker) Sender {
	return Sender{
		linker: linker,
	}
}

func (s *Sender) Send(store []byte, counter int) {
	s.linker <- NewLinker(store, counter)
}

func CreateCommunicationChanel() chan Linker {
	return make(chan Linker)
}
