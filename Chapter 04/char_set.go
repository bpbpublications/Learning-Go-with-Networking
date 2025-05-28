  package main
  
  import (
    "fmt"
    "unicode/utf8"
  )
  
  func main() {
    // ASCII
    asciiStr := "Hello, World!" // ASCII encoded string
    asciiBytes := []byte(asciiStr)
    fmt.Println("ASCII String:", asciiStr)
    fmt.Println("ASCII Bytes:", asciiBytes)
    
    // ISO-8859-1 (Latin-1)
    isoStr := "Caf�" // ISO-8859-1 encoded string
    isoBytes := []byte(isoStr)
    fmt.Println("ISO-8859-1 String:", isoStr)
    fmt.Println("ISO-8859-1 Bytes:", isoBytes)
    
    // UTF-8
    utfStr := "?????" // UTF-8 encoded string
    utfBytes := []byte(utfStr)
    fmt.Println("UTF-8 String:", utfStr)
    fmt.Println("UTF-8 Bytes:", utfBytes)
    
    // String length
    fmt.Println("ASCII String Length:", len(asciiStr))
    fmt.Println("ISO-8859-1 String Length:", len(isoStr))
    fmt.Println("UTF-8 String Length:", utf8.RuneCountInString(utfStr))
  }