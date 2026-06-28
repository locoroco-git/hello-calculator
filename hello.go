import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	{
		// 1. Grab the port from the environment
		port := os.Getenv("PORT")

		// 2. If it is empty, default to 8080 for local development
		if port == "" {
			port = "8080"
		}

		// 3. Bind your server to that port
		fmt.Println("Server is running on port:", port)
		http.ListenAndServe(":"+port, nil)
	}

	fmt.Println("Hello World")
}
