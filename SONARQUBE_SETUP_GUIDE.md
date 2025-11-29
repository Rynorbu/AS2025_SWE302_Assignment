# SonarQube Cloud Setup Guide
**Date:** 2025-11-29  
**Project:** RealWorld Conduit (Backend + Frontend)

---

## Overview

This guide walks you through setting up SonarQube Cloud for both the backend (Go) and frontend (React) projects.

---

## Prerequisites

- [x] GitHub repository: AS2025_SWE302_Assignment
- [x] GitHub account: Rynorbu
- [ ] SonarQube Cloud account
- [ ] Admin access to GitHub repository

---

## Step 1: Create SonarQube Cloud Account

### 1.1 Sign Up
1. Go to: https://sonarqube.cloud/
2. Click **"Sign up"** or **"Log in with GitHub"**
3. Choose: **"Sign up with GitHub"**
4. Authorize SonarQube Cloud to access your GitHub account
5. Accept the terms and conditions

### 1.2 Create Organization
1. After login, you'll be prompted to create an organization
2. Choose: **"Import an organization from GitHub"**
3. Select your GitHub account: **Rynorbu**
4. Choose organization plan: **"Free Plan"** (perfect for this assignment)
5. Create organization

---

## Step 2: Setup Backend Project (Go)

### 2.1 Add New Project
1. In SonarQube Cloud, click **"Analyze new project"**
2. Select: **"AS2025_SWE302_Assignment"** repository
3. Click **"Set Up"**

### 2.2 Configure Analysis Method
1. Choose: **"With GitHub Actions"** (recommended) OR **"Other CI"**
2. For this assignment, we'll use **"Locally"** for quick results

