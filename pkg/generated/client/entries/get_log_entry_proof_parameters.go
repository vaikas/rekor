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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetLogEntryProofParams creates a new GetLogEntryProofParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLogEntryProofParams() *GetLogEntryProofParams {
	return &GetLogEntryProofParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLogEntryProofParamsWithTimeout creates a new GetLogEntryProofParams object
// with the ability to set a timeout on a request.
func NewGetLogEntryProofParamsWithTimeout(timeout time.Duration) *GetLogEntryProofParams {
	return &GetLogEntryProofParams{
		timeout: timeout,
	}
}

// NewGetLogEntryProofParamsWithContext creates a new GetLogEntryProofParams object
// with the ability to set a context for a request.
func NewGetLogEntryProofParamsWithContext(ctx context.Context) *GetLogEntryProofParams {
	return &GetLogEntryProofParams{
		Context: ctx,
	}
}

// NewGetLogEntryProofParamsWithHTTPClient creates a new GetLogEntryProofParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLogEntryProofParamsWithHTTPClient(client *http.Client) *GetLogEntryProofParams {
	return &GetLogEntryProofParams{
		HTTPClient: client,
	}
}

/* GetLogEntryProofParams contains all the parameters to send to the API endpoint
   for the get log entry proof operation.

   Typically these are written to a http.Request.
*/
type GetLogEntryProofParams struct {

	/* EntryUUID.

	   the UUID of the entry for which the inclusion proof information should be returned
	*/
	EntryUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get log entry proof params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLogEntryProofParams) WithDefaults() *GetLogEntryProofParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get log entry proof params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLogEntryProofParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get log entry proof params
func (o *GetLogEntryProofParams) WithTimeout(timeout time.Duration) *GetLogEntryProofParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get log entry proof params
func (o *GetLogEntryProofParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get log entry proof params
func (o *GetLogEntryProofParams) WithContext(ctx context.Context) *GetLogEntryProofParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get log entry proof params
func (o *GetLogEntryProofParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get log entry proof params
func (o *GetLogEntryProofParams) WithHTTPClient(client *http.Client) *GetLogEntryProofParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get log entry proof params
func (o *GetLogEntryProofParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntryUUID adds the entryUUID to the get log entry proof params
func (o *GetLogEntryProofParams) WithEntryUUID(entryUUID string) *GetLogEntryProofParams {
	o.SetEntryUUID(entryUUID)
	return o
}

// SetEntryUUID adds the entryUuid to the get log entry proof params
func (o *GetLogEntryProofParams) SetEntryUUID(entryUUID string) {
	o.EntryUUID = entryUUID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLogEntryProofParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param entryUUID
	if err := r.SetPathParam("entryUUID", o.EntryUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
