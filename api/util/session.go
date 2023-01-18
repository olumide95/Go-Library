package util

import (
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Session struct {
	storage sessions.Session
}

func NewSession(c *gin.Context) Session {
	return Session{sessions.Default(c)}
}

func (s *Session) SetSessionData(key string, data string) bool {
	s.storage.Set(key, data)
	if err := s.storage.Save(); err != nil {
		log.Printf("error in saving session data: %s", err)
		return false
	}

	return true
}

func (s *Session) GetSessionData(key string) interface{} {
	data := s.storage.Get(key)
	if err := s.storage.Save(); err != nil {
		log.Printf("error in getting session data: %s", err)
		return nil
	}
	return data
}

func (s *Session) RemoveSessionData(key string) bool {
	s.storage.Delete(key)
	if err := s.storage.Save(); err != nil {
		log.Printf("error in removing session data: %s", err)
		return false
	}

	return true
}

func (s *Session) SetFlashMessage(message string) bool {
	s.storage.AddFlash(message)
	if err := s.storage.Save(); err != nil {
		log.Printf("error in SetFlashMessage saving session: %s", err)
		return false
	}

	return true
}

func (s *Session) GetFlashMessage() []interface{} {
	flashes := s.storage.Flashes()
	if len(flashes) != 0 {
		if err := s.storage.Save(); err != nil {
			log.Printf("error in flashes saving session: %s", err)
			return nil
		}
	}
	return flashes
}
