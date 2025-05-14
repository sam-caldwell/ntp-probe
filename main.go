// file: main.go
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/beevik/ntp"
)

// queryNTP queries the NTP server at the given target and port
// and returns the NTP response or an error.
func queryNTP(target string, port int) (*ntp.Response, error) {
	host := fmt.Sprintf("%s:%d", target, port)
	resp, err := ntp.Query(host)
	if err != nil {
		return nil, err
	}
	if err := resp.Validate(); err != nil {
		return nil, err
	}
	return resp, nil
}

// main parses command-line flags and queries the specified NTP server.
func main() {
	target := flag.String("target", "", "NTP server IP address or hostname")
	port := flag.Int("port", 123, "NTP server port")
	flag.Parse()

	if *target == "" {
		log.Fatal("error: -target is required")
	}

	resp, err := queryNTP(*target, *port)
	if err != nil {
		log.Fatalf("failed to query NTP server: %v", err)
	}

	fmt.Printf("Time:           %v\n", resp.Time)
	fmt.Printf("Clock Offset:   %v\n", resp.ClockOffset)
	fmt.Printf("Precision:      %v\n", resp.Precision)
	fmt.Printf("Root Delay:     %v\n", resp.RootDelay)
	fmt.Printf("Root Dispersion:%v\n", resp.RootDispersion)
	fmt.Printf("Stratum:        %d\n", resp.Stratum)
}
