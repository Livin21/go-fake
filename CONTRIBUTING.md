# Contributing to go-fake

Thank you for your interest in contributing to go-fake! 🎉

## Development Setup

1. **Fork and clone the repository:**
   ```bash
   git clone https://github.com/livin21/go-fake.git
   cd go-fake
   ```

2. **Install dependencies:**
   ```bash
   make deps
   ```

3. **Build and test:**
   ```bash
   make build
   make test
   ```

## Development Workflow

1. **Create a feature branch:**
   ```bash
   git checkout -b feature/amazing-feature
   ```

2. **Make your changes and add tests**

3. **Test your changes:**
   ```bash
   make test
   make build
   ./bin/go-fake --version
   ```

4. **Commit with conventional commits:**
   ```bash
   git commit -m "feat: add amazing new feature"
   git commit -m "fix: resolve issue with data generation"
   git commit -m "docs: update README with new examples"
   ```

5. **Push and create PR:**
   ```bash
   git push origin feature/amazing-feature
   # Create pull request on GitHub
   ```

## Adding New Data Types

1. **Add generator function** to `pkg/faker/providers.go`:
   ```go
   func GenerateCustomType() string {
       // Your implementation
       return "custom_value"
   }
   ```

2. **Add pattern detection** to `internal/generator/intelligent.go`:
   ```go
   exactPatterns["custom_field"] = "customtype"
   partialPatterns["custom"] = "customtype"
   ```

3. **Add case** to `generateValueByType()` in the same file:
   ```go
   case "customtype":
       return faker.GenerateCustomType()
   ```

4. **Update help text** in `cmd/generate/main.go`

5. **Add tests** for the new functionality

## Release Process

**For Maintainers:**

1. **Update version and changelog:**
   ```bash
   # Update CHANGELOG.md with new features
   # Version will be updated automatically by release script
   ```

2. **Create release:**
   ```bash
   ./scripts/release.sh v1.2.0
   ```

3. **GitHub Actions will automatically:**
   - Build binaries for all platforms
   - Run tests
   - Create GitHub release
   - Upload binaries with checksums

## Code Style

- Follow standard Go conventions
- Use `gofmt` to format code
- Add comments for exported functions
- Write tests for new functionality
- Keep functions focused and small

## Testing

- Add unit tests for new functionality
- Test with various schema formats
- Verify AI integration works correctly
- Test format override functionality

## Documentation

- Update README.md for new features
- Add examples for new data types
- Update help text and command documentation
- Include AI integration examples where relevant

## Questions?

- Open an issue for discussion
- Check existing issues and PRs
- Follow the code of conduct

Thank you for contributing! 🚀
