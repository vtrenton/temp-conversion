package main

import (
  "fmt"
  "os"
  "strconv"
)

func main() {
  // if osargs are of length less than one
  // let the user know how to use the program andexit
  // note that arg 0 is the name of the program - so 2 is the magic num of len
  if len(os.Args) < 2 {
    fmt.Println("Welcome to the fahrenheit coversion program!")
    fmt.Println("This binary requires you run it with a single argument (only one for now)")
    fmt.Println("This argument needs to be a type float64 value (signed)")
    fmt.Println("The argument is the value you want to convert in fahrenheit")
    fmt.Println("The program will output the respective values in Celsius and Kelvin")
    fmt.Println("Please run this program with only one argument to begin")
    os.Exit(1)
  // else if os args are greater than one - use only value at arg 1 and give user a message
  } else if len(os.Args) > 2 {
      fmt.Println("One argument at a time pls")
      fmt.Printf("Only producing output for %s\n", os.Args[1])
  }

  // Attempt to convert the user provided string to a float64
  // if there is an error catch and exit 1
  f, err := strconv.ParseFloat(os.Args[1], 64)
  if err != nil {
    fmt.Println("That doesn't seem like a number to me")
    fmt.Println("I'll just take this input and convert the first char to it's decimal value :^)")
    // lol nested type conversion
    f = float64(os.Args[1][0])
    fmt.Printf("your input temp is now %.0ff\n", f)
  }

  // validate number is not less than -459.67f, -273.15c or absoute zero and not greater than a trillion degress
  if f < -459.67 {
    fmt.Println("Below absolute zero 🥶")
    os.Exit(0)
  } else if f >= 1000000000000 {
    fmt.Printf("Too hot for this program to handle! 🫠")
    os.Exit(0)
  }

  c, k := convert(f)
  
  // Print the Celsius and Kelvin results
  fmt.Printf("Celsuis: %.2fc\n", c)
  fmt.Printf("Kelvin: %.2fk\n", k)
}

// takes in f which is the user input fahrenheit value
func convert (f float64) (float64, float64) {
  c := (f - 32) * 5/9
  k := c + 273.15
  return c, k
}
