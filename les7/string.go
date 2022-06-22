package main

import (
    "fmt"
    "strconv"
    "strings"
)

func main() {
    s := "1 2 3 4 5"
    var t = []int32{}
    ss := strings.Fields(s)
    for _, v := range ss {
        a, _ := strconv.Atoi(v)
        t = append(t, int32(a))
    }

    if len(t) == 1 {
        fmt.Println(t[0])
    } else {
        max, min := t[0], t[0]
        for _, val := range t {
            if val > max {
                max = val
            }
            if val < min {
                min = val
            }
        }
        fmt.Println(max, min)
    }
}
