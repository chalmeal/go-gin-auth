/*
* app sessions
 */
package sessions

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type SessionInfo struct {
	UserId string
	State  string
}

func (si *SessionInfo) GetUserId(c *gin.Context) interface{} {
	sess := sessions.Default(c)
	return sess.Get("user_id")
}

func (si *SessionInfo) GetState(c *gin.Context) interface{} {
	sess := sessions.Default(c)
	return sess.Get("state")
}

func (si *SessionInfo) Set(c *gin.Context, s *SessionInfo) {
	sess := sessions.Default(c)
	sess.Set("user_id", s.UserId)
	sess.Set("state", s.State)
	sess.Save()
}

func (si *SessionInfo) Clear(c *gin.Context) {
	sess := sessions.Default(c)
	sess.Clear()
	sess.Save()
}
