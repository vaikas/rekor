// Code generated by go-swagger; DO NOT EDIT.

// /*
// Copyright The Rekor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */
//

package entries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/sigstore/rekor/pkg/generated/models"
)

// GetLogEntryByUUIDReader is a Reader for the GetLogEntryByUUID structure.
type GetLogEntryByUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLogEntryByUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLogEntryByUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetLogEntryByUUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetLogEntryByUUIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLogEntryByUUIDOK creates a GetLogEntryByUUIDOK with default headers values
func NewGetLogEntryByUUIDOK() *GetLogEntryByUUIDOK {
	return &GetLogEntryByUUIDOK{}
}

/* GetLogEntryByUUIDOK describes a response with status code 200, with default header values.

the entry in the transparency log requested
*/
type GetLogEntryByUUIDOK struct {
	Payload models.LogEntry
}

func (o *GetLogEntryByUUIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/log/entries/{entryUUID}][%d] getLogEntryByUuidOK  %+v", 200, o.Payload)
}
func (o *GetLogEntryByUUIDOK) GetPayload() models.LogEntry {
	return o.Payload
}

func (o *GetLogEntryByUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogEntryByUUIDNotFound creates a GetLogEntryByUUIDNotFound with default headers values
func NewGetLogEntryByUUIDNotFound() *GetLogEntryByUUIDNotFound {
	return &GetLogEntryByUUIDNotFound{}
}

/* GetLogEntryByUUIDNotFound describes a response with status code 404, with default header values.

The content requested could not be found
*/
type GetLogEntryByUUIDNotFound struct {
}

func (o *GetLogEntryByUUIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/log/entries/{entryUUID}][%d] getLogEntryByUuidNotFound ", 404)
}

func (o *GetLogEntryByUUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLogEntryByUUIDDefault creates a GetLogEntryByUUIDDefault with default headers values
func NewGetLogEntryByUUIDDefault(code int) *GetLogEntryByUUIDDefault {
	return &GetLogEntryByUUIDDefault{
		_statusCode: code,
	}
}

/* GetLogEntryByUUIDDefault describes a response with status code -1, with default header values.

There was an internal error in the server while processing the request
*/
type GetLogEntryByUUIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get log entry by UUID default response
func (o *GetLogEntryByUUIDDefault) Code() int {
	return o._statusCode
}

func (o *GetLogEntryByUUIDDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/log/entries/{entryUUID}][%d] getLogEntryByUUID default  %+v", o._statusCode, o.Payload)
}
func (o *GetLogEntryByUUIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLogEntryByUUIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
