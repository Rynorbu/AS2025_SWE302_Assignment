# Assignment 2: Task Tracker

**Student:** [Your Name]  
**Started:** [Date/Time]  
**Deadline:** November 30, 2025, 11:59 PM  
**Status:** Not Started

---

## ⏱️ Time Tracking

| Phase | Estimated | Actual | Status |
|-------|-----------|--------|--------|
| Setup | 30-45 min | | ⬜ |
| Snyk Scans | 2-3 hours | | ⬜ |
| SonarQube | 2-3 hours | | ⬜ |
| ZAP Testing | 4-5 hours | | ⬜ |
| Fixes | 2-3 hours | | ⬜ |
| Documentation | 1-2 hours | | ⬜ |
| **Total** | **12-16 hrs** | | |

---

## 📋 Phase 1: Setup & Prerequisites

### Environment Setup
- [ ] Installed Node.js and npm
- [ ] Installed Go
- [ ] Installed Git
- [ ] Installed Snyk CLI: `npm install -g snyk`
- [ ] Created Snyk account at https://snyk.io/
- [ ] Authenticated Snyk: `snyk auth`
- [ ] Setup SonarQube Cloud account
- [ ] Downloaded OWASP ZAP
- [ ] Ran setup script: `.\assignment2_script.ps1 -Task setup`

### Application Verification
- [ ] Backend runs successfully: `cd golang-gin-realworld-example-app; go run hello.go`
- [ ] Frontend runs successfully: `cd react-redux-realworld-example-app; npm start`
- [ ] Can access frontend at http://localhost:4100
- [ ] Can access backend API at http://localhost:8080/api

**Notes:**
```
[Add any setup issues or notes here]
```

---

## 📋 Phase 2: SAST with Snyk

### Backend Scan
- [ ] Navigated to backend directory
- [ ] Ran: `snyk test`
- [ ] Generated JSON report: `snyk test --json > snyk-backend-report.json`
- [ ] Ran: `snyk test --all-projects`
- [ ] Ran: `snyk monitor`
- [ ] Created `snyk-backend-analysis.md`
- [ ] Documented vulnerability summary
- [ ] Analyzed critical/high severity issues
- [ ] Listed affected dependencies

**Findings Summary:**
- Total vulnerabilities: ___
- Critical: ___
- High: ___
- Medium: ___
- Low: ___

### Frontend Scan
- [ ] Navigated to frontend directory
- [ ] Ran: `snyk test`
- [ ] Generated dependency report: `snyk test --json > snyk-frontend-report.json`
- [ ] Ran: `snyk code test`
- [ ] Generated code report: `snyk code test --json > snyk-code-report.json`
- [ ] Ran: `snyk monitor`
- [ ] Created `snyk-frontend-analysis.md`
- [ ] Documented dependency vulnerabilities
- [ ] Analyzed code vulnerabilities
- [ ] Identified React-specific issues

**Findings Summary:**
- Total vulnerabilities: ___
- Critical: ___
- High: ___
- XSS issues: ___
- Hardcoded secrets: ___

### Remediation
- [ ] Created `snyk-remediation-plan.md`
- [ ] Prioritized critical issues (score > 7.0)
- [ ] Listed high priority issues (4.0-7.0)
- [ ] Documented medium/low issues
- [ ] Planned dependency updates

### Fixes Implementation
- [ ] Fixed vulnerability 1: ________________
- [ ] Fixed vulnerability 2: ________________
- [ ] Fixed vulnerability 3: ________________
- [ ] Updated `package.json` with new versions
- [ ] Updated `go.mod` with new versions
- [ ] Tested application after updates
- [ ] Re-ran Snyk scans to verify fixes
- [ ] Took screenshots of Snyk dashboard
- [ ] Created `snyk-fixes-applied.md`

**Notes:**
```
[Add details of fixes applied]
```

---

## 📋 Phase 3: SAST with SonarQube

### Backend Analysis
- [ ] Connected GitHub repo to SonarQube Cloud
- [ ] Configured backend project
- [ ] Triggered analysis (via GitHub Action or manual)
- [ ] Waited for analysis to complete
- [ ] Reviewed Quality Gate status
- [ ] Captured dashboard screenshot
- [ ] Captured issues screenshot
- [ ] Captured security hotspots screenshot
- [ ] Captured code coverage screenshot
- [ ] Created `sonarqube-backend-analysis.md`

**Metrics:**
- Lines of Code: ___
- Quality Gate: Pass/Fail
- Bugs: ___
- Vulnerabilities: ___
- Code Smells: ___
- Security Hotspots: ___
- Technical Debt: ___

### Frontend Analysis
- [ ] Configured frontend project in SonarQube
- [ ] Triggered analysis
- [ ] Waited for analysis to complete
- [ ] Reviewed Quality Gate status
- [ ] Captured dashboard screenshot
- [ ] Captured issues screenshot
- [ ] Created `sonarqube-frontend-analysis.md`

**Metrics:**
- Lines of Code: ___
- Quality Gate: Pass/Fail
- Bugs: ___
- Vulnerabilities: ___
- Code Smells: ___
- Security Hotspots: ___

