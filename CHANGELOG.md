# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- 📊 **Verbose Logging** with `-verbose` flag for detailed execution visibility
- ⏱️ **Performance Timing** for data generation operations
- 📈 **Progress Tracking** for multi-table generation
- 🔍 **Step-by-step Execution Logging** with INFO/DEBUG levels
- GitHub release pipeline with multi-platform binaries
- Cross-platform build support (Linux, Windows, macOS - AMD64/ARM64)

## [1.1.0] - 2025-08-05

### Added
- 🤖 **AI-Enhanced Field Inference** with OpenAI integration
- 🧠 **Intelligent Pattern Matching** for 40+ data types
- 🔄 **Format Override** option (`-format json|csv`)
- 🎯 **Extended Data Types**: Added 20+ new specialized data types
- 📁 **Directory-based Output** for multi-table schemas
- 🔗 **Relationship Constraints** with foreign key support
- ⚡ **Performance Optimizations** and enhanced error handling

### Enhanced
- Smart field type detection with multi-layered inference
- Comprehensive help text with all supported data types
- Better CSV header handling
- Improved data generation algorithms

### Fixed
- CSV header generation issues
- Schema type resolution conflicts
- Memory usage optimizations

## [1.0.0] - Initial Release

### Added
- Basic fake data generation from JSON and SQL schemas
- Support for common data types (string, int, float, bool, date, etc.)
- Multi-table schema support
- Command-line interface
- JSON and CSV output formats

---

## Release Process

To create a new release:

1. **Add your changes to the [Unreleased] section** in this CHANGELOG.md file
   - Follow the existing format with appropriate sections (Added, Enhanced, Fixed, etc.)
   - Use descriptive entries with emojis for visual clarity
   - Ensure the [Unreleased] section has meaningful content before releasing

2. **Run the release script**: `./scripts/release.sh v1.x.x`
   - The script will automatically validate the CHANGELOG.md
   - Extract release notes from the [Unreleased] section
   - Update the CHANGELOG.md by moving [Unreleased] content to the new version
   - Update version in `cmd/generate/main.go`
   - Run tests to ensure everything works
   - Create git commit and tag with extracted release notes
   - Create backup files for safety

3. **The GitHub Actions workflow will automatically:**
   - Build binaries for all platforms (Linux, Windows, macOS - AMD64/ARM64)
   - Extract release notes from the updated CHANGELOG.md
   - Create GitHub release with the extracted content
   - Upload binaries and checksums
   - Make the release publicly available

### Safety Features:
- **Validation**: Script validates CHANGELOG.md before proceeding
- **Backups**: Automatic backup of CHANGELOG.md before modification
- **Rollback**: Failed operations restore from backup
- **Fallbacks**: GitHub Actions falls back to generic release notes if extraction fails
- **Testing**: Full test suite runs before release creation

### CHANGELOG.md Format Requirements:
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

**Important**: Always ensure the [Unreleased] section contains meaningful content before running a release.

## Version Format

This project uses [Semantic Versioning](https://semver.org/):
- **MAJOR** version for incompatible API changes
- **MINOR** version for backwards-compatible functionality additions  
- **PATCH** version for backwards-compatible bug fixes
