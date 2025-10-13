#!/bin/bash

# PayGateway Setup Script
# Description: Complete project setup and initialization
# Usage: ./setup.sh [--env development|production]

set -euo pipefail

# Configuration
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Default values
ENVIRONMENT="development"
VERBOSE=false

# Logging functions
log_info() {
    echo -e "${GREEN}[INFO]${NC} $1" >&2
}

log_warn() {
    echo -e "${YELLOW}[WARN]${NC} $1" >&2
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1" >&2
}

echo -e "${GREEN}"
echo "╔══════════════════════════════════════════╗"
echo "║          PayGateway Setup Script         ║"
echo "╚══════════════════════════════════════════╝"
echo -e "${NC}"

log_info "Starting project setup..."

# Copying environment files
log_info "Setting up environment files..."
if [[ -f ".env.example" ]]; then
    cp .env.example .env
    log_info "✓ Copied .env.example to .env"
fi

if [[ -f ".env.example.development" ]]; then
    cp .env.example.development .env.development
    log_info "✓ Copied .env.example.development to .env.development"
fi

# Installing npm dependencies
log_info "Installing Node.js dependencies..."
npm install
log_info "✓ Root dependencies installed"

# Installing frontend dependencies
if [[ -d "frontend" ]]; then
    log_info "Installing frontend dependencies..."
    cd frontend && npm install && cd ..
    log_info "✓ Frontend dependencies installed"
fi

# Initializing Husky
log_info "Setting up Git hooks with Husky..."
if command -v npx &> /dev/null; then
    npx husky init
    echo "npx lint-staged" > .husky/pre-commit
    log_info "✓ Husky initialized with pre-commit hooks"
else
    log_warn "npx not found, skipping Husky setup"
fi

# Setup Docker network
log_info "Setting up Docker network..."
if command -v docker &> /dev/null; then
    if ! docker network ls | grep -q "paygateway-network"; then
        docker network create paygateway-network 2>/dev/null || true
        log_info "✓ Created Docker network: paygateway-network"
    else
        log_info "✓ Docker network 'paygateway-network' already exists"
    fi
else
    log_warn "Docker not found, skipping network setup"
fi

# Make scripts executable
log_info "Setting up script permissions..."
if [[ -d "scripts" ]]; then
    chmod +x scripts/*.sh 2>/dev/null || true
    log_info "✓ Made scripts executable"
fi

echo
log_info "Setup completed successfully! 🎉"
echo
echo -e "${GREEN}Next steps:${NC}"
echo "1. Review and customize environment variables in .env files"
echo "2. Start the development environment:"
echo -e "   ${BLUE}docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d${NC}"
echo "3. Or start specific service:"
echo -e "   ${BLUE}docker-compose -f docker-compose.yml -f docker-compose.dev.yml up rust-processor${NC}"
echo "4. Access the services:"
echo "   • Frontend: http://localhost:8081"
echo "   • Go API: http://localhost:8080"
echo "   • Swagger Docs: http://localhost:8080/swagger/index.html"
echo
echo -e "${GREEN}Happy coding! 🚀${NC}"