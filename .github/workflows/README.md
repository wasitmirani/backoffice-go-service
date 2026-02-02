# GitHub Actions Workflows

This directory contains GitHub Actions workflows for CI/CD automation.

## Workflows

### CI (`ci.yaml`)
- Runs on every push and pull request
- Performs linting, testing, building, and security scanning
- Builds for multiple platforms
- Generates test coverage reports

### CD (`cd.yaml`)
- Runs on pushes to main/master and on tags
- Builds and pushes Docker images
- Deploys to staging and production environments
- Creates GitHub releases

### Release (`release.yaml`)
- Runs when a GitHub release is published
- Builds binaries for all platforms
- Uploads release assets

## Usage

Workflows run automatically based on triggers. You can also manually trigger the CD workflow from the Actions tab.

## Configuration

- Update workflow files to match your deployment targets
- Configure GitHub secrets for deployment credentials
- Set up environment protection rules for staging/production

