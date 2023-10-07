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
    fmt.Println("The arguement is the value you want to convert in fahrenheit")
    fmt.Println("The program will output the respective values in Celsius and Kelvin")
    fmt.Println("Please run this program with only one arguement to begin")
    os.Exit(1)
  } else {
    if len(os.Args) > 2 {
      fmt.Println("one arguement at a time pls")
      fmt.Printf("only producing output for %s\n", os.Args[1])
    }
    arg, err := strconv.ParseFloat(os.Args[1], 64)
    if err != nil {
      fmt.Println("ya dun goofd")
      os.Exit(1)
    }
    fmt.Printf("%.2f\n", arg)
 }

  // else if os args are greater than one - use only value at arg 1 and give user a message

  // call to conversion function
//  convert(f)
}

// takes in
//func convert (f float64) (float64, float64) {
	// validate number signed and not less than -273.15c or absoute zero
	// to save space we need to restrict any number over a million to just print 'a lot' <- the approach here might need to be thought about some
//}
