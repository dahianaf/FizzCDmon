package main

import "fmt"

func main(){
  for i:= 1; i <=100 ; i++{
   if i%3 == 0 {
     if i%5 == 0{
       fmt.Print("CDmon\n")
     }else {
       fmt.Print("CD\n")
     }
   }else if i%5 == 0{
     fmt.Printf("mon\n")
   }else{
     fmt.Printf("%d\n", i)
   }
  }
}

