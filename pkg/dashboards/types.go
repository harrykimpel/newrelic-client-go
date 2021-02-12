// Code generated by tutone: DO NOT EDIT
package dashboards

import (
	"github.com/harrykimpel/newrelic-client-go/pkg/entities"
	"github.com/harrykimpnewrelic-client-go/pkg/nrdb"
	"github.com/harrykimp/pkg/nrtime"
)

// DashboardCreateErrorType - List of expected error types that can be thrown by a dashboard create operation
type DashboardCreateErrorType string

var DashboardCreateErrorTypeTypes = struct {
	// Invalid input error
	INVALID_INPUT DashboardCreateErrorType
}{
	// Invalid input error
	INVALID_INPUT: "INVALID_INPUT",
}

// DashboardDeleteErrorType - Expected error type that can be thrown by a Dashboard delete operation
type DashboardDeleteErrorType string

var DashboardDeleteErrorTypeTypes = struct {
	// Dashboard not found in the system
	DASHBOARD_NOT_FOUND DashboardDeleteErrorType
	// User is not allowed to execute the operation
	FORBIDDEN_OPERATION DashboardDeleteErrorType
}{
	// Dashboard not found in the system
	DASHBOARD_NOT_FOUND: "DASHBOARD_NOT_FOUND",
	// User is not allowed to execute the operation
	FORBIDDEN_OPERATION: "FORBIDDEN_OPERATION",
}

// DashboardDeleteResultStatus - Result status from a Dashboard delete operation.
type DashboardDeleteResultStatus string

var DashboardDeleteResultStatusTypes = struct {
	// FAILURE.
	FAILURE DashboardDeleteResultStatus
	// SUCCESS.
	SUCCESS DashboardDeleteResultStatus
}{
	// FAILURE.
	FAILURE: "FAILURE",
	// SUCCESS.
	SUCCESS: "SUCCESS",
}

// DashboardUpdateErrorType - List of expected error types that can be thrown by a dashboard update operation
type DashboardUpdateErrorType string

var DashboardUpdateErrorTypeTypes = struct {
	// User is not allowed to execute the operation
	FORBIDDEN_OPERATION DashboardUpdateErrorType
	// Invalid input error
	INVALID_INPUT DashboardUpdateErrorType
}{
	// User is not allowed to execute the operation
	FORBIDDEN_OPERATION: "FORBIDDEN_OPERATION",
	// Invalid input error
	INVALID_INPUT: "INVALID_INPUT",
}

// DashboardAreaWidgetConfigurationInput - Configuration for visualization type 'viz.area'
type DashboardAreaWidgetConfigurationInput struct {
	// nrql queries
	NRQLQueries []DashboardWidgetNRQLQueryInput `json:"nrqlQueries,omitempty"`
}

// DashboardBarWidgetConfigurationInput - Configuration for visualization type 'viz.bar'
type DashboardBarWidgetConfigurationInput struct {
	// nrql queries
	NRQLQueries []DashboardWidgetNRQLQueryInput `json:"nrqlQueries,omitempty"`
}

// DashboardBillboardWidgetConfigurationInput - Configuration for visualization type 'viz.billboard'
type DashboardBillboardWidgetConfigurationInput struct {
	// nrql queries
	NRQLQueries []DashboardWidgetNRQLQueryInput `json:"nrqlQueries,omitempty"`
	// thresholds
	Thresholds []DashboardBillboardWidgetThresholdInput `json:"thresholds,omitempty"`
}

// DashboardBillboardWidgetThresholdInput - Billboard widget threshold input.
type DashboardBillboardWidgetThresholdInput struct {
	// alert severity.
	AlertSeverity entities.DashboardAlertSeverity `json:"alertSeverity,omitempty"`
	// value.
	Value float64 `json:"value,omitempty"`
}

// DashboardCreateError - Error information to add as first citizen in create operations
type DashboardCreateError struct {
	// Error description
	Description string `json:"description,omitempty"`
	// Error type
	Type DashboardCreateErrorType `json:"type"`
}

// DashboardCreateResult - Create mutation results
type DashboardCreateResult struct {
	// Dashboard creation result
	EntityResult DashboardEntityResult `json:"entityResult,omitempty"`
	// Expected errors while processing request
	Errors []DashboardCreateError `json:"errors,omitempty"`
}

// DashboardDeleteError - Expected error that can be thrown by a Dashboard delete operation
type DashboardDeleteError struct {
	// Error description
	Description string `json:"description,omitempty"`
	// Error type
	Type DashboardDeleteErrorType `json:"type"`
}

// DashboardDeleteResult - Result from a Dashboard delete operation.
type DashboardDeleteResult struct {
	// Expected errors while processing request
	Errors []DashboardDeleteError `json:"errors,omitempty"`
	// The status of the attempted delete.
	Status DashboardDeleteResultStatus `json:"status,omitempty"`
}

// DashboardEntityResult - Public schema - `DashboardEntity` result representation for mutations. It's a subset of the `DashboardEntity` that inherits from the Entity type, but a complete different type.
type DashboardEntityResult struct {
	// Account ID.
	AccountID int `json:"accountId,omitempty"`
	// Dashboard creation timestamp.
	CreatedAt nrtime.DateTime `json:"createdAt,omitempty"`
	// Dashboard description.
	Description string `json:"description,omitempty"`
	// Unique entity identifier.
	GUID entities.EntityGUID `json:"guid,omitempty"`
	// Dashboard name.
	Name string `json:"name,omitempty"`
	// Dashboard owner
	Owner entities.DashboardOwnerInfo `json:"owner,omitempty"`
	// Dashboard pages.
	Pages []entities.DashboardPage `json:"pages,omitempty"`
	// Dashboard permissions configuration.
	Permissions entities.DashboardPermissions `json:"permissions,omitempty"`
	// Dashboard update timestamp.
	UpdatedAt nrtime.DateTime `json:"updatedAt,omitempty"`
}

