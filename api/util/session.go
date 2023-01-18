package util

import (
	"log"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Session struct {
	storage sessions.Session
}

func NewSession(c *gin.Context) Session {
	return Session{sessions.Default(c)}
}

func (s *Session) SetSessionData(data string) bool {
	s.storage.Set(os.Getenv("SESSION_SECRET"), data)
	if err := s.storage.Save(); err != nil {
		log.Printf("error in saving session data: %s", err)
		return false
	}

	return true
}

func (s *Session) GetSessionData(message string) interface{} {
	data := s.storage.Get(message)
	if err := s.storage.Save(); err != nil {
		log.Printf("error in getting session data: %s", err)
		return nil
	}
	return data
}

func (s *Session) RemoveSessionData(message string) bool {
	s.storage.AddFlash(message)
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

func (s *Session) GetFlashMessage(c *gin.Context) []interface{} {
	flashes := s.storage.Flashes()
	if len(flashes) != 0 {
		if err := s.storage.Save(); err != nil {
			log.Printf("error in flashes saving session: %s", err)
			return nil
		}
	}
	return flashes
}
