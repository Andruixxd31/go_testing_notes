package developers

import "fmt"


type Developer struct {
    Name string
    Age int
}

func FilterUnique(developers []Developer) []string {
    var unq []string
    check := make(map[string]int)
    
    for _, developer := range developers{
        check[developer.Name] = 1
    }

    for name := range check{
        unq = append(unq, name)
    }

    return unq
}

func main(){
    fmt.Println("Starting Tests")
}
