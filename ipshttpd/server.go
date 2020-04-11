// ipshttpd 是ip search httpd服务器，支持通过http查询ip相关信息
package ipshttpd

import (
	"fmt"
	"github.com/tkstorm/ip-search/ipsearch"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
)

func Main() {
	registeRoute()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		select {
		case <-c:
			log.Println("ipshttpd is shutdown")
		}
		os.Exit(1)
	}()
	//log.Printf("ipshttpd listen on http://%s, %s", listen, version2.Version())
	//log.Fatalln(http.ListenAndServe(listen, nil))
}

// 帮助信息
var usageFormat = `Version %s
Usage:
	//search current client ip information
	curl localhost:8680/ips

	//search for target ip information  
	curl localhost:8680/ips?ip=targetIp	
`

func helpMessage(w http.ResponseWriter, r *http.Request) {
	helpMsg := fmt.Sprintf(usageFormat, version2.Version())
	_, _ = fmt.Fprintf(w, helpMsg)
}

func ipsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RequestURI)
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	// request search ip pick
	ip := r.FormValue("ip")
	if ip == "" {
		// host ip
		hostIp, _, err := net.SplitHostPort(r.RemoteAddr)
		if err == nil && hostIp != "127.0.0.1" {
			ip = hostIp
		}
		// by proxy request
		realIp := r.Header.Get("X-Real-IP")
		if realIp != "" {
			ip = realIp
		}
		ip = "myip"
	}
	// ip search
	ips := ipsearch.NewIps(false, r.FormValue("proxy"), 0)
	ipsRs, err := ips.Search(ip)
	if err != nil {
		http.Error(w, fmt.Sprintf("ip serach error: %s", err), http.StatusInternalServerError)
	}

	// out by json format
	msg, err := ipsRs.Render("json")
	if err != nil {
		http.Error(w, fmt.Sprintf("ip serach message show error: %s", err), http.StatusInternalServerError)
	}

	log.Println(msg)
	_, _ = fmt.Fprint(w, msg)
}