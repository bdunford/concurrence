package concurrence

import ( 
    "sync"
    "time"
)


type Action struct {
    fn func(interface{})(interface{})
    parms interface{}
}

//TODO Success to accept Return + Error to be all go like. 
type Threader struct {
    Threads int
    Success func(interface{})
    work []Action
    mutex *sync.Mutex
    waitGroup sync.WaitGroup
    waitForWork bool
}


func (t *Threader) nextWorkItem() *Action {
    
    t.mutex.Lock()
    if len(t.work) > 0 {
        var w Action
        w, t.work = t.work[len(t.work)-1], t.work[:len(t.work)-1]
        t.mutex.Unlock()
        return &w
    }
    
    t.mutex.Unlock()
    if t.waitForWork  {
        time.Sleep(time.Duration(200) * time.Millisecond)
        return t.nextWorkItem()
    }

    return nil

}

func (t *Threader) doWork() {
    for {
        w := t.nextWorkItem()
        if w == nil {
            t.waitGroup.Done()
            break
        }
        
        o := w.fn(w.parms)

        t.mutex.Lock() 
        t.Success(o)
        t.mutex.Unlock()
    }

}

func (t *Threader) Add(fn func(interface{})(interface{}), parms interface{}) {
    t.AddAction(Action{fn,parms})
}

func (t *Threader) AddAction(a Action) {
    t.work = append(t.work,a)
}

func (t *Threader) Start() {
    
    t.mutex = &sync.Mutex{}

    for i := 1; i <= t.Threads; i++ {
        t.waitGroup.Add(1)
        go t.doWork()
    }

    if !t.waitForWork { 
        t.waitGroup.Wait()
    }
}

func (t *Threader) StartAsync() {
    t.waitForWork = true
    t.Start()
}

func (t *Threader) Finish() {
    t.waitForWork = false
    t.waitGroup.Wait()
}

