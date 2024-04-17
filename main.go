package main

import (
  "fmt"
  "log"
  "os"
  "time"

  "github.com/hashicorp/memberlist"
)

func main() {
  /* Create the initial memberlist from a safe configuration.
     Please reference the godoc for other default config types.
     http://godoc.org/github.com/hashicorp/memberlist#Config
  */
  //list, err := memberlist.Create(memberlist.DefaultLANConfig())

  config := getConfig("192.168.1.222")
  list, err := memberlist.Create(config)
  if err != nil {
    panic("Failed to create memberlist: " + err.Error())
  }

  // Join an existing cluster by specifying at least one known member.
  /*
  _, err = list.Join([]string{"192.168.1.223:7946"})
  if err != nil {
    panic("Failed to join cluster: " + err.Error())
  }
  */

  for true {
    // Ask for members of the cluster
    for _, member := range list.Members() {
      fmt.Printf("Member [%s] with IP [%s]\n", member.Name, member.Addr)
    }

    time.Sleep(10 * time.Second)
  }

  // Continue doing whatever you need, memberlist will maintain membership
  // information in the background. Delegates can be used for receiving
  // events when members join or leave.
}

func getConfig(bindAddr string) *memberlist.Config {
  config := memberlist.DefaultLANConfig()
  if len(bindAddr) == 0 {
    bindAddr = "0.0.0.0"
  }
	config.BindAddr = bindAddr
	config.Name = fmt.Sprintf("Peer %s", bindAddr)
	config.BindPort = 0 // choose free port
	//config.RequireNodeNames = true
  //config.PushPullInterval = 5 * time.Second
	config.Logger = log.New(os.Stderr, config.Name, log.LstdFlags)

	return config
}
