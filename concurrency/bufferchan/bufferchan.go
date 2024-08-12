package bufferchan

import (
    "fmt"
)

func send(ch chan string, message string) {
    ch <- message
}

func RunBufferChan() {
    size := 2
    ch := make(chan string, size)
    send(ch, "one")
    send(ch, "two")
    go send(ch, "three")
    go send(ch, "four")
    fmt.Println("All data sent to the channel ...")

    for i := 0; i < 4; i++ {
        fmt.Println(<-ch)
    }

    fmt.Println("Done!")
}