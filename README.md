# Building a Simple Products API in Go – A Beginner’s Guide

## 1. Title & Objective

### Title: 
Getting Started with Go – Building a Simple E-Commerce Products API
### Objective:
This toolkit documents my learning journery with G0(Golang) and demonstrates how to create a minimal, functional Products API for an e-commerce system.
### Why I Chose Go:
I chose Go because it is fast, simple, and beginner-friendly, making it ideal for building efficient backend services. It is also one of the most widely used languages in DevOps and cloud engineering, which aligns with my current career path. Since I recently transitioned into DevOps, I wanted to explore how Go compares to other languages commonly used in the field especially Python both in terms of performance and ease of building scalable services. 

### End Goal:
Build and run a simple API endpoint that returns a list of products in JSON format.

## 2. Quick Summary of the Technology
Go programming language is an open-source, compiled programming language created by Google. It's a fast, statistically compiled language that feels like a dynamically typed interpreted language. It also excels in handling concurrent tasks through its lightweight goroutines and channels.
### Where it is used:
#### Go is popular for 
    Cloud & Network Services,
    DevOps & Site Reliability Engineering (SRE) 
    Web Development (Backend) 
    Distributed Systems
    Command-Line Tools
    Data Science & AI (to some extent)
### Real-WOrld Example:
Netflix utilizes Go in parts of its server architecture.

## 3. System Requirements
### Operating Systems:
    Windows(via WSL)
    Linux
    MacOs
### Tools/Editors:
    Visucal Studio Code
    Git
    Terminal access

### Required Packages
    Go 1.22+


### 4. Installation and Setup Instructions 
#### Step 1 - Download Go
##### Go to the official website:
      👉   https://go.dev/dl/
#### Step 2 — Install Go 
##### Assuming your file is in ~/Dowloads:

      cd ~/Downloads
      sudo rm -rf /usr/local/go
      sudo tar -C /usr/local -xzf go*.tar.gz
#### Step 3 — Add Go to PATH
     echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.bashrc
     source ~/.bashrc
#### Step 4 — Verify
     go version
#### Expected output:
     go version go1.22.x linux/amd64
#### Create workspace
     mkdir -p ~/go-projects/products-api
     cd ~/go-projects/products-api
     go mod init products-api

### Minimal Working Example
This example creates a simple HTTP server with a /products endpoint.
        package main

        import (
            "encoding/json"
            "net/http"
        )

        type Product struct {
            ID    int     `json:"id"`
            Name  string  `json:"name"`
            Price float64 `json:"price"`
        }

        func productsHandler(w http.ResponseWriter, r *http.Request) {
            products := []Product{
                {ID: 1, Name: "Laptop", Price: 1200.50},
                {ID: 2, Name: "Headphones", Price: 150.99},
                {ID: 3, Name: "Keyboard", Price: 89.90},
            }

            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(products)
        }

        func main() {
            http.HandleFunc("/products", productsHandler)
            http.ListenAndServe(":8080", nil)
        }
What the API does
Starts a server on port 8080
##### Run the Go file
      go run main.go
###### Returns JSON list of products when visiting:
       👉 http://localhost:8080/products

### Expected Output
        [
            {"id":1,"name":"Laptop","price":1200.5},
            {"id":2,"name":"Headphones","price":150.99},
            {"id":3,"name":"Keyboard","price":89.9}
        ]

## 6. AI Prompt Journal 
Below are the prompts I used during development and what I learned from them.
### Prompt used:
    Create a simple Products API using Go that returns a list of products in JSON format
### Link to the Curriculum for the Prompt
    https://training.moringaschool.com/courses/2/pages/learning-a-new-programming-language-with-ai?module_item_id=66
### AI Response Summary
    The AI provided a complete example of a Go HTTP server with a /products endpoint, explained how to define structs, encode JSON, and start the server
### Helpful Part of the Response
    The AI helped me scaffold a working API and explained how to structure the Product struct and JSON response.
### Evaluation of Helpfulness
    Very helpful — it gave me a working API quickly and helped me understand Go’s HTTP handling and JSON encoding.

## 7. Common Issues & Fixes
### Issue: 
    go: command not found
##### Cause: 
      PATH not correctly updated.
       Fix:
          export PATH=$PATH:/usr/local/go/bin
          source ~/.bashrc
### Error
JSON encoding errors
##### Cause:
      Unexported fields in struct.
      Fix:
      Ensure struct fields start with capital letters.
      type Product struct {
            ID int `json:"id"`   // correct
        }

## 8. References
### Official Documentation

    Go Documentation → https://go.dev/doc/

    Go Packages → https://pkg.go.dev/

    Go by Example → https://gobyexample.com/

### Video Tutorials

    “Go Crash Course” – Traversy Media

    “Learn Go in 1 Hour” – freeCodeCamp

### Extra Learning

    https://tutorialedge.net/golang

    https://golangbridge.org

