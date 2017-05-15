package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/bozaro/tech-db-forum/generated/models"
)

// ForumGetUsersReader is a Reader for the ForumGetUsers structure.
type ForumGetUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ForumGetUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewForumGetUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewForumGetUsersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewForumGetUsersOK creates a ForumGetUsersOK with default headers values
func NewForumGetUsersOK() *ForumGetUsersOK {
	return &ForumGetUsersOK{}
}

/*ForumGetUsersOK handles this case with default header values.

Информация о пользователях форума.

*/
type ForumGetUsersOK struct {
	Payload models.Users
}

func (o *ForumGetUsersOK) Error() string {
	return fmt.Sprintf("[GET /forum/{slug}/users][%d] forumGetUsersOK  %+v", 200, o.Payload)
}

func (o *ForumGetUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewForumGetUsersNotFound creates a ForumGetUsersNotFound with default headers values
func NewForumGetUsersNotFound() *ForumGetUsersNotFound {
	return &ForumGetUsersNotFound{}
}

/*ForumGetUsersNotFound handles this case with default header values.

Форум отсутсвует в системе.

*/
type ForumGetUsersNotFound struct {
}

func (o *ForumGetUsersNotFound) Error() string {
	return fmt.Sprintf("[GET /forum/{slug}/users][%d] forumGetUsersNotFound ", 404)
}

func (o *ForumGetUsersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
