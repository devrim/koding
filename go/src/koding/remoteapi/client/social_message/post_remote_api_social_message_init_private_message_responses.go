package social_message

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPISocialMessageInitPrivateMessageReader is a Reader for the PostRemoteAPISocialMessageInitPrivateMessage structure.
type PostRemoteAPISocialMessageInitPrivateMessageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPISocialMessageInitPrivateMessageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPISocialMessageInitPrivateMessageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPISocialMessageInitPrivateMessageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPISocialMessageInitPrivateMessageOK creates a PostRemoteAPISocialMessageInitPrivateMessageOK with default headers values
func NewPostRemoteAPISocialMessageInitPrivateMessageOK() *PostRemoteAPISocialMessageInitPrivateMessageOK {
	return &PostRemoteAPISocialMessageInitPrivateMessageOK{}
}

/*PostRemoteAPISocialMessageInitPrivateMessageOK handles this case with default header values.

Request processed succesfully
*/
type PostRemoteAPISocialMessageInitPrivateMessageOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPISocialMessageInitPrivateMessageOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialMessage.initPrivateMessage][%d] postRemoteApiSocialMessageInitPrivateMessageOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPISocialMessageInitPrivateMessageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPISocialMessageInitPrivateMessageUnauthorized creates a PostRemoteAPISocialMessageInitPrivateMessageUnauthorized with default headers values
func NewPostRemoteAPISocialMessageInitPrivateMessageUnauthorized() *PostRemoteAPISocialMessageInitPrivateMessageUnauthorized {
	return &PostRemoteAPISocialMessageInitPrivateMessageUnauthorized{}
}

/*PostRemoteAPISocialMessageInitPrivateMessageUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPISocialMessageInitPrivateMessageUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPISocialMessageInitPrivateMessageUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialMessage.initPrivateMessage][%d] postRemoteApiSocialMessageInitPrivateMessageUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPISocialMessageInitPrivateMessageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
