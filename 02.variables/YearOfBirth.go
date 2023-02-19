package main

import (
  "fmt"
  "os"
  "strconv"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Printf("What year were you born?: ")
  scanner.Scan()
  input, err := strconv.ParseInt(scanner.Text(),10,64)
  fmt.Printf("You will be %d years old at the end of 2020", 2020 - input)
  
}
