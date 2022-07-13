package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		items := []string{
			"Dragan",
			"Amer",
			"Danijela",
			"Rade",
			"Leta",
		}

		d1 := []byte("hello go")

		err := os.WriteFile("/data/dat1", d1, 0644)
		if err != nil {
			fmt.Println("Can not create /data/dat1")
		}

		err = os.WriteFile("/tmp/dat2", d1, 0644)
		if err != nil {
			fmt.Println("Can not create /tmp/dat2")
		}

		var tmp string
		for i := 0; i < 10; i++ {
			left := int8(rand.Intn(len(items)))
			right := int8(rand.Intn(len(items)))
			tmp = items[left]
			items[left] = items[right]
			items[right] = tmp
		}

		fmt.Printf("Next: %v\n", items)
		for i, n := range items {
			out := fmt.Sprintf("%v: %v\n", i+1, n)
			if _, err := writer.Write([]byte(out)); err != nil {
				fmt.Println("Server Error")
				writer.WriteHeader(503)
				return
			}
		}
	})

	fmt.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Server Error\n%v\n", err)
	}
}

func removeItem(items []string, idx int8) []string {
	return append(items[:idx], items[idx+1:]...)
}
