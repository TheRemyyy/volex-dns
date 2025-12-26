package dnsserver

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/TheRemyyy/volex-dns/internal/cache"
	"github.com/TheRemyyy/volex-dns/internal/config"
	"github.com/TheRemyyy/volex-dns/internal/upstream"
	"github.com/miekg/dns"
)

type Server struct {
	Config   *config.Config
	Cache    *cache.MemoryCache
	Upstream *upstream.Client
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		Config:   cfg,
		Cache:    cache.NewMemoryCache(),
		Upstream: upstream.NewClient(cfg.Upstreams, cfg.UpstreamTimeoutMs),
	}
}

func (s *Server) Start() error {
	dns.HandleFunc(".", s.handleDNS)
	server := &dns.Server{Addr: s.Config.BindAddr, Net: "udp"}
	log.Printf("Starting Volex DNS on %s", s.Config.BindAddr)
	return server.ListenAndServe()
}

func (s *Server) handleDNS(w dns.ResponseWriter, r *dns.Msg) {
	q := r.Question[0]

	// Self-PTR check
	if q.Qtype == dns.TypePTR && q.Name == "1.0.0.127.in-addr.arpa." {
		log.Printf("[SELF-PTR] Responding with server name for 127.0.0.1")
		msg := new(dns.Msg)
		msg.SetReply(r)
		msg.Authoritative = true
		rr, err := dns.NewRR(fmt.Sprintf("%s PTR %s.", q.Name, s.Config.ServerName))
		if err == nil {
			msg.Answer = append(msg.Answer, rr)
			w.WriteMsg(msg)
			return
		}
	}

	start := time.Now()
	cacheKey := fmt.Sprintf("%s:%d:%d", q.Name, q.Qtype, q.Qclass)

	// Check Cache
	if entry, found := s.Cache.Get(cacheKey); found {
		log.Printf("[CACHE HIT] %s", q.Name)
		cachedMsg := *entry
		cachedMsg.Id = r.Id
		w.WriteMsg(&cachedMsg)
		return
	}
	log.Printf("[CACHE MISS] %s", q.Name)

	r.SetEdns0(4096, true)

	// Query Upstream
	resp, err := s.Upstream.Query(r)
	if err != nil {
		log.Printf("Error querying upstreams for %s: %v", q.Name, err)
		dns.HandleFailed(w, r)
		return
	}

	// Cache Result
	if resp.Rcode == dns.RcodeSuccess && (len(resp.Answer) > 0 || len(resp.Ns) > 0) {
		ttl := getMinTTL(resp)
		if ttl < s.Config.CacheMinTtlSec {
			ttl = s.Config.CacheMinTtlSec
		}
		if ttl > s.Config.CacheMaxTtlSec {
			ttl = s.Config.CacheMaxTtlSec
		}

		s.Cache.Set(cacheKey, resp, ttl)
		log.Printf("[CACHED] %s with TTL %d sec", q.Name, ttl)
	}

	resp.Id = r.Id
	w.WriteMsg(resp)

	duration := time.Since(start)
	log.Printf("[RESOLVED] %s in %.3f ms", q.Name, float64(duration.Nanoseconds())/1e6)
}

func getMinTTL(msg *dns.Msg) uint32 {
	var minTTL uint32 = math.MaxUint32
	if len(msg.Answer) == 0 && len(msg.Ns) > 0 {
		for _, rr := range msg.Ns {
			if soa, ok := rr.(*dns.SOA); ok {
				return soa.Header().Ttl
			}
		}
	}
	for _, rr := range msg.Answer {
		if rr.Header().Ttl < minTTL {
			minTTL = rr.Header().Ttl
		}
	}
	return minTTL
}
