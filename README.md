# Risk-Management-Service
This project is a RESTful web service for managing risks. It is designed to showcase production-quality code, supporting basic CRUD operations for risks.

# Problem Statement

## OVERVIEW
As the Risky Plumbers team works primarily in the creation of web services, this
assignment is intended to allow you to demonstrate your abilities towards this end.
This assignment can be placed in a GitHub repository (please use a GH generic name for
the repo) or e-mailed to the recruiter in a zip file.
Please treat this as if you were writing production quality code.

## TASK
The task is to create a Risk application which:
- Listens on port 8080 for standard HTTP traffic (not HTTPS); and
- Can return a list of Risk given a GET to /v1/risks on the endpoint; and
- Can create a new Risk given a POST to /v1/risks on the endpoint; and
- Can retrieve an individual Risk given a GET to /v1/risks/{id} .

## Frameworks used
1. **go-gin**: A lightweight and fast web framework for building RESTful APIs in Go, with excellent support for routing and middleware.
2. **wire**: A dependency injection tool for Go, used to manage and wire dependencies in a clean and efficient manner.
3. **ctxzap**: A logging library that integrates zap with Go’s context to enable structured and context-aware logging.
4. **ozzo-validation**: A Go library for validating structs, variables, and fields with expressive and reusable validation rules.
5. **enumer**: For enum handling and validations.

