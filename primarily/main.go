package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "math"
)

/**
* https://www.hackerrank.com/challenges/ctci-big-o/problem
* Verifies if n is a prime number
**/
func primality(n int32) string {
    if(n == 1) {
        return "Not prime"
    }
    limit := int32(math.Sqrt(float64(n)))  
    for i:=int32(2); i <= limit; i++ {
        if(n%i == 0) {
            return "Not prime"
        }
    }

    return "Prime"
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    pTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    p := int32(pTemp)

    for pItr := 0; pItr < int(p); pItr++ {
        nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
        checkError(err)
        n := int32(nTemp)

        result := primality(n)

        fmt.Fprintf(writer, "%s\n", result)
    }

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
