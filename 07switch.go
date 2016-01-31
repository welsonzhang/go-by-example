package main

import "fmt"
import "time"

func main(){
   
   i:=2
   fmt.Println("write",i,"as")
   switch i{
   case 1:
      fmt.Println("one")
   case 2:
      fmt.Println("two")
   case 3:
      fmt.Println("three")
   }
   //t := time.Now().Weekday()
   //fmt.Println(t)

   switch time.Now().Weekday(){
   case time.Saturday,time.Sunday:
       fmt.Println("It's the weekend")
   default:
       fmt.Println("It's the weekday")    
   }
 
   t :=time.Now().Hour()
   fmt.Println(t)
   switch  {
   case t <12:
        fmt.Println("It's before noon")
   default:
        fmt.Println("It's after noon")
   }

}

