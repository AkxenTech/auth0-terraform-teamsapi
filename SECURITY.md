# Security Guidelines

## Overview
This document outlines security best practices for the Auth0 Terraform Teams API project.

## Critical Security Rules

### 1. Never Commit Sensitive Files
The following files should NEVER be committed to the repository:
- `*.tfvars` - Contains variable values including API tokens
- `*.tfstate` - Contains sensitive infrastructure state
- `*.tfstate.backup` - Backup files containing sensitive data
- `.env` files - Environment variables
- Any files containing API tokens, passwords, or secrets

### 2. Sensitive Data Handling
- Use Terraform variables for sensitive inputs
- Mark sensitive outputs with `sensitive = true`
- Use environment variables or secure secret management
- Never hardcode credentials in source code

### 3. Git History Security
- Use `git filter-repo` to remove sensitive files from history
- Regularly audit git history for secrets
- Use pre-commit hooks to prevent accidental commits

## Incident Response

### If Secrets Are Accidentally Committed:
1. **IMMEDIATELY** rotate/revoke the exposed credentials
2. Use `git filter-repo` to remove from history:
   ```bash
   git filter-repo --path path/to/sensitive/file --invert-paths --force
   ```
3. Force push to remote repository
4. Notify team members to update their local copies

### Example Commands:
```bash
# Remove specific file from entire history
git filter-repo --path terraform.tfvars --invert-paths --force

# Remove multiple sensitive files
git filter-repo --path-glob "*.tfvars" --invert-paths --force
git filter-repo --path-glob "*.tfstate*" --invert-paths --force
```

## Security Checklist

Before committing:
- [ ] No API tokens in code
- [ ] No passwords in code
- [ ] No sensitive data in comments
- [ ] All sensitive outputs marked with `sensitive = true`
- [ ] Using variables for all sensitive inputs
- [ ] .gitignore properly configured

## Recent Security Incidents

### 2025-05-27: Exposed API Token and Client Secret
- **Issue**: Real API token and client secret committed to git history
- **Resolution**: Used git filter-repo to remove from entire history
- **Prevention**: Enhanced .gitignore and security documentation

## Contact
For security issues, contact the project maintainers immediately.
