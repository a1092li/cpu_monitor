## Contributing to CPU Monitor
We love your input! We want to make contributing to CPU Monitor as easy and transparent as possible, whether it's:
- Reporting a bug
- Discussing the current state of the code
- Submitting a fix
- Proposing new features
- Becoming a maintainer

## Development Process
We use GitHub to host code, to track issues and feature requests, as well as accept pull requests.
1. Fork the repo and create your branch from main
2. Make your changes
3. Add or update tests if necessary
4. Ensure the test suite passes
5. Make sure your code lints
6. Issue that pull request!

## Pull Request Process
1. Update the README.md with details of changes if applicable
2. Update the documentation if you're changing functionality
3. The PR will be merged once you have the sign-off of at least one maintainer

## Any Contributions You Make Will Be Under the MIT License
In short, when you submit code changes, your submissions are understood to be under the same MIT License that covers the project.

## Code of Conduct
### Our Pledge
We as members, contributors, and leaders pledge to make participation in our community a harassment-free experience for everyone.

### Our Standards
Examples of behavior that contributes to a positive environment:
- Using welcoming and inclusive language
- Being respectful of differing viewpoints and experiences
- Gracefully accepting constructive criticism
- Focusing on what is best for the community
- Showing empathy towards other community members

## Style Guide
### Go Code
Please follow standard Go conventions:
- Run gofmt on your code before committing
- Follow Go Code Review Comments
- Write clear comments for exported functions and types
- Use descriptive variable names

### Commit Messages
- Use the present tense ("Add feature" not "Added feature")
- Use the imperative mood ("Move cursor to..." not "Moves cursor to...")
- Limit the first line to 72 characters or less
- Reference issues and pull requests liberally after the first line

### Example Commit Message
```text
Add support for custom refresh intervals

- Add -i flag for interval configuration
- Update help text with new option
- Add validation for interval values

Fixes #123
```
## Setting Up Development Environment
1. Fork the repository on GitHub
2. Clone your fork locally:

```bash
git clone https://github.com/yourusername/cpu-monitor.git
cd cpu-monitor
```
3. Add upstream remote:

```bash
git remote add upstream https://github.com/originalusername/cpu-monitor.git
```
4. Create a feature branch:

```bash
git checkout -b feature/amazing-feature
```
5. Make your changes and test:

```bash
make build
make test
```
6. Push to your fork and submit a Pull Request

## Testing
Please add tests for any new functionality. To run the test suite:
```bash
make test
``
## Reporting Bugs
We use GitHub issues to track public bugs. Report a bug by opening a new issue.

## Write Good Bug Reports
Great bug reports tend to have:
- A quick summary and/or background
- Steps to reproduce
* Be specific!
* Give sample code if you can
- What you expected would happen
- What actually happens
- Notes (possibly including why you think this might be happening, or stuff you tried that didn't work)

## Feature Requests
We love feature requests! Please open an issue with:
- Use case description
- Proposed solution
- Alternative solutions considered
- Additional context

## License
By contributing, you agree that your contributions will be licensed under the project's MIT License.

## Questions?
Feel free to open an issue with your question or contact the maintainers directly.
Thank you for contributing to CPU Monitor! 🎉