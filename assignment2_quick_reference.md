# Assignment 2: Quick Reference Guide

## 📚 Overview
This guide provides quick commands and references for completing Assignment 2.

---

## 🔧 Initial Setup Commands

### 1. Install Snyk
```powershell
npm install -g snyk
snyk auth
```

### 2. Verify Prerequisites
```powershell
node --version
npm --version
go version
git --version
snyk --version
```

### 3. Install Frontend Dependencies
```powershell
cd react-redux-realworld-example-app
npm install
```

---

## 🔍 SAST - Snyk Commands

### Backend (Go)
```powershell
cd golang-gin-realworld-example-app

# Basic scan
snyk test

# Generate JSON report
snyk test --json > snyk-backend-report.json

# Scan all projects
snyk test --all-projects

# Monitor project (upload to dashboard)
snyk monitor

# View help
snyk test --help
```

### Frontend (React)
```powershell
cd react-redux-realworld-example-app

# Dependency scan
snyk test

# Generate dependency report
snyk test --json > snyk-frontend-report.json

# Code analysis (SAST)
snyk code test

# Generate code analysis report
snyk code test --json > snyk-code-report.json

# Monitor project
snyk monitor
```

### Useful Snyk Options
```powershell
# Show only high and critical
snyk test --severity-threshold=high

# Scan specific file
snyk test --file=package.json

# Skip unresolved dependencies
snyk test --skip-unresolved

# Output as JSON
snyk test --json

# Scan for license issues
snyk test --show-vulnerable-paths=all
```

---

## 🔍 SAST - SonarQube Setup

### Cloud Setup (Recommended)
1. Go to: https://sonarqube.cloud/
2. Sign up with GitHub account
3. Create new organization
4. Click "Analyze new project"
5. Select GitHub repository
6. Follow integration instructions
7. Install SonarQube GitHub Action (automatic analysis)

### Key URLs
- SonarQube Cloud: https://sonarqube.cloud/
- Documentation: https://docs.sonarsource.com/sonarqube-cloud/

### What to Capture
- Dashboard screenshot
- Issues tab screenshot
- Security Hotspots screenshot
- Code coverage screenshot
- Quality Gate status

---

## 🔍 DAST - OWASP ZAP Setup

### Download ZAP
- Website: https://www.zaproxy.org/download/
- Choose installer for Windows
- Install and launch application

### Start Applications for Testing
```powershell
# Terminal 1 - Backend
cd golang-gin-realworld-example-app
go run hello.go

# Terminal 2 - Frontend
cd react-redux-realworld-example-app
npm start
```

### Create Test Account
1. Navigate to: http://localhost:4100
2. Click "Sign up"
3. Email: `security-test@example.com`
4. Username: `securitytest`
5. Password: `SecurePass123!`
6. Create 2-3 sample articles
7. Add some comments

---

## 🔍 DAST - ZAP Testing Guide

### Passive Scan
1. Open OWASP ZAP
2. Automated Scan → Enter URL: `http://localhost:4100`
3. Select "Use traditional spider"
4. Click "Attack"
5. Wait for completion
6. Export report: Report → Generate HTML Report

### Active Scan (Authenticated)

#### Step 1: Create Context
1. Right-click on `http://localhost:4100` in Sites tree
2. Include in Context → New Context
3. Name: "Conduit Authenticated"
4. Add to Include in Context: `http://localhost:4100.*`
5. Add to Include in Context: `http://localhost:8080/api.*`

#### Step 2: Configure Authentication
1. Right-click Context → Edit Context
2. Go to Authentication tab
3. Method: "JSON-based Authentication"
4. Login URL: `http://localhost:8080/api/users/login`
5. Login Request POST Data:
```json
{
  "user": {
    "email": "security-test@example.com",
    "password": "SecurePass123!"
  }
}
```
6. Username parameter: `email`
7. Password parameter: `password`
8. Logged in regex: `.*token.*`
9. Logged out regex: Leave empty

#### Step 3: Configure Session Management
1. Stay in Context → Session Management tab
2. Method: "HTTP Authentication"
3. Go to Context → Session Properties
4. Add session token: `Authorization`
5. Token location: Header
6. Token value from response: `user.token`

