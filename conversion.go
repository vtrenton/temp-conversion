package main

import (
  "fmt"
  "os"
//  "math"
  "strconv"
)

func main() {
  //DEBUG
  //fmt.Println(len(os.Args))

  // if osargs are of length less than 1 prompt the user
  // note that arg 0 is the name of the program - so 2 is the magic num of len
  if len(os.Args) < 2 {
    fmt.Println("Welcome to the fahrenheit coversion program!")
    fmt.Println("This binary requires you run it with a single arguement (only one for now)")
    fmt.Println("This argument needs to be a type float64 value (signed)")
    fmt.Println("The argument is the value you want to convert in fahrenheit")
    fmt.Println("The program will output the respective values in Celsius and Kelvin")
    fmt.Println("Please run this program with only one arguement to begin")
    os.Exit(1)
  // else if os args are greater than one - use only value at arg 1 and give user a message
  } else if len(os.Args) > 2 {
      fmt.Println("one argument at a time pls")
      fmt.Printf("only producing output for %s\n", os.Args[1])
  }

  // Attempt to convert the user provided string to a float64
  // if there is an error catch and exit 1
  f, err := strconv.ParseFloat(os.Args[1], 64)
  if err != nil {
    fmt.Println("That doesn't seem like a number to me")
    os.Exit(1)
  }

  // validate number is not less than -459.67f, -273.15c or absoute zero and not greater than a trillion degress
  if f <= -459.68 || f >= 1000000000000 {
    fmt.Printf("Out of range")
  }

  // DEBUG
  //fmt.Printf("%.2f\n", f)
  
  c, k := convert(f)
  
  // give us some output
  fmt.Printf("Celsuis: %.2fc\n", c)
  fmt.Printf("Kelvin: %.2fc\n", k)
}

// takes in
func convert (f float64) (float64, float64) {
  c := (f - 32) * 5/9
  k := c + 273.15
  return c, k
}
