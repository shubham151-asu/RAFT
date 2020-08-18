package main

import
(
 "testing"
 "reflect"
  pb "../../server/gRPC"
)



func TestLogMatching(t *testing.T){
    baseLog := LogMatching(1)

    testCases := []struct{
        Name string
        ExpectedLogs []*pb.LogsResponseLogEntry
    }{
        {
            Name: "Verify LogMatching on Server2",
            ExpectedLogs : baseLog ,
        },
        {
            Name: "Verify LogMatching on Server3",
            ExpectedLogs : baseLog,
        },
    }
    for i,tc := range testCases {
        t.Log(tc.Name)
	    logs := LogMatching(i+2)
	    if reflect.DeepEqual(logs,tc.ExpectedLogs){
	        t.Log("Logs matched with Server1")
	    } else {
	        t.Errorf("Logs did not matched Server1")
	    }
    }
}

