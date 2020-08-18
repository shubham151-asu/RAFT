package main

import
(
 "strings"
 "sync"
 "log"
)



var lock sync.Mutex

func ConsistencyCheck()(string){
    log.Printf("Inside Consistency Check")
    c := Client{}
    c.InitializeCache()
    defer c.StopClientConnection()
    x := c.Get("x")
    log.Printf("Values of x %v",x.String())
    lock.Lock()
    if strings.EqualFold(x.GetResult(), "0"){
        c.Put("y","1")
    }
    lock.Unlock()
    y := c.Get("y")
    log.Printf("Values of y %v",y.String())
    lock.Lock()
    if strings.EqualFold(y.GetResult(), "0"){
        c.Put("x","1")
    }
    lock.Unlock()
    return x.GetResult() + y.GetResult()
}
