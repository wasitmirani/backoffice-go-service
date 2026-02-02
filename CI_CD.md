# CI/CD Documentation

## Overview

This project uses GitHub Actions for Continuous Integration (CI) and Continuous Deployment (CD). The workflows are configured to automatically test, build, and deploy the application.

## Workflows

### 1. CI Workflow (`.github/workflows/ci.yaml`)

Runs on every push and pull request to `main`, `develop`, or `master` branches.

**Jobs:**
- **Lint**: Runs `golangci-lint` to check code quality
- **Test**: Runs tests with PostgreSQL and MySQL services
- **Build**: Builds the application for multiple platforms (Linux, Windows, macOS)
- **Security**: Runs `gosec` security scanner
- **Docker Build**: Builds Docker image for testing

**Features:**
- Parallel job execution
- Test coverage reporting to Codecov
- Multi-platform builds
- Docker image caching

### 2. CD Workflow (`.github/workflows/cd.yaml`)

Runs on pushes to `main`/`master` branches and on tag releases.

**Jobs:**
- **Build and Push**: Builds and pushes Docker image to GitHub Container Registry
- **Deploy Staging**: Deploys to staging environment
- **Deploy Production**: Deploys to production environment (on tags)
- **Release**: Creates GitHub release with binaries

**Features:**
- Automatic Docker image tagging
- Environment-based deployments
- Manual workflow dispatch
- Release creation

### 3. Release Workflow (`.github/workflows/release.yaml`)

Runs when a GitHub release is published.

**Features:**
- Builds binaries for all platforms
- Creates checksums
- Uploads release assets

## Setup

### 1. GitHub Secrets

Configure these secrets in your GitHub repository settings:

```
GITHUB_TOKEN          # Automatically provided by GitHub Actions
```

For deployments, you may need additional secrets:
```
KUBECONFIG            # Kubernetes configuration (if using K8s)
DOCKER_REGISTRY_TOKEN # Docker registry token (if using external registry)
```

### 2. Environment Protection

Set up environment protection rules in GitHub:
- **Staging**: Require approval for deployments
- **Production**: Require approval and restrict to specific branches

### 3. Branch Protection

Enable branch protection rules:
- Require status checks to pass
- Require pull request reviews
- Require up-to-date branches

## Usage

### Running CI Locally

```bash
# Install tools
make install-tools

# Run linting
make lint

# Run tests
make test

# Run all CI checks
make ci
```

### Manual Deployment

1. Go to Actions tab in GitHub
2. Select "CD" workflow
3. Click "Run workflow"
4. Choose environment (staging/production)
5. Click "Run workflow"

### Creating a Release

1. Create a new tag:
   ```bash
   git tag -a v1.0.0 -m "Release version 1.0.0"
   git push origin v1.0.0
   ```

2. GitHub Actions will automatically:
   - Build Docker image
   - Deploy to production
   - Create GitHub release

## Docker

### Building Locally

```bash
# Build Docker image
make docker-build

# Run Docker container
make docker-run

# Or manually
docker build -t backoffice-service:latest .
docker run -p 8080:8080 --env-file .env backoffice-service:latest
```

### Docker Image Tags

Images are automatically tagged:
- `latest`: Latest build from main branch
- `main-<sha>`: Specific commit SHA
- `v1.0.0`: Version tags
- `v1.0`: Major.minor version

## Testing

### Local Testing

```bash
# Run all tests
make test

# Run with coverage
make test-coverage

# Run verbose tests
make test-verbose
```

### CI Testing

Tests run automatically on:
- Every push to main/develop branches
- Every pull request
- Uses PostgreSQL and MySQL services
- Generates coverage reports

## Security Scanning

### Local Scanning

```bash
# Run security scan
make security
```

### CI Scanning

- Runs `gosec` on every CI run
- Results uploaded as artifacts
- Can be integrated with security dashboards

## Dependabot

Dependabot is configured to:
- Check for updates weekly (Mondays at 9 AM)
- Update Go modules
- Update Docker images
- Update GitHub Actions

Configuration: `.github/dependabot.yaml`

## Monitoring

### Build Status

Check build status:
- GitHub Actions tab
- Status badges in README

### Deployment Status

Monitor deployments:
- GitHub Environments
- Deployment logs in Actions

## Troubleshooting

### CI Failures

1. Check Actions tab for error details
2. Review logs for specific job failures
3. Run tests locally: `make test`
4. Run linting locally: `make lint`

### Deployment Failures

1. Check environment secrets are configured
2. Verify deployment permissions
3. Review deployment logs
4. Test Docker image locally

### Common Issues

**Issue**: Tests fail in CI but pass locally
- **Solution**: Ensure database services are properly configured

**Issue**: Docker build fails
- **Solution**: Check Dockerfile and .dockerignore

**Issue**: Deployment fails
- **Solution**: Verify environment secrets and permissions

## Best Practices

1. **Always run CI checks locally** before pushing
2. **Use feature branches** for development
3. **Require PR reviews** before merging
4. **Tag releases** for production deployments
5. **Monitor CI/CD** for failures
6. **Keep dependencies updated** via Dependabot

## Workflow Files

- `.github/workflows/ci.yaml` - Continuous Integration
- `.github/workflows/cd.yaml` - Continuous Deployment
- `.github/workflows/release.yaml` - Release automation
- `.github/dependabot.yaml` - Dependency updates
- `Dockerfile` - Container image definition
- `.dockerignore` - Docker build exclusions

