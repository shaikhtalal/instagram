SHELL := /bin/bash -o pipefail


.PHONY: help

help:
	@echo "Usage: make <TARGET>"
	@echo ""
	@echo "Available targets are:"
	@echo ""
	@echo "    backend                         Run the backend service"
	@echo "    frontend                        Run the frontend service"
	@echo ""


.PHONY: backend
backend:
	@echo "starting backend service!"
	@go run main.go serve 


.PHONY: frontend
frontend:
	@echo "starting frontend service!"
	@cd frontend && npm install && npm run serve 