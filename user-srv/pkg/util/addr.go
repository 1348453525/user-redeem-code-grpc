package util

import (
	"fmt"
	"net"
)

func GetFreePort(ip string, port int) (int, error) {
	// 检查指定端口是否可用
	addr := fmt.Sprintf("%s:%d", ip, port)
	l, err := net.Listen("tcp", addr)

	// 指定端口被占用，随机分配
	if err != nil {
		l, err = net.Listen("tcp", fmt.Sprintf("%s:%d", ip, 0))
		if err != nil {
			return 0, err
		}
		defer l.Close()

		// fmt.Printf("%T\n", l.Addr())  // 输出：*net.TCPAddr
		if tcpAddr, ok := l.Addr().(*net.TCPAddr); ok {
			return tcpAddr.Port, nil
		}
		return 0, fmt.Errorf("failed to get port from address")
	}

	defer l.Close()
	return port, nil // 端口可用，返回原端口
}
