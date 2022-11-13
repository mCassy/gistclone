#!/bin/bash
echo "Generating swagger docs..."
swag init --parseDependency --parseInternal
echo "Building app..."
go build