package main

// import (
// 	"crypto/rand"
// 	"encoding/base64"
// 	"fmt"
// 	"html/template"
// 	"log"
// 	"net/http"
// 	"net/url"
// 	"sync"
// 	"time"
// )

// // session管理设计涉及的几个因素:
// // 1.全局session管理器
// // 2.保证sessionid的唯一性
// // 3.为每个客户关联一个session
// // 4.session的存储(可以存储到文件、内存、数据库)
// // 5.session过期处理

// type Session interface {
// 	Set(key, value interface{}) error // set session value
// 	Get(key interface{}) interface{}  // get session value
// 	Delete(key interface{}) error     // delete session value
// 	SessionID() string                // back current sessionID
// }

// type Provider interface {
// 	SessionInit(sid string) (Session, error)
// 	SessionRead(sid string) (Session, error)
// 	SessionDestroy(sid string) error
// 	SessionGC(maxLifeTime int64)
// }

// // 全局session管理器
// type Manager struct {
// 	cookieName  string
// 	lock        sync.Mutex
// 	provider    Provider
// 	maxLifeTime int64
// }

// var globalSessions *Manager
// var provides = make(map[string]Provider)

// func NewManager(provideName, cookieName string, maxLifeTime int64) (*Manager, error) {
// 	provider, ok := provides[provideName]
// 	if !ok {
// 		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", provideName)
// 	}
// 	return &Manager{provider: provider, cookieName: cookieName, maxLifeTime: maxLifeTime}, nil
// }

// // Register makes a session provide available by the provided name.
// // If Register is called twice with the same name or if driver is nil,
// // it panics.
// func Register(name string, provider Provider) {
// 	if provider == nil {
// 		panic("session: Register provider is nil")
// 	}
// 	if _, dup := provides[name]; dup {
// 		panic("session: Register called twice for provider " + name)
// 	}
// 	provides[name] = provider
// }

// func (manager *Manager) sessionId() string {
// 	b := make([]byte, 32)
// 	if _, err := rand.Read(b); err != nil {
// 		return ""
// 	}
// 	return base64.URLEncoding.EncodeToString(b)
// }

// func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
// 	manager.lock.Lock()
// 	defer manager.lock.Unlock()
// 	cookie, err := r.Cookie(manager.cookieName)
// 	if err != nil || cookie.Value == "" {
// 		sid := manager.sessionId()
// 		session, _ = manager.provider.SessionInit(sid)
// 		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.maxLifeTime)}
// 		http.SetCookie(w, &cookie)
// 	} else {
// 		sid, _ := url.QueryUnescape(cookie.Value)
// 		session, _ = manager.provider.SessionRead(sid)
// 	}
// 	return
// }

// //Destroy sessionid
// func (manager *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
// 	cookie, err := r.Cookie(manager.cookieName)
// 	if err != nil || cookie.Value == "" {
// 		return
// 	} else {
// 		manager.lock.Lock()
// 		defer manager.lock.Unlock()
// 		manager.provider.SessionDestroy(cookie.Value)
// 		expiration := time.Now()
// 		cookie := http.Cookie{Name: manager.cookieName, Path: "/", HttpOnly: true, Expires: expiration, MaxAge: -1}
// 		http.SetCookie(w, &cookie)
// 	}
// }

// func login(w http.ResponseWriter, r *http.Request) {
// 	sess := globalSessions.SessionStart(w, r)
// 	r.ParseForm()
// 	if r.Method == "GET" {
// 		t, _ := template.ParseFiles("login.gtpl")
// 		w.Header().Set("Content-Type", "text/html")
// 		t.Execute(w, sess.Get("username"))
// 	} else {
// 		sess.Set("username", r.Form["username"])
// 		http.Redirect(w, r, "/", 302)
// 	}
// }

// func count(w http.ResponseWriter, r *http.Request) {
// 	sess := globalSessions.SessionStart(w, r)
// 	createtime := sess.Get("createtime")
// 	if createtime == nil {
// 		sess.Set("createtime", time.Now().Unix())
// 	} else if (createtime.(int64) + 360) < (time.Now().Unix()) {
// 		globalSessions.SessionDestroy(w, r)
// 		sess = globalSessions.SessionStart(w, r)
// 	}
// 	ct := sess.Get("countnum")
// 	if ct == nil {
// 		sess.Set("countnum", 1)
// 	} else {
// 		sess.Set("countnum", (ct.(int) + 1))
// 	}
// 	t, _ := template.ParseFiles("count.gtpl")
// 	w.Header().Set("Content-Type", "text/html")
// 	t.Execute(w, sess.Get("countnum"))
// }

// func (manager *Manager) GC() {
// 	manager.lock.Lock()
// 	defer manager.lock.Unlock()
// 	manager.provider.SessionGC(manager.maxLifeTime)
// 	time.AfterFunc(time.Duration(manager.maxLifeTime), func() { manager.GC() })
// }

// func init() {
// 	globalSessions, _ = NewManager("memory", "gosessionid", 3600)
// 	go globalSessions.GC()
// }

// func mysession() {
// 	http.HandleFunc("/login", login)         // 设置访问的路由
// 	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
// 	if err != nil {
// 		log.Fatal("ListenAndServer:", err)
// 	}
// }
