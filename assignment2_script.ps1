# Assignment 2: SAST & DAST Automation Script
# Created: November 29, 2025
# Purpose: Automate security testing workflows for Assignment 2

param(
    [Parameter(Mandatory=$false)]
    [ValidateSet('setup', 'snyk-backend', 'snyk-frontend', 'snyk-all', 'zap-passive', 'verify', 'all')]
    [string]$Task = 'all'
)

# Color output functions
function Write-Success { param($Message) Write-Host "✓ $Message" -ForegroundColor Green }
function Write-Info { param($Message) Write-Host "ℹ $Message" -ForegroundColor Cyan }
function Write-Warning { param($Message) Write-Host "⚠ $Message" -ForegroundColor Yellow }
function Write-Error-Custom { param($Message) Write-Host "✗ $Message" -ForegroundColor Red }
function Write-Section { param($Message) Write-Host "`n========== $Message ==========" -ForegroundColor Magenta }

# Configuration
$ROOT_DIR = Get-Location
$BACKEND_DIR = Join-Path $ROOT_DIR "golang-gin-realworld-example-app"
$FRONTEND_DIR = Join-Path $ROOT_DIR "react-redux-realworld-example-app"
$REPORTS_DIR = Join-Path $ROOT_DIR "security-reports"

# Create reports directory
function Initialize-ReportsDirectory {
    Write-Section "Initializing Reports Directory"
    
    if (-not (Test-Path $REPORTS_DIR)) {
        New-Item -ItemType Directory -Path $REPORTS_DIR -Force | Out-Null
        Write-Success "Created reports directory: $REPORTS_DIR"
    } else {
        Write-Info "Reports directory already exists"
    }
    
    # Create subdirectories
    $subdirs = @('snyk', 'sonarqube', 'zap', 'screenshots')
    foreach ($dir in $subdirs) {
        $path = Join-Path $REPORTS_DIR $dir
        if (-not (Test-Path $path)) {
            New-Item -ItemType Directory -Path $path -Force | Out-Null
            Write-Success "Created subdirectory: $dir"
        }
    }
}

# Check prerequisites
function Test-Prerequisites {
    Write-Section "Checking Prerequisites"
    
    $prerequisites = @{
        'node' = 'Node.js'
        'npm' = 'npm'
        'go' = 'Go'
        'git' = 'Git'
    }
    
    $allPresent = $true
    
    foreach ($cmd in $prerequisites.Keys) {
        try {
            $null = Get-Command $cmd -ErrorAction Stop
            Write-Success "$($prerequisites[$cmd]) is installed"
        } catch {
            Write-Error-Custom "$($prerequisites[$cmd]) is NOT installed"
            $allPresent = $false
        }
    }
    
    # Check Snyk
    try {
        $null = Get-Command snyk -ErrorAction Stop
        Write-Success "Snyk CLI is installed"
    } catch {
        Write-Warning "Snyk CLI is NOT installed. Run: npm install -g snyk"
        $allPresent = $false
    }
    
    # Check if directories exist
    if (Test-Path $BACKEND_DIR) {
        Write-Success "Backend directory found"
    } else {
        Write-Error-Custom "Backend directory not found: $BACKEND_DIR"
        $allPresent = $false
    }
    
    if (Test-Path $FRONTEND_DIR) {
        Write-Success "Frontend directory found"
    } else {
        Write-Error-Custom "Frontend directory not found: $FRONTEND_DIR"
        $allPresent = $false
    }
    
    return $allPresent
}

