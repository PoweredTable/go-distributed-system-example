package client

import (
	"bufio"
	"fmt"
	"go-client-server/common"
	"net"
	"sync"
)

type Client struct {
	routines int
	servers  []string
}

func NewClient(routines int, servers []string) *Client {
	return &Client{
		routines: routines,
		servers:  servers,
	}
}

func (c *Client) ProcessText(text string) {
	var wg sync.WaitGroup
	wg.Add(c.routines)
	ch := make(chan float64, c.routines)

	parts := common.SplitText(text, c.routines)

	s := 0
	for i := 0; i < c.routines; i++ {

		go func(index int, part string, server string) {
			fmt.Printf("server: %v | part: %v\n", part, server)
			defer wg.Done()

			conn, err := net.Dial("tcp", server)
			if err != nil {
				fmt.Println("net dial error:", err)
				return
			}
			defer conn.Close()

			requestPacket := common.Packet{
				ID:      string(i),
				Type:    common.REQUEST,
				Message: "countText",
				Body: map[string]interface{}{
					"text": part,
				},
			}
			bytesPacket, err := requestPacket.ToBytes()
			if err != nil {
				fmt.Println("bytes err:", err)
				return
			}
			conn.Write(append(bytesPacket, '\n'))

			n, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				fmt.Println("read error:", err)
				return
			}
			packet, err := common.PacketFromBytes([]byte(n))
			if err != nil {
				fmt.Println("error converting bytes to packet: ", err)
				return
			}
			count, ok := packet.Body["count"].(float64)
			fmt.Println(count)
			if !ok {
				fmt.Println("error: count is not a float64")
				return
			}
			ch <- count

		}(i, parts[i], c.servers[s])

		s++
		if s >= (len(c.servers)) {
			s = 0
		}
	}
	wg.Wait()
	close(ch)

	var result float64

	for v := range ch {
		result += v
	}
	// apresenta o resultado final
	fmt.Println(result)

}
