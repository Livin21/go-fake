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

# Function to extract unreleased section from CHANGELOG.md
extract_unreleased_notes() {
    local changelog_file="CHANGELOG.md"
    
    if [[ ! -f "$changelog_file" ]]; then
        print_error "CHANGELOG.md not found"
        return 1
    fi
    
    # Extract content between ## [Unreleased] and the next ## section
    local temp_file=$(mktemp)
    local in_unreleased=false
    local found_unreleased=false
    
    while IFS= read -r line; do
        if [[ "$line" =~ ^##[[:space:]]\[Unreleased\] ]]; then
            in_unreleased=true
            found_unreleased=true
            continue
        elif [[ "$line" =~ ^##[[:space:]] ]] && [[ "$in_unreleased" == true ]]; then
            # Found next section, stop
            break
        elif [[ "$in_unreleased" == true ]]; then
            echo "$line" >> "$temp_file"
        fi
    done < "$changelog_file"
    
    if [[ "$found_unreleased" == false ]]; then
        print_error "No [Unreleased] section found in CHANGELOG.md"
        rm -f "$temp_file"
        return 1
    fi
    
    # Check if unreleased section has content
    if [[ ! -s "$temp_file" ]] || ! grep -q '[^[:space:]]' "$temp_file"; then
        print_error "Unreleased section in CHANGELOG.md is empty"
        rm -f "$temp_file"
        return 1
    fi
    
    cat "$temp_file"
    rm -f "$temp_file"
}

# Function to update CHANGELOG.md for release
update_changelog() {
    local version="$1"
    local changelog_file="CHANGELOG.md"
    local backup_file="${changelog_file}.backup"
    local temp_file=$(mktemp)
    
    # Create backup
    cp "$changelog_file" "$backup_file"
    
    # Get current date in YYYY-MM-DD format
    local release_date=$(date +%Y-%m-%d)
    
    # Process the changelog
    local in_unreleased=false
    local processed_unreleased=false
    
    while IFS= read -r line; do
        if [[ "$line" =~ ^##[[:space:]]\[Unreleased\] ]]; then
            # Replace [Unreleased] with version and date, then add new [Unreleased] section
            echo "## [Unreleased]" >> "$temp_file"
            echo "" >> "$temp_file"
            echo "## [${version#v}] - $release_date" >> "$temp_file"
            in_unreleased=true
            processed_unreleased=true
            continue
        elif [[ "$line" =~ ^##[[:space:]] ]] && [[ "$in_unreleased" == true ]]; then
            # Found next section after unreleased
            in_unreleased=false
            echo "$line" >> "$temp_file"
        else
            echo "$line" >> "$temp_file"
        fi
    done < "$changelog_file"
    
    if [[ "$processed_unreleased" == false ]]; then
        print_error "Failed to process [Unreleased] section in CHANGELOG.md"
        rm -f "$temp_file"
        return 1
    fi
    
    # Replace original file
    mv "$temp_file" "$changelog_file"
    
    print_success "CHANGELOG.md updated for release $version"
    return 0
}

# Function to create release notes from unreleased section
create_release_notes() {
    local version="$1"
    local notes_file="$2"
    
    print_status "Extracting release notes from CHANGELOG.md"
    
    # Extract unreleased content
    local unreleased_content
    if ! unreleased_content=$(extract_unreleased_notes); then
        print_error "Failed to extract unreleased notes from CHANGELOG.md"
        return 1
    fi
    
    # Create formatted release notes
    cat > "$notes_file" << EOF
## 🚀 go-fake $version

$unreleased_content

### Download Instructions:

**Linux:**
\`\`\`bash
# AMD64
wget https://github.com/Livin21/go-fake/releases/download/$version/go-fake-linux-amd64
chmod +x go-fake-linux-amd64
sudo mv go-fake-linux-amd64 /usr/local/bin/go-fake

# ARM64
wget https://github.com/Livin21/go-fake/releases/download/$version/go-fake-linux-arm64
chmod +x go-fake-linux-arm64
sudo mv go-fake-linux-arm64 /usr/local/bin/go-fake
\`\`\`

**Windows:**
Download the appropriate \`.exe\` file from the release assets.

**macOS:**
\`\`\`bash
# Intel Macs
wget https://github.com/Livin21/go-fake/releases/download/$version/go-fake-darwin-amd64
chmod +x go-fake-darwin-amd64
sudo mv go-fake-darwin-amd64 /usr/local/bin/go-fake

# Apple Silicon Macs
wget https://github.com/Livin21/go-fake/releases/download/$version/go-fake-darwin-arm64
chmod +x go-fake-darwin-arm64
sudo mv go-fake-darwin-arm64 /usr/local/bin/go-fake
\`\`\`

**Using Go:**
\`\`\`bash
go install github.com/Livin21/go-fake/cmd/generate@$version
\`\`\`

### Usage Examples:
\`\`\`bash
# Basic usage
go-fake -schema examples/sample.json -rows 100 -output data.json

# Use AI enhancement (requires OPENAI_API_KEY)
export OPENAI_API_KEY="your-key"
go-fake -schema schema.json -ai -output enhanced_data.json

# Override output format
go-fake -schema schema.sql -format json -output data.json

# Enable verbose logging
go-fake -schema schema.json -rows 50 -verbose
\`\`\`

### Verify Downloads:
All binaries are signed and checksums are provided in \`checksums.txt\`. Verify your download:
\`\`\`bash
sha256sum go-fake-* && cat checksums.txt
\`\`\`
EOF
    
    print_success "Release notes created from CHANGELOG.md"
    return 0
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

# Validate CHANGELOG.md before proceeding
print_status "Validating CHANGELOG.md"
if ! extract_unreleased_notes > /dev/null; then
    print_error "CHANGELOG.md validation failed. Please ensure:"
    print_error "  1. CHANGELOG.md exists"
    print_error "  2. It contains an [Unreleased] section"  
    print_error "  3. The [Unreleased] section has content"
    exit 1
fi

print_success "CHANGELOG.md validation passed"

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

# Create release notes from CHANGELOG.md before updating it
RELEASE_NOTES_FILE=$(mktemp)
print_status "Creating release notes from CHANGELOG.md"
if ! create_release_notes "$VERSION" "$RELEASE_NOTES_FILE"; then
    print_error "Failed to create release notes"
    rm -f "$RELEASE_NOTES_FILE"
    exit 1
fi

# Update CHANGELOG.md
print_status "Updating CHANGELOG.md for release $VERSION"
if ! update_changelog "$VERSION"; then
    print_error "Failed to update CHANGELOG.md"
    print_error "Backup available at CHANGELOG.md.backup"
    rm -f "$RELEASE_NOTES_FILE"
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
    # Restore CHANGELOG.md backup
    if [[ -f "CHANGELOG.md.backup" ]]; then
        mv CHANGELOG.md.backup CHANGELOG.md
        print_status "CHANGELOG.md restored from backup"
    fi
    rm -f "$RELEASE_NOTES_FILE"
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

# Commit version change and changelog update
print_status "Committing version change and CHANGELOG.md update" 
git add cmd/generate/main.go CHANGELOG.md
git commit -m "chore: bump version to $VERSION and update CHANGELOG.md"

# Create release tag with notes from CHANGELOG.md
print_status "Creating git tag $VERSION with release notes from CHANGELOG.md"

# Extract just the content part for the git tag (without the header and download instructions)
TAG_MESSAGE="Release $VERSION

$(extract_unreleased_notes)"

git tag -a "$VERSION" -m "$TAG_MESSAGE"

print_status "Pushing changes and tag to remote"
git push origin $CURRENT_BRANCH
git push origin $VERSION

# Clean up temporary files
rm -f "$RELEASE_NOTES_FILE"
rm -f CHANGELOG.md.backup

print_success "Release $VERSION created successfully!"
print_status "GitHub Actions will now:"
print_status "  1. Build binaries for all platforms"
print_status "  2. Run tests"
print_status "  3. Create GitHub release with release notes from CHANGELOG.md"
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
