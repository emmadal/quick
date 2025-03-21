// Package quick provides a minimalistic and high-performance web framework for Go.
//
// This file contains example implementations demonstrating different functionalities
// of the Quick framework, including route handling, middleware usage, and configuration management.
package quick

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/jeffotoni/quick/middleware/cors"
)

// This function is named ExampleGetDefaultConfig()
// it with the Examples type.
func ExampleGetDefaultConfig() {
	// Get the default configuration settings
	result := GetDefaultConfig()

	// Print individual configuration values
	fmt.Printf("BodyLimit: %d\n", result.BodyLimit)           // Maximum request body size
	fmt.Printf("MaxBodySize: %d\n", result.MaxBodySize)       // Maximum allowed body size for requests
	fmt.Printf("MaxHeaderBytes: %d\n", result.MaxHeaderBytes) // Maximum size for request headers
	fmt.Printf("RouteCapacity: %d\n", result.RouteCapacity)   // Maximum number of registered routes
	fmt.Printf("MoreRequests: %d\n", result.MoreRequests)     // Maximum concurrent requests allowed

	// Print the entire configuration struct
	fmt.Println(result)

	// Out put: BodyLimit: 2097152, MaxBodySize: 2097152, MaxHeaderBytes: 1048576, RouteCapacity: 1000, MoreRequests: 290

}

// This function is named ExampleNew()
// it with the Examples type.
func ExampleNew() {
	// Start Quick instance
	q := New()

	// Define a simple GET route
	q.Get("/", func(c *Ctx) error {
		// Set response header
		c.Set("Content-Type", "text/plain")

		// Return a text response
		return c.Status(200).String("Quick in action ❤️!")
	})

	// Simulate a request to the defined route for testing
	res, _ := q.QuickTest("GET", "/", nil)
	fmt.Println(res.BodyStr())

	// Out put: Quick in action ❤️!

}

// This function is named ExampleQuick_Use()
// it with the Examples type.
func ExampleQuick_Use() {
	// Start Quick instance
	q := New()

	// Apply CORS middleware to allow cross-origin requests
	q.Use(cors.New())

	// Define a route that will be affected by the middleware
	q.Get("/use", func(c *Ctx) error {
		// Set response header
		c.Set("Content-Type", "text/plain")

		// Return response with middleware applied
		return c.Status(200).String("Quick in action com middleware ❤️!")
	})

	// Simulate a request for testing
	res, _ := q.QuickTest("GET", "/use", nil)
	fmt.Println(res.BodyStr())

	// Out put: Quick in action com middleware ❤️!

}

// This function is named ExampleQuick_Get()
// it with the Examples type.
func ExampleQuick_Get() {
	// Start Quick instance
	q := New()

	// Define a GET route with a handler function
	q.Get("/hello", func(c *Ctx) error {
		// Return a simple text response
		return c.Status(200).String("Olá, mundo!")
	})

	// Simulate a GET request to the route
	res, _ := q.QuickTest("GET", "/hello", nil)
	fmt.Println(res.BodyStr())

	// Out put: Olá, mundo!
}

// This function is named ExampleQuick_Post()
// it with the Examples type.
func ExampleQuick_Post() {
	// Start Quick instance
	q := New()

	// Define a POST route
	q.Post("/create", func(c *Ctx) error {
		// Return response indicating resource creation
		return c.Status(201).String("Recurso criado!")
	})

	// Simulate a POST request for testing
	res, _ := q.QuickTest("POST", "/create", nil)
	fmt.Println(res.BodyStr())

	// Out put: Recurso criado!
}

// This function is named ExampleQuick_Put()
// it with the Examples type.
func ExampleQuick_Put() {
	// Start Quick instance
	q := New()

	// Define a PUT route
	q.Put("/update", func(c *Ctx) error {
		// Return response indicating resource update
		return c.Status(200).String("Recurso atualizado!")
	})

	// Simulate a PUT request for testing
	res, _ := q.QuickTest("PUT", "/update", nil)

	fmt.Println(res.BodyStr())

	// Out put: Recurso atualizado!
}

// This function is named ExampleQuick_Delete()
// it with the Examples type.
func ExampleQuick_Delete() {
	// Start Quick instance
	q := New()

	// Define a DELETE route
	q.Delete("/delete", func(c *Ctx) error {
		// Return response indicating resource deletion
		return c.Status(200).String("Recurso deletado!")
	})

	// Simulate a DELETE request for testing
	res, _ := q.QuickTest("DELETE", "/delete", nil)

	fmt.Println(res.BodyStr())

	// Out put: Recurso deletado!
}

// This function is named ExampleQuick_ServeHTTP()
// it with the Examples type.
func ExampleQuick_ServeHTTP() {
	// Start Quick instance
	q := New()

	// Define a route with a dynamic parameter
	q.Get("/users/:id", func(c *Ctx) error {
		// Retrieve the parameter and return it in the response
		return c.Status(200).String("User Id: " + c.Params["id"])
	})

	// Simulate a request with a user ID
	res, _ := q.QuickTest("GET", "/users/42", nil)

	// Print the response status and body
	fmt.Println(res.StatusCode())
	fmt.Println(res.BodyStr())

	// Out put:	200, 42
}

