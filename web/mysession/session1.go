package main

import (
	"container/list"
	"html/template"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/astaxie/session"
	_ "github.com/astaxie/session/providers/memory"
)

var pder = &Provider{list: list.New()}

type SessionStore struct {
	sid          string                      // session id
	timeAccessed time.Time                   // 最后访问时间
	value        map[interface{}]interface{} // session里面存储的值
}

func (st *SessionStore) Set(key, value interface{}) error {
	st.value[key] = value
	pder.SessionUpdate(st.sid)
	return nil
}

func (st *SessionStore) Get(key interface{}) interface{} {
	pder.SessionUpdate(st.sid)
	if v, ok := st.value[key]; ok {
		return v
	}
	return nil
}

func (st *SessionStore) Delete(key interface{}) error {
	delete(st.value, key)
	pder.SessionUpdate(st.sid)
	return nil
}

func (st *SessionStore) SessionID() string {
	return st.sid
}

type Provider struct {
	lock     sync.Mutex
	sessions map[string]*list.Element // 用来存储在内存
	list     *list.List               // 用来做gc
}

func (pder *Provider) SessionInit(sid string) (session.Session, error) {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	v := make(map[interface{}]interface{}, 0)
	newsess := &SessionStore{sid: sid, timeAccessed: time.Now(), value: v}
	element := pder.list.PushBack(newsess)
	pder.sessions[sid] = element
	return newsess, nil
}

func (pder *Provider) SessionRead(sid string) (session.Session, error) {
	if element, ok := pder.sessions[sid]; ok {
		return element.Value.(*SessionStore), nil
	} else {
		sess, err := pder.SessionInit(sid)
		return sess, err
	}
	return nil, nil
}

func (pder *Provider) SessionDestroy(sid string) error {
	if element, ok := pder.sessions[sid]; ok {
		delete(pder.sessions, sid)
		pder.list.Remove(element)
		return nil
	}
	return nil
}

func (pder *Provider) SessionGC(maxlifetime int64) {
	pder.lock.Lock()
	defer pder.lock.Unlock()

	for {
		element := pder.list.Back()
		if element == nil {
			break
		}
		if (element.Value.(*SessionStore).timeAccessed.Unix() + maxlifetime) < time.Now().Unix() {
			pder.list.Remove(element)
			delete(pder.sessions, element.Value.(*SessionStore).sid)
		} else {
			break
		}
	}
}

func (pder *Provider) SessionUpdate(sid string) error {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	if element, ok := pder.sessions[sid]; ok {
		element.Value.(*SessionStore).timeAccessed = time.Now()
		pder.list.MoveToFront(element)
		return nil
	}
	return nil
}

var globalSessions *session.Manager

//然后在init函数中初始化
func init() {
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	go globalSessions.GC()
}

func count(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	key := "countnum"
	ct := sess.Get(key)
	if ct == nil {
		sess.Set(key, 1)
	} else {
		sess.Set(key, (ct.(int) + 1))
	}

	t, _ := template.ParseFiles("count.gtpl")
	w.Header().Set("Content-Type", "text/html")
	t.Execute(w, sess.Get(key))
}

func session1() {
	http.HandleFunc("/count", count)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
