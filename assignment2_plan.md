# Assignment 2 Implementation Plan
## Static & Dynamic Application Security Testing (SAST & DAST)

**Created:** November 29, 2025  
**Deadline:** November 30, 2025, 11:59 PM  
**Total Points:** 100

---

## 📋 Overview

This assignment requires performing comprehensive security testing on the RealWorld Conduit application using:
- **SAST Tools:** Snyk (dependency & code scanning) and SonarQube (code quality & security)
- **DAST Tool:** OWASP ZAP (dynamic security testing)

---

## 🎯 Phase-by-Phase Implementation Plan

### **PHASE 1: Environment Setup & Prerequisites** (30-45 minutes)

#### 1.1 Install Required Tools
- [ ] Install Snyk CLI globally: `npm install -g snyk`
- [ ] Create free Snyk account at https://snyk.io/
- [ ] Authenticate Snyk: `snyk auth`
- [ ] Setup SonarQube Cloud at https://docs.sonarsource.com/sonarqube-cloud/getting-started/github
- [ ] Download OWASP ZAP from https://www.zaproxy.org/download/

#### 1.2 Verify Applications Work
- [ ] Test backend: `cd golang-gin-realworld-example-app; go run hello.go`
- [ ] Test frontend: `cd react-redux-realworld-example-app; npm install; npm start`
- [ ] Verify both communicate properly

---

### **PHASE 2: SAST with Snyk** (2-3 hours)

#### Task 1.2: Backend Security Scan (Go)
**Time Estimate:** 45 minutes

**Actions:**
1. Navigate to backend directory
2. Run basic scan: `snyk test`
3. Generate JSON report: `snyk test --json > snyk-backend-report.json`
4. Test all projects: `snyk test --all-projects`
5. Monitor project: `snyk monitor`

**Deliverable:** `snyk-backend-analysis.md`
- Total vulnerabilities count
- Breakdown by severity (Critical, High, Medium, Low)
- Affected dependencies list
- Detailed analysis of Critical/High issues with CVE numbers
- Direct vs transitive dependencies
- Recommended fix/upgrade paths

#### Task 1.3: Frontend Security Scan (React)
**Time Estimate:** 45 minutes

**Actions:**
1. Navigate to frontend directory
2. Run dependency scan: `snyk test`
3. Generate dependency report: `snyk test --json > snyk-frontend-report.json`
4. Run code analysis: `snyk code test`
5. Generate code report: `snyk code test --json > snyk-code-report.json`
6. Monitor project: `snyk monitor`

**Deliverable:** `snyk-frontend-analysis.md`
- Dependency vulnerabilities summary
- Code vulnerabilities from `snyk code test`
- XSS vulnerabilities
- Hardcoded secrets
- React-specific issues (dangerouslySetInnerHTML, etc.)

#### Task 1.4: Remediation Plan
**Time Estimate:** 30 minutes

**Deliverable:** `snyk-remediation-plan.md`
- Critical issues (severity > 7.0) with immediate fix plan
- High priority issues (4.0-7.0) with remediation approach
- Medium/Low priority issues documented for future
- Dependency update strategy
- Testing plan after upgrades

#### Task 1.5: Implementation and Verification
**Time Estimate:** 1-2 hours

**Actions:**
1. Fix at least 3 critical/high severity vulnerabilities
2. Update vulnerable dependencies in `package.json` and `go.mod`
3. Test application still works after updates
4. Run Snyk again to verify fixes
5. Take screenshots of Snyk dashboard

**Deliverable:** `snyk-fixes-applied.md`
- Issues fixed list
- Changes made (code and dependency updates)
- Before/after Snyk scan results comparison
- Screenshots showing improvement

---

### **PHASE 3: SAST with SonarQube** (2-3 hours)

#### Task 2.1-2.2: Backend Analysis
**Time Estimate:** 1 hour

**Actions:**
1. Setup SonarQube Cloud via GitHub integration
2. Configure backend project in SonarQube
3. Run analysis on backend codebase
4. Review dashboard and export findings

