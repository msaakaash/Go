package main
import "fmt"

type Person struct{
    Name string
    Age int
}

func main() {
    p1 := Person{"Aakaash",20}
    fmt.Println(p1)
    
    fmt.Println("Name:",p1.Name)
    fmt.Println("Age:",p1.Age)
    
    p2 := Person{Age:53,Name:"Mani Suresh"}
    fmt.Println(p2)
    
}