package main

import (
  "fmt"
)

func main() {
  // fmt.Println("Hello world")
  text := "Hello world"

  letter_map := map_letters(text)

  has_del, del := find_delimeter(letter_map)

  if has_del {
    fmt.Printf("This %s\n", del)
  } else {
    fmt.Println("Sorry, no delimeter could be found.")
    return
  }




}

func map_letters(text string) []int {
  letter_map := make([]int, 256)
  for i := 0; i < len(text); i++ {
    letter_map[text[i]] = 1
  }
  return letter_map
}

func find_delimeter(used []int) (bool, string) {
  for i := 0; i < 256; i++ {
    val := used[i]
    if val == 0 { return true, string(i) }
  }
  return false, ""
}
