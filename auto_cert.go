package httpslan

// func START() {

// certManager := autocert.Manager{
// 	Prompt: autocert.AcceptTOS,
// 	Cache:  autocert.DirCache("certs"),
// }

// var mode string
// var port = "443"
// if s.AppInProduction() {
// 	mode = "Produccion"
// } else {
// 	port = s.Port()
// 	mode = "Desarrollo"
// }

// srv := &http.Server{
// 	Addr:    ":" + port,
// 	Handler: mux, //conexion

// 	TLSConfig: &tls.Config{
// 		Rand:           nil,
// 		GetCertificate: certManager.GetCertificate,
// 	},

// 	// ReadTimeout es la duración máxima para leer la solicitud completa, incluido el cuerpo.
// 	ReadTimeout: 0, // 1 * time.Minute, 0 no hay tiempo de espera
// 	// WriteTimeout es la duración máxima antes de que se agoten las escrituras de la respuesta. Se restablece cada vez que se lee el encabezado de una nueva solicitud. Al igual que ReadTimeout, no permite que los controladores tomen decisiones por solicitud.
// 	WriteTimeout:   30 * time.Minute,
// 	MaxHeaderBytes: 1 << 20,
// }

// fmt.Printf(" *** Plataforma %v ver.%v Modo %v ***\n", s.AppName(), s.AppVersion(), mode)
// fmt.Printf("Servidor Iniciado en Puerto: %v\n", port)

// Vale la pena señalar que certManager.HTTPHandler(nil)redirige todo el tráfico HTTP a HTTPS automáticamente,
// pero esto se puede personalizar reemplazando el parámetro nil con un controlador HTTP.
// go http.ListenAndServe(":80", certManager.HTTPHandler(nil))
// log.Fatal(srv.ListenAndServeTLS("", ""))

// }
