package main

import (
	"fmt"
	"sync"
)

var (
	balance int
	mutex   sync.Mutex
)

func withdraw(value int, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()

	fmt.Printf("Withdrawing amount: %d with balance: %d \n", value, balance)
	balance -= value

	mutex.Unlock()
}

func deposit(value int, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()

	fmt.Printf("Depositing amount: %d with balance: %d \n", value, balance)
	balance += value

	mutex.Unlock()
}

func main() {
	balance = 1000
	var wg sync.WaitGroup

	wg.Add(2)
	go withdraw(500, &wg)
	go deposit(700, &wg)
	wg.Wait()
	fmt.Printf("New balance: %d \n", balance)
}
