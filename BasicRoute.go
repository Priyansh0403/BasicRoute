package main
// These are namespace needed to execute below set of instructions
// fmt : It means formatter 
// log : used to capture the logs
// net/http or net : Used to get how webservice will be routed
import (
	"fmt"
	"log"
	"net/http"
	"net"
    "time"
)

func handler(w http.ResponseWriter, req *http.Request) {
    ip, _, err := net.SplitHostPort(req.RemoteAddr)
    if err != nil {
        fmt.Fprintf(w, "Current IP Address: %s", req.RemoteAddr)
    }

    

    // This will only be defined when site is accessed via non-anonymous proxy
    // and takes precedence over RemoteAddr
    // Header.Get is case-insensitive
    forward := req.Header.Get("X-Forwarded-For")
    t := time.Now()
    zone, _ :=t.Zone()
    if forward!=""{
		fmt.Fprintf(w,"Current IP Address:%s : Location:%s  : Timezone:%s", forward,t.Location(),zone)
        
		return
	}
	userIP := net.ParseIP(ip)
    if userIP == nil {
        fmt.Fprintf(w, "Current IP Address:%s : Location:%s  : Timezone:%s", req.RemoteAddr,t.Location(),zone)
        
        return
    }
	fmt.Fprintf(w, "Current IP Address:%s: Location:%s  : Timezone:%s", userIP,t.Location(),zone)
}


// Entry point of application being called
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))  
    
}

