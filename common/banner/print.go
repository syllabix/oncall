package banner

import "fmt"

// Print the application banner to the console
func Print() {
	fmt.Printf(`
 ██████╗ ███╗   ██╗     ██████╗ █████╗ ██╗     ██╗         
██╔═══██╗████╗  ██║    ██╔════╝██╔══██╗██║     ██║         
██║   ██║██╔██╗ ██║    ██║     ███████║██║     ██║         
██║   ██║██║╚██╗██║    ██║     ██╔══██║██║     ██║         
╚██████╔╝██║ ╚████║    ╚██████╗██║  ██║███████╗███████╗    
 ╚═════╝ ╚═╝  ╚═══╝     ╚═════╝╚═╝  ╚═╝╚══════╝╚══════╝    															   
simple incident response management
alpha version - v0.0.1

`)
}
