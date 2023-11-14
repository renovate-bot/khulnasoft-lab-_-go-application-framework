package configuration

const (
	API_URL                         string = "vulnmap_api"                   // AKA "endpoint" in the config file
	AUTHENTICATION_SUBDOMAINS       string = "internal_auth_subdomain"       // array of additional subdomains to add authentication for
	AUTHENTICATION_ADDITIONAL_URLS  string = "internal_additional_auth_urls" // array of additional urls to add authentication for
	AUTHENTICATION_TOKEN            string = "vulnmap_token"
	AUTHENTICATION_BEARER_TOKEN     string = "vulnmap_oauth_token"
	INTEGRATION_NAME                string = "vulnmap_integration_name"
	INTEGRATION_VERSION             string = "vulnmap_integration_version"
	INTEGRATION_ENVIRONMENT         string = "vulnmap_integration_environment"
	INTEGRATION_ENVIRONMENT_VERSION string = "vulnmap_integration_environment_version"
	ANALYTICS_DISABLED              string = "vulnmap_disable_analytics"
	ORGANIZATION                    string = "org"
	DEBUG                           string = "debug"
	DEBUG_FORMAT                    string = "debug_format"
	INSECURE_HTTPS                  string = "insecure"
	PROXY_AUTHENTICATION_MECHANISM  string = "proxy_auth"
	CACHE_PATH                      string = "vulnmap_cache_path"
	WORKFLOW_USE_STDIO              string = "internal_wflstdio"
	RAW_CMD_ARGS                    string = "internal_raw_cmd_args"
	WEB_APP_URL                     string = "internal_vulnmap_app"
	INPUT_DIRECTORY                 string = "targetDirectory"
	ADD_TRUSTED_CA_FILE             string = "internal_additional_trusted_ca_file" // pem file location containing additional CAs to trust
	FIPS_ENABLED                    string = "gofips"
	WORKING_DIRECTORY               string = "internal_working_dir"
	IS_FEDRAMP                      string = "internal_is_fedramp"
	UNKNOWN_ARGS                    string = "internal_unknown_arguments" // arguments unknown to the current application but maybe relevant for delegated application calls

	// feature flags
	FF_OAUTH_AUTH_FLOW_ENABLED string = "internal_vulnmap_oauth_enabled"
)
