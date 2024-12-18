// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_virtual_serial_number

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

// NewPcloudVirtualserialnumberGetParams creates a new PcloudVirtualserialnumberGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudVirtualserialnumberGetParams() *PcloudVirtualserialnumberGetParams {
	return &PcloudVirtualserialnumberGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudVirtualserialnumberGetParamsWithTimeout creates a new PcloudVirtualserialnumberGetParams object
// with the ability to set a timeout on a request.
func NewPcloudVirtualserialnumberGetParamsWithTimeout(timeout time.Duration) *PcloudVirtualserialnumberGetParams {
	return &PcloudVirtualserialnumberGetParams{
		timeout: timeout,
	}
}

// NewPcloudVirtualserialnumberGetParamsWithContext creates a new PcloudVirtualserialnumberGetParams object
// with the ability to set a context for a request.
func NewPcloudVirtualserialnumberGetParamsWithContext(ctx context.Context) *PcloudVirtualserialnumberGetParams {
	return &PcloudVirtualserialnumberGetParams{
		Context: ctx,
	}
}

// NewPcloudVirtualserialnumberGetParamsWithHTTPClient creates a new PcloudVirtualserialnumberGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudVirtualserialnumberGetParamsWithHTTPClient(client *http.Client) *PcloudVirtualserialnumberGetParams {
	return &PcloudVirtualserialnumberGetParams{
		HTTPClient: client,
	}
}

/*
PcloudVirtualserialnumberGetParams contains all the parameters to send to the API endpoint

	for the pcloud virtualserialnumber get operation.

	Typically these are written to a http.Request.
*/
type PcloudVirtualserialnumberGetParams struct {

	/* VirtualSerialNumber.

	   Virtual Serial Number
	*/
	VirtualSerialNumber string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud virtualserialnumber get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudVirtualserialnumberGetParams) WithDefaults() *PcloudVirtualserialnumberGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud virtualserialnumber get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudVirtualserialnumberGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud virtualserialnumber get params
func (o *PcloudVirtualserialnumberGetParams) WithTimeout(timeout time.Duration) *PcloudVirtualserialnumberGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud virtualserialnumber get params
func (o *PcloudVirtualserialnumberGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud virtualserialnumber get params
func (o *PcloudVirtualserialnumberGetParams) WithContext(ctx context.Context) *PcloudVirtualserialnumberGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud virtualserialnumber get params
func (o *PcloudVirtualserialnumberGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud virtualserialnumber get params
func (o *PcloudVirtualserialnumberGetParams) WithHTTPClient(client *http.Client) *PcloudVirtualserialnumberGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud virtualserialnumber get params
func (o *PcloudVirtualserialnumberGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVirtualSerialNumber adds the virtualSerialNumber to the pcloud virtualserialnumber get params
func (o *PcloudVirtualserialnumberGetParams) WithVirtualSerialNumber(virtualSerialNumber string) *PcloudVirtualserialnumberGetParams {
	o.SetVirtualSerialNumber(virtualSerialNumber)
	return o
}

// SetVirtualSerialNumber adds the virtualSerialNumber to the pcloud virtualserialnumber get params
func (o *PcloudVirtualserialnumberGetParams) SetVirtualSerialNumber(virtualSerialNumber string) {
	o.VirtualSerialNumber = virtualSerialNumber
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudVirtualserialnumberGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param virtual_serial_number
	if err := r.SetPathParam("virtual_serial_number", o.VirtualSerialNumber); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
