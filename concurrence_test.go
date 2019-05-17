package concurrence

import "testing"
import "fmt"
import "time"


func TestThreaderStart(t *testing.T) {
    start := time.Now()

    sum := 0

    onSucess := func(o interface{}) {
        sum += o.(p).i
    }

    tdz := Threader{Threads: 4, Success: onSucess}
    tdz.Add(theWork,p{1,"One Second Delay"})
    tdz.Add(theWork,p{2,"Two Second Delay"})
    tdz.Add(theWork,p{3,"Three Second Delay"})
    tdz.Add(theWork,p{4,"Four Second Delay"})
    
    tdz.Start()

    elapsed := int(time.Since(start).Seconds())
    
    if elapsed >= sum {
        t.Errorf("Expected duration to be less then delay, Elapsed: %d Delay: %d",
            sum,
            elapsed,
        )
    }
}


func TestThreaderStartAsync(t *testing.T) {
    start := time.Now()

    sum := 0
    
    
    onSucess := func(o interface{}) {
        sum += o.(p).i
    }

    tdz := Threader{Threads: 10, Success: onSucess}
    tdz.StartAsync()
    
    for i := 1; i < 11; i++ {
        tdz.Add(theWork,p{1,fmt.Sprintf("%d", i)})
	}
    
    tdz.Finish()
   
    elapsed := int(time.Since(start).Seconds())
    
    if elapsed >= sum {
        t.Errorf("Expected duration to be less then delay, Elapsed: %d Delay: %d",
            sum,
            elapsed,
        )
    }
}

func delay(d int) {
    time.Sleep(time.Duration(d) * time.Second)
}

type p struct {
    i int
    msg string
}

func theWork(mi interface{}) interface{} {
    time.Sleep(time.Duration(mi.(p).i) * time.Second)
    return mi
}