### Security Hotspots
- [ ] Reviewed each security hotspot
- [ ] Assessed real vs false positives
- [ ] Evaluated risk levels
- [ ] Created `security-hotspots-review.md`

### Improvements
- [ ] Fixed identified bugs
- [ ] Resolved security vulnerabilities
- [ ] Reduced code complexity
- [ ] Removed code duplications
- [ ] Re-ran SonarQube analysis
- [ ] Verified improvements
- [ ] Created `sonarqube-improvements.md`
- [ ] Documented before/after metrics

**Notes:**
```
[Add details of improvements made]
```

---

## 📋 Phase 4: DAST with OWASP ZAP

### Preparation
- [ ] Opened OWASP ZAP application
- [ ] Started backend server
- [ ] Started frontend server
- [ ] Verified both are running
- [ ] Navigated to http://localhost:4100
- [ ] Registered test user: security-test@example.com
- [ ] Created 2-3 sample articles
- [ ] Added some comments
- [ ] Documented test credentials

**Test Account:**
- Email: security-test@example.com
- Password: SecurePass123!

### Passive Scan
- [ ] Selected "Automated Scan" in ZAP
- [ ] Entered URL: http://localhost:4100
- [ ] Selected "Use traditional spider"
- [ ] Started scan
- [ ] Waited for completion (_____ minutes)
- [ ] Reviewed alerts
- [ ] Exported HTML report: `zap-passive-report.html`
- [ ] Created `zap-passive-scan-analysis.md`

**Findings:**
- Total alerts: ___
- High: ___
- Medium: ___
- Low: ___
- Informational: ___

### Active Scan (Authenticated)
- [ ] Created context: "Conduit Authenticated"
- [ ] Added URLs to context
- [ ] Configured JSON-based authentication
- [ ] Set login URL and request body
- [ ] Configured token extraction
- [ ] Set Authorization header
- [ ] Added test user
- [ ] Enabled user for context
- [ ] Spider with authenticated user
- [ ] Started active scan
- [ ] Waited for completion (_____ minutes)
- [ ] Exported HTML report: `zap-active-report.html`
- [ ] Exported XML report: `zap-active-report.xml`
- [ ] Exported JSON report: `zap-active-report.json`
- [ ] Created `zap-active-scan-analysis.md`

**Findings:**
- Total vulnerabilities: ___
- Critical: ___
- High: ___
- Medium: ___
- Low: ___

**Vulnerabilities Found:**
- [ ] SQL Injection: ___
- [ ] Cross-Site Scripting (XSS): ___
- [ ] Security Misconfiguration: ___
- [ ] Sensitive Data Exposure: ___
- [ ] Broken Authentication: ___
- [ ] IDOR: ___
- [ ] Missing Access Control: ___
- [ ] CSRF: ___
- [ ] Known Vulnerable Components: ___
- [ ] Unvalidated Redirects: ___

### API Security Testing
- [ ] Tested authentication bypass
- [ ] Tested authorization flaws
- [ ] Tested input validation (SQL injection, XSS)
- [ ] Tested rate limiting
- [ ] Checked for information disclosure
- [ ] Created `zap-api-security-analysis.md`

**API Issues Found:**
```
[List API-specific vulnerabilities]
```

### Security Fixes
- [ ] Prioritized critical/high vulnerabilities
- [ ] Fixed vulnerability 1: ________________
- [ ] Fixed vulnerability 2: ________________
- [ ] Fixed vulnerability 3: ________________
- [ ] Tested fixes locally
- [ ] Created `zap-fixes-applied.md`

### Security Headers
- [ ] Added security headers to backend (hello.go)
- [ ] Implemented X-Frame-Options
- [ ] Implemented X-Content-Type-Options
- [ ] Implemented X-XSS-Protection
- [ ] Implemented Strict-Transport-Security
- [ ] Implemented Content-Security-Policy
- [ ] Implemented Referrer-Policy
- [ ] Tested headers are present
- [ ] Took screenshots in ZAP
- [ ] Created `security-headers-analysis.md`

### Final Verification
- [ ] Re-ran ZAP passive scan
- [ ] Re-ran ZAP active scan
- [ ] Compared before/after results
- [ ] Verified critical/high issues resolved
- [ ] Documented remaining issues
- [ ] Created `final-security-assessment.md`

**Notes:**
```
[Add details of final assessment]
```

---

## 📋 Phase 5: Documentation & Submission

### Analysis Documents
- [ ] `snyk-backend-analysis.md` - Complete with all sections
- [ ] `snyk-frontend-analysis.md` - Complete with all sections
- [ ] `snyk-remediation-plan.md` - Prioritized plan
- [ ] `snyk-fixes-applied.md` - Detailed fixes
- [ ] `sonarqube-backend-analysis.md` - With screenshots
- [ ] `sonarqube-frontend-analysis.md` - With screenshots
- [ ] `security-hotspots-review.md` - Each hotspot reviewed
- [ ] `sonarqube-improvements.md` - Before/after comparison
- [ ] `zap-passive-scan-analysis.md` - All findings documented
- [ ] `zap-active-scan-analysis.md` - Comprehensive analysis
- [ ] `zap-api-security-analysis.md` - API-specific issues
- [ ] `zap-fixes-applied.md` - Fixes with evidence
- [ ] `security-headers-analysis.md` - Headers explained
- [ ] `final-security-assessment.md` - Complete assessment