// DashboardInput - Dashboard input
type DashboardInput struct {
	// Dashboard description.
	Description string `json:"description,omitempty"`
	// Dashboard name.
	Name string `json:"name"`
	// Dashboard page input.
	Pages []DashboardPageInput `json:"pages,omitempty"`
	// Dashboard permissions configuration.
	Permissions entities.DashboardPermissions `json:"permissions"`
}

// DashboardLineWidgetConfigurationInput - Configuration for visualization type 'viz.line'
type DashboardLineWidgetConfigurationInput struct {
	// nrql queries
	NRQLQueries []DashboardWidgetNRQLQueryInput `json:"nrqlQueries,omitempty"`
}

// DashboardMarkdownWidgetConfigurationInput - Configuration for visualization type 'viz.markdown'
type DashboardMarkdownWidgetConfigurationInput struct {
	// Markdown content of the widget
	Text string `json:"text"`
}

// DashboardPageInput - Page input
type DashboardPageInput struct {
	// Page description.
	Description string `json:"description,omitempty"`
	// Unique entity identifier of the Page to be updated. When null, it means a new Page will be created.
	GUID entities.EntityGUID `json:"guid,omitempty"`
	// Page name.
	Name string `json:"name"`
	// Page widgets.
	Widgets []DashboardWidgetInput `json:"widgets,omitempty"`
}

// DashboardPieWidgetConfigurationInput - Configuration for visualization type 'viz.pie'
type DashboardPieWidgetConfigurationInput struct {
	// nrql queries
	NRQLQueries []DashboardWidgetNRQLQueryInput `json:"nrqlQueries,omitempty"`
}

// DashboardTableWidgetConfigurationInput - Configuration for visualization type 'viz.table'
type DashboardTableWidgetConfigurationInput struct {
	// nrql queries
	NRQLQueries []DashboardWidgetNRQLQueryInput `json:"nrqlQueries,omitempty"`
}

// DashboardUpdateError - Error information to add as first citizen in update operations
type DashboardUpdateError struct {
	// Error description
	Description string `json:"description,omitempty"`
	// Error type
	Type DashboardUpdateErrorType `json:"type"`
}

// DashboardUpdateResult - Update mutation results
type DashboardUpdateResult struct {
	// Dashboard update result
	EntityResult DashboardEntityResult `json:"entityResult,omitempty"`
	// Expected errors while processing request
	Errors []DashboardUpdateError `json:"errors,omitempty"`
}

// DashboardWidgetConfigurationInput - Typed configuration for known visualizations. At most one may be populated.
type DashboardWidgetConfigurationInput struct {
	// Configuration for visualization type 'viz.area'
	Area *DashboardAreaWidgetConfigurationInput `json:"area,omitempty"`
	// Configuration for visualization type 'viz.bar'
	Bar *DashboardBarWidgetConfigurationInput `json:"bar,omitempty"`
	// Configuration for visualization type 'viz.billboard'
	Billboard *DashboardBillboardWidgetConfigurationInput `json:"billboard,omitempty"`
	// Configuration for visualization type 'viz.line'
	Line *DashboardLineWidgetConfigurationInput `json:"line,omitempty"`
	// Configuration for visualization type 'viz.markdown'
	Markdown *DashboardMarkdownWidgetConfigurationInput `json:"markdown,omitempty"`
	// Configuration for visualization type 'viz.pie'
	Pie *DashboardPieWidgetConfigurationInput `json:"pie,omitempty"`
	// Configuration for visualization type 'viz.table'
	Table *DashboardTableWidgetConfigurationInput `json:"table,omitempty"`
}

// DashboardWidgetInput - Widget Input
type DashboardWidgetInput struct {
	// Typed configuration for the widget
	Configuration DashboardWidgetConfigurationInput `json:"configuration,omitempty"`
	// id.
	ID string `json:"id,omitempty"`
	// layout
	Layout DashboardWidgetLayoutInput `json:"layout,omitempty"`
	// Related entities. Currently only supports Dashboard entities, but may allow other cases in the future.
	LinkedEntityGUIDs []entities.EntityGUID `json:"linkedEntityGuids"`
	// Untyped scalar of configuration for the widget
	RawConfiguration entities.DashboardWidgetRawConfiguration `json:"rawConfiguration,omitempty"`
	// title
	Title string `json:"title,omitempty"`
	// Specifies how this widget will be visualized. If null, the WidgetConfigurationInput will be used to determine the visualization.
	Visualization DashboardWidgetVisualizationInput `json:"visualization,omitempty"`
}

// DashboardWidgetLayoutInput - Widget layout input
type DashboardWidgetLayoutInput struct {
	// column.
	Column int `json:"column,omitempty"`
	// height.
	Height int `json:"height,omitempty"`
	// row.
	Row int `json:"row,omitempty"`
	// width.
	Width int `json:"width,omitempty"`
}

// DashboardWidgetNRQLQueryInput -
type DashboardWidgetNRQLQueryInput struct {
	// accountId
	AccountID int `json:"accountId"`
	// NRQL formatted query
	Query nrdb.NRQL `json:"query"`
}

// DashboardWidgetVisualizationInput - visualization configuration
type DashboardWidgetVisualizationInput struct {
	// Nerdpack artifact ID
	ID string `json:"id,omitempty"`
}

// Float - The `Float` scalar type represents signed double-precision fractional
// values as specified by
// [IEEE 754](http://en.wikipedia.org/wiki/IEEE_floating_point).
type Float string