# Setup environment
function Invoke-Setup {
    Write-Section "Setting Up Environment"
    
    # Initialize reports directory
    Initialize-ReportsDirectory
    
    # Check prerequisites
    if (-not (Test-Prerequisites)) {
        Write-Error-Custom "Prerequisites check failed. Please install missing tools."
        return $false
    }
    
    # Install Snyk if not present
    try {
        $null = Get-Command snyk -ErrorAction Stop
        Write-Info "Snyk already installed"
    } catch {
        Write-Info "Installing Snyk CLI..."
        npm install -g snyk
        if ($LASTEXITCODE -eq 0) {
            Write-Success "Snyk CLI installed successfully"
        } else {
            Write-Error-Custom "Failed to install Snyk CLI"
            return $false
        }
    }
    
    # Check Snyk authentication
    Write-Info "Checking Snyk authentication..."
    $snykAuth = snyk auth --status 2>&1
    if ($LASTEXITCODE -ne 0) {
        Write-Warning "Snyk is not authenticated"
        Write-Info "Please run: snyk auth"
        Write-Info "This will open a browser for authentication"
        return $false
    } else {
        Write-Success "Snyk is authenticated"
    }
    
    # Install frontend dependencies if needed
    Write-Info "Checking frontend dependencies..."
    Push-Location $FRONTEND_DIR
    if (-not (Test-Path "node_modules")) {
        Write-Info "Installing frontend dependencies..."
        npm install
        if ($LASTEXITCODE -eq 0) {
            Write-Success "Frontend dependencies installed"
        } else {
            Write-Error-Custom "Failed to install frontend dependencies"
            Pop-Location
            return $false
        }
    } else {
        Write-Success "Frontend dependencies already installed"
    }
    Pop-Location
    
    Write-Success "Setup completed successfully!"
    return $true
}

# Run Snyk on Backend
function Invoke-SnykBackend {
    Write-Section "Running Snyk Security Scan on Backend (Go)"
    
    Push-Location $BACKEND_DIR
    
    try {
        # Run basic test
        Write-Info "Running Snyk test on backend..."
        snyk test --severity-threshold=low
        
        # Generate JSON report
        Write-Info "Generating JSON report..."
        $reportPath = Join-Path $REPORTS_DIR "snyk\snyk-backend-report.json"
        snyk test --json > $reportPath
        
        if (Test-Path $reportPath) {
            Write-Success "Backend JSON report saved: $reportPath"
        }
        
        # Test all projects
        Write-Info "Testing all projects..."
        snyk test --all-projects --severity-threshold=low
        
        # Monitor project (uploads to Snyk dashboard)
        Write-Info "Monitoring project (uploading to Snyk dashboard)..."
        snyk monitor
        
        Write-Success "Backend Snyk scan completed!"
        Write-Info "Next step: Analyze results and create snyk-backend-analysis.md"
        
    } catch {
        Write-Error-Custom "Error running Snyk on backend: $_"
    } finally {
        Pop-Location
    }
}

# Run Snyk on Frontend
function Invoke-SnykFrontend {
    Write-Section "Running Snyk Security Scan on Frontend (React)"
    
    Push-Location $FRONTEND_DIR
    
    try {
        # Run dependency test
        Write-Info "Running Snyk dependency test on frontend..."
        snyk test --severity-threshold=low
        
        # Generate dependency report
        Write-Info "Generating dependency JSON report..."
        $depReportPath = Join-Path $REPORTS_DIR "snyk\snyk-frontend-report.json"
        snyk test --json > $depReportPath
        
        if (Test-Path $depReportPath) {
            Write-Success "Frontend dependency report saved: $depReportPath"
        }
        
        # Run code analysis
        Write-Info "Running Snyk code analysis..."
        snyk code test
        
        # Generate code analysis report
        Write-Info "Generating code analysis JSON report..."
        $codeReportPath = Join-Path $REPORTS_DIR "snyk\snyk-code-report.json"
        snyk code test --json > $codeReportPath
        
        if (Test-Path $codeReportPath) {
            Write-Success "Frontend code analysis report saved: $codeReportPath"
        }
        
        # Monitor project
        Write-Info "Monitoring project (uploading to Snyk dashboard)..."
        snyk monitor
        
        Write-Success "Frontend Snyk scan completed!"
        Write-Info "Next step: Analyze results and create snyk-frontend-analysis.md"
        
    } catch {
        Write-Error-Custom "Error running Snyk on frontend: $_"
    } finally {
        Pop-Location
    }
}

