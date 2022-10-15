// Code generated by go-swagger; DO NOT EDIT.

package api_security

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

// NewDeleteAPISecurityAPICatalogIDParams creates a new DeleteAPISecurityAPICatalogIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAPISecurityAPICatalogIDParams() *DeleteAPISecurityAPICatalogIDParams {
	return &DeleteAPISecurityAPICatalogIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPISecurityAPICatalogIDParamsWithTimeout creates a new DeleteAPISecurityAPICatalogIDParams object
// with the ability to set a timeout on a request.
func NewDeleteAPISecurityAPICatalogIDParamsWithTimeout(timeout time.Duration) *DeleteAPISecurityAPICatalogIDParams {
	return &DeleteAPISecurityAPICatalogIDParams{
		timeout: timeout,
	}
}

// NewDeleteAPISecurityAPICatalogIDParamsWithContext creates a new DeleteAPISecurityAPICatalogIDParams object
// with the ability to set a context for a request.
func NewDeleteAPISecurityAPICatalogIDParamsWithContext(ctx context.Context) *DeleteAPISecurityAPICatalogIDParams {
	return &DeleteAPISecurityAPICatalogIDParams{
		Context: ctx,
	}
}

// NewDeleteAPISecurityAPICatalogIDParamsWithHTTPClient creates a new DeleteAPISecurityAPICatalogIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAPISecurityAPICatalogIDParamsWithHTTPClient(client *http.Client) *DeleteAPISecurityAPICatalogIDParams {
	return &DeleteAPISecurityAPICatalogIDParams{
		HTTPClient: client,
	}
}

/* DeleteAPISecurityAPICatalogIDParams contains all the parameters to send to the API endpoint
   for the delete API security API catalog ID operation.

   Typically these are written to a http.Request.
*/
type DeleteAPISecurityAPICatalogIDParams struct {

	// CatalogID.
	//
	// Format: uuid
	CatalogID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete API security API catalog ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAPISecurityAPICatalogIDParams) WithDefaults() *DeleteAPISecurityAPICatalogIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete API security API catalog ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAPISecurityAPICatalogIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete API security API catalog ID params
func (o *DeleteAPISecurityAPICatalogIDParams) WithTimeout(timeout time.Duration) *DeleteAPISecurityAPICatalogIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete API security API catalog ID params
func (o *DeleteAPISecurityAPICatalogIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete API security API catalog ID params
func (o *DeleteAPISecurityAPICatalogIDParams) WithContext(ctx context.Context) *DeleteAPISecurityAPICatalogIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete API security API catalog ID params
func (o *DeleteAPISecurityAPICatalogIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete API security API catalog ID params
func (o *DeleteAPISecurityAPICatalogIDParams) WithHTTPClient(client *http.Client) *DeleteAPISecurityAPICatalogIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete API security API catalog ID params
func (o *DeleteAPISecurityAPICatalogIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCatalogID adds the catalogID to the delete API security API catalog ID params
func (o *DeleteAPISecurityAPICatalogIDParams) WithCatalogID(catalogID strfmt.UUID) *DeleteAPISecurityAPICatalogIDParams {
	o.SetCatalogID(catalogID)
	return o
}

// SetCatalogID adds the catalogId to the delete API security API catalog ID params
func (o *DeleteAPISecurityAPICatalogIDParams) SetCatalogID(catalogID strfmt.UUID) {
	o.CatalogID = catalogID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPISecurityAPICatalogIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param catalogId
	if err := r.SetPathParam("catalogId", o.CatalogID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}