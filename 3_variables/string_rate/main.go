package main

import (
	"fmt"
	"time"
)

// 参考：
// https://mp.weixin.qq.com/s?__biz=MzAxMTA4Njc0OQ==&mid=2651440991&idx=1&sn=d12e95a9369e33f9920db86dc133292b&chksm=80bb19adb7cc90bb5eaa4038083348544f00a627b6132c30ebd72a6d87606b39c595cdb548a0&scene=126&sessionid=1596070333&key=f87c13d2d4a2ca88fffa8ec796ea314a244118f5c1a972791ae42ba0716bccb4689f203b43a0ac0b018f29c5c3e39f4c4f39053a75f777d0d23d1e21cd1ae18ab19ce3166c7ef149c915cd06a3898089&ascene=1&uin=MTM5NzM0ODgw&devicetype=Windows+10+x64&version=62090529&lang=zh_CN&exportkey=AS%2FZpNw3WO9M52exOqpRDgs%3D&pass_ticket=Iz65VF3yla2GMadu3KVkhEFi%2FS%2BgXt9OP0B6gthTp0E%3D

const (
	FIRST  = "WHAT THE"
	SECOND = "F*CK"
)

func main() {
	var s string
	go func() {
		i := 1
		for {
			i = 1 - i
			if i == 0 {
				s = FIRST
			} else {
				s = SECOND
			}
			time.Sleep(100)
		}
	}()

	for {
		fmt.Println(s)
		time.Sleep(100)
	}
}
