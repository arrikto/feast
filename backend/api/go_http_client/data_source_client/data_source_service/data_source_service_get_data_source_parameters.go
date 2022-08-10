// Code generated by go-swagger; DO NOT EDIT.

package data_source_service

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

// NewDataSourceServiceGetDataSourceParams creates a new DataSourceServiceGetDataSourceParams object
// with the default values initialized.
func NewDataSourceServiceGetDataSourceParams() *DataSourceServiceGetDataSourceParams {
	var ()
	return &DataSourceServiceGetDataSourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDataSourceServiceGetDataSourceParamsWithTimeout creates a new DataSourceServiceGetDataSourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDataSourceServiceGetDataSourceParamsWithTimeout(timeout time.Duration) *DataSourceServiceGetDataSourceParams {
	var ()
	return &DataSourceServiceGetDataSourceParams{

		timeout: timeout,
	}
}

// NewDataSourceServiceGetDataSourceParamsWithContext creates a new DataSourceServiceGetDataSourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewDataSourceServiceGetDataSourceParamsWithContext(ctx context.Context) *DataSourceServiceGetDataSourceParams {
	var ()
	return &DataSourceServiceGetDataSourceParams{

		Context: ctx,
	}
}

// NewDataSourceServiceGetDataSourceParamsWithHTTPClient creates a new DataSourceServiceGetDataSourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDataSourceServiceGetDataSourceParamsWithHTTPClient(client *http.Client) *DataSourceServiceGetDataSourceParams {
	var ()
	return &DataSourceServiceGetDataSourceParams{
		HTTPClient: client,
	}
}

/*DataSourceServiceGetDataSourceParams contains all the parameters to send to the API endpoint
for the data source service get data source operation typically these are written to a http.Request
*/
type DataSourceServiceGetDataSourceParams struct {

	/*Name*/
	Name *string
	/*Project*/
	Project *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the data source service get data source params
func (o *DataSourceServiceGetDataSourceParams) WithTimeout(timeout time.Duration) *DataSourceServiceGetDataSourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the data source service get data source params
func (o *DataSourceServiceGetDataSourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the data source service get data source params
func (o *DataSourceServiceGetDataSourceParams) WithContext(ctx context.Context) *DataSourceServiceGetDataSourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the data source service get data source params
func (o *DataSourceServiceGetDataSourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the data source service get data source params
func (o *DataSourceServiceGetDataSourceParams) WithHTTPClient(client *http.Client) *DataSourceServiceGetDataSourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the data source service get data source params
func (o *DataSourceServiceGetDataSourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the data source service get data source params
func (o *DataSourceServiceGetDataSourceParams) WithName(name *string) *DataSourceServiceGetDataSourceParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the data source service get data source params
func (o *DataSourceServiceGetDataSourceParams) SetName(name *string) {
	o.Name = name
}

// WithProject adds the project to the data source service get data source params
func (o *DataSourceServiceGetDataSourceParams) WithProject(project *string) *DataSourceServiceGetDataSourceParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the data source service get data source params
func (o *DataSourceServiceGetDataSourceParams) SetProject(project *string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *DataSourceServiceGetDataSourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

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
