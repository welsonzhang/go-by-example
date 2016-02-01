/*
A common requirement in programs is getting the number of seconds, milliseconds, or nanoseconds since the Unix epoch. Here's how to do it in Go.
*/
package main

import "fmt"
import "time"

func main(){
   
  now :=time.Now()
  secs :=now.Unix()
  nanos :=now.UnixNano()
  fmt.Println(now)

  millis :=nanos / 1000000
  fmt.Println(secs)
  fmt.Println(millis)
  fmt.Println(nanos)
  
  fmt.Println(time.Unix(secs,0))
  fmt.Println(time.Unix(0,nanos))
}
/*
2016-02-01 18:50:28.530169971 +0800 HKT
1454323828
1454323828530
1454323828530169971
2016-02-01 18:50:28 +0800 HKT
2016-02-01 18:50:28.530169971 +0800 HKT
*/
