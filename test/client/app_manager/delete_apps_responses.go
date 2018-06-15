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

// DeleteAppsReader is a Reader for the DeleteApps structure.
type DeleteAppsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAppsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteAppsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteAppsOK creates a DeleteAppsOK with default headers values
func NewDeleteAppsOK() *DeleteAppsOK {
	return &DeleteAppsOK{}
}

/*DeleteAppsOK handles this case with default header values.

DeleteAppsOK delete apps o k
*/
type DeleteAppsOK struct {
	Payload *models.OpenpitrixDeleteAppsResponse
}

func (o *DeleteAppsOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps][%d] deleteAppsOK  %+v", 200, o.Payload)
}

func (o *DeleteAppsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixDeleteAppsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
