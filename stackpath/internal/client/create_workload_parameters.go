// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/terraform-providers/terraform-provider-stackpath/stackpath/internal/models"
)

// NewCreateWorkloadParams creates a new CreateWorkloadParams object
// with the default values initialized.
func NewCreateWorkloadParams() *CreateWorkloadParams {
	var ()
	return &CreateWorkloadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateWorkloadParamsWithTimeout creates a new CreateWorkloadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateWorkloadParamsWithTimeout(timeout time.Duration) *CreateWorkloadParams {
	var ()
	return &CreateWorkloadParams{

		timeout: timeout,
	}
}

// NewCreateWorkloadParamsWithContext creates a new CreateWorkloadParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateWorkloadParamsWithContext(ctx context.Context) *CreateWorkloadParams {
	var ()
	return &CreateWorkloadParams{

		Context: ctx,
	}
}

// NewCreateWorkloadParamsWithHTTPClient creates a new CreateWorkloadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateWorkloadParamsWithHTTPClient(client *http.Client) *CreateWorkloadParams {
	var ()
	return &CreateWorkloadParams{
		HTTPClient: client,
	}
}

/*CreateWorkloadParams contains all the parameters to send to the API endpoint
for the create workload operation typically these are written to a http.Request
*/
type CreateWorkloadParams struct {

	/*Body*/
	Body *models.V1CreateWorkloadRequest
	/*StackID
	  The ID of the stack to create the workload in

	*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create workload params
func (o *CreateWorkloadParams) WithTimeout(timeout time.Duration) *CreateWorkloadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create workload params
func (o *CreateWorkloadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create workload params
func (o *CreateWorkloadParams) WithContext(ctx context.Context) *CreateWorkloadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create workload params
func (o *CreateWorkloadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create workload params
func (o *CreateWorkloadParams) WithHTTPClient(client *http.Client) *CreateWorkloadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create workload params
func (o *CreateWorkloadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create workload params
func (o *CreateWorkloadParams) WithBody(body *models.V1CreateWorkloadRequest) *CreateWorkloadParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create workload params
func (o *CreateWorkloadParams) SetBody(body *models.V1CreateWorkloadRequest) {
	o.Body = body
}

// WithStackID adds the stackID to the create workload params
func (o *CreateWorkloadParams) WithStackID(stackID string) *CreateWorkloadParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the create workload params
func (o *CreateWorkloadParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateWorkloadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param stack_id
	if err := r.SetPathParam("stack_id", o.StackID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
