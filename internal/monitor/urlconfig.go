package monitor

// URLConfig representa la configuración de una URL para el monitoreo
type URLConfig struct {
	URL            string // URL a monitorear
	Timeout        int    // Tiempo de espera en segundos para la solicitud HTTP
	ExpectedStatus []int  // Códigos de estado esperados
}
