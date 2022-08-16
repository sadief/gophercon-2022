package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/serve-customer/", ServeCustomer)
	http.ListenAndServe(":8080", nil)
}

func ServeCustomer(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	numCustomers, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/serve-customer/"))
	if err != nil || numCustomers == 0 {
		numCustomers = 1
	}

	count := 0
	for i := 0; i < numCustomers; i++ {
		go MakeCoffee()
		count++
	}

	timeTaken := time.Since(start)
	log.Printf("Took %s to serve coffee to %v customer(s)", timeTaken, count)
}

func MakeCoffee() {
	PayForCoffee()
	MakeEspresso()
	SteamMilk()
}

func PayForCoffee() {
	time.Sleep(2 * time.Second)
	log.Printf("Coffee paid for 💰")
}

func MakeEspresso() {
	time.Sleep(2 * time.Second)
	log.Printf("Espresso made ☕️")
}

func SteamMilk() {
	time.Sleep(2 * time.Second)
	log.Printf("Milk steamed 🥛")
}
