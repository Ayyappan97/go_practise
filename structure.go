package main 
  
import "fmt"
  

type Address struct { 
    Name    string 
    city    string 
    Pincode int
} 
  
func main() { 
  
    
    var a Address  
    fmt.Println(a) 
  
    
    a1 := Address{"ayyappan", "chennai", 600001} 
  
    fmt.Println("Address1: ", a1) 
  
    
    a2 := Address{Name: "Arun", city: "trichy", 
                                 Pincode: 600009} 
  
    fmt.Println("Address2: ", a2) 
  
    
    a3 := Address{Name: "raj"} 
    fmt.Println("Address3: ", a3) 
} 