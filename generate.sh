#!/bin/bash

# Script to generate Go types and server interfaces from OpenAPI specification

set -e

echo "Generating OpenAPI types and interfaces..."

# Sletter generated mappen og lager eny tom generated mappe
rm -r generated
mkdir -p generated

# Generate types (DTOs)
echo "Generating types..."
go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest \
  --package=generated \
  --generate=types \
  api/openapi.yaml > generated/openapi_types.go

# Generate main Chi server interface (for backwards compatibility)
echo "Generating main Chi server interface..."
go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest \
  --package=generated \
  --generate=chi-server \
  api/openapi.yaml > generated/openapi_chi_server.go

# Generate domain-specific handler interfaces automatically
echo "Generating domain-specific handler interfaces..."

# Extract tags from OpenAPI spec
TAGS=$(grep -A1 "tags:" api/openapi.yaml | grep -E "^\s*-\s+" | sed 's/^[[:space:]]*-[[:space:]]*//' | sort | uniq)

for tag in $TAGS; do
    echo "Generating ${tag} handler interface..."
    
    # Create separate directory for each domain
    mkdir -p "generated/${tag}"
    
    # Generate interface for this tag in its own package
    go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest \
      --package="${tag}" \
      --generate=chi-server \
      --include-tags="${tag}" \
      --o="generated/${tag}/handler.go" \
      api/openapi.yaml
done

echo "Generation complete!"
echo "Generated files:"
echo "  - generated/openapi_types.go (DTOs and request/response types)"
echo "  - generated/openapi_chi_server.go (Main Chi server interface)"

# List generated tag-specific interfaces
for tag in $TAGS; do
    echo "  - generated/${tag}/handler.go (${tag} domain interface)"
done
