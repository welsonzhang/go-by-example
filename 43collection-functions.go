/*
We often need our programs to perform operations on collections of data, like selecting all items that satisfy a given predicate or mapping all items to a new collection with a custom function.
In some languages it's idiomatic to use generic data structures and algorithms. Go does not support generics; in Go it's common to provide collection functions if and when they are specifically needed for your program and data types.
Here are some example collection functions for slices of strings. You can use these examples to build your own functions. Note that in some cases it may be clearest to just inline the collection-manipulating code directly, instead of creating and calling a helper function.
*/

package main

import "strings"
import "fmt"

func Index(vs []string,t string) int{
   for i,v :=range vs{
       if v==t{
          return i
     }
   }
   return -1
}

func Include(vs []string,t string) bool{
     return Index(vs,t) >=0 
}

func Any(vs []string,f func(string) bool) bool{
  for _,v:=range vs{
       if f(v) {
           return true  
       }
  }
  return false
}

func All(vs []string,f func(string) bool) bool{
   for _,v:=range vs{
      if !f(v){
         return false
      }
   }
   return true
}

func Filter(vs []string,f func(string) bool)[]string{
     vsf :=make([]string,0)
     for _,v:=range vs{
          if f(v){
           vsf=append(vsf,v)
       }
     }
   return vsf
}

func Map(vs []string,f func(string) string)[]string{
   vsm :=make([]string,len(vs))
    for i,v :=range vs{
        vsm[i]=f(v)
    }
   return vsm
}

func main(){

   var strs=[]string{"peach","apple","pear","plum"}

   fmt.Println(Index(strs,"pear"))

   fmt.Println(Include(strs,"grape"))

   fmt.Println(Any(strs,func(v string)bool{
         return strings.HasPrefix(v,"p")
   }))

   fmt.Println(Filter(strs,func(v string)bool{
         return strings.Contains(v,"e")
   }))

   fmt.Println(Map(strs,strings.ToUpper))

}
