// Code generated by go-swagger; DO NOT EDIT.

package feature_view_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewFeatureViewServiceListFeatureViewsParams creates a new FeatureViewServiceListFeatureViewsParams object
// with the default values initialized.
func NewFeatureViewServiceListFeatureViewsParams() *FeatureViewServiceListFeatureViewsParams {
	var ()
	return &FeatureViewServiceListFeatureViewsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFeatureViewServiceListFeatureViewsParamsWithTimeout creates a new FeatureViewServiceListFeatureViewsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFeatureViewServiceListFeatureViewsParamsWithTimeout(timeout time.Duration) *FeatureViewServiceListFeatureViewsParams {
	var ()
	return &FeatureViewServiceListFeatureViewsParams{

		timeout: timeout,
	}
}

// NewFeatureViewServiceListFeatureViewsParamsWithContext creates a new FeatureViewServiceListFeatureViewsParams object
// with the default values initialized, and the ability to set a context for a request
func NewFeatureViewServiceListFeatureViewsParamsWithContext(ctx context.Context) *FeatureViewServiceListFeatureViewsParams {
	var ()
	return &FeatureViewServiceListFeatureViewsParams{

		Context: ctx,
	}
}

// NewFeatureViewServiceListFeatureViewsParamsWithHTTPClient creates a new FeatureViewServiceListFeatureViewsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFeatureViewServiceListFeatureViewsParamsWithHTTPClient(client *http.Client) *FeatureViewServiceListFeatureViewsParams {
	var ()
	return &FeatureViewServiceListFeatureViewsParams{
		HTTPClient: client,
	}
}

/*FeatureViewServiceListFeatureViewsParams contains all the parameters to send to the API endpoint
for the feature view service list feature views operation typically these are written to a http.Request
*/
type FeatureViewServiceListFeatureViewsParams struct {

	/*Project*/
	Project *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the feature view service list feature views params
func (o *FeatureViewServiceListFeatureViewsParams) WithTimeout(timeout time.Duration) *FeatureViewServiceListFeatureViewsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the feature view service list feature views params
func (o *FeatureViewServiceListFeatureViewsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the feature view service list feature views params
func (o *FeatureViewServiceListFeatureViewsParams) WithContext(ctx context.Context) *FeatureViewServiceListFeatureViewsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the feature view service list feature views params
func (o *FeatureViewServiceListFeatureViewsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the feature view service list feature views params
func (o *FeatureViewServiceListFeatureViewsParams) WithHTTPClient(client *http.Client) *FeatureViewServiceListFeatureViewsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the feature view service list feature views params
func (o *FeatureViewServiceListFeatureViewsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProject adds the project to the feature view service list feature views params
func (o *FeatureViewServiceListFeatureViewsParams) WithProject(project *string) *FeatureViewServiceListFeatureViewsParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the feature view service list feature views params
func (o *FeatureViewServiceListFeatureViewsParams) SetProject(project *string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *FeatureViewServiceListFeatureViewsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Project != nil {

		// query param project
		var qrProject string
		if o.Project != nil {
			qrProject = *o.Project
		}
		qProject := qrProject
		if qProject != "" {
			if err := r.SetQueryParam("project", qProject); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}