package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "5.3" )

func BkK7NA84Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rK2M8I46Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qa6iNT8LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0jCsJTGMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s4ScoNoDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q12NFtAfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WBdvgvJIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ycQUw8iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8wUItLJXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BoAlVz8lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XRTT1LdcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ickFh3MOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KKMqov0xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aLJVcrP0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LeSFEyKJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9YflesXHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q96ik0ihWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uAjWAo4DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PPASVKDyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yed2Hk1EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5rXkfa8qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ec8kymPdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FUPcR3omWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yOqGiRCCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s9Xz76gnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JLAChSgxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FI2u2FWbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rMSvbweSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eHvLFrulWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XdZbm4rOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OnJodLDOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Wwy8CwAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oZ8JiLhMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QWqxJYk1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pAIVoWBLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mXII81klWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GyZQ4018Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bmzfP5HeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t6C3QgZjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K30RWC5jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ABvVLZzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n7czZmAVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mqr4XNsUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hjAIV6wZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VnFNN4bQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WGzafXXlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IpvJv9hsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8mlUxFw2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xvwqTCh0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DnYnH8kwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C2c0EH7OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EsXDTrPJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TIFO07gdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6wpa73f3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J0hZLFdYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rk8yUkMHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iZa7ri1JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Omgk4D6OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TVMsZbEpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r19kicjNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KPJs8P5XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JqPKKTHUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LCjMHe0TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SUQwYmm1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B2wqQclTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dObQy5m1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dI3DVq9MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RIqWICl6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0SAZ1IVPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QIU8HgUQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mSqKKBrkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CrZHBDomWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fbCNObBbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yWt7NxxXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rxWSOiO9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gDOwOjKHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 36AGdVZTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BuNj0VPWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QG3b2CEZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tC2LG07vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func smyQ4ccRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bst56nvLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jn1b5mOIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Yypr67zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Iz5vC96xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pk5jjr3MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HPxdHHxgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uAcVpNapWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W7zNYxI5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bGhZ6JFkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gdCJ1qOiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LVjxJslrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7XxfPeozWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 74zrFIb2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9sI5RI43Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wz8WbtGsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qyYaDDN9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dnu6yrEMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OlOC3q7mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wJBH5s3zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MKHBB8sqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EiKuoamCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D6VAngVlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZHTwzq4IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lMOZyL38Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EyvIPp7IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PJyiJ5tjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yhNUuNeKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B7dEO8DeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LLIFRJbsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BUchAx1YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nlLojLgPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wzGmIc2PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YFIWxXtdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fPLhJX6JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xI1237MFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IYtYRyemWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OBLwDFA7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KFKy4U0KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QNt5OYR8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2M8vX9wGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OoM6dQZ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IrTLC7a3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cd92rN0oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ub0CnYycWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b95fkjd8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ilTjpm7JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1VxwBw3oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tdP7hRlyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OWeRHfdwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wlIj7jyaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jfcspdmbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1b4EtiyjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y7fTmNORWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wkkg7bqaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uLCSuJahWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZlXY3f5OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DURWNb22Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1mfML1SLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cDHk9lV1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1HF8de28Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G0l1SF8tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x8cwp8UPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0IE8SakNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q7duiPHQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eKa8i4I3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1jGf6OlQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jqo2cxy3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vPOqGqq0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bzbVyeMvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FRVsr9VSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZbxZ2hjKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m1gAvKFrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OhMDeSuiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pHiZbgMhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sEXHx4TjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func scXdAdZ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pVcfskI9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1yDPZEqxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s1FV3mVhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GbTSYYcwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HAYoVZkQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func op6Xt8zYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CmaBLjYoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QZ1nuPA1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j4I9vKSXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WYqIrBkrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m5VnkJFwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wrtm5iUrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RD0rzlb2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pk0XUhQNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MiutFGcXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XjgxGZ95Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0FMtk1ywWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DGATCQrpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BQOlmML4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HDRWqVCXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ipcsIq8dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cPEWEVkoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hi5zOyiqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R0Dpz6VJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rSjTiMduWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qE6SzxmPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2oKLyY3LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jXFtdHkhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 06Gea3xAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jUD31CY7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func io4yi6vmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dakrOE0PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u353ECQDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L0eAOly1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vI6kzHgDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dvF9KEMDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0xjbvNxdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SY0C3TeJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lYRuCI3LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7BbN0RFkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M3QWezQnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BtrBsCs4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VC0ErUaPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EagnOBFnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H3tDkWcbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dTGOV10IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dOcg3W0lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hjdwgeKdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cy4E6ifFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fvzJqFPoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J9GMLvoWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yOvDSVrbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jedSwsC8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QpWaZdofWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P9WzSZPVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MDtwow20Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7UC5IM0xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b9MlPbYdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Kk2s9irWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wQWpSNrFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ED7yNacWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WLLza0pRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AZxdLl9NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func peViCSGKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VdsQ0rzJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lxOIxFgPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fqt3KCRmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DQwggvf9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kWIbtrkAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xa0U2i6GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6zCUL7ATWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qmmcN6wVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0KvI9BJJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Aa3QKk7TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SS69Aon4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7yfs61daWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3yBvPQDFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SLiXnSoWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mRTeeTzuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eJRJVvtOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nlPUpnguWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DSdrFZIMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9oV2AP5oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rAszp5GuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XUeZZn7pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h2asriPkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w3GGGLJHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RhELZcLuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R3ZHjq8vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

