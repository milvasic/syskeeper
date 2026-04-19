# Security Policy

## Supported Versions

The following versions of syskeeper currently receive security updates:

| Version        | Supported |
| -------------- | --------- |
| `main`         | ✅ Yes    |
| Older releases | ❌ No     |

We recommend always running the latest version from the `main` branch until a formal release cycle is established.

## Reporting a Vulnerability

**Please do not report security vulnerabilities through public GitHub issues.**

If you discover a security vulnerability, please use one of the following private disclosure channels:

1. **GitHub Private Security Advisory** (preferred): Open a draft advisory at  
   https://github.com/milvasic/syskeeper/security/advisories/new
2. **Email**: Contact the project maintainer directly via the email address on their [GitHub profile](https://github.com/milvasic).

### What to Include

To help us triage and resolve the issue quickly, please include as much of the following as possible:

- A clear description of the vulnerability and its potential impact
- Steps to reproduce the issue (proof-of-concept code is welcome)
- Affected component(s): `keeper` (server), `keepee` (agent), or both
- Any proposed mitigation or fix

## Responsible Disclosure Process

1. **Report**: Submit your report using one of the channels above.
2. **Acknowledgement**: We will acknowledge receipt within **5 business days**.
3. **Assessment**: We will triage and assess severity within **10 business days**.
4. **Fix & Coordination**: We will work on a fix and coordinate a release. For critical issues, we may ask you to keep the vulnerability confidential until the fix is released.
5. **Disclosure**: Once a fix is released, we will publicly acknowledge the reporter (unless you prefer to remain anonymous) and document the CVE if applicable.

## Scope

The following are in scope:

- Authentication and authorization bypasses (API key handling, agent registration)
- Remote code execution or privilege escalation in either `keeper` or `keepee`
- SQL injection or data exposure via the PostgreSQL-backed `keeper`
- Insecure communication between `keepee` agents and the `keeper` server
- Sensitive data leakage (system information, credentials) via the dashboard or API

The following are **out of scope**:

- Vulnerabilities in third-party dependencies (please report those upstream)
- Issues in development/test environments not reproducible in production
- Social engineering attacks

## Security Best Practices for Deployers

- Always run `keeper` behind a reverse proxy (e.g., nginx/Caddy) with TLS termination.
- Use strong, randomly generated API keys for each registered `keepee` agent.
- Restrict network access so that only known `keepee` hosts can reach the `keeper` API.
- Keep the Go and .NET runtimes up to date.
- Follow the principle of least privilege when configuring the PostgreSQL user for `keeper`.
