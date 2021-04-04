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

// GetLogEntryProofReader is a Reader for the GetLogEntryProof structure.
type GetLogEntryProofReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLogEntryProofReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLogEntryProofOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetLogEntryProofNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetLogEntryProofDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLogEntryProofOK creates a GetLogEntryProofOK with default headers values
func NewGetLogEntryProofOK() *GetLogEntryProofOK {
	return &GetLogEntryProofOK{}
}

/* GetLogEntryProofOK describes a response with status code 200, with default header values.

Information needed for a client to compute the inclusion proof
*/
type GetLogEntryProofOK struct {
	Payload *models.InclusionProof
}

func (o *GetLogEntryProofOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/log/entries/{entryUUID}/proof][%d] getLogEntryProofOK  %+v", 200, o.Payload)
}
func (o *GetLogEntryProofOK) GetPayload() *models.InclusionProof {
	return o.Payload
}

func (o *GetLogEntryProofOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InclusionProof)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogEntryProofNotFound creates a GetLogEntryProofNotFound with default headers values
func NewGetLogEntryProofNotFound() *GetLogEntryProofNotFound {
	return &GetLogEntryProofNotFound{}
}

/* GetLogEntryProofNotFound describes a response with status code 404, with default header values.

The content requested could not be found
*/
type GetLogEntryProofNotFound struct {
}

func (o *GetLogEntryProofNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/log/entries/{entryUUID}/proof][%d] getLogEntryProofNotFound ", 404)
}

func (o *GetLogEntryProofNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLogEntryProofDefault creates a GetLogEntryProofDefault with default headers values
func NewGetLogEntryProofDefault(code int) *GetLogEntryProofDefault {
	return &GetLogEntryProofDefault{
		_statusCode: code,
	}
}

/* GetLogEntryProofDefault describes a response with status code -1, with default header values.

There was an internal error in the server while processing the request
*/
type GetLogEntryProofDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get log entry proof default response
func (o *GetLogEntryProofDefault) Code() int {
	return o._statusCode
}

func (o *GetLogEntryProofDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/log/entries/{entryUUID}/proof][%d] getLogEntryProof default  %+v", o._statusCode, o.Payload)
}
func (o *GetLogEntryProofDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLogEntryProofDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
