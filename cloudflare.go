package main

import(


  //"github.com/urfave/cli"
  "os"
  //"io"
//  "net/icmp"
  "log"
  //"strings"
  //"syscall"
  "time"
 "net"
"golang.org/x/net/icmp"
"golang.org/x/net/ipv4"

  "fmt"


)



//https://hackernoon.com/building-a-network-command-line-interface-in-go-fd57b31df3fe
//helpful url!
func main(){



      //ListenPacket listens for responses
      //func ListenPacket(network, address string) (*PacketConn, error)
      //ListenPacket("ip4:icmp", "192.168.0.1")
      //It should report loss and RTT times for each sent message.

      //ex


      var userInput = os.Args[1]
      //userInput will = the users host name or ip address for pint

      var count = 0
      for{
        //convert string ip address to net.IPv4 address or hostname to ip if not
        userNet, err := net.ResolveIPAddr("ip4", userInput)


        response, _ := icmp.ListenPacket("ip4:icmp", userInput) //listen to responses and user address

          if err != nil{
            fmt.Println(err)

        }
        defer response.Close()


          //set up echo request to send to address
          message := icmp.Message{
                 Type: ipv4.ICMPTypeEcho, Code: 0,
                 Body: &icmp.Echo{
                     ID: os.Getpid() & 0xffff, Seq: count,
                     Data: []byte("echo requests"),
                 },
             }
             count = count + 1
            messByte, err := message.Marshal(nil)

            if err != nil{
              log.Fatal(err)
            }



            start := time.Now() //timer to display latency when received
//userNet = net.Addr(userInput)


            fmt.Println(userNet)
            _, err = response.WriteTo(messByte, userNet) //sending the echo requests

            if err != nil {
              fmt.Println(err)
            }


            reply := make([]byte, 1500)
          	read, replier, err := response.ReadFrom(reply) //reading response
          	if err != nil {
          		  fmt.Println(err)
          	}

            stop := time.Since(start) //stopping timer to get latency

            fmt.Println("Latency for message: " + stop.String())


          	mess, err := icmp.ParseMessage(1, reply[:read])
          	if err != nil {
          		  fmt.Println(err)
	          }


          //  fmt.Println(mess.Body.Len(2))


            switch mess.Type {
            case ipv4.ICMPTypeEchoReply:
            		fmt.Println("got reflection from %v", replier)
            	default:
            		fmt.Println("got %+v; want echo reply", mess)
            	}

      }


}
