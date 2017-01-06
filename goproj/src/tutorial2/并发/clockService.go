/*
一个顺序执行的时钟服务器，它会每隔一秒钟将当前时间写到客户端
*/

package main


import (
        "io"
        "log"
        "net"
        "time"
)

func handleConn(c net.Conn) {
        defer c.Close()
        for {
                _, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
                if err != nil {
                        return
                }
                time.Sleep(1 * time.Second)
        }
}
func main() {
        listener, err := net.Listen("tcp", "localhost:9999")
        if err != nil {
                log.Fatal(err)
        }
        for {
                conn, err := listener.Accept()
                if err != nil {
                        log.Print(err)
                        continue
                }
                handleConn(conn)
        }
}
