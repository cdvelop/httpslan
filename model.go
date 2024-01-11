package httpslan

type httpsLan struct {
	app_domain         string
	certs_dir          string
	publicCertFilePath string
	privateKeyFilePath string

	built_cert bool
}
