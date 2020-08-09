package main

import
(
 "testing"
 "strings"
 "sync"
 "log"
)

var lock sync.Mutex


func (c *Client) ConsistencyCheck()(string){
    x := c.Get("x")
    lock.Lock()
    if !strings.EqualFold(x.GetResult(), "0"){
        c.Put("y","1")
    }
    lock.Unlock()
    y := c.Get("y")
    lock.Lock()
    if !strings.EqualFold(x.GetResult(), "0"){
        c.Put("x","1")
    }
    lock.Unlock()
    return x.GetResult() + y.GetResult()
}

func ConsistencyTest(t *testing.T){
    client := Client{}
    client.InitializeCache()
    defer client.StopClientConnection()
    response := client.Put("x","0")
    if response!=nil{
        log.Printf("Test : ConsistencyCheck : Unable to set x")
    }
    response = client.Put("y","0")
    if response!=nil{
        log.Printf("Test : ConsistencyCheck : Unable to set y")
    }

    testCases := []struct{
        Name string
        UnExpected string
    }{
        {
            Name:  "Test Consistency Thread T1",
            UnExpected : "11",
        },
        {
            Name:  "Test Consistency Thread T2",
            UnExpected : "11",
        },
    }
    for _,tc := range testCases {
        tc := tc
        t.Run(tc.Name,func (t *testing.T){
            t.Parallel()
            res := client.ConsistencyCheck()
            if res=="11"{
                t.Errorf("ConsistencyCheck : found an unexpected value %v ",res)
            } else{
                t.Log("ConsistencyCheck : consistency check passed")
            }
        })
    }
}

