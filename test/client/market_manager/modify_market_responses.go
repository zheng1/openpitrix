// Code generated by go-swagger; DO NOT EDIT.

package market_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// ModifyMarketReader is a Reader for the ModifyMarket structure.
type ModifyMarketReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyMarketReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewModifyMarketOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewModifyMarketOK creates a ModifyMarketOK with default headers values
func NewModifyMarketOK() *ModifyMarketOK {
	return &ModifyMarketOK{}
}

/*ModifyMarketOK handles this case with default header values.

A successful response.
*/
type ModifyMarketOK struct {
	Payload *models.OpenpitrixModifyMarketResponse
}

func (o *ModifyMarketOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/markets][%d] modifyMarketOK  %+v", 200, o.Payload)
}

func (o *ModifyMarketOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixModifyMarketResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}