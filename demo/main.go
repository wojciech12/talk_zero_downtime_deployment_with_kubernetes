package main

import (
	"net/http"
	"os"
	"sync"

	"github.com/microdevs/missy/data"
	"github.com/microdevs/missy/log"
	"github.com/microdevs/missy/service"
)

type Handler struct {
	readyMu  sync.Mutex
	healthMu sync.Mutex

	// liveness in k8s
	isHealthy bool

	// readiness in k8s
	isReady bool

	nodeName string
}

func main() {

	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	log.Println("Starting zero demo service")
	s := service.New("zero-demo")
	s.Host = "0.0.0.0"
	s.Port = "8090"

	var handler Handler
	handler.isHealthy = true
	handler.isReady = true
	handler.nodeName = name

	s.UnsafeHandleFunc("/healthz", handler.Health).Methods(http.MethodGet)
	s.UnsafeHandleFunc("/ready", handler.Readiness).Methods(http.MethodGet)

	// for simplification
	s.UnsafeHandleFunc("/doUnHealthz", handler.DoUnHealthy).Methods(http.MethodGet)
	s.UnsafeHandleFunc("/doNotReady", handler.DoNotReady).Methods(http.MethodGet)

	s.UnsafeHandleFunc("/doHealthz", handler.DoHealthy).Methods(http.MethodGet)
	s.UnsafeHandleFunc("/doReady", handler.DoReady).Methods(http.MethodGet)

	s.Start()
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	h.healthMu.Lock()
	defer h.healthMu.Unlock()
	if h.isHealthy == true {
		data.MarshalWithCode(w, r, "OK. Node: "+h.nodeName, 200)
	} else {
		data.MarshalWithCode(w, r, "Not working. Node: "+h.nodeName, 500)
	}
}

func (h *Handler) Readiness(w http.ResponseWriter, r *http.Request) {
	if h.isReady {
		data.MarshalWithCode(w, r, "OK. Node: "+h.nodeName, 200)
	} else {
		data.MarshalWithCode(w, r, "Not Ready. Node: "+h.nodeName, 500)
	}
}

func (h *Handler) DoUnHealthy(w http.ResponseWriter, r *http.Request) {
	h.makeUnHealthy()
}

func (h *Handler) makeUnHealthy() {
	h.healthMu.Lock()
	defer h.healthMu.Unlock()
	log.Printf("Switch to unhealth!")
	h.isHealthy = false
}

func (h *Handler) DoHealthy(w http.ResponseWriter, r *http.Request) {
	h.healthMu.Lock()
	defer h.healthMu.Unlock()
	h.isHealthy = true
}

func (h *Handler) DoNotReady(w http.ResponseWriter, r *http.Request) {
	h.makeNotReady()
}

func (h *Handler) makeNotReady() {
	h.readyMu.Lock()
	defer h.readyMu.Unlock()
	log.Printf("Inform other services, I am not ready to serve reaquests!")
	h.isReady = false
}

func (h *Handler) DoReady(w http.ResponseWriter, r *http.Request) {
	h.readyMu.Lock()
	defer h.readyMu.Unlock()
	h.isReady = true
}
