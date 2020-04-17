package main

import(


  //"github.com/urfave/cli"
  "os"
  //"io"
  "net/icmp"
  "log"
  //"strings"
  //"syscall"
 "net"

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


      for{


          response, err := icmp.ListenPacket("ipv4", userInput) //listen to responses and user address
          if err != nil{
            fmt.Println(err)
          }
          defer response.Close()


          //set up echo request to send to address
          message := icmp.Message{
                 Type: ipv4.ICMPTypeEcho, Code: 0,
                 Body: &icmp.Echo{
                     ID: os.Getpid() & 0xffff, Seq: 1,
                     Data: []byte("echo requests"),
                 },
             }
            b, err := message.Marshal(nil)
            if err != nil{
              log.Fatal(err)
            }

            start := time.Now() //timer to display latency when received
            n, err := c.WriteTo(b, dst)


      }


}
