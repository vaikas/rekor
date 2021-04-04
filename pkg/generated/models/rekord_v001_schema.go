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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RekordV001Schema Rekor v0.0.1 Schema
//
// Schema for Rekord object
//
// swagger:model rekordV001Schema
type RekordV001Schema struct {

	// data
	// Required: true
	Data *RekordV001SchemaData `json:"data"`

	// Arbitrary content to be included in the verifiable entry in the transparency log
	ExtraData interface{} `json:"extraData,omitempty"`

	// signature
	// Required: true
	Signature *RekordV001SchemaSignature `json:"signature"`
}

// Validate validates this rekord v001 schema
func (m *RekordV001Schema) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignature(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RekordV001Schema) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *RekordV001Schema) validateSignature(formats strfmt.Registry) error {

	if err := validate.Required("signature", "body", m.Signature); err != nil {
		return err
	}

	if m.Signature != nil {
		if err := m.Signature.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("signature")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this rekord v001 schema based on the context it is used
func (m *RekordV001Schema) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSignature(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RekordV001Schema) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if m.Data != nil {
		if err := m.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *RekordV001Schema) contextValidateSignature(ctx context.Context, formats strfmt.Registry) error {

	if m.Signature != nil {
		if err := m.Signature.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("signature")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RekordV001Schema) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RekordV001Schema) UnmarshalBinary(b []byte) error {
	var res RekordV001Schema
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RekordV001SchemaData Information about the content associated with the entry
//
// swagger:model RekordV001SchemaData
type RekordV001SchemaData struct {

	// Specifies the content inline within the document
	// Format: byte
	Content strfmt.Base64 `json:"content,omitempty"`

	// hash
	Hash *RekordV001SchemaDataHash `json:"hash,omitempty"`

	// Specifies the location of the content; if this is specified, a hash value must also be provided
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this rekord v001 schema data
func (m *RekordV001SchemaData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RekordV001SchemaData) validateHash(formats strfmt.Registry) error {
	if swag.IsZero(m.Hash) { // not required
		return nil
	}

	if m.Hash != nil {
		if err := m.Hash.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "hash")
			}
			return err
		}
	}

	return nil
}

func (m *RekordV001SchemaData) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("data"+"."+"url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this rekord v001 schema data based on the context it is used
func (m *RekordV001SchemaData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHash(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RekordV001SchemaData) contextValidateHash(ctx context.Context, formats strfmt.Registry) error {

	if m.Hash != nil {
		if err := m.Hash.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "hash")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RekordV001SchemaData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RekordV001SchemaData) UnmarshalBinary(b []byte) error {
	var res RekordV001SchemaData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RekordV001SchemaDataHash Specifies the hash algorithm and value for the content
//
// swagger:model RekordV001SchemaDataHash
type RekordV001SchemaDataHash struct {

	// The hashing function used to compute the hash value
	// Required: true
	// Enum: [sha256]
	Algorithm *string `json:"algorithm"`

	// The hash value for the content
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this rekord v001 schema data hash
func (m *RekordV001SchemaDataHash) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlgorithm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var rekordV001SchemaDataHashTypeAlgorithmPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["sha256"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rekordV001SchemaDataHashTypeAlgorithmPropEnum = append(rekordV001SchemaDataHashTypeAlgorithmPropEnum, v)
	}
}

const (

	// RekordV001SchemaDataHashAlgorithmSha256 captures enum value "sha256"
	RekordV001SchemaDataHashAlgorithmSha256 string = "sha256"
)

// prop value enum
func (m *RekordV001SchemaDataHash) validateAlgorithmEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rekordV001SchemaDataHashTypeAlgorithmPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RekordV001SchemaDataHash) validateAlgorithm(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"hash"+"."+"algorithm", "body", m.Algorithm); err != nil {
		return err
	}

	// value enum
	if err := m.validateAlgorithmEnum("data"+"."+"hash"+"."+"algorithm", "body", *m.Algorithm); err != nil {
		return err
	}

	return nil
}

func (m *RekordV001SchemaDataHash) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"hash"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this rekord v001 schema data hash based on context it is used
func (m *RekordV001SchemaDataHash) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RekordV001SchemaDataHash) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RekordV001SchemaDataHash) UnmarshalBinary(b []byte) error {
	var res RekordV001SchemaDataHash
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RekordV001SchemaSignature Information about the detached signature associated with the entry
//
// swagger:model RekordV001SchemaSignature
type RekordV001SchemaSignature struct {

	// Specifies the content of the signature inline within the document
	// Format: byte
	Content strfmt.Base64 `json:"content,omitempty"`

	// Specifies the format of the signature
	// Enum: [pgp minisign x509 ssh]
	Format string `json:"format,omitempty"`

	// public key
	PublicKey *RekordV001SchemaSignaturePublicKey `json:"publicKey,omitempty"`

	// Specifies the location of the signature
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this rekord v001 schema signature
func (m *RekordV001SchemaSignature) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var rekordV001SchemaSignatureTypeFormatPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["pgp","minisign","x509","ssh"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rekordV001SchemaSignatureTypeFormatPropEnum = append(rekordV001SchemaSignatureTypeFormatPropEnum, v)
	}
}

const (

	// RekordV001SchemaSignatureFormatPgp captures enum value "pgp"
	RekordV001SchemaSignatureFormatPgp string = "pgp"

	// RekordV001SchemaSignatureFormatMinisign captures enum value "minisign"
	RekordV001SchemaSignatureFormatMinisign string = "minisign"

	// RekordV001SchemaSignatureFormatX509 captures enum value "x509"
	RekordV001SchemaSignatureFormatX509 string = "x509"

	// RekordV001SchemaSignatureFormatSSH captures enum value "ssh"
	RekordV001SchemaSignatureFormatSSH string = "ssh"
)

// prop value enum
func (m *RekordV001SchemaSignature) validateFormatEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rekordV001SchemaSignatureTypeFormatPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RekordV001SchemaSignature) validateFormat(formats strfmt.Registry) error {
	if swag.IsZero(m.Format) { // not required
		return nil
	}

	// value enum
	if err := m.validateFormatEnum("signature"+"."+"format", "body", m.Format); err != nil {
		return err
	}

	return nil
}

