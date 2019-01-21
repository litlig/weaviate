/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@creativesoftwarefdn.org
 */ // Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateBatchingActionsCreateReader is a Reader for the WeaviateBatchingActionsCreate structure.
type WeaviateBatchingActionsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateBatchingActionsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateBatchingActionsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewWeaviateBatchingActionsCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateBatchingActionsCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateBatchingActionsCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewWeaviateBatchingActionsCreateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateBatchingActionsCreateOK creates a WeaviateBatchingActionsCreateOK with default headers values
func NewWeaviateBatchingActionsCreateOK() *WeaviateBatchingActionsCreateOK {
	return &WeaviateBatchingActionsCreateOK{}
}

/*WeaviateBatchingActionsCreateOK handles this case with default header values.

Actions created.
*/
type WeaviateBatchingActionsCreateOK struct {
	Payload []*models.ActionsGetResponse
}

func (o *WeaviateBatchingActionsCreateOK) Error() string {
	return fmt.Sprintf("[POST /batching/actions][%d] weaviateBatchingActionsCreateOK  %+v", 200, o.Payload)
}

func (o *WeaviateBatchingActionsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateBatchingActionsCreateAccepted creates a WeaviateBatchingActionsCreateAccepted with default headers values
func NewWeaviateBatchingActionsCreateAccepted() *WeaviateBatchingActionsCreateAccepted {
	return &WeaviateBatchingActionsCreateAccepted{}
}

/*WeaviateBatchingActionsCreateAccepted handles this case with default header values.

Successfully received.
*/
type WeaviateBatchingActionsCreateAccepted struct {
	Payload []*models.ActionsGetResponse
}

func (o *WeaviateBatchingActionsCreateAccepted) Error() string {
	return fmt.Sprintf("[POST /batching/actions][%d] weaviateBatchingActionsCreateAccepted  %+v", 202, o.Payload)
}

func (o *WeaviateBatchingActionsCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateBatchingActionsCreateUnauthorized creates a WeaviateBatchingActionsCreateUnauthorized with default headers values
func NewWeaviateBatchingActionsCreateUnauthorized() *WeaviateBatchingActionsCreateUnauthorized {
	return &WeaviateBatchingActionsCreateUnauthorized{}
}

/*WeaviateBatchingActionsCreateUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateBatchingActionsCreateUnauthorized struct {
}

func (o *WeaviateBatchingActionsCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /batching/actions][%d] weaviateBatchingActionsCreateUnauthorized ", 401)
}

func (o *WeaviateBatchingActionsCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateBatchingActionsCreateForbidden creates a WeaviateBatchingActionsCreateForbidden with default headers values
func NewWeaviateBatchingActionsCreateForbidden() *WeaviateBatchingActionsCreateForbidden {
	return &WeaviateBatchingActionsCreateForbidden{}
}

/*WeaviateBatchingActionsCreateForbidden handles this case with default header values.

The used API-key has insufficient permissions.
*/
type WeaviateBatchingActionsCreateForbidden struct {
}

func (o *WeaviateBatchingActionsCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /batching/actions][%d] weaviateBatchingActionsCreateForbidden ", 403)
}

func (o *WeaviateBatchingActionsCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateBatchingActionsCreateUnprocessableEntity creates a WeaviateBatchingActionsCreateUnprocessableEntity with default headers values
func NewWeaviateBatchingActionsCreateUnprocessableEntity() *WeaviateBatchingActionsCreateUnprocessableEntity {
	return &WeaviateBatchingActionsCreateUnprocessableEntity{}
}

/*WeaviateBatchingActionsCreateUnprocessableEntity handles this case with default header values.

Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?
*/
type WeaviateBatchingActionsCreateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateBatchingActionsCreateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /batching/actions][%d] weaviateBatchingActionsCreateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *WeaviateBatchingActionsCreateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*WeaviateBatchingActionsCreateBody weaviate batching actions create body
swagger:model WeaviateBatchingActionsCreateBody
*/
type WeaviateBatchingActionsCreateBody struct {

	// actions
	Actions []*models.ActionCreate `json:"actions"`

	// If `async` is true, return a 202 with the new ID of the Action. You will receive this response before the persistence of the data is confirmed. If `async` is false, you will receive confirmation after the persistence of the data is confirmed. The value of `async` defaults to false.
	Async bool `json:"async,omitempty"`

	// Define which fields need to be returned. Default value is ALL
	Fields []*string `json:"fields"`
}

// Validate validates this weaviate batching actions create body
func (o *WeaviateBatchingActionsCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFields(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *WeaviateBatchingActionsCreateBody) validateActions(formats strfmt.Registry) error {

	if swag.IsZero(o.Actions) { // not required
		return nil
	}

	for i := 0; i < len(o.Actions); i++ {
		if swag.IsZero(o.Actions[i]) { // not required
			continue
		}

		if o.Actions[i] != nil {
			if err := o.Actions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "actions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var weaviateBatchingActionsCreateBodyFieldsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ALL","@class","schema","key","actionId","creationTimeUnix"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		weaviateBatchingActionsCreateBodyFieldsItemsEnum = append(weaviateBatchingActionsCreateBodyFieldsItemsEnum, v)
	}
}

func (o *WeaviateBatchingActionsCreateBody) validateFieldsItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, weaviateBatchingActionsCreateBodyFieldsItemsEnum); err != nil {
		return err
	}
	return nil
}

func (o *WeaviateBatchingActionsCreateBody) validateFields(formats strfmt.Registry) error {

	if swag.IsZero(o.Fields) { // not required
		return nil
	}

	for i := 0; i < len(o.Fields); i++ {
		if swag.IsZero(o.Fields[i]) { // not required
			continue
		}

		// value enum
		if err := o.validateFieldsItemsEnum("body"+"."+"fields"+"."+strconv.Itoa(i), "body", *o.Fields[i]); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *WeaviateBatchingActionsCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *WeaviateBatchingActionsCreateBody) UnmarshalBinary(b []byte) error {
	var res WeaviateBatchingActionsCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
