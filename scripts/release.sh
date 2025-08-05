#!/bin/bash
# Release script for go-fake

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if version is provided
if [ -z "$1" ]; then
    print_error "Usage: $0 <version>"
    print_error "Example: $0 v1.2.0"
    exit 1
fi

VERSION=$1

# Validate version format
if [[ ! $VERSION =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    print_error "Version must be in format v1.2.3"
    exit 1
fi

print_status "Preparing release $VERSION"

# Check if we're on main/master branch
CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
if [[ "$CURRENT_BRANCH" != "main" && "$CURRENT_BRANCH" != "master" ]]; then
    print_warning "You're not on main/master branch (current: $CURRENT_BRANCH)"
    read -p "Continue anyway? (y/N): " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        print_error "Aborted"
        exit 1
    fi
fi

# Check for uncommitted changes
if ! git diff-index --quiet HEAD --; then
    print_error "You have uncommitted changes. Please commit or stash them first."
    exit 1
fi

# Update version in main.go
print_status "Updating version in main.go"
sed -i.bak "s/const version = \".*\"/const version = \"$VERSION\"/" cmd/generate/main.go
rm cmd/generate/main.go.bak

# Run tests
print_status "Running tests"
if ! make test; then
    print_error "Tests failed. Aborting release."
    exit 1
fi

print_success "Tests passed"

# Build local release to test
print_status "Building release binaries locally"
if ! make release-local; then
    print_error "Release build failed"
    exit 1
fi

print_success "Release binaries built successfully"

# Test the binary
print_status "Testing built binary"
./build/go-fake-linux-amd64 --version
if [ $? -ne 0 ]; then
    print_error "Built binary test failed"
    exit 1
fi

print_success "Built binary test passed"

# Commit version change
print_status "Committing version change"
git add cmd/generate/main.go
git commit -m "chore: bump version to $VERSION"

# Create and push tag
print_status "Creating git tag $VERSION"
git tag -a "$VERSION" -m "Release $VERSION

Features:
- AI-Enhanced Field Inference with OpenAI integration
- Intelligent Pattern Matching for 40+ data types  
- Format Override for JSON/CSV conversion
- Relationship Constraints with foreign key support
- Directory-based Output for multi-table schemas
- High Performance fake data generation

See full changelog: https://github.com/Livin21/go-fake/releases/tag/$VERSION"

print_status "Pushing changes and tag to remote"
git push origin $CURRENT_BRANCH
git push origin $VERSION

print_success "Release $VERSION created successfully!"
print_status "GitHub Actions will now:"
print_status "  1. Build binaries for all platforms"
print_status "  2. Run tests"
print_status "  3. Create GitHub release with binaries"
print_status "  4. Make binaries available for download"
print_status ""
print_status "Release will be available at:"
print_status "  https://github.com/Livin21/go-fake/releases/tag/$VERSION"
print_status ""
print_status "Users can install with:"
print_status "  go install github.com/Livin21/go-fake/cmd/generate@$VERSION"

# Clean up build directory
make clean

print_success "Release process completed! 🎉"