# Run all Snyk scans
function Invoke-SnykAll {
    Invoke-SnykBackend
    Start-Sleep -Seconds 2
    Invoke-SnykFrontend
    
    Write-Section "Snyk Scan Summary"
    Write-Info "All Snyk scans completed!"
    Write-Info "Reports saved in: $REPORTS_DIR\snyk\"
    Write-Info ""
    Write-Info "Next steps:"
    Write-Info "1. Review JSON reports in security-reports/snyk/"
    Write-Info "2. Create snyk-backend-analysis.md"
    Write-Info "3. Create snyk-frontend-analysis.md"
    Write-Info "4. Create snyk-remediation-plan.md"
    Write-Info "5. Implement fixes and create snyk-fixes-applied.md"
}

# Verify applications are running
function Test-Applications {
    Write-Section "Verifying Application Health"
    
    # Check backend
    Write-Info "Checking backend (http://localhost:8080)..."
    try {
        $response = Invoke-WebRequest -Uri "http://localhost:8080/api/tags" -Method GET -UseBasicParsing -TimeoutSec 5
        if ($response.StatusCode -eq 200) {
            Write-Success "Backend is running and responding"
        }
    } catch {
        Write-Warning "Backend is NOT responding. Make sure it's running with: cd golang-gin-realworld-example-app; go run hello.go"
    }
    
    # Check frontend
    Write-Info "Checking frontend (http://localhost:4100)..."
    try {
        $response = Invoke-WebRequest -Uri "http://localhost:4100" -Method GET -UseBasicParsing -TimeoutSec 5
        if ($response.StatusCode -eq 200) {
            Write-Success "Frontend is running and responding"
        }
    } catch {
        Write-Warning "Frontend is NOT responding. Make sure it's running with: cd react-redux-realworld-example-app; npm start"
    }
}

# Create template analysis files
function New-AnalysisTemplates {
    Write-Section "Creating Analysis Document Templates"
    
    # Snyk Backend Analysis Template
    $snykBackendTemplate = @"
# Snyk Backend Security Analysis
**Date:** $(Get-Date -Format "yyyy-MM-dd")
**Project:** RealWorld Conduit - Backend (Go)

## 1. Vulnerability Summary

### Overall Statistics
- **Total Vulnerabilities:** [NUMBER]
- **Critical:** [NUMBER]
- **High:** [NUMBER]
- **Medium:** [NUMBER]
- **Low:** [NUMBER]

### Affected Dependencies
| Package | Version | Vulnerabilities |
|---------|---------|-----------------|
| [name]  | [ver]   | [count]         |

---

## 2. Critical/High Severity Issues

### Vulnerability 1: [NAME]
- **Severity:** [Critical/High]
- **Package:** [package-name]
- **Current Version:** [version]
- **CVE:** [CVE-YYYY-XXXXX]
- **CWE:** [CWE-XXX]
- **Description:** [Description of vulnerability]
- **Exploit Scenario:** [How this could be exploited]
- **Recommended Fix:** [Upgrade to version X.Y.Z / Apply patch / Workaround]
- **CVSS Score:** [X.X]

---

## 3. Dependency Analysis

### Direct Dependencies
- [List direct dependencies with versions]

### Transitive Dependencies
- [List vulnerable transitive dependencies]

### Outdated Dependencies
- [List packages that need updating]

### License Issues
- [Any license-related concerns]

---

## 4. Recommendations

### Immediate Actions
1. [Action item 1]
2. [Action item 2]

### Short-term Actions
1. [Action item 1]
2. [Action item 2]

### Long-term Actions
1. [Action item 1]
2. [Action item 2]

---

## 5. References
- Snyk Dashboard: [URL]
- Report File: snyk-backend-report.json
"@

    # Snyk Frontend Analysis Template
    $snykFrontendTemplate = @"
# Snyk Frontend Security Analysis
**Date:** $(Get-Date -Format "yyyy-MM-dd")
**Project:** RealWorld Conduit - Frontend (React)

## 1. Dependency Vulnerabilities

### Overall Statistics
- **Total Vulnerabilities:** [NUMBER]
- **Critical:** [NUMBER]
- **High:** [NUMBER]
- **Medium:** [NUMBER]
- **Low:** [NUMBER]

### Vulnerable npm Packages
| Package | Version | Severity | Fix Available |
|---------|---------|----------|---------------|
| [name]  | [ver]   | [sev]    | [yes/no]      |

---

## 2. Code Vulnerabilities (Snyk Code)

### Code Security Issues
| Issue | Location | Severity | Description |
|-------|----------|----------|-------------|
| [type]| [file:line] | [sev] | [desc]   |

### XSS Vulnerabilities
- [List any XSS issues found]

### Hardcoded Secrets
- [List any hardcoded secrets/credentials]

### Insecure Cryptography
- [List crypto-related issues]

---

## 3. React-Specific Issues

### Dangerous Props
- [dangerouslySetInnerHTML usage]

### Client-side Security
- [localStorage security issues]
- [Session management issues]

### Component Security Concerns
- [Vulnerable component patterns]

---

## 4. Upgrade Recommendations

### Critical Upgrades
1. [Package]: [current] → [recommended]

### Important Upgrades
1. [Package]: [current] → [recommended]

---

## 5. References
- Snyk Dashboard: [URL]
- Dependency Report: snyk-frontend-report.json
- Code Analysis Report: snyk-code-report.json
"@

    # Save templates
    $templates = @{
        'snyk-backend-analysis.md' = $snykBackendTemplate
        'snyk-frontend-analysis.md' = $snykFrontendTemplate
    }
    
    foreach ($file in $templates.Keys) {
        $path = Join-Path $ROOT_DIR $file
        if (-not (Test-Path $path)) {
            $templates[$file] | Out-File -FilePath $path -Encoding UTF8
            Write-Success "Created template: $file"
        } else {
            Write-Info "Template already exists: $file"
        }
    }
}

