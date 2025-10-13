# 📜 PayGateway Scripts

<div align="center">

**Automation and utility scripts for PayGateway project**

*Collection of helpful scripts for development, deployment, and maintenance*

[![Shell](https://img.shields.io/badge/Shell-Scripts-4EAA25?style=flat-square&logo=gnu-bash&logoColor=white)](https://www.gnu.org/software/bash/)
[![Docker](https://img.shields.io/badge/Docker-Automation-2496ED?style=flat-square&logo=docker&logoColor=white)](https://www.docker.com/)

</div>

## 📋 Overview

This directory contains **automation scripts** and **utility tools** to streamline development, testing, and deployment of the PayGateway system. These scripts help maintain consistency across different environments and reduce manual setup overhead.

---

## 📋 Table of Contents

<details>
<summary><strong>📁 Available Scripts</strong></summary>

### Setup & Installation Scripts

#### 🚀 `setup.sh`
**Purpose:** Complete project setup and initialization

**Features:**
- Environment file configuration
- Dependency installation
- Database initialization
- Docker network setup
- Permission configuration

**Usage:**
```bash
# Make executable and run
chmod +x scripts/setup.sh
./scripts/setup.sh

# Run with specific environment
./scripts/setup.sh --env development
./scripts/setup.sh --env production
```

#### 🔧 `install-deps.sh`
**Purpose:** Install all project dependencies

**Features:**
- Go module download
- Rust crate installation
- Node.js package installation
- System dependency checks

**Usage:**
```bash
./scripts/install-deps.sh
```

### Development Scripts

#### 🏃 `dev-start.sh`
**Purpose:** Start development environment

**Features:**
- Infrastructure services startup
- Database migrations
- Service health checks
- Development server launching

**Usage:**
```bash
# Start all services
./scripts/dev-start.sh

# Start specific service
./scripts/dev-start.sh --service frontend
./scripts/dev-start.sh --service go-api
./scripts/dev-start.sh --service rust-processor
```

#### 🛑 `dev-stop.sh`
**Purpose:** Stop development environment

**Usage:**
```bash
./scripts/dev-stop.sh
```

### Database Scripts

#### 🗄️ `db-migrate.sh`
**Purpose:** Database migration management

**Features:**
- Run pending migrations
- Rollback migrations
- Migration status checking
- Database reset functionality

**Usage:**
```bash
# Run migrations
./scripts/db-migrate.sh up

# Rollback last migration
./scripts/db-migrate.sh down

# Check migration status
./scripts/db-migrate.sh status

# Reset database (development only)
./scripts/db-migrate.sh reset
```

#### 🔄 `db-seed.sh`
**Purpose:** Database seeding with test data

**Usage:**
```bash
./scripts/db-seed.sh
```

### Testing Scripts

#### 🧪 `run-tests.sh`
**Purpose:** Execute comprehensive test suite

**Features:**
- Unit tests for all services
- Integration tests
- End-to-end tests
- Coverage reporting

**Usage:**
```bash
# Run all tests
./scripts/run-tests.sh

# Run specific test type
./scripts/run-tests.sh --type unit
./scripts/run-tests.sh --type integration
./scripts/run-tests.sh --type e2e

# Generate coverage report
./scripts/run-tests.sh --coverage
```

### Deployment Scripts

#### 🚀 `deploy.sh`
**Purpose:** Production deployment automation

**Features:**
- Docker image building
- Service deployment
- Health check verification
- Rollback capability

**Usage:**
```bash
# Deploy to staging
./scripts/deploy.sh --env staging

# Deploy to production
./scripts/deploy.sh --env production

# Rollback deployment
./scripts/deploy.sh --rollback
```

#### 🏗️ `build.sh`
**Purpose:** Build all project components

**Usage:**
```bash
# Build all services
./scripts/build.sh

# Build specific service
./scripts/build.sh --service go-api
./scripts/build.sh --service rust-processor
./scripts/build.sh --service frontend
```

### Maintenance Scripts

#### 🧹 `cleanup.sh`
**Purpose:** Clean up development environment

**Features:**
- Remove unused Docker containers/images
- Clear log files
- Clean build artifacts
- Reset development data

**Usage:**
```bash
# Full cleanup
./scripts/cleanup.sh

# Partial cleanup
./scripts/cleanup.sh --docker-only
./scripts/cleanup.sh --logs-only
```

#### 📊 `health-check.sh`
**Purpose:** System health monitoring

**Features:**
- Service availability checking
- Database connectivity
- Message queue status
- Performance metrics

**Usage:**
```bash
./scripts/health-check.sh
```

</details>

<details>
<summary><strong>🔧 Script Configuration</strong></summary>

### Environment Variables

Scripts use the following environment variables:

```bash
# Project Configuration
PROJECT_NAME="paygateway"
PROJECT_ROOT="/path/to/payment-gateway-challenge"

# Service Ports
FRONTEND_PORT=8081
GO_API_PORT=8080
POSTGRES_PORT=5432
REDIS_PORT=6379
RABBITMQ_PORT=5672
RABBITMQ_MANAGEMENT_PORT=15672

# Database Configuration
DB_HOST=localhost
DB_USER=paygateway_user
DB_PASSWORD=paygateway_pass
DB_NAME=paygateway_db

# Docker Configuration
DOCKER_NETWORK=paygateway-network
DOCKER_COMPOSE_FILE=docker-compose.yml
DOCKER_COMPOSE_DEV_FILE=docker-compose.dev.yml

# Build Configuration
BUILD_ENV=development
RUST_LOG=info
GIN_MODE=debug
```

### Configuration Files

#### `.env.scripts`
Create this file to override default script configurations:

```bash
# Custom script configuration
VERBOSE_OUTPUT=true
AUTO_MIGRATE=true
SKIP_HEALTH_CHECKS=false
DEFAULT_ENVIRONMENT=development
```

</details>

<details>
<summary><strong>📝 Script Usage Examples</strong></summary>

### Complete Project Setup

```bash
# First time setup
git clone https://github.com/EduardoMG12/payment-gateway-challenge.git
cd payment-gateway-challenge

# Run complete setup
chmod +x scripts/*.sh
./scripts/setup.sh

# Start development environment
./scripts/dev-start.sh
```

### Development Workflow

```bash
# Daily development startup
./scripts/dev-start.sh

# Run tests during development
./scripts/run-tests.sh --type unit --watch

# Database operations
./scripts/db-migrate.sh up
./scripts/db-seed.sh

# Cleanup at end of day
./scripts/dev-stop.sh
```

### Testing Workflow

```bash
# Complete test suite
./scripts/run-tests.sh

# Specific service testing
./scripts/run-tests.sh --service go-api --coverage
./scripts/run-tests.sh --service rust-processor --bench

# Integration testing
./scripts/run-tests.sh --type integration --verbose
```

### Deployment Workflow

```bash
# Build for production
./scripts/build.sh --env production

# Deploy to staging
./scripts/deploy.sh --env staging --health-check

# Promote to production
./scripts/deploy.sh --env production --backup

# Monitor deployment
./scripts/health-check.sh --continuous
```

</details>

<details>
<summary><strong>🛠️ Creating Custom Scripts</strong></summary>

### Script Template

```bash
#!/bin/bash

# Script: template.sh
# Description: Template for new scripts
# Usage: ./template.sh [options]

set -euo pipefail  # Exit on error, undefined vars, pipe failures

# Configuration
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"

# Load common functions
source "$SCRIPT_DIR/common.sh"

# Default values
VERBOSE=false
DRY_RUN=false

# Functions
usage() {
    cat << EOF
Usage: $0 [OPTIONS]

Description of what this script does.

OPTIONS:
    -h, --help      Show this help message
    -v, --verbose   Enable verbose output
    -n, --dry-run   Show what would be done without executing
    
EXAMPLES:
    $0 --verbose
    $0 --dry-run

EOF
}

main() {
    local option=""
    
    # Parse arguments
    while [[ $# -gt 0 ]]; do
        case $1 in
            -h|--help)
                usage
                exit 0
                ;;
            -v|--verbose)
                VERBOSE=true
                shift
                ;;
            -n|--dry-run)
                DRY_RUN=true
                shift
                ;;
            *)
                log_error "Unknown option: $1"
                usage
                exit 1
                ;;
        esac
    done
    
    # Script logic here
    log_info "Script execution completed"
}

# Execute main function
main "$@"
```

### Common Functions (`common.sh`)

```bash
#!/bin/bash

# Common functions for PayGateway scripts

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

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

log_debug() {
    if [[ "${VERBOSE:-false}" == "true" ]]; then
        echo -e "${BLUE}[DEBUG]${NC} $1" >&2
    fi
}

# Utility functions
check_command() {
    if ! command -v "$1" &> /dev/null; then
        log_error "Required command '$1' is not installed"
        return 1
    fi
}

check_file() {
    if [[ ! -f "$1" ]]; then
        log_error "Required file '$1' does not exist"
        return 1
    fi
}

check_docker_running() {
    if ! docker info &> /dev/null; then
        log_error "Docker is not running"
        return 1
    fi
}

wait_for_service() {
    local service_name="$1"
    local host="$2"
    local port="$3"
    local max_attempts="${4:-30}"
    
    log_info "Waiting for $service_name to be ready..."
    
    for ((i=1; i<=max_attempts; i++)); do
        if nc -z "$host" "$port" &> /dev/null; then
            log_info "$service_name is ready"
            return 0
        fi
        
        if [[ $i -eq $max_attempts ]]; then
            log_error "$service_name failed to start within timeout"
            return 1
        fi
        
        sleep 2
    done
}
```

### Best Practices

#### Script Guidelines
1. **Use `set -euo pipefail`** for error handling
2. **Include usage documentation** with `--help`
3. **Validate prerequisites** before execution
4. **Provide verbose output** options
5. **Include dry-run mode** for destructive operations
6. **Use consistent logging** functions
7. **Handle interruption signals** gracefully

#### Security Considerations
1. **Avoid hardcoded secrets** in scripts
2. **Use environment variables** for configuration
3. **Validate input parameters** thoroughly
4. **Set appropriate file permissions** (755 for scripts)
5. **Use `shellcheck`** for static analysis

</details>

---

## 🚀 Getting Started with Scripts

### Initial Setup

```bash
# Make all scripts executable
find scripts -name "*.sh" -exec chmod +x {} \;

# Run initial setup
./scripts/setup.sh

# Verify installation
./scripts/health-check.sh
```

### Common Workflows

#### New Developer Onboarding
```bash
./scripts/setup.sh --new-developer
./scripts/dev-start.sh
./scripts/run-tests.sh --smoke-test
```

#### Daily Development
```bash
./scripts/dev-start.sh
# ... development work ...
./scripts/run-tests.sh --quick
./scripts/dev-stop.sh
```

#### Production Deployment
```bash
./scripts/build.sh --env production
./scripts/run-tests.sh --full-suite
./scripts/deploy.sh --env production
```

## 🤝 Contributing

When adding new scripts:

1. Follow the established template and conventions
2. Include comprehensive help documentation
3. Add error handling and validation
4. Test scripts in different environments
5. Update this README with new script documentation

## 📄 License

These scripts are part of the PayGateway project and follow the same [MIT License](../LICENSE).

---

<div align="center">

**Part of the PayGateway ecosystem**

[🏠 Main Project](../README.md) • [🌐 Frontend](../frontend/README.md) • [🚀 Go API](../go-api/README.md) • [⚡ Rust Processor](../rust-processor/README.md)

</div>
