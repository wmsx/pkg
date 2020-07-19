package gin

import (
	g "github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/rbcervilla/redisstore"
	"net/http"
)

const (
	SessionKey = "sid"
)

var (
	store *redisstore.RedisStore
)

type Session struct {
	s *sessions.Session
	c *g.Context
}

func NewSession(c *g.Context) (*Session, error) {
	var (
		s   *sessions.Session
		err error
	)
	if s, err = store.Get(c.Request, SessionKey); err != nil {
		return nil, err
	}
	return &Session{s: s, c: c}, nil
}

func (s *Session) SaveMenger(id int64, name string) {
	s.s.Values["id"] = id
	s.s.Values["name"] = name
}

func (s *Session) GetMengerId() int64 {
	return s.s.Values["id"].(int64)
}

func (s *Session) Save() error {
	return sessions.Save(s.c.Request, s.c.Writer)
}

func (s *Session) Remove() error {
	s.s.Options.MaxAge = -1
	return sessions.Save(s.c.Request, s.c.Writer)
}

func AuthWrapper(handler g.HandlerFunc) g.HandlerFunc {
	return func(c *g.Context) {
		var (
			s   *sessions.Session
			err error
		)
		if s, err = store.Get(c.Request, SessionKey); err != nil {
			return
		}
		if s == nil {
			c.JSON(http.StatusOK, Response{
				Code:    UnauthorizedErrorCode,
				Message: "未登录",
			})
			return
		}
		handler(c)
	}
}


