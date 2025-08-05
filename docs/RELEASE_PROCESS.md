# Automated Release Process

This document explains the enhanced automated release process for go-fake.

## Overview

The release process now automatically reads from `CHANGELOG.md` and handles all the formatting and distribution automatically. No more manual release note writing!

## How It Works

### 1. Development Workflow
As you develop features, add them to the `[Unreleased]` section in `CHANGELOG.md`:

```markdown
## [Unreleased]

### Added
- 📊 **New Feature** with detailed description
- 🎯 **Another Feature** explaining the benefit

### Enhanced  
- Improved existing functionality
- Better performance in specific areas

### Fixed
- Bug fixes and corrections
```

### 2. Release Creation
When ready to release, simply run:

```bash
./scripts/release.sh v1.2.0
```

The script will:
1. ✅ Validate `CHANGELOG.md` has content in `[Unreleased]` section
2. 📝 Extract release notes from `[Unreleased]` section
3. 🔄 Update `CHANGELOG.md` by moving content to versioned section
4. 🔢 Update version in `cmd/generate/main.go`
5. 🧪 Run full test suite
6. 📦 Build and test release binaries
7. 📤 Create git commit and tag with extracted release notes
8. 🚀 Push to GitHub

### 3. GitHub Actions
Once the tag is pushed, GitHub Actions will:
1. 🏗️ Build binaries for all platforms (Linux, Windows, macOS - AMD64/ARM64)
2. 📋 Extract the same release notes from `CHANGELOG.md`
3. 🎁 Create GitHub release with formatted notes and download instructions
4. 📎 Upload binaries and checksums
5. 🌐 Make the release publicly available

## Safety Features

### Validation
- **CHANGELOG.md Validation**: Ensures file exists and has content in `[Unreleased]`
- **Git Status Check**: Prevents release with uncommitted changes
- **Branch Check**: Warns if not on main/master branch
- **Test Suite**: Full test suite must pass before release

### Backup and Recovery
- **Automatic Backup**: `CHANGELOG.md.backup` created before modification
- **Rollback on Failure**: Failed operations restore from backup
- **Safe Exit**: Script exits cleanly on any validation failure

### Fallback Mechanisms
- **GitHub Actions Fallback**: If CHANGELOG extraction fails, uses generic release notes
- **Error Recovery**: Clear error messages guide you to fix issues
- **Dry Run Capability**: Functions can be tested individually

## CHANGELOG.md Format

### Required Structure
```markdown
## [Unreleased]

### Added
- New features

### Enhanced
- Improvements to existing features

### Fixed
- Bug fixes
```

### Supported Sections
- `Added` - New features
- `Enhanced` - Improvements to existing functionality  
- `Fixed` - Bug fixes
- `Changed` - Changes in existing functionality
- `Deprecated` - Soon-to-be removed features
- `Removed` - Now removed features
- `Security` - Security improvements

### Best Practices
- Use emojis for visual clarity (📊 🎯 🔧 🐛)
- Write clear, user-focused descriptions
- Group related changes together
- Always keep `[Unreleased]` section updated during development

## Example Release Flow

```bash
# 1. During development, update CHANGELOG.md
git add CHANGELOG.md
git commit -m "docs: add new feature to changelog"

# 2. When ready to release
./scripts/release.sh v1.3.0

# 3. Script output:
# [INFO] Preparing release v1.3.0
# [INFO] Validating CHANGELOG.md
# [SUCCESS] CHANGELOG.md validation passed
# [INFO] Creating release notes from CHANGELOG.md
# [SUCCESS] Release notes created from CHANGELOG.md
# [INFO] Updating CHANGELOG.md for release v1.3.0
# [SUCCESS] CHANGELOG.md updated for release v1.3.0
# [INFO] Updating version in main.go
# [INFO] Running tests
# [SUCCESS] Tests passed
# [INFO] Building release binaries locally
# [SUCCESS] Release binaries built successfully
# [INFO] Testing built binary
# [SUCCESS] Built binary test passed
# [INFO] Committing version change and CHANGELOG.md update
# [INFO] Creating git tag v1.3.0 with release notes from CHANGELOG.md
# [INFO] Pushing changes and tag to remote
# [SUCCESS] Release v1.3.0 created successfully!

# 4. GitHub Actions automatically:
# - Builds all platform binaries
# - Creates GitHub release with extracted notes
# - Makes download links available
```

## Troubleshooting

### Common Issues

**"CHANGELOG.md validation failed"**
- Ensure `CHANGELOG.md` exists
- Check that `[Unreleased]` section exists
- Verify the section has actual content (not just whitespace)

**"You have uncommitted changes"**
- Commit or stash your changes before releasing
- Use `git status` to see what's uncommitted

**"Tests failed. Aborting release"**
- Fix failing tests before releasing
- Run `make test` to see detailed test output

**"Release build failed"**
- Check build dependencies
- Ensure `make release-local` works manually

### Manual Recovery

If something goes wrong:
1. Check for `CHANGELOG.md.backup` - restore if needed
2. Check git log for any commits that need reverting
3. Delete any problematic git tags: `git tag -d v1.x.x`
4. Fix the issue and run the release script again

## Benefits

- 🚀 **Faster Releases**: No manual release note writing
- 📝 **Consistent Format**: Automated formatting and structure
- 🔒 **Safety**: Multiple validation layers prevent errors
- 📋 **Complete Automation**: From CHANGELOG to GitHub release
- 🎯 **Developer Focus**: Spend time on features, not release management
- 📚 **Better Documentation**: Forces good changelog practices
