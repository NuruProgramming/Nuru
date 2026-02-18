# Contributing to Nuru

Thank you for your interest in contributing. This document explains how to get started and what we expect from contributors.

## Table of contents

- [Code of conduct](#code-of-conduct)
- [Design policy](#design-policy)
- [What we accept](#what-we-accept)
- [Getting started](#getting-started)
- [Git workflow](#git-workflow)
- [Before submitting](#before-submitting)
- [Pull requests](#pull-requests)
- [Reporting bugs](#reporting-bugs)
- [Suggesting features](#suggesting-features)
- [Documentation](#documentation)
- [Security](#security)
- [Getting help](#getting-help)

## Code of conduct

By participating in this project, you agree to uphold our [Code of Conduct](CODE_OF_CONDUCT.md).

## Design policy

New features must follow our design policy: we keep a small set of **global builtins** (I/O, type conversion, range), use **modules** for domain and system features (faili, crypto, mfumo, hisabati, etc.), and use **methods** for operations on values (strings, arrays, dicts, files). When adding functionality, prefer methods or modules over new builtins. See [docs/POLICY.md](docs/POLICY.md) for the full guide.

## What we accept

We welcome:

- **Code** – bug fixes, new features (aligned with [docs/POLICY.md](docs/POLICY.md)), and tests
- **Documentation** – improvements and translations (English and Kiswahili) in `repl/docs/`
- **Bug reports** – via the Bug Report issue template
- **Feature requests** – via the Feature Request issue template

We prefer extending existing modules and adding methods on types over introducing new global builtins.

## Getting started

1. Clone the repository.
2. Ensure you have Go 1.21 or later.
3. Run `go mod tidy` to fetch dependencies.
4. Run the project or tests (see [Before submitting](#before-submitting)).

For full installation and usage, see the [README](README.md).

## Git workflow

We use a simple branch workflow (GitHub Flow):

- **Base branch:** `main`. Create your branch from `main`.
- **Pull requests:** Open your PR against `main`.
- **CI:** Tests run on push to `main` and `dev`. Your PR will run tests when you push.

Work on a descriptive branch (e.g. `fix-regex-error` or `docs-strings-methods`), then open a pull request to `main`.

## Before submitting

Run the test suite locally so CI can pass:

```bash
make test
```

You can also run the smoke example:

```bash
make smoke
```

## Pull requests

1. Use our [pull request template](.github/PULL_REQUEST_TEMPLATE): describe what your PR does and why.
2. **Title format:** `module: description` (e.g. `time: fix date format`). Use present tense, not past.
3. For new functions, add documentation and tests.
4. For bug fixes, include a short explanation of the issue and the fix.
5. Run `make test` before submitting.

Maintainers aim to review PRs within about 24 hours. You do not need to add labels.

## Reporting bugs

Use the [Bug Report](https://github.com/NuruProgramming/Nuru/issues/new?template=bug_report.yml) issue template. Include:

- A clear description of the bug
- Expected vs actual behavior
- Steps to reproduce (and a minimal code sample if possible)
- Your Nuru version and environment (OS, etc.)

Check [existing issues](https://github.com/NuruProgramming/Nuru/issues) first to avoid duplicates.

## Suggesting features

Use the [Feature Request](https://github.com/NuruProgramming/Nuru/issues/new?template=feature-request.yml) issue template. Describe the feature, your use case, and any proposed solution. Check existing issues and [docs/POLICY.md](docs/POLICY.md) to see where the feature would fit (builtin, module, or method).

## Documentation

Documentation lives under [repl/docs/](repl/docs/). We have two languages:

- **English:** `repl/docs/en/`
- **Kiswahili:** `repl/docs/sw/`

All docs are in Markdown. When you add or change language features, please update the relevant docs and consider both languages.

## Security

If you believe you have found a security vulnerability, please see [SECURITY.md](SECURITY.md) for how to report it.

## Getting help

- **Community:** [Telegram – Nuru Programming Chat](https://t.me/NuruProgrammingChat)
- **Issues:** [GitHub Issues](https://github.com/NuruProgramming/Nuru/issues) for bugs, features, and questions

Thanks for contributing.