**Deliverable:** `sonarqube-backend-analysis.md`
- Quality Gate status (Pass/Fail)
- Code metrics (LOC, duplications, complexity)
- Issues by category (Bugs, Vulnerabilities, Code Smells, Security Hotspots)
- Detailed vulnerability analysis with OWASP/CWE references
- Code quality ratings (Maintainability, Reliability, Security)
- Technical debt estimation
- Screenshots of dashboard, issues, security hotspots, coverage

#### Task 2.3: Frontend Analysis
**Time Estimate:** 1 hour

**Actions:**
1. Configure frontend project in SonarQube
2. Run analysis on React codebase
3. Review JavaScript/React specific issues

**Deliverable:** `sonarqube-frontend-analysis.md`
- Quality Gate status
- JavaScript/React specific issues (anti-patterns, JSX security)
- Security vulnerabilities (XSS, weak crypto, etc.)
- Code smells (duplicated code, complex functions)
- Best practices violations
- Screenshots of dashboard and findings

#### Task 2.4: Security Hotspot Review
**Time Estimate:** 45 minutes

**Deliverable:** `security-hotspots-review.md`
- Each security hotspot with location, OWASP category, impact
- Risk assessment (real vulnerability? exploit scenario? risk level)
- Remediation recommendations

#### Task 2.5: Improvements Implementation
**Time Estimate:** 1-2 hours

**Actions:**
1. Fix identified bugs and vulnerabilities
2. Reduce code complexity where possible
3. Remove code duplications
4. Run SonarQube again to verify improvements

**Deliverable:** `sonarqube-improvements.md`
- Before/after metrics comparison
- Issues resolved
- Quality rating improvements

---

### **PHASE 4: DAST with OWASP ZAP** (4-5 hours)

#### Task 3.1-3.2: Setup and Preparation
**Time Estimate:** 30 minutes

**Actions:**
1. Launch OWASP ZAP application
2. Start backend server: `cd golang-gin-realworld-example-app; go run hello.go`
3. Start frontend server: `cd react-redux-realworld-example-app; npm start`
4. Navigate to http://localhost:4100
5. Register test user: `security-test@example.com` / `SecurePass123!`
6. Create sample articles and comments for testing
7. Document test credentials

#### Task 3.3: Passive Scan
**Time Estimate:** 30 minutes

**Actions:**
1. Open ZAP, select "Automated Scan"
2. Target URL: `http://localhost:4100`
3. Use traditional spider
4. Enable passive scan
5. Wait for scan to complete
6. Export report: `zap-passive-report.html`

**Deliverable:** `zap-passive-scan-analysis.md`
- Alerts summary (total, High/Medium/Low/Info breakdown)
- High priority findings with URLs, descriptions, CWE/OWASP refs
- Common issues (missing security headers, cookie issues, info disclosure, CORS)
- Screenshots and HTML report

#### Task 3.4: Active Scan (Authenticated)
**Time Estimate:** 2-3 hours (scan takes 30+ minutes)

**Actions:**
1. **Configure ZAP Context:**
   - Name: "Conduit Authenticated"
   - Include URLs: `http://localhost:4100.*` and `http://localhost:8080/api.*`

2. **Setup Authentication:**
   - Method: JSON-based
   - Login URL: `http://localhost:8080/api/users/login`
   - Request body:
     ```json
     {
       "user": {
         "email": "security-test@example.com",
         "password": "SecurePass123!"
       }
     }
     ```
   - Extract token: `user.token`
   - Header: `Authorization: Token {token}`

3. **Configure User and Session:**
   - Add user with test credentials
   - Enable for context
   - Set session management to HTTP Authentication Header

4. **Run Active Scan:**
   - Spider with authenticated user
   - Run active scan on spidered URLs
   - Use "OWASP Top 10" scan policy
   - Medium intensity (to save time)

5. **Export Reports:**
   - HTML: `zap-active-report.html`
   - XML: `zap-active-report.xml`
   - JSON: `zap-active-report.json`

