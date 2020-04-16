package main

import(


  "github.com/urfave/cli"
  "os"
  //"io"
  "log"
  //"strings"
  //"syscall"
  //"fmt"

)



//https://hackernoon.com/building-a-network-command-line-interface-in-go-fd57b31df3fe
//helpful url!
func main(){

    app:= cli.NewApp()

    app.Name = "Ping CLI Application"
    app.Usage = "Enter a Hostname or IP Address, then sends and recieves echo requests and echo reply at said address."


    app.Commands = []*cli.Command{
      {
        Name: "Hostname",
        Action:
      //  Usage: "Pings will be sent and recieved at Hostname"





      },

      {
        Name: "IP Address",

      },




    }
      err := app.Run(os.Args)
      if err != nil{
        log.Fatal(err)
      }
      //fmt.Println(app)

}
