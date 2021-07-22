package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/miekg/dns"
)

type dnsServer struct {
}

var chaos *rand.Rand

func chaosRanged(bottom, top int32) int32 {
	return bottom + (chaos.Int31() % top)
}

func chaosDo(chances int) bool {
	return chaos.Int()%chances == 0
}

func (srv *dnsServer) ServeDNS(rw dns.ResponseWriter, r *dns.Msg) {
	for _, q := range r.Question {
		if q.Qtype == dns.TypeA {
			fmt.Printf("A QUERY=%+v\n", q)
			jitter := chaosRanged(1, 5000)
			if jitter >= 1000 {
				fmt.Printf("jitter=%d\n", jitter)
			}
			td := time.NewTimer(time.Millisecond * time.Duration(jitter))
			r := &dns.Msg{
				MsgHdr: dns.MsgHdr{
					Id:            r.Id,
					Response:      true,
					Authoritative: true,
				},
			}
			r.Question = append(r.Question, q)
			<-td.C
			r.Answer = append(r.Answer, &dns.A{
				Hdr: dns.RR_Header{
					Name:   q.Name,
					Rrtype: dns.TypeA,
					Class:  dns.ClassINET,
					Ttl:    3600,
				},
				A: net.ParseIP("127.0.0.1"),
			})
			data, _ := r.Pack()
			if chaosDo(10) == true {
				fmt.Println("truncated packet")
				data = data[0:chaosRanged(1, int32(len(data)-1))]
			}
			if chaosDo(10) == true {
				fmt.Println("dropped packet")
			} else {
				rw.Write(data)
			}
		}
	}
}

func main() {
	chaos = rand.New(rand.NewSource(time.Now().Unix()))
	srv := &dnsServer{}
	dns.ListenAndServe("0.0.0.0:53", "tcp", srv)
}