# Main execution
function Start-Assignment2Script {
    Write-Host @"
╔════════════════════════════════════════════════════════════════╗
║                                                                ║
║        Assignment 2: SAST & DAST Automation Script            ║
║        Static & Dynamic Application Security Testing          ║
║                                                                ║
╚════════════════════════════════════════════════════════════════╝
"@ -ForegroundColor Cyan

    switch ($Task) {
        'setup' {
            if (Invoke-Setup) {
                New-AnalysisTemplates
            }
        }
        'snyk-backend' {
            Invoke-SnykBackend
        }
        'snyk-frontend' {
            Invoke-SnykFrontend
        }
        'snyk-all' {
            Invoke-SnykAll
        }
        'verify' {
            Test-Applications
        }
        'all' {
            Write-Info "Running complete workflow..."
            if (Invoke-Setup) {
                New-AnalysisTemplates
                Start-Sleep -Seconds 2
                Test-Applications
                Start-Sleep -Seconds 2
                Invoke-SnykAll
            }
        }
        default {
            Write-Error-Custom "Unknown task: $Task"
        }
    }
    
    Write-Host "`n"
}

# Run the script
Start-Assignment2Script

# Usage Instructions
Write-Host @"

╔════════════════════════════════════════════════════════════════╗
║                     USAGE INSTRUCTIONS                         ║
╚════════════════════════════════════════════════════════════════╝

To run specific tasks, use:

    .\assignment2_script.ps1 -Task setup          # Setup environment
    .\assignment2_script.ps1 -Task snyk-backend   # Scan backend only
    .\assignment2_script.ps1 -Task snyk-frontend  # Scan frontend only
    .\assignment2_script.ps1 -Task snyk-all       # Scan both
    .\assignment2_script.ps1 -Task verify         # Check apps running
    .\assignment2_script.ps1 -Task all            # Complete workflow

Before running scans:
1. Authenticate Snyk: snyk auth
2. Ensure both applications can run successfully

For OWASP ZAP testing:
1. Download ZAP from: https://www.zaproxy.org/download/
2. Start backend: cd golang-gin-realworld-example-app; go run hello.go
3. Start frontend: cd react-redux-realworld-example-app; npm start
4. Use ZAP GUI for manual testing

For SonarQube:
1. Setup at: https://docs.sonarsource.com/sonarqube-cloud/getting-started/github
2. Configure both projects in SonarQube Cloud
3. Follow integration instructions

Reports are saved in: $REPORTS_DIR

"@ -ForegroundColor Yellow
