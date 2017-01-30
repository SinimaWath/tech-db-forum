package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Thread Ветка обсуждения на форуме.
//
// swagger:model Thread
type Thread struct {

	// Пользователь, создавший данную тему.
	Author string `json:"author,omitempty"`

	// Дата создания ветки на форуме.
	Created strfmt.DateTime `json:"created,omitempty"`

	// Форум, в котором расположена данная ветка обсуждения.
	// Read Only: true
	Forum string `json:"forum,omitempty"`

	// Идентификатор ветки обсуждения.
	// Read Only: true
	ID int32 `json:"id,omitempty"`

	// Описание ветки обсуждения.
	Message string `json:"message,omitempty"`

	// Человекопонятный URL (https://ru.wikipedia.org/wiki/%D0%A1%D0%B5%D0%BC%D0%B0%D0%BD%D1%82%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%B8%D0%B9_URL).
	// В данной структуре slug опционален и не может быть числом.
	//
	// Read Only: true
	// Pattern: ^(\d|\w|-|_)*(\w|-|_)(\d|\w|-|_)*$
	Slug string `json:"slug,omitempty"`

	// Заголовок ветки обсуждения.
	Title string `json:"title,omitempty"`

	// Кол-во голосов непосредственно за данное сообщение форума.
	// Read Only: true
	Votes int32 `json:"votes,omitempty"`
}

// Validate validates this thread
func (m *Thread) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSlug(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Thread) validateSlug(formats strfmt.Registry) error {

	if swag.IsZero(m.Slug) { // not required
		return nil
	}

	if err := validate.Pattern("slug", "body", string(m.Slug), `^(\d|\w|-|_)*(\w|-|_)(\d|\w|-|_)*$`); err != nil {
		return err
	}

	return nil
}