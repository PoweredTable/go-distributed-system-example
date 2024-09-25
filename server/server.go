package server

import (
	"bufio"
	"fmt"
	"net"
	"go-client-server/common"
)

type TCPServer struct {
	listenAddr string
	ln         net.Listener
	quit_ch    chan struct{}
}

func (s *TCPServer) acceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("accept error: ", err)
			continue
		}
		fmt.Println("new conn: ", conn.RemoteAddr())
		go s.handleConnection(conn)
	}
}
 
func (s *TCPServer) handleConnection(conn net.Conn) {
	defer conn.Close()
	n, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("read error:", err)
		return
	}
	// Convert the string to bytes and use FromBytes
	packet, err := common.PacketFromBytes([]byte(n))
	if err != nil {
		fmt.Println("error converting bytes to packet: ", err)
		return
	}
	fmt.Println(packet)
	handlePacket(conn, *packet)
}

func handlePacket(conn net.Conn, packet common.Packet) {
	// Dispatch based on the packet message
	switch packet.Message {
	case "countText":
		// Extract the text field from the packet body
		text, ok := packet.Body["text"].(string)
		if !ok {
			fmt.Println("Invalid body format for countWord")
			return
		}
		wordCount := common.CountWords(text)
		fmt.Printf("Value: %v, Type: %T\n", wordCount,wordCount)

		// Send a RESPONSE back to the client with the word count
		responsePacket := common.Packet{
			ID:      packet.ID, // Use the same ID as the request
			Type:    common.RESPONSE,
			Message: "textCount",
			Body: map[string]interface{}{
				"count": wordCount,
			},
		}
		err := sendPacket(conn, responsePacket)
		if err != nil {
			fmt.Println("Error sending response packet:", err)
		}
	default:
		fmt.Println("Unknown request message:", packet.Message)
	}
}

func sendPacket(conn net.Conn, packet common.Packet) error {
	packetBytes, err := packet.ToBytes()
	if err != nil {
		return fmt.Errorf("error converting packet to bytes: %w", err)
	}
	_, err = conn.Write(append(packetBytes, '\n')) // append newline for server's ReadString
	return err
}

func (s *TCPServer) Start() error {
	fmt.Printf("Connecting to %s ... ", s.listenAddr)
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		fmt.Println("error!")
		return err
	}
	defer ln.Close()
	s.ln = ln
	fmt.Println("success!")

	go s.acceptLoop()

	<-s.quit_ch
	return nil
}

func NewTCPServer(listenAddr string) *TCPServer {
	return &TCPServer{
		listenAddr: listenAddr,
		quit_ch:    make(chan struct{}),
	}
}
