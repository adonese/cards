package main

import (
    "io/ioutil"
    "github.com/google/uuid"
    "flag"
)


var fname = flag.String("f", "data.txt", "file to write data to")
var count = flag.Int("c", 5000, "number of outputted uuids")

func generate(count int) []byte {
    var res []byte
    
    for i := 0; i <= count; i++ {
        uid := uuid.New().String()
        data := []byte(uid + "\n")
        res = append(res, data...)
    }
    return res
}

func main(){
    flag.Parse()
    res := generate(*count)
    ioutil.WriteFile(*fname, res, 0775)
}
