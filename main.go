package main

import (
  "flag"
  "fmt"
  "net"
  "sync"
  "time"
)

func main() {
  start := time.Now()
  ip := flag.String("ip", "", "ip")
  startPort := flag.Int("start-port", 1, "start port")
  endPort := flag.Int("eend-port", 65535, "end port")
  flag.Parse()

  scanPort(start, ip, startPort, endPort)

}

func scanPort(start time.Time, ip *string, startPort *int, endPort *int) {

  //ip := ""
  //startPort := 1
  //endPort := 65535
  var wg sync.WaitGroup
  for i := *startPort; i < *endPort+1; i++ {
    wg.Add(1)
    go func(ip string, port int) {
      defer wg.Done()
      address := fmt.Sprintf("%s:%d", ip, port)
      conn, err := net.Dial("tcp", address)
      if err != nil {
        //直接返回
        return
      }
      fmt.Printf("%d open\n", port)
      conn.Close()
    }(*ip, i)
  }
  //等待程序程序退出
  wg.Wait()
  // 计算时间
  elapsed := time.Since(start) / 1e9
  fmt.Printf("%d seconds\n", elapsed)
}
