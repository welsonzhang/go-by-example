/*
We can use channels to synchronize execution across goroutines. Here is an example of using a blocking receive to wait for a goroutine to finish.
*/
package main

import "fmt"
import "time"

func worker(done chan bool){
   fmt.Print("working...")
   time.Sleep(time.Second)
   fmt.Println("done")
   
   done <- true
}

func main(){

  done :=make(chan bool,1)
  go worker(done)

  <-done

}
