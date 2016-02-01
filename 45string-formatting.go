/*
Go offers excellent support for string formatting in the printf tradition. Here are some examples of common string formatting tasks.
*/

package main

import "fmt"
import "os"

type point struct{
  x,y int
}
var pf=fmt.Printf

func main(){
    p :=point{1,2}
    pf("%v\n",p)

    pf("%+v\n",p)
   
    pf("%#v\n",p)

    pf("%T\n",p)

    pf("%t\n",true)

    pf("%d\n",123)

    pf("%b\n",14)
  
    pf("%c\n",33)
  
    pf("%x\n",456)

    pf("%f\n",78.9)

    pf("%e\n",123400000.0)
    pf("%E\n",123400000.0)

    pf("%s\n","\"string\"")
   
    pf("%q\n","\"string\"")
  
    pf("%x\n","hex this")

    pf("%p\n",&p)

    pf("|%6d|%6d|\n",12,345)


    pf("|%6.2f|%6.2f|\n",1.2,3.45)

    pf("|%-6.2f|%-6.2f|\n",1.2,3.45)

    pf("|%6s|%6s|\n","foo","b")

    pf("|%-6s|%-6s|\n","foo","b")

    s:=fmt.Sprintf("a %s","string")
    fmt.Println(s)

    fmt.Fprintf(os.Stderr,"an %s\n","error")
}

/*
$ go run string-formatting.go
{1 2}
{x:1 y:2}
main.point{x:1, y:2}
main.point
true
123
1110
!
1c8
78.900000
1.234000e+08
1.234000E+08
"string"
"\"string\""
6865782074686973
0x42135100
|    12|   345|
|  1.20|  3.45|
|1.20  |3.45  |
|   foo|     b|
|foo   |b     |
a string
an error
*/
