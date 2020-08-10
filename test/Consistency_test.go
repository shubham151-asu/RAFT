package main

import
(
 "testing"
 "log"
)



func TestConsistency(t *testing.T){
    client := Client{}
    client.InitializeCache()
    defer client.StopClientConnection()
    response := client.Put("x","0")
    if response==nil{
        log.Printf("Test : ConsistencyCheck : Unable to set x")
    }
    response = client.Put("y","0")
    if response==nil{
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
	log.Printf("Starting threads %v",tc.Name)
        t.Run(tc.Name,func (t *testing.T){
            t.Parallel()
            
            res := ConsistencyCheck()
            if res=="11"{
                t.Errorf("ConsistencyCheck : found an unexpected value %v ",res)
            } else{
                t.Log("ConsistencyCheck : consistency check passed")
            }
        })
    }
}

