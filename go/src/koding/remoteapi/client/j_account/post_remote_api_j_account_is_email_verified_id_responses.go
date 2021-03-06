package j_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJAccountIsEmailVerifiedIDReader is a Reader for the PostRemoteAPIJAccountIsEmailVerifiedID structure.
type PostRemoteAPIJAccountIsEmailVerifiedIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJAccountIsEmailVerifiedIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJAccountIsEmailVerifiedIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJAccountIsEmailVerifiedIDOK creates a PostRemoteAPIJAccountIsEmailVerifiedIDOK with default headers values
func NewPostRemoteAPIJAccountIsEmailVerifiedIDOK() *PostRemoteAPIJAccountIsEmailVerifiedIDOK {
	return &PostRemoteAPIJAccountIsEmailVerifiedIDOK{}
}

/*PostRemoteAPIJAccountIsEmailVerifiedIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJAccountIsEmailVerifiedIDOK struct {
	Payload *models.JAccount
}

func (o *PostRemoteAPIJAccountIsEmailVerifiedIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JAccount.isEmailVerified/{id}][%d] postRemoteApiJAccountIsEmailVerifiedIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJAccountIsEmailVerifiedIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
