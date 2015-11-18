package utils

import (
	"github.com/gorilla/sessions"
	"net/http"
	//"os"
)

type Session struct {
	realSession *sessions.Session
	w           http.ResponseWriter
	r           *http.Request
}

var sessionStore *sessions.CookieStore

func GetSession(w http.ResponseWriter, r *http.Request) (session *Session) {
	rs, _ := sessionStore.Get(r, "session-name")
	if rs == nil {
		Errorf("Error get Session")
	}
	rs.Options = &sessions.Options{
		Path: "/",
	}
	return &Session{rs, w, r}
}

func (s *Session) Get(key string) string {
	if s.realSession == nil {
		Warnf("Real session is nil")
		return ""
	}
	Debugf("Session id in get: %s", s.realSession.ID)
	value, _ := s.realSession.Values[key]
	if value == nil {
		return ""
	} else {
		return value.(string)
	}
}

func (s *Session) Set(key, value string) {
	if s.realSession == nil {
		Warnf("Real session is nil")
		return
	}
	s.realSession.Values[key] = value
	if err := s.realSession.Save(s.r, s.w); err != nil {
		Errorf("Error save session %s, %s: %s", key, value, err.Error())
	}
	Debugf("Session id in set: %s", s.realSession.ID)
}
