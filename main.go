package main

import (
	"html/template"
	"log" // Agregado para manejar logs
	"net"
	"net/http"
	"os" // Agregado para manejar variables de entorno
)

type PageData struct {
	LocalIP string
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "Unknown"
	}
	for _, addr := range addrs {
		// Verifica si la dirección es una IP y no un localhost
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return "Unknown"
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Obtén la IP local
	ip := getLocalIP()

	// Define la estructura de datos para la plantilla
	data := PageData{
		LocalIP: ip,
	}

	// Carga y ejecuta la plantilla
	tmpl := template.Must(template.ParseFiles("index.html"))
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error ejecutando la plantilla: %v\n", err) // Agregado para manejar errores de ejecución
	}
}

func main() {
	port := os.Getenv("PORT") // Agregado para obtener el puerto desde una variable de entorno
	if port == "" {
		port = "8080"                                     // Establece un puerto predeterminado si no está configurado
		log.Println("Usando puerto predeterminado: 8080") // Agregado para mostrar el puerto utilizado
	}

	http.HandleFunc("/", handler)

	log.Printf("Servidor iniciado en puerto %s\n", port) // Agregado para loguear el inicio del servidor
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Error iniciando el servidor: %v\n", err) // Agregado para manejar errores al iniciar el servidor
	}
}
