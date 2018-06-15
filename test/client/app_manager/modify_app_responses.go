// Code generated by go-swagger; DO NOT EDIT.

package app_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "openpitrix.io/openpitrix/test/models"
)

// ModifyAppReader is a Reader for the ModifyApp structure.
type ModifyAppReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyAppReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewModifyAppOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewModifyAppOK creates a ModifyAppOK with default headers values
func NewModifyAppOK() *ModifyAppOK {
	return &ModifyAppOK{}
}

/*ModifyAppOK handles this case with default header values.

ModifyAppOK modify app o k
*/
type ModifyAppOK struct {
	Payload *models.OpenpitrixModifyAppResponse
}

func (o *ModifyAppOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/apps][%d] modifyAppOK  %+v", 200, o.Payload)
}

func (o *ModifyAppOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixModifyAppResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
