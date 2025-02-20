func main() {
  x := make([]int, 0, 21)
  x = append(x, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
  fmt.Println(x)
  x = append(x, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21)
  fmt.Println(x)
}