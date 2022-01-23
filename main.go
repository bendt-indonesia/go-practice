package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)


func main() {
	newCar := car{1,"March","Nissan"}
	fmt.Println(newCar)
}

func randomNumber(max int) int {
	//seed value
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return r.Intn(max)
}

func array() {
	//Semua data di array itu harus sama tipe datanya.
	//Misalnya Array of Integers, Array of Strings
	numbers := []int {3,5,6}
	//names := []string {"Ben","Walandow"}

	//Add item to array
	numbers = append(numbers,7)
	fmt.Println(numbers) //[3 5 6 7]
	fmt.Println(numbers[:1]) //[3] => dari index awal(0) sampai < index 1
	fmt.Println(numbers[:2]) //[3 5]
	fmt.Println(numbers[2:]) //[6 7] => dati index 2 sampai index terakhir
	fmt.Println(numbers[1:]) //[5 6 7]
	fmt.Println(len(numbers))

	//Shuffle Arry
	for i := range numbers {
		max_index := len(numbers) - 1;
		newPosition := randomNumber(max_index)
		numbers[i], numbers[newPosition] = numbers[newPosition], numbers[i]
	}


	fmt.Println(numbers) //[3 5 6 7]
}

func looping() {
	texts := []string{"ben","walandow","cung"}
	for idx, text := range texts {
		fmt.Println( idx, text )
	}
}

func writeFile() {
	text := []string{"ben","walandow"}
	joined := strings.Join(text, " ")
	ioutil.WriteFile("name.txt", []byte(joined),0666)
}

func readFile() {
	bs, err := ioutil.ReadFile("name.txt")
	if err != nil {
		fmt.Println("Error: ",err)
		os.Exit(1)
	}

	fmt.Println(string(bs));


}
