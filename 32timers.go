/*
We often want to execute Go code at some point in the future, or repeatedly at some interval. Go's built-in timer and ticker features make both of these tasks easy. We'll look first at timers and then at tickers.
*/

package main

import "time"
import "fmt"

func main(){

  timer1 :=time.NewTimer(time.Second *2)
  //The <-timer1.C blocks on the timer's channel C until it sends a value indicating that the timer expired.
  <- timer1.C
  fmt.Println("Timer 1 expired")

  timer2 := time.NewTimer(time.Second)
  go func(){
    <- timer2.C
    fmt.Println("Timer 2 expired")
  }()
  
  stop2 :=timer2.Stop()
  if stop2 {
     fmt.Println("Timer 2 Stopped")
  }
  //The first timer will expire ~2s after we start the program, but the second should be stopped before it has a chance to expire.
}
