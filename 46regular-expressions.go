/*
Go offers built-in support for regular expressions. Here are some examples of common regexp-related tasks in Go.
*/

package main

import "bytes"
import "fmt"
import "regexp"

var p=fmt.Println

func main(){

   match,_:=regexp.MatchString("p([a-z]+)ch","peach")
   p(match)
  
   r,_:=regexp.Compile("p([a-z]+)ch")


   p(r.MatchString("peach"))

   p(r.FindString("peach punch"))

   p(r.FindStringIndex("peach punch"))

   p(r.FindStringSubmatch("peach punch"))

   p(r.FindAllString("peach punch pinch",-1))

   p(r.FindAllStringSubmatchIndex("peach punch pinch",-1))
   
   p(r.FindAllString("peach punch pinch",2))

   p(r.Match([]byte("peach")))

   r= regexp.MustCompile("p([a-z]+)ch")

   p(r)

   p(r.ReplaceAllString("a peach","<fruit>"))

   in :=[]byte("a peach")

   out :=r.ReplaceAllFunc(in,bytes.ToUpper)
   p(string(out))
}
/*
	
$ go run regular-expressions.go 
true
true
peach
[0 5]
[peach ea]
[0 5 1 3]
[peach punch pinch]
[[0 5 1 3] [6 11 7 9] [12 17 13 15]]
[peach punch]
true
p([a-z]+)ch
a <fruit>
a PEACH
*/
