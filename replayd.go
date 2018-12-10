// replayd.go to read in config, e.g. listen port and then do GET and POST/PUT

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-ini/ini"
	"github.com/gorilla/mux"
)

func CfgInit() string {
	cfg, err := ini.Load("replayd.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	// Classic read of values, default section can be represented as empty string
	var AppMod string = cfg.Section("").Key("app_mode").String()
	var TmpData string = cfg.Section("paths").Key("tmp_data").String()
	// Let's do some candidate value limitation
	var HttpProtocol string = cfg.Section("server").Key("protocol").In("http", []string{"http", "https"})
	// Value read that is not in candidates will be discarded and fall back to given default value
	var SmtpProtocol string = cfg.Section("server").Key("protocol").In("smtp", []string{"imap", "smtp"})
	// Try out auto-type conversion
	var HttpPort string = cfg.Section("server").Key("http_port").String()
	var EnforceDomain = cfg.Section("server").Key("enforce_domain").MustBool(false)
	// // select new debug_level value
	var DebugLevel string = cfg.Section("debug").Key("debug_level").In("none", []string{"finest", "fine", "normal", "none"})

	//Print out all ini values
	fmt.Println("App Mode:", AppMod)
	fmt.Println("Tmp Data Path:", TmpData)
	fmt.Println("Server Protocol:", HttpProtocol)
	fmt.Println("Email Protocol:", SmtpProtocol)
	fmt.Println("Port Number:", HttpPort)
	fmt.Printf("Enforce Domain: (%[1]T) %[1]v\n", EnforceDomain)
	fmt.Println("Debug Level:", DebugLevel)

	// Now, make some changes and save it
	cfg.Section("").Key("app_mode").SetValue("production")
	cfg.SaveTo("my.ini.local")
	return HttpPort
}

// add Person struct
type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {

	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	people[0] = person
	json.NewEncoder(w).Encode(people)
}

func main() {
	var HttpPort = CfgInit()
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "Walter", Lastname: "Lee", Address: &Address{City: "San Francisco", State: "CA"}})

	//single person payload
	router.HandleFunc("/", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/", CreatePersonEndpoint).Methods("PUT")

	// form the proper listen port before ListenAndServe
	HttpPort = ":" + HttpPort
	//	fmt.Println("Port Number:", HttpPort)

	log.Fatal(http.ListenAndServe(HttpPort, router))
	//  log.Fatal(http.ListenAndServe(":8081", router))
}
