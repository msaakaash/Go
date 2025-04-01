package main
import (
    "fmt"
    "sort"
    )

type Students struct {
    Name string
    Age int
    id int
}

func removeById(st []Students,id int) []Students{
    newS := []Students{}
    for _,s := range st {
        if s.id != id {
            newS = append(newS,s)
        }
    }
    return newS
}

func add(st [] Students,s Students) [] Students{
    return append(st,s)
}

func main() {
    students := []Students{
       {"Freaky Liza",20,303},
       {"Aakaash",20,301},
       {Age:20,Name:"Aiyyappan",id:302},
    }
   fmt.Println(students)
   sort.Slice(students,func (i,j int) bool{
       return students[i].id<students[j].id
   })
   fmt.Println("Sorted by id:")
   
   for _,s := range students{
       fmt.Println(s)
   }
   // deleting an element by id
   students = removeById(students,301)
   
	fmt.Println("After removing ID 301:")
	fmt.Println(students)
	
	//adding an element
	newRecord := Students{"Aakaash",20,301}
	students = add(students,newRecord)
   fmt.Println(students)
  
}