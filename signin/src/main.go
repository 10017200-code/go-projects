//Login, signup, and logout GUI (or api, if i want this tot go further) that logs passwords and holds SHA256 encrypted passwords
//in a XML markdown file. Passwords are stored and read from XML and log is a simple ASCII file with nothing fancy.
//Creates a socckewt/ddaemon (whatever idk) that apps can call to login or check what user is logged in.
package main

import (
  "fmt"
  "crypo/sha256"
  "enccoding/hex"
  "encoding/xml"
  "os"

  //These are needed for the socket.
  "log"
  "net"
  "os/signal"
  "syscall" 
)

func main() {
  //Here, start daemon\socket that programs can call to fetch which profile is logged in.

  prog := os.Args[0]
  cmd := os.Args[1]
  switch cmd {
    case "login":
      uname, passwd, hash := ""
      //Login code here.
      //Asks for username with fmt.Scanln(&uname)
      //Asks for password (2.0 will try to not echo password) with fmt.Scanln(&passwd)
      //Encrypts passwd with sha256.Sum256() and encoding/hex is used to make it a string to input to the xml
      //Checks username first then checks passwd hash with stored hash for said user
    case "signup":
      uname, passwd, hash := ""
      // Same thing as login but instead of reading stores it in the .xml password document
      // Exits program after and tells user to log in with their credentials by running the program again

    default:
      fmt.Println("signin v1.0")
      fmt.Println("Usage: signin <option> [username] [password]")
      fmt.Println("Options:")
      fmt.Println("signin\nsignout\nsignup")
      fmt.Println("Username and password fields are exected to fully function in v2.0")
    }
}
