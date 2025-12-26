package upstream

import (
	"fmt"
	"sync"
	"time"

	"github.com/miekg/dns"
)

type Client struct {
	Upstreams []string
	Timeout   time.Duration
}

func NewClient(upstreams []string, timeoutMs int) *Client {
	return &Client{
		Upstreams: upstreams,
		Timeout:   time.Duration(timeoutMs) * time.Millisecond,
	}
}

func (uc *Client) Query(r *dns.Msg) (*dns.Msg, error) {
	c := &dns.Client{
		Net:     "tcp-tls",
		Timeout: uc.Timeout,
	}

	respChan := make(chan *dns.Msg, len(uc.Upstreams))
	errChan := make(chan error, len(uc.Upstreams))
	var wg sync.WaitGroup

	for _, upstream := range uc.Upstreams {
		wg.Add(1)
		go func(server string) {
			defer wg.Done()
			resp, _, err := c.Exchange(r, server)
			if err != nil {
				errChan <- err
				return
			}
			if resp != nil && resp.Rcode != dns.RcodeServerFailure {
				respChan <- resp
			}
		}(upstream)
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case resp := <-respChan:
		return resp, nil
	case <-done:
		select {
		case err := <-errChan:
			return nil, fmt.Errorf("all upstreams failed, last error: %v", err)
		default:
			return nil, fmt.Errorf("all upstreams failed without a specific error")
		}
	}
}