#### Step 4: Add User
1. Go to Users tab in Context
2. Add user:
   - Username: `security-test@example.com`
   - Password: `SecurePass123!`
3. Enable user

#### Step 5: Run Active Scan
1. Right-click on site → Attack → Active Scan
2. User: Select `security-test@example.com`
3. Policy: OWASP Top 10
4. Start Scan
5. Wait 30-60 minutes for completion

### Export ZAP Reports
```
Report → Generate HTML Report → Save as zap-active-report.html
Report → Generate XML Report → Save as zap-active-report.xml
Report → Generate JSON Report → Save as zap-active-report.json
```

---

## 🔧 Security Fixes

### Add Security Headers to Backend

Edit `golang-gin-realworld-example-app/hello.go`:

```go
// After creating router, before defining routes:
router.Use(func(c *gin.Context) {
    // Prevent clickjacking
    c.Header("X-Frame-Options", "DENY")
    
    // Prevent MIME sniffing
    c.Header("X-Content-Type-Options", "nosniff")
    
    // Enable XSS protection
    c.Header("X-XSS-Protection", "1; mode=block")
    
    // Force HTTPS
    c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
    
    // Content Security Policy
    c.Header("Content-Security-Policy", "default-src 'self'")
    
    // Referrer Policy
    c.Header("Referrer-Policy", "no-referrer-when-downgrade")
    
    // Permissions Policy
    c.Header("Permissions-Policy", "geolocation=(), microphone=(), camera=()")
    
    c.Next()
})
```

### Update Vulnerable Dependencies

#### Backend (Go)
```powershell
cd golang-gin-realworld-example-app

# Check for updates
go list -u -m all

# Update specific package
go get -u github.com/package/name@version

# Update all dependencies
go get -u ./...

# Tidy up
go mod tidy
```

#### Frontend (React)
```powershell
cd react-redux-realworld-example-app

# Check for updates
npm outdated

# Update specific package
npm install package-name@version

# Update all packages (careful with breaking changes)
npm update

# Use npm-check-updates for major versions
npx npm-check-updates
npx npm-check-updates -u
npm install
```

---

## 📊 Analysis Document Templates

### Structure for Analysis Documents

#### Snyk Analysis
```markdown
# Snyk [Backend/Frontend] Analysis

## 1. Vulnerability Summary
- Total: X
- Critical: X
- High: X
- Medium: X
- Low: X

## 2. Critical/High Issues
### Issue 1
- Severity:
- Package:
- CVE:
- Description:
- Fix:

## 3. Recommendations
```

#### SonarQube Analysis
```markdown
# SonarQube [Backend/Frontend] Analysis

## 1. Quality Gate Status
- Pass/Fail:
- Conditions:

## 2. Metrics
- Lines of Code:
- Coverage:
- Duplications:
- Complexity:

## 3. Issues
- Bugs:
- Vulnerabilities:
- Code Smells:
- Security Hotspots:

## 4. Screenshots
```

#### ZAP Analysis
```markdown
# OWASP ZAP [Passive/Active] Scan Analysis

## 1. Summary
- Total Alerts:
- High: X
- Medium: X
- Low: X

## 2. Critical Findings
### Finding 1
- Alert:
- Risk:
- URLs:
- CWE:
- Description:
- Solution:

## 3. Evidence
```

---

## 🎯 API Endpoints to Test

### Authentication
```
POST   /api/users                 # Register
POST   /api/users/login           # Login
GET    /api/user                  # Current User
PUT    /api/user                  # Update User
```

### Profiles
```
GET    /api/profiles/:username
POST   /api/profiles/:username/follow
DELETE /api/profiles/:username/follow
```

### Articles
```
GET    /api/articles              # List
POST   /api/articles              # Create
GET    /api/articles/feed         # Feed
GET    /api/articles/:slug        # Get
PUT    /api/articles/:slug        # Update
DELETE /api/articles/:slug        # Delete
POST   /api/articles/:slug/favorite
DELETE /api/articles/:slug/favorite
```

