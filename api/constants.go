package api

const (
	// Status constants denotes the status of the monitors.
	Down               Status = 0
	Up                 Status = 1
	Trouble            Status = 2
	Critical           Status = 3
	Suspended          Status = 5
	Maintenance        Status = 7
	Discovery          Status = 9
	ConfigurationError Status = 10

	// Type of the Site24x7 resource.
	URL      MonitorType = "URL"
	SSL_CERT MonitorType = "SSL_CERT"
	RESTAPI  MonitorType = "RESTAPI"
	AMAZON   MonitorType = "AMAZON"
)