#### Option A: Local Analysis (Quick - Recommended for Assignment)
1. Choose: **"Locally"**
2. Select project key name: `AS2025_SWE302_Assignment-backend`
3. Click **"Generate a token"**
   - Token name: `swe302-backend-token`
   - Copy and save the token (you won't see it again!)
4. Select build technology: **"Other"** (for Go)
5. Copy the provided command

**Example command:**
```bash
sonar-scanner \
  -Dsonar.projectKey=AS2025_SWE302_Assignment-backend \
  -Dsonar.sources=golang-gin-realworld-example-app \
  -Dsonar.host.url=https://sonarqube.cloud \
  -Dsonar.token=YOUR_TOKEN_HERE
```

#### Option B: GitHub Actions (Better for Production)
1. Choose: **"With GitHub Actions"**
2. Follow the setup wizard:
   - Add SONAR_TOKEN to GitHub Secrets
   - Create `.github/workflows/sonarqube.yml`
   - Push changes to trigger analysis

### 2.3 Install SonarScanner (for Local Analysis)

**For Windows:**
1. Download: https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-4.8.0.2856-windows.zip
2. Extract to: `C:\sonar-scanner`
3. Add to PATH:
   ```powershell
   $env:PATH = "C:\sonar-scanner\bin;$env:PATH"
   ```
4. Verify: `sonar-scanner --version`

**Alternative - Use Docker:**
```bash
docker pull sonarsource/sonar-scanner-cli
```

### 2.4 Create sonar-project.properties for Backend

Create file: `golang-gin-realworld-example-app/sonar-project.properties`

```properties
# Project identification
sonar.projectKey=AS2025_SWE302_Assignment-backend
sonar.projectName=RealWorld Conduit Backend
sonar.projectVersion=1.0

# Path configuration
sonar.sources=.
sonar.exclusions=**/*_test.go,**/vendor/**,**/*.pb.go

# Go specific
sonar.go.coverage.reportPaths=coverage.out
sonar.go.tests.reportPaths=test-report.json

# SonarQube server
sonar.host.url=https://sonarqube.cloud
sonar.organization=rynorbu

# Authentication (use environment variable)
# sonar.token will be provided via command line or environment
```

### 2.5 Run Backend Analysis

```bash
cd golang-gin-realworld-example-app

# Generate coverage report
$env:CGO_ENABLED=1
go test ./... -coverprofile=coverage.out

# Run SonarQube analysis
sonar-scanner -Dsonar.token=YOUR_TOKEN_HERE
```

---

## Step 3: Setup Frontend Project (React)

### 3.1 Add Frontend as New Project
1. Click **"Analyze new project"** again
2. Select the same repository or add as separate project
3. Project key: `AS2025_SWE302_Assignment-frontend`

### 3.2 Create sonar-project.properties for Frontend

Create file: `react-redux-realworld-example-app/sonar-project.properties`

```properties
# Project identification
sonar.projectKey=AS2025_SWE302_Assignment-frontend
sonar.projectName=RealWorld Conduit Frontend
sonar.projectVersion=1.0

# Path configuration
sonar.sources=src
sonar.exclusions=**/node_modules/**,**/*.test.js,**/*.spec.js,**/coverage/**,**/build/**

# JavaScript/TypeScript specific
sonar.javascript.lcov.reportPaths=coverage/lcov.info
sonar.testExecutionReportPaths=test-report.xml

# SonarQube server
sonar.host.url=https://sonarqube.cloud
sonar.organization=rynorbu
```

### 3.3 Run Frontend Analysis

```bash
cd react-redux-realworld-example-app

# Generate coverage report
npm test -- --coverage --watchAll=false

# Run SonarQube analysis
sonar-scanner -Dsonar.token=YOUR_TOKEN_HERE
```

---

## Step 4: View Results in Dashboard

### 4.1 Access Dashboard
1. Go to: https://sonarqube.cloud/
2. Click on your organization: **Rynorbu**
3. You should see both projects:
   - AS2025_SWE302_Assignment-backend
   - AS2025_SWE302_Assignment-frontend

### 4.2 What to Look For

**For Each Project, Review:**

1. **Overview Tab:**
   - Quality Gate status (Passed/Failed)
   - Reliability rating (A-E)
   - Security rating (A-E)
   - Maintainability rating (A-E)
   - Coverage percentage
   - Duplications percentage

2. **Issues Tab:**
   - Bugs count and details
   - Vulnerabilities count and details
   - Code Smells count and details
   - Severity breakdown (Blocker, Critical, Major, Minor, Info)

3. **Security Hotspots Tab:**
   - Security-sensitive code to review
   - OWASP Top 10 categorization
   - Taint analysis results

4. **Measures Tab:**
   - Lines of code
   - Complexity metrics
   - Technical debt
   - Detailed metrics

---

## Step 5: Capture Screenshots

### Required Screenshots for Backend:
1. **Dashboard Overview** - showing Quality Gate and ratings
2. **Issues List** - showing bugs, vulnerabilities, code smells
3. **Security Hotspots** - if any found
4. **Code Tab** - showing specific code with issues highlighted
5. **Measures** - showing metrics and complexity

### Required Screenshots for Frontend:
1. **Dashboard Overview** - showing Quality Gate and ratings
2. **Issues List** - focusing on JavaScript/React issues
3. **Security Hotspots** - especially XSS-related
4. **Code Tab** - showing React components with issues
5. **Measures** - showing metrics

### How to Take Screenshots:
```powershell
# Use Windows Snipping Tool or
# Win + Shift + S for screen capture
# Save to: security-reports/screenshots/sonarqube-backend-*.png
```

---

## Step 6: Analysis Documents to Create

After analysis is complete, create these documents:

### Backend Analysis Document
File: `sonarqube-backend-analysis.md`

**Must include:**
- Quality Gate status
- Reliability rating and bugs found
- Security rating and vulnerabilities
- Maintainability rating and code smells
- Security hotspots details
- Complexity metrics
- Technical debt
- Screenshots

### Frontend Analysis Document
File: `sonarqube-frontend-analysis.md`

**Must include:**
- Quality Gate status
- JavaScript/React specific issues
- Security vulnerabilities
- Code smells
- Component complexity
- Screenshots

### Security Hotspots Review
File: `security-hotspots-review.md`

**Must include:**
- Each hotspot identified
- Risk assessment
- Real vs false positive
- Remediation recommendations

---

## Troubleshooting

### Issue: Token Authentication Failed
**Solution:**
```bash
# Set token as environment variable
$env:SONAR_TOKEN="your-token-here"
sonar-scanner -Dsonar.token=$env:SONAR_TOKEN
```

### Issue: SonarScanner Not Found
**Solution:**
```powershell
# Add to PATH
$env:PATH = "C:\sonar-scanner\bin;$env:PATH"

# Verify
sonar-scanner --version
```

### Issue: Go Coverage Not Showing
**Solution:**
```bash
# Ensure CGO is enabled
$env:CGO_ENABLED=1

# Generate coverage in correct format
go test ./... -coverprofile=coverage.out

# Verify coverage file exists
ls coverage.out
```

### Issue: Frontend Coverage Not Showing
**Solution:**
```bash
# Generate coverage in lcov format
npm test -- --coverage --watchAll=false --coverageReporters=lcov

# Verify coverage folder exists
ls coverage/lcov.info
```

---

## Quick Reference Commands

### Backend Analysis
```bash
cd golang-gin-realworld-example-app
$env:CGO_ENABLED=1
go test ./... -coverprofile=coverage.out
sonar-scanner -Dsonar.token=YOUR_TOKEN
```

### Frontend Analysis
```bash
cd react-redux-realworld-example-app
npm test -- --coverage --watchAll=false
sonar-scanner -Dsonar.token=YOUR_TOKEN
```

### View Dashboard
```
https://sonarqube.cloud/dashboard?id=AS2025_SWE302_Assignment-backend
https://sonarqube.cloud/dashboard?id=AS2025_SWE302_Assignment-frontend
```

---

## Expected Timeline

- **SonarQube Account Setup:** 5-10 minutes
- **Backend Project Setup:** 10-15 minutes
- **Backend Analysis Run:** 5-10 minutes
- **Frontend Project Setup:** 10-15 minutes
- **Frontend Analysis Run:** 5-10 minutes
- **Screenshot Capture:** 10-15 minutes
- **Total:** ~60 minutes

---

## Next Steps

After completing SonarQube setup and analysis:

1. ✅ Review dashboard results
2. ✅ Take all required screenshots
3. ✅ Create backend analysis document
4. ✅ Create frontend analysis document
5. ✅ Create security hotspots review
6. ✅ Implement high-priority fixes
7. ✅ Re-run analysis to verify improvements

---

## Resources

- **SonarQube Cloud:** https://sonarqube.cloud/
- **Documentation:** https://docs.sonarsource.com/sonarqube-cloud/
- **GitHub Integration:** https://docs.sonarsource.com/sonarqube-cloud/getting-started/github/
- **Go Analysis:** https://docs.sonarsource.com/sonarqube-cloud/enriching/languages/go/
- **JavaScript Analysis:** https://docs.sonarsource.com/sonarqube-cloud/enriching/languages/javascript/

---

**Note:** Since SonarQube Cloud requires manual browser interaction, please complete the setup steps above and let me know when you've:
1. Created your SonarQube Cloud account
2. Run the analysis for both projects
3. Have access to the dashboard

Then I can help you document the findings!
