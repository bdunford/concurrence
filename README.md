# concurrence
Helping to do many things at the same time.

### Installing
```
$ go get github.com/bdunford/concurrence
```

### Usage

```
package main

import (
  "time"
  "fmt"
  "github.com/bdunford/concurrence"
)

func main() {
  ThreaderStart()
  ThreaderStartAsync()
}

func ThreaderStart() {

    start := time.Now()
    sum := 0
    //Function accepting dynamic parms for receiving response when work is complete.
    onSucess := func(o interface{}) {
        sum += o.(p).i
        fmt.Printf("%s\n",o.(p).msg)
    }

    tdz := concurrence.Threader{Threads: 4, Success: onSucess}
    tdz.Add(theWork,p{4,"Four Second Delay"})
    tdz.Add(theWork,p{3,"Three Second Delay"})
    tdz.Add(theWork,p{2,"Two Second Delay"})
    tdz.Add(theWork,p{1,"One Second Delay"})
    tdz.Start()

    elapsed := int(time.Since(start).Seconds())
    fmt.Printf("Calculated Sum: %d\n",sum)
    fmt.Printf("Time Elapsed: %d\n",elapsed)
}

func ThreaderStartAsync() {
    start := time.Now()
    sum := 0
    //Function accepting dynamic parms for receiving response when work is complete.
    onSucess := func(o interface{}) {
        sum += o.(p).i
        fmt.Printf("%s\n",o.(p).msg)
    }
    //Instantiate Threader
    tdz := concurrence.Threader{Threads: 10, Success: onSucess}
    //Started and Waiting for work.
    tdz.StartAsync()

    for i := 1; i < 10; i++ {
        //Adding tasks or work to the Threader
        tdz.Add(theWork,p{1,fmt.Sprintf("Thread ID: %d One second delay", i)})
	}

    tdz.Finish()

    elapsed := int(time.Since(start).Seconds())
    fmt.Printf("Calculated Sum: %d\n",sum)
    fmt.Printf("Time Elapsed: %d\n",elapsed)
}

//Struct for containing input Parameters accepted by the Work Function
type p struct {
    i int
    msg string
}

//Function for preforming the work. Receives params dynamicly as interface{}
func theWork(mi interface{}) interface{} {
    time.Sleep(time.Duration(mi.(p).i) * time.Second)
    return mi
}

```
