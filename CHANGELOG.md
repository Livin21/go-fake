# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- GitHub release pipeline with multi-platform binaries
- Automated CI/CD with GitHub Actions
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

1. Update version in `cmd/generate/main.go`
2. Update this CHANGELOG.md
3. Run the release script: `./scripts/release.sh v1.x.x`
4. The GitHub Actions workflow will automatically:
   - Build binaries for all platforms
   - Create GitHub release
   - Upload binaries and checksums

## Version Format

This project uses [Semantic Versioning](https://semver.org/):
- **MAJOR** version for incompatible API changes
- **MINOR** version for backwards-compatible functionality additions  
- **PATCH** version for backwards-compatible bug fixes
