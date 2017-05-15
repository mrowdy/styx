package resource

import (
	"errors"
	"net/http"
	"time"
	"strings"

	"github.com/fetzi/styx/model"
	"github.com/fetzi/styx/storage"
	"github.com/google/uuid"
	"github.com/manyminds/api2go"
)

type MailResource struct {
	MailStatusStorage *storage.MailStatusStorage
}

func (s MailResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	mail, ok := obj.(model.Mail)

	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	mail.ID = uuid.New().String()

	var from []string
	var to [] string

	for _, client := range mail.Clients {
		switch client.Type {
			case "to":
				to = append(to, client.Email)
			case "from":
				from = append(from, client.Email)
		}
	}

	mailStatus := model.MailStatus{
		MailID: mail.ID,
		Subject: mail.Subject,
		From: strings.Join(from, ", "),
		To: strings.Join(to, ", "),
		Created: time.Now().Unix(),
		Sent: 0,
	}

	s.MailStatusStorage.Insert(mailStatus)

	return &Response{Res: mailStatus, Code: http.StatusCreated}, nil
}

func (s MailResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	return &Response{}, api2go.NewHTTPError(errors.New("find one not allowed"), "find one not allowed", http.StatusBadRequest)
}

func (s MailResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	return &Response{}, api2go.NewHTTPError(errors.New("delete not allowed"), "delete not allowed", http.StatusBadRequest)
}

func (s MailResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	return &Response{}, api2go.NewHTTPError(errors.New("update not allowed"), "update not allowed", http.StatusBadRequest)
}