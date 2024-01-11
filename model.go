package httpslan

type httpsLan struct {
	app_domain         string
	certs_dir          string
	publicCertFilePath string
	privateKeyFilePath string

	built_cert bool

	schedule
}

type schedule interface {
	AddFuncToSchedule(schedule, description string, fun any, args ...any) (err string)
}
