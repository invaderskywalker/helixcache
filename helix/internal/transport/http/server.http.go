package http

import (
	"helix/helix/internal/cache"
)

type Transport struct {
	port  string
	cache cache.ICache
}

func CreateTransport(port string, cache cache.ICache) *Transport {
	return &Transport{port: port, cache: cache}
}

// func (t *Transport) Init() {
// 	http.HandleFunc("/set", t.setHandler)
// 	http.HandleFunc("/get", t.getHandler)
// 	http.HandleFunc("/del", t.delHandler)

// 	http.ListenAndServe(t.port, nil)
// }

// func (t *Transport) setHandler(w http.ResponseWriter, r *http.Request) {
// 	key := r.URL.Query().Get("key")
// 	value := r.URL.Query().Get("value")

// 	t.cache.Set(key, []byte(value)) // using interface only
// 	w.Write([]byte("OK"))
// }

// func (t *Transport) getHandler(w http.ResponseWriter, r *http.Request) {
// 	key := r.URL.Query().Get("key")
// 	value, found := t.cache.Get(key)
// 	if !found {
// 		w.Write([]byte("NOT_FOUND"))
// 		return
// 	}
// 	w.Write(value)
// }

// func (t *Transport) delHandler(w http.ResponseWriter, r *http.Request) {
// 	key := r.URL.Query().Get("key")
// 	t.cache.Delete(key)
// 	w.Write([]byte("OK"))
// }
