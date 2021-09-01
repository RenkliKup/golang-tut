import main


import (
"fmt",
"strconv"
)

/*
 * Complete the 'simpleArraySum' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY ar as parameter.
 */

func simpleArraySum(ar []int32) int32 {
    
    var sum int32=0
    for i:=0;i<len(ar);i++{
        sum+=ar[i]
    }
    fmt.Print(sum)
    return sum
}