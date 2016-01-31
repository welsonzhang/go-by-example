/*
Closing a channel indicates that no more values will be sent on it. This can be useful to communicate completion to the channel's receivers.
*/

package main

import "fmt"

func main(){
  jobs :=make(chan int,5)
  done :=make(chan bool)

 go func(){
  for{
    j,more := <- jobs
    if more {
       fmt.Println("receieved job",j)    
    }else{
       fmt.Println("receieved all jobs")
       done <- true
       return
    }
   }
 }()

  for j:=1; j<=3;j++{
      jobs <-j
      fmt.Println("send job",j)
  }
  close(jobs)  //This will let more is not true ,then the goroutine will quit
  fmt.Println("sent all jobs")

}
