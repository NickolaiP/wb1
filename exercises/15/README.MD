var justString string

func someFunc() {
    v := createHugeString(1 << 10)
    buf := make([]byte, 100)
    copy(buf, v)
    justString = string(buf)
}

func main() {
    someFunc()
}
