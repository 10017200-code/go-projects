// Login, signup, and logout GUI (or api, if i want this tot go further) that logs passwords and holds SHA256 encrypted passwords
// in a XML markdown file. Passwords are stored and read from XML and log is a simple ASCII file with nothing fancy.
// Creates a socckewt/ddaemon (whatever idk) that apps can call to login or check what user is logged in.
package main

import (
	"crypto/sha256"
	"encoding/hex"
	//"encoding/xml" Commented out til needed
	"fmt"
	"os"

	//These are needed for the socket. Commented out til needed
	//"log"
	//"net"
	//"os/signal"
	//"syscall"
)

//TODO:
//First step, hash function that takes in a string and hashes it in SHA256 and returns a string
//Second step, write the basics for the XML writer and reader

func hash(input string) string{
  h := sha256.New()
  h.Write([]byte(input))
  out := hex.EncodeToString(h.Sum(nil))
  return out
} //WIP

func main() {
  //Here, start daemon\socket that programs can call to fetch which profile is logged in.

  //prog := os.Args[0]
  cmd := os.Args[1]
  switch cmd {
    case "login":
      //TODO:
      //Hash userids
      var userid, passwd, hashed string
      fmt.Print("userid: ")
      fmt.Scanln(&userid)
      fmt.Print("passwd: ")
      fmt.Scanln(&passwd)
      hashed = hash(passwd)
      fmt.Printf("[ DEBUG ] userid=%s passwd=%s hashed=%s\n", userid, passwd, hashed)
      //Login code here.
      //Asks for username with fmt.Scanln(&uname)
      //Asks for password (2.0 will try to not echo password) with fmt.Scanln(&passwd)
      //Encrypts passwd with sha256.Sum256() and encoding/hex is used to make it a string to input to the xml
      //Checks username first then checks passwd hash with stored hash for said user
    case "signup":
      var uname, passwd, hash string
      fmt.Println(uname, passwd, hash) //Temp
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
