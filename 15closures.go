package main

import "fmt"

func intSeq() func() int{
   i:=0
   return func() int{
    i+=1
    return i
  }

}
//This is function will return a functions
func main(){
   nextInt :=intSeq()
    
   fmt.Println(nextInt())
   fmt.Println(nextInt())
   fmt.Println(nextInt())

   newInts :=intSeq()
   fmt.Println(newInts())
 
}
