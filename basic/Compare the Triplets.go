package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'compareTriplets' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func compareTriplets(a []int32, b []int32) []int32 {
    var bob=0
    var alice=0
    for i:=0;i<len(a);i++{
        if a[i]>b[i]{
            alice+=1
        }else if b[i]>a[i]{
            bob+=1
        }else{
            continue
        }
    }