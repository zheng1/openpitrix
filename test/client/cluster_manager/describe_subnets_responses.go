// Code generated by go-swagger; DO NOT EDIT.

package cluster_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "openpitrix.io/openpitrix/test/models"
)

// DescribeSubnetsReader is a Reader for the DescribeSubnets structure.
type DescribeSubnetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeSubnetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDescribeSubnetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDescribeSubnetsOK creates a DescribeSubnetsOK with default headers values
func NewDescribeSubnetsOK() *DescribeSubnetsOK {
	return &DescribeSubnetsOK{}
}

/*DescribeSubnetsOK handles this case with default header values.

DescribeSubnetsOK describe subnets o k
*/
type DescribeSubnetsOK struct {
	Payload *models.OpenpitrixDescribeSubnetsResponse
}

func (o *DescribeSubnetsOK) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/subnets][%d] describeSubnetsOK  %+v", 200, o.Payload)
}

func (o *DescribeSubnetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixDescribeSubnetsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