### Comments
```
GET    /api/articles/:slug/comments
POST   /api/articles/:slug/comments
DELETE /api/articles/:slug/comments/:id
```

### Tags
```
GET    /api/tags
```

---

## 🧪 Testing Scenarios

### Test for SQL Injection
```
# Try in article title, username, tags
' OR '1'='1
'; DROP TABLE users--
1' UNION SELECT NULL--
```

### Test for XSS
```
# Try in article content, comments, bio
<script>alert('XSS')</script>
<img src=x onerror=alert('XSS')>
<svg onload=alert('XSS')>
```

### Test for Authorization
```
# Try to access/modify other users' resources
GET /api/articles/other-user-article
PUT /api/articles/other-user-article
DELETE /api/articles/other-user-article
```

### Test for Authentication Bypass
```
# Try accessing protected endpoints without token
GET /api/user
# (without Authorization header)

# Try with invalid token
Authorization: Token invalid-token-here
```

---

## 📋 Deliverables Checklist

### Snyk Deliverables
- [ ] snyk-backend-analysis.md
- [ ] snyk-frontend-analysis.md
- [ ] snyk-remediation-plan.md
- [ ] snyk-fixes-applied.md
- [ ] snyk-backend-report.json
- [ ] snyk-frontend-report.json
- [ ] snyk-code-report.json
- [ ] Screenshots of Snyk dashboard

### SonarQube Deliverables
- [ ] sonarqube-backend-analysis.md
- [ ] sonarqube-frontend-analysis.md
- [ ] security-hotspots-review.md
- [ ] sonarqube-improvements.md
- [ ] Screenshots of all dashboards

### ZAP Deliverables
- [ ] zap-passive-scan-analysis.md
- [ ] zap-active-scan-analysis.md
- [ ] zap-api-security-analysis.md
- [ ] zap-fixes-applied.md
- [ ] security-headers-analysis.md
- [ ] final-security-assessment.md
- [ ] zap-passive-report.html
- [ ] zap-active-report.html
- [ ] zap-active-report.xml
- [ ] zap-active-report.json

### Summary
- [ ] ASSIGNMENT_2_REPORT.md
- [ ] All code changes committed
- [ ] Updated package.json / go.mod

---

## ⚡ Quick Tips

### Time-Saving Tips
1. Run Snyk scans first while reading documentation
2. Setup SonarQube Cloud early (analysis takes time)
3. Create test account before starting ZAP
4. Take screenshots as you go
5. Document findings immediately

### Common Issues
1. **Snyk auth fails**: Clear cache with `snyk config clear`
2. **ZAP can't authenticate**: Double-check JSON format in login request
3. **Active scan too slow**: Reduce scan policy or use fewer rules
4. **Go dependencies fail**: Run `go mod tidy` first
5. **npm install fails**: Delete node_modules and package-lock.json, retry

### Best Practices
1. Always test application after making fixes
2. Keep backup of working code
3. Fix one issue at a time
4. Document what breaks and why
5. Use version control (git) for changes

---

## 🔗 Important Links

### Tools
- Snyk: https://snyk.io/
- SonarQube Cloud: https://sonarqube.cloud/
- OWASP ZAP: https://www.zaproxy.org/

### Documentation
- Snyk Docs: https://docs.snyk.io/
- SonarQube Docs: https://docs.sonarsource.com/
- ZAP Docs: https://www.zaproxy.org/docs/
- OWASP Top 10: https://owasp.org/www-project-top-ten/
- CWE Database: https://cwe.mitre.org/

### Tutorials
- ZAP Authentication: https://www.zaproxy.org/docs/desktop/start/features/authentication/
- Snyk Code: https://docs.snyk.io/products/snyk-code
- SonarQube Cloud: https://docs.sonarsource.com/sonarqube-cloud/

---

## 🆘 Getting Help

If stuck:
1. Check tool documentation
2. Review error messages carefully
3. Search for specific error on GitHub/Stack Overflow
4. Check if services are running (backend/frontend)
5. Verify authentication tokens
6. Review this quick reference guide

---

**Remember:** The goal is to learn security testing, not to achieve perfection. Document your findings honestly and focus on understanding the vulnerabilities.

Good luck! 🚀
