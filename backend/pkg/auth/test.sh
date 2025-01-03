#! /bin/bash

# Test login request
curl -X POST http://localhost:8000/login -H "Content-Type: application/json" -d '{"name": "John", "password": "test123"}'
