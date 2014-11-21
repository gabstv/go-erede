package erede

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"
	"time"
)

type WebService struct {
	user        string
	password    string
	environment int
	verbose     false
}

func NewWebService(user, pw string) *WebService {
	return &WebService{user, pw, 1, false}
}

func (ws *WebService) URL() string {
	if ws.environment == ProductionEnv {
		return URLPROD
	}
	return URLDEV
}

//TODO: proper fulfill and pre methods
