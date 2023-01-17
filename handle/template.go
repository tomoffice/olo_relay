package handle

import (
	"log"
	"net/http"
	"relay/firing"
	"strings"
)

type Handler interface {
	Config(token string, targetName ...string)
	GetFunc(token string) func(w http.ResponseWriter, r *http.Request)
}
type Handle struct {
	token      string
	targetName []string
}

// set auth token to avoid stranger access | set what key want to return value
func (h *Handle) Config(token string, targetName ...string) {
	h.token = token
	h.targetName = targetName
}

// test title
func (h Handle) GetFunc(notifyToken string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		switch r.Method {
		case "POST":
			token := strings.Join(r.Header["Bearer"], "")
			if token != h.token {
				log.Println(r.Header, "some body trying our server")
				w.Write([]byte("go away stop doing this"))
				return
			}

			var recvdata firing.RecvDataer = firing.New()
			postMap, err := recvdata.GetValueMap(r, h.targetName...)
			if err != nil {
				log.Println(err)
				return
			}
			recvdata.Set(postMap)
			recvdata.CallCommand(notifyToken)
			w.Write([]byte("200 OK"))
			return
		default:
			w.Write([]byte("wrong protocol"))
			return
		}
	}

}
func New() *Handle {
	return &Handle{}
}
