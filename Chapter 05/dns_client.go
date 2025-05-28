package main

import (
    "fmt"
    "github.com/miekg/dns"
)

func main() {
    // Specify the DNS server to query
    server := "8.8.8.8:53"

    // Domain name to resolve
    domain := "www.example.com"

    // Create a DNS client
    client := new(dns.Client)

    // Create a DNS request message
    msg := new(dns.Msg)
    msg.SetQuestion(domain+".", dns.TypeA)

    // Send the DNS query to the specified DNS server
    response, _, err := client.Exchange(msg, server)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Parse and print the IP addresses associated with the domain
    for _, answer := range response.Answer {
        if a, ok := answer.(*dns.A); ok {
            fmt.Printf("IP Address: %s\n", a.A.String())
        }
    }
}