### Reports & Exports
- [ ] `snyk-backend-report.json`
- [ ] `snyk-frontend-report.json`
- [ ] `snyk-code-report.json`
- [ ] `zap-passive-report.html`
- [ ] `zap-active-report.html`
- [ ] `zap-active-report.xml`
- [ ] `zap-active-report.json`

### Screenshots
- [ ] Snyk dashboard - backend
- [ ] Snyk dashboard - frontend
- [ ] SonarQube dashboard - backend
- [ ] SonarQube issues - backend
- [ ] SonarQube security hotspots - backend
- [ ] SonarQube coverage - backend
- [ ] SonarQube dashboard - frontend
- [ ] SonarQube issues - frontend
- [ ] ZAP passive scan results
- [ ] ZAP active scan results
- [ ] ZAP security headers verification
- [ ] Before/after comparisons

### Code Changes
- [ ] All security fixes committed
- [ ] Updated `package.json` (frontend)
- [ ] Updated `go.mod` (backend)
- [ ] Security headers in `hello.go`
- [ ] Application tested and working

### Summary Report
- [ ] Created `ASSIGNMENT_2_REPORT.md`
- [ ] Executive summary written
- [ ] SAST findings summarized
- [ ] DAST findings summarized
- [ ] Remediation actions documented
- [ ] Remaining risks listed
- [ ] Lessons learned included
- [ ] Professional formatting applied

### Final Review
- [ ] All deliverables present
- [ ] All screenshots captured
- [ ] All reports exported
- [ ] Documentation is clear
- [ ] Formatting is professional
- [ ] No placeholders or TODOs
- [ ] Grammar and spelling checked
- [ ] Ready for submission

---

## 📊 Points Checklist (100 Total)

| Component | Points | Status |
|-----------|--------|--------|
| Snyk Backend Analysis | 8 | ⬜ |
| Snyk Frontend Analysis | 8 | ⬜ |
| SonarQube Backend | 8 | ⬜ |
| SonarQube Frontend | 8 | ⬜ |
| SonarQube Improvements | 10 | ⬜ |
| ZAP Passive Scan | 8 | ⬜ |
| ZAP Active Scan | 15 | ⬜ |
| ZAP API Testing | 10 | ⬜ |
| Security Fixes | 15 | ⬜ |
| Security Headers | 5 | ⬜ |
| Documentation | 5 | ⬜ |
| **TOTAL** | **100** | |

---

## 🎯 Quality Checklist

### Analysis Documents Quality
- [ ] Vulnerability counts are accurate
- [ ] Severity ratings are correct
- [ ] CVE/CWE numbers included
- [ ] OWASP categories referenced
- [ ] Descriptions are clear
- [ ] Remediation steps are specific
- [ ] Evidence/screenshots included

### Fixes Quality
- [ ] At least 3 critical/high issues fixed
- [ ] Application still works after fixes
- [ ] Fixes are properly tested
- [ ] Before/after evidence provided
- [ ] Code changes are documented

### Documentation Quality
- [ ] Professional formatting
- [ ] Clear and concise
- [ ] No spelling/grammar errors
- [ ] Proper markdown syntax
- [ ] Tables formatted correctly
- [ ] Code blocks properly formatted
- [ ] Screenshots are clear and relevant

---

## 📝 Notes & Issues

### Issues Encountered
```
[Document any problems you faced and how you solved them]
```

### Key Learnings
```
[Document important things you learned]
```

### Time Spent
```
Phase 1 (Setup): ___ hours
Phase 2 (Snyk): ___ hours
Phase 3 (SonarQube): ___ hours
Phase 4 (ZAP): ___ hours
Phase 5 (Documentation): ___ hours
Total: ___ hours
```

---

## ✅ Submission Checklist

### Pre-Submission
- [ ] All files are in correct locations
- [ ] All deliverables are complete
- [ ] All screenshots are included
- [ ] All reports are exported
- [ ] Code changes are committed
- [ ] Everything is properly named
- [ ] No missing information
- [ ] Quality review completed

### Submission
- [ ] Verified deadline: November 30, 2025, 11:59 PM
- [ ] Prepared submission package
- [ ] Double-checked all requirements
- [ ] Submitted assignment
- [ ] Received submission confirmation

**Submission Time:** _______________

---

## 🎉 Completion

**Assignment Status:** ⬜ Not Started | ⬜ In Progress | ⬜ Completed | ⬜ Submitted

**Final Grade (when received):** _____/100

**Instructor Feedback:**
```
[Add feedback here when received]
```

---

*Keep this file updated as you progress through the assignment!*
