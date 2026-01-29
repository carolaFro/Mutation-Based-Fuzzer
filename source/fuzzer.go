package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

//variable declaration
var (
	prng_seed         uint32
	num_of_iterations uint32
	size              int64
	file              string
)

// this is a default seed taken from the seed0 file provided by the staff.
func get_default() []byte {
	return []byte{16, 0, 0, 0, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 10}
}

func main() {
	file = "seed"
	pOf32 := uint64(1) << 32

	switch len(os.Args) {
	case 2:
		// config file mode
		configPath := os.Args[1]
		data, err := os.ReadFile(configPath)
		if err != nil {
			log.Fatal(err)
		}

		fields := strings.Fields(string(data))
		if len(fields) < 2 {
			log.Fatal("Config must contain: <prng_seed> <num_of_iters>")
		}

		num1, err1 := strconv.ParseUint(fields[0], 10, 64)
		num2, err2 := strconv.ParseUint(fields[1], 10, 64)
		if err1 != nil || err2 != nil {
			log.Fatal("Config values must be integers")
		}

		if num1 == 0 || num1 >= pOf32 || num2 == 0 || num2 >= pOf32 {
			log.Fatal("Values must be 32-bit positive integers")
		}

		prng_seed = uint32(num1)
		num_of_iterations = uint32(num2)

	case 3:
		// two args
		num1, err1 := strconv.ParseUint(os.Args[1], 10, 64)
		num2, err2 := strconv.ParseUint(os.Args[2], 10, 64)
		if err1 != nil || err2 != nil {
			log.Fatal("You need two 32 bit integers")
		}

		if num1 == 0 || num1 >= pOf32 || num2 == 0 || num2 >= pOf32 {
			log.Fatal("Values must be 32-bit positive integers")
		}

		prng_seed = uint32(num1)
		num_of_iterations = uint32(num2)

	default:
		log.Fatal("Usage: ./fuzzer <config_file> OR ./fuzzer <prng_seed> <num_of_iters>")
	}
	//cheking if file exists
	size = 0
	f, err := os.Stat(file)
	if err == nil {
		size = f.Size()
	} else { //if no fine then I call the default seed (from seed0)
		size = int64(len(get_default()))
	}
	seed := make([]byte, size)
	seed, err = os.ReadFile(file)
	if err != nil {
		seed = get_default()
	}
	// main loop
	for i := 0; uint32(i) < num_of_iterations; i++ {
		rand.Seed(int64(prng_seed)) //this seeds all use of rand()
		if i%500 == 0 {             //checking for every 500 iters to grow the input by 10 bytes
			for j := 0; j < 10; j++ {
				seed = append(seed, byte(rand.Intn(255)))
			}
		}
		for k := 0; k < len(seed); k++ { //this is where the mutation happens
			p := rand.Intn(100)
			if p <= 13 { //prob of 13%
				seed[k] = byte(rand.Intn(255))
			}
		}
	}
	out := string(seed[:]) //from slice to string
	fmt.Println(out)       //printing final result

}