// This function is named ExampleQuick_GetRoute()
// it with the Examples type.
func ExampleQuick_GetRoute() {
	// Start Quick instance
	q := New()

	// Define multiple routes
	q.Get("/users/:id", func(c *Ctx) error {
		return c.Status(200).String("User ID: " + c.Params["id"])
	})
	q.Post("/users", func(c *Ctx) error {
		return c.Status(201).String("User created")
	})

	// Get a list of all registered routes
	routes := q.GetRoute()

	// Print the total number of routes
	fmt.Println(len(routes))

	// Iterate over the routes and print their method and pattern
	for _, route := range routes {
		fmt.Println(route.Method, route.Pattern)
	}

	// Out put: 2, GET /users/:id, POST /users
}

// This function is named ExampleQuick_Listen()
// it with the Examples type.
func ExampleQuick_Listen() {
	// Start Quick instance
	q := New()

	// Define a simple route
	q.Get("/", func(c *Ctx) error {
		return c.Status(200).String("Hello, Quick!")
	})

	// Start the server and listen on port 8080
	err := q.Listen(":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
	// Out put:
	// (This function starts a server and does not return an output directly)
}

// This function is named ExampleQuick_Patch()
//
//	it with the Examples type.
func ExampleQuick_Patch() {
	// Start Quick instance
	q := New()

	// Define a PATCH route
	q.Patch("/update", func(c *Ctx) error {
		return c.Status(200).String("PATCH request received")
	})

	// Simulate a PATCH request to "/update"
	res, _ := q.QuickTest("PATCH", "/update", nil)

	// Print the response status and body
	fmt.Println("Status:", res.StatusCode())

	// Out put: Status: 200
}

// This function is named ExampleQuick_Options()
//	it with the Examples type.

func ExampleQuick_Options() {
	// Start Quick instance
	q := New()

	// Define an OPTIONS route
	q.Options("/resource", func(c *Ctx) error {
		c.Set("Allow", "GET, POST, OPTIONS")
		return c.Status(204).Send(nil) // No Content response
	})

	// Simulate an OPTIONS request to "/resource"
	res, _ := q.QuickTest("OPTIONS", "/resource", nil)

	// Print the response status
	fmt.Println("Status:", res.StatusCode())

	// Out put: Status: 204
}

// This function is named ExampleQuick_Static()
//
//	it with the Examples type.
func ExampleQuick_Static() {
	// Quick Start
	q := New()

	// Uncomment this block if using embedded static files
	/*
		//go:embed static/*
		var staticFiles embed.FS
		q.Static("/static", staticFiles)
	*/

	// Serve static files from the "static" directory
	q.Static("/static", "./static")

	// Serve a specific file on the root path
	q.Get("/", func(c *Ctx) error {
		c.File("./static/index.html")
		return nil
	})

	// Simulate a request to a static file (e.g., "/static/index.html")
	res, _ := q.QuickTest("GET", "/static/index.html", nil)

	// Print the response status
	fmt.Println("Status:", res.StatusCode())

	// Out put: Status: 200
}

// This function is named ExampleQuick_Shutdown()
//
//	it with the Examples type.
func ExampleQuick_Shutdown() {
	// Create a new Quick instance
	q := New()

	// Define a GET route with a handler function
	q.Get("/", func(c *Ctx) error {
		// Return a simple text response
		return c.SendString("Server is running!")
	})

	// Simulate a GET request to the route and capture the response
	resp, _ := q.QuickTest("GET", "/", nil)

	// Print only the response body to match GoDoc expectations
	fmt.Println(resp.BodyStr())

	// Simulate server shutdown immediately after starting
	err := q.Shutdown()

	// Print the shutdown status
	if err == nil {
		fmt.Println("Server shut down successfully.")
	} else {
		fmt.Println("Error shutting down server:", err)
	}

	// Out put:
	// Server is running!
	// Server shut down gracefully.
}

// This function is named ExampleMaxBytesReader()
//
//	it with the Examples type.
func ExampleMaxBytesReader() {
	// Start Quick framework instance
	q := New()

	// Set max request body size to 1KB (1024 bytes)
	const maxBodySize = 1024

	// Simulate a request using Quick's test utility with a payload exceeding 1KB
	oversizedBody := make([]byte, 2048) // 2KB of data to exceed the limit
	for i := range oversizedBody {
		oversizedBody[i] = 'A'
	}

	// Define a route with MaxBytesReader for extra validation
	q.Post("/v1/user/maxbody/max", func(c *Ctx) error {
		c.Set("Content-Type", "application/json")
		// Apply MaxBytesReader for additional size enforcement
		c.Request.Body = MaxBytesReader(c.Response, c.Request.Body, maxBodySize)

		// Read request body
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			log.Printf("Error reading request body: %v", err)
			return c.Status(http.StatusRequestEntityTooLarge).String("Request body too large")
		}

		return c.Status(http.StatusOK).Send(body)
	})

	// Simulate a request using Quick's test utility
	res, _ := q.Qtest(QuickTestOptions{
		Method:  MethodPost,
		URI:     "/v1/user/maxbody/max",
		Headers: map[string]string{"Content-Type": "application/json"},
		Body:    oversizedBody, // Convert string to []byte
	})

	// Print response status and body for verification
	fmt.Println(res.StatusCode()) // Expecting: 413 (Payload too large)
	fmt.Println(res.BodyStr())    // Expecting: "Request body too large"

	// Out put:
	// 413
	// Request body too large
}
