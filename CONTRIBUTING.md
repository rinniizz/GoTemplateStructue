# Contributing to Go Template Structure

Thank you for your interest in contributing! This document provides guidelines for contributing to this project.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Workflow](#development-workflow)
- [Coding Standards](#coding-standards)
- [Commit Message Guidelines](#commit-message-guidelines)
- [Pull Request Process](#pull-request-process)
- [Testing Guidelines](#testing-guidelines)

## Code of Conduct

- Be respectful and inclusive
- Welcome newcomers and encourage diverse contributions
- Focus on constructive criticism
- Respect differing viewpoints and experiences

## Getting Started

### Prerequisites

- Go 1.24.0 or higher
- PostgreSQL 15+
- Redis 7+ (optional but recommended)
- Git

### Setup Development Environment

1. **Fork and Clone the Repository**
   ```bash
   git clone https://github.com/YOUR_USERNAME/FinHub-BackEnd.git
   cd FinHub-BackEnd
   ```

2. **Install Dependencies**
   ```bash
   go mod download
   ```

3. **Setup Environment Variables**
   ```bash
   cp .env.example .env
   # Edit .env with your local configuration
   ```

4. **Run Database Migrations**
   ```bash
   # See migrations/README.md for detailed instructions
   migrate -path migrations -database "YOUR_DB_URL" up
   ```

5. **Run Development Server**
   ```bash
   # Windows
   dev.bat
   
   # Linux/Mac
   make dev
   ```

## Development Workflow

### Branch Naming Convention

Use descriptive branch names with the following prefixes:

- `feature/` - New features (e.g., `feature/add-user-authentication`)
- `bugfix/` - Bug fixes (e.g., `bugfix/fix-login-error`)
- `hotfix/` - Urgent production fixes (e.g., `hotfix/security-patch`)
- `refactor/` - Code refactoring (e.g., `refactor/improve-user-service`)
- `docs/` - Documentation updates (e.g., `docs/update-readme`)
- `test/` - Test additions or modifications (e.g., `test/add-user-service-tests`)

### Workflow Steps

1. **Create a new branch**
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make your changes**
   - Write code following our [Coding Standards](#coding-standards)
   - Add tests for new functionality
   - Update documentation if needed

3. **Test your changes**
   ```bash
   # Run tests
   go test ./...
   
   # Run with coverage
   go test -cover ./...
   
   # Run linter
   golangci-lint run
   ```

4. **Commit your changes**
   - Follow [Commit Message Guidelines](#commit-message-guidelines)
   ```bash
   git add .
   git commit -m "feat: add user profile endpoint"
   ```

5. **Push to your fork**
   ```bash
   git push origin feature/your-feature-name
   ```

6. **Create a Pull Request**
   - Go to the original repository
   - Click "New Pull Request"
   - Select your branch
   - Fill in the PR template

## Coding Standards

### Go Style Guide

Follow the official [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments) and [Effective Go](https://go.dev/doc/effective_go).

### Key Principles

1. **Clean Architecture**
   - Separate concerns into layers (domain, repository, service, handler)
   - Dependencies point inward (handler → service → repository)
   - Keep domain models independent

2. **Code Organization**
   ```
   internal/
   ├── domain/       # Business entities and interfaces
   ├── repository/   # Data access layer
   ├── service/      # Business logic layer
   ├── handler/      # HTTP handlers
   ├── middleware/   # HTTP middlewares
   └── config/       # Configuration
   ```

3. **Naming Conventions**
   - Use camelCase for variables and functions
   - Use PascalCase for exported types and functions
   - Use descriptive names (avoid abbreviations)
   - Interfaces should describe behavior (e.g., `UserRepository`, not `IUserRepository`)

4. **Error Handling**
   ```go
   // ✅ Good
   if err != nil {
       return fmt.Errorf("failed to create user: %w", err)
   }
   
   // ❌ Bad
   if err != nil {
       return err
   }
   ```

5. **Logging**
   ```go
   // Use structured logging
   logger.Info("User created", "user_id", user.ID, "email", user.Email)
   
   // Not plain strings
   logger.Info("User created: " + user.Email)
   ```

### Code Quality Checklist

- [ ] No compiler warnings or errors
- [ ] All tests pass
- [ ] Code coverage maintained or improved
- [ ] No commented-out code
- [ ] No TODO comments (create issues instead)
- [ ] Proper error handling
- [ ] Logging added for important operations
- [ ] Documentation updated

## Commit Message Guidelines

We follow [Conventional Commits](https://www.conventionalcommits.org/) specification.

### Format

```
<type>(<scope>): <subject>

<body>

<footer>
```

### Types

- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, missing semi-colons, etc.)
- `refactor`: Code refactoring
- `perf`: Performance improvements
- `test`: Adding or updating tests
- `chore`: Maintenance tasks (dependencies, build process, etc.)
- `ci`: CI/CD changes

### Examples

```bash
# Feature
git commit -m "feat(auth): add JWT token refresh endpoint"

# Bug fix
git commit -m "fix(user): resolve duplicate email validation issue"

# Documentation
git commit -m "docs(readme): update installation instructions"

# With body and footer
git commit -m "feat(api): add pagination support

Add pagination to user list endpoint with page and limit parameters.

Closes #123"
```

### Best Practices

- Use imperative mood ("add" not "added" or "adds")
- Keep subject line under 72 characters
- Capitalize first letter
- No period at the end of subject line
- Reference issues in footer (e.g., "Closes #123", "Fixes #456")

## Pull Request Process

### Before Submitting

1. ✅ All tests pass locally
2. ✅ Code follows style guidelines
3. ✅ Documentation is updated
4. ✅ Commit messages follow conventions
5. ✅ No merge conflicts with main branch

### PR Title Format

Follow the same format as commit messages:

```
feat(auth): add OAuth2 provider support
```

### PR Description Template

```markdown
## Description
Brief description of changes

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Breaking change
- [ ] Documentation update

## How Has This Been Tested?
Describe testing steps

## Checklist
- [ ] My code follows the project's style guidelines
- [ ] I have performed a self-review
- [ ] I have commented my code where necessary
- [ ] I have updated the documentation
- [ ] My changes generate no new warnings
- [ ] I have added tests that prove my fix/feature works
- [ ] New and existing unit tests pass locally
- [ ] Any dependent changes have been merged

## Screenshots (if applicable)
Add screenshots here

## Related Issues
Closes #(issue number)
```

### Review Process

1. **Automated Checks**
   - CI/CD pipeline must pass
   - Code coverage should not decrease
   - Linting must pass

2. **Code Review**
   - At least one maintainer approval required
   - Address all review comments
   - Re-request review after making changes

3. **Merge**
   - Use "Squash and merge" for feature branches
   - Use "Rebase and merge" for hotfixes
   - Delete branch after merging

## Testing Guidelines

### Test Structure

```go
func TestFunctionName(t *testing.T) {
    // Setup
    mockRepo := new(MockRepository)
    service := NewService(mockRepo)
    
    // Test cases
    t.Run("Success case", func(t *testing.T) {
        // Arrange
        mockRepo.On("GetByID", 1).Return(expectedResult, nil)
        
        // Act
        result, err := service.GetByID(1)
        
        // Assert
        assert.NoError(t, err)
        assert.Equal(t, expectedResult, result)
        mockRepo.AssertExpectations(t)
    })
    
    t.Run("Error case", func(t *testing.T) {
        // Similar structure
    })
}
```

### Test Coverage Goals

- **Overall**: Aim for 70%+ coverage
- **Critical paths**: 90%+ coverage (auth, payments, etc.)
- **New features**: Must include tests
- **Bug fixes**: Add regression tests

### Running Tests

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Run specific package
go test ./internal/service/...

# Run specific test
go test -run TestUserService_GetProfile ./test/...
```

## Questions?

- Create an issue for bugs or feature requests
- Use GitHub Discussions for questions
- Check existing issues and discussions first

---

**Thank you for contributing! 🎉**
