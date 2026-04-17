# AI Guidelines for syskeeper

This document describes the policy on AI-assisted contributions to the **syskeeper** project. These guidelines apply to all contributors, whether they are human, AI-assisted, or a combination of both.

---

## Acceptable Use of AI Tools

AI coding assistants (such as GitHub Copilot, ChatGPT, Claude, Gemini, or similar tools) may be used to assist with contributions to syskeeper, provided the following conditions are met:

1. **You understand and own the code you submit.** AI-generated code must be reviewed, understood, and tested by the contributor before submission. You are responsible for everything in your PR — "the AI wrote it" is not an acceptable explanation for a bug or security issue.
2. **AI-generated code meets the same quality bar.** All code, regardless of origin, must pass tests, follow the project's coding conventions, and be reviewed by a maintainer.
3. **Security-sensitive areas require extra care.** Do not blindly accept AI-generated code in authentication, authorization, cryptography, input validation, or database query construction. These areas require human understanding and deliberate review.

---

## Disclosure Requirements

Contributors are encouraged (but not required) to disclose AI assistance in their PR descriptions. A simple note is sufficient:

> "Parts of this PR were drafted with AI assistance (GitHub Copilot) and reviewed by the author."

Disclosure helps reviewers calibrate their review depth and fosters trust in the community.

---

## Review Expectations for AI-Generated Code

Reviewers should apply additional scrutiny to code flagged as AI-generated, or to code that shows signs of AI generation (unusual verbosity, boilerplate-heavy patterns, or logic that doesn't quite fit the context):

- Verify correctness of logic, not just surface-level formatting.
- Watch for hallucinated APIs or library methods that do not exist.
- Check that error handling is appropriate and not silently swallowed.
- Confirm test coverage is meaningful, not just token assertions.

---

## What AI Tools Should Not Do

- **Generate secrets, credentials, or private keys** — never commit AI-generated secrets.
- **Replace human judgment on architecture decisions** — ADRs and significant design choices must reflect deliberate human reasoning.
- **Produce copyrighted content verbatim** — do not use AI to reproduce licensed code or documentation from third-party sources without proper attribution.

---

## GitHub Copilot Specific Notes

Project-specific Copilot instructions are maintained in [`.github/copilot-instructions.md`](.github/copilot-instructions.md). These instructions provide Copilot with context about the project's architecture, naming conventions, and coding style to produce more relevant suggestions.

---

## Questions

If you are unsure whether your use of AI tooling is appropriate for a given contribution, open a discussion or ask in the PR. We prefer transparency over guessing.