func (m *RekordV001SchemaSignature) validatePublicKey(formats strfmt.Registry) error {
	if swag.IsZero(m.PublicKey) { // not required
		return nil
	}

	if m.PublicKey != nil {
		if err := m.PublicKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("signature" + "." + "publicKey")
			}
			return err
		}
	}

	return nil
}

func (m *RekordV001SchemaSignature) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("signature"+"."+"url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this rekord v001 schema signature based on the context it is used
func (m *RekordV001SchemaSignature) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePublicKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RekordV001SchemaSignature) contextValidatePublicKey(ctx context.Context, formats strfmt.Registry) error {

	if m.PublicKey != nil {
		if err := m.PublicKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("signature" + "." + "publicKey")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RekordV001SchemaSignature) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RekordV001SchemaSignature) UnmarshalBinary(b []byte) error {
	var res RekordV001SchemaSignature
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RekordV001SchemaSignaturePublicKey The public key that can verify the signature
//
// swagger:model RekordV001SchemaSignaturePublicKey
type RekordV001SchemaSignaturePublicKey struct {

	// Specifies the content of the public key inline within the document
	// Format: byte
	Content strfmt.Base64 `json:"content,omitempty"`

	// Specifies the location of the public key
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this rekord v001 schema signature public key
func (m *RekordV001SchemaSignaturePublicKey) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RekordV001SchemaSignaturePublicKey) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("signature"+"."+"publicKey"+"."+"url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this rekord v001 schema signature public key based on context it is used
func (m *RekordV001SchemaSignaturePublicKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RekordV001SchemaSignaturePublicKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RekordV001SchemaSignaturePublicKey) UnmarshalBinary(b []byte) error {
	var res RekordV001SchemaSignaturePublicKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