**Deliverable:** `zap-active-scan-analysis.md`
- Vulnerability summary with OWASP Top 10 mapping
- Critical/High severity vulnerabilities (for each: name, risk, URLs, CWE, OWASP category, description, attack details, evidence, impact, remediation)
- Expected findings checklist (SQL Injection, XSS, Security Misconfiguration, etc.)
- API security issues
- Frontend security issues
- Export all reports

#### Task 3.5: API Security Testing
**Time Estimate:** 1 hour

**Actions:**
1. Test all API endpoints manually in ZAP
2. Test for authentication bypass (no token, invalid token)
3. Test for authorization flaws (access other users' resources)
4. Test input validation (SQL injection, XSS, XXE, command injection)
5. Test rate limiting (brute force attempts)
6. Check for information disclosure (verbose errors, stack traces)

**Deliverable:** `zap-api-security-analysis.md`
- API-specific findings with endpoint, method, vulnerability, POC request/response, risk assessment

#### Task 3.6: Fix Implementation
**Time Estimate:** 2-3 hours

**Actions:**
1. Prioritize critical/high vulnerabilities
2. Implement fixes in backend and frontend code
3. Test fixes locally
4. Document all changes made

**Deliverable:** `zap-fixes-applied.md`
- Vulnerability fixed
- Code changes made
- Files modified
- Testing performed
- Before/after evidence

#### Task 3.7: Security Headers Implementation
**Time Estimate:** 30 minutes

**Actions:**
1. Add security headers middleware to backend (Go):
   ```go
   router.Use(func(c *gin.Context) {
       c.Header("X-Frame-Options", "DENY")
       c.Header("X-Content-Type-Options", "nosniff")
       c.Header("X-XSS-Protection", "1; mode=block")
       c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
       c.Header("Content-Security-Policy", "default-src 'self'")
       c.Next()
   })
   ```
2. Test headers are present
3. Take screenshots in ZAP

**Deliverable:** `security-headers-analysis.md`
- Code implementing headers
- Explanation of each header
- Screenshots showing headers

#### Task 3.8: Final Verification Scan
**Time Estimate:** 1 hour

**Actions:**
1. Run full ZAP scan again (passive + active)
2. Compare before/after results
3. Verify critical/high issues resolved
4. Document remaining issues with justification

**Deliverable:** `final-security-assessment.md`
- Before/after vulnerability counts
- Risk score improvement
- Outstanding issues and mitigation plan
- Screenshots of final ZAP report
- Security posture assessment

---

### **PHASE 5: Documentation and Submission** (1-2 hours)

#### Create Summary Report

**Deliverable:** `ASSIGNMENT_2_REPORT.md`

**Structure:**
1. **Executive Summary**
   - Overview of testing performed
   - Key statistics (vulnerabilities found, fixed, remaining)
   - Overall security posture improvement

2. **SAST Findings Summary**
   - Snyk results overview (backend & frontend)
   - SonarQube results overview (backend & frontend)
   - Key vulnerabilities identified

3. **DAST Findings Summary**
   - ZAP passive scan highlights
   - ZAP active scan highlights
   - API security findings

4. **Remediation Actions Taken**
   - List of all fixes implemented
   - Dependencies updated
   - Code changes made

5. **Remaining Risks**
   - Unresolved vulnerabilities with justification
   - Recommended future actions
   - Risk mitigation strategies

6. **Lessons Learned**
   - Key takeaways from the assignment
   - Tools comparison
   - Best practices identified

#### Organize All Files

**Checklist:**
- [ ] All markdown analysis files created
- [ ] All JSON/HTML/XML reports exported
- [ ] Code changes committed
- [ ] Screenshots captured and referenced
- [ ] Summary report completed

---

## 📦 Final Submission Checklist

### SAST Reports (Snyk)
- [ ] `snyk-backend-analysis.md`
- [ ] `snyk-frontend-analysis.md`
- [ ] `snyk-remediation-plan.md`
- [ ] `snyk-fixes-applied.md`
- [ ] `snyk-backend-report.json`
- [ ] `snyk-frontend-report.json`
- [ ] `snyk-code-report.json`

### SAST Reports (SonarQube)
- [ ] `sonarqube-backend-analysis.md`
- [ ] `sonarqube-frontend-analysis.md`
- [ ] `security-hotspots-review.md`
- [ ] `sonarqube-improvements.md`
- [ ] Screenshots of all dashboards

### DAST Reports (OWASP ZAP)
- [ ] `zap-passive-scan-analysis.md`
- [ ] `zap-active-scan-analysis.md`
- [ ] `zap-api-security-analysis.md`
- [ ] `zap-fixes-applied.md`
- [ ] `security-headers-analysis.md`
- [ ] `final-security-assessment.md`
- [ ] `zap-passive-report.html`
- [ ] `zap-active-report.html`
- [ ] `zap-active-report.xml`
- [ ] `zap-active-report.json`

### Code Changes
- [ ] Modified backend files with security fixes
- [ ] Modified frontend files with security fixes
- [ ] Updated `package.json` with dependency updates
- [ ] Updated `go.mod` with dependency updates

### Summary
- [ ] `ASSIGNMENT_2_REPORT.md`

---

## ⚠️ Important Notes

### Time Management
- **Total Estimated Time:** 12-16 hours
- **Deadline:** November 30, 2025, 11:59 PM
- **Recommendation:** Start immediately, as ZAP active scans can take 30+ minutes

### Common Pitfalls to Avoid
1. ❌ Running scans without proper authentication
2. ❌ Ignoring false positives without investigation
3. ❌ Fixing symptoms instead of root causes
4. ❌ Not testing fixes
5. ❌ Upgrading dependencies without testing for breaking changes
6. ❌ Applying fixes that break application functionality

### Testing Strategy
- Test application after EVERY fix
- Keep a backup of working code
- Make incremental changes
- Document what breaks and why

### Key Resources
- [OWASP Top 10](https://owasp.org/www-project-top-ten/)
- [Snyk Documentation](https://docs.snyk.io/)
- [SonarQube Documentation](https://docs.sonarqube.org/)
- [OWASP ZAP Documentation](https://www.zaproxy.org/docs/)
- [CWE Database](https://cwe.mitre.org/)

---

## 🎯 Success Criteria

### Minimum Requirements (to pass)
- All scans completed (Snyk, SonarQube, ZAP)
- At least 3 critical/high vulnerabilities fixed
- All deliverable documents created
- Summary report submitted

### Excellence Criteria (for high grade)
- Comprehensive analysis of all findings
- Most critical/high vulnerabilities fixed
- Clear documentation with screenshots
- Measurable security improvements
- Professional reporting

---

## 📊 Grading Breakdown

| Component | Points | Focus Area |
|-----------|--------|------------|
| Snyk Backend Analysis | 8 | Thorough vulnerability documentation |
| Snyk Frontend Analysis | 8 | Code and dependency analysis |
| SonarQube Backend | 8 | Bugs, vulnerabilities, code smells |
| SonarQube Frontend | 8 | Quality and security issues |
| SonarQube Improvements | 10 | Measurable quality improvement |
| ZAP Passive Scan | 8 | Complete scan documentation |
| ZAP Active Scan | 15 | Authenticated scan, all vulnerabilities |
| ZAP API Testing | 10 | API-specific vulnerabilities |
| Security Fixes | 15 | Critical issues fixed and verified |
| Security Headers | 5 | All recommended headers |
| Documentation | 5 | Clear, professional reporting |
| **TOTAL** | **100** | |

---

## 🚀 Getting Started

**Next Steps:**
1. Review this plan thoroughly
2. Ensure you understand each phase
3. Set up your working environment
4. Begin with Phase 1 (Environment Setup)
5. Work through phases sequentially
6. Document as you go (don't leave it to the end!)

**Questions to Consider Before Starting:**
- Do I have all tools installed?
- Do both applications run successfully?
- Do I have accounts for Snyk and SonarQube Cloud?
- Have I allocated enough time for this assignment?
- Do I understand what each tool does?

---

Good luck! Remember: Security testing is about finding issues to make applications safer, not about achieving a perfect score. Document findings honestly and focus on learning.
