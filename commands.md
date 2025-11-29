# Screenshot Commands Guide

This document contains all the commands you need to run to capture screenshots for Assignment 1 evidence.

---

## 📋 Prerequisites

### Backend Setup

**IMPORTANT: Install GCC for Windows (required for SQLite/CGO)**

The backend uses SQLite which requires CGO and a C compiler. You have two options:

**Option 1: Install TDM-GCC (Recommended - Easiest)**
1. Download TDM-GCC from: wnload/https://jmeubank.github.io/tdm-gcc/do
2. Install the 64-bit version (tdm64-gcc-10.3.0-2.exe or later)
3. During installation, make sure "Add to PATH" is checked
4. Restart PowerShell after installation
5. Verify: `gcc --version`

**Option 2: Install MinGW-w64**
1. Download from: https://www.mingw-w64.org/downloads/
2. Add MinGW bin folder to PATH
3. Restart PowerShell
4. Verify: `gcc --version`

**After GCC Installation:**
```powershell
cd c:\Users\HP\OneDrive\Desktop\swe302_assignments-master\golang-gin-realworld-example-app
$env:CGO_ENABLED=1
```

### Frontend Setup
```powershell
cd c:\Users\HP\OneDrive\Desktop\swe302_assignments-master\react-redux-realworld-example-app

# Install testing dependencies (first time only)
npm install --save-dev @testing-library/react @testing-library/jest-dom @testing-library/user-event
```

---

## 🔴 Backend Screenshots

### Screenshot 1: Run All Backend Tests (Verbose Output)
**Command:**
```powershell
cd c:\Users\HP\OneDrive\Desktop\swe302_assignments-master\golang-gin-realworld-example-app
$env:CGO_ENABLED=1; go test ./... -v
```
**What to capture:** Terminal output showing all test names and PASS/FAIL status

---

### Screenshot 2: Backend Test Coverage Summary
**Command:**
```powershell
cd c:\Users\HP\OneDrive\Desktop\swe302_assignments-master\golang-gin-realworld-example-app
$env:CGO_ENABLED=1; go test ./... -cover
```
**What to capture:** Terminal output showing coverage percentage for each package (common, users, articles)

---

### Screenshot 3: Generate Coverage Profile
**Command:**
```powershell
cd c:\Users\HP\OneDrive\Desktop\swe302_assignments-master\golang-gin-realworld-example-app
$env:CGO_ENABLED=1; go test ./... -coverprofile=coverage.out
```
**What to capture:** Terminal output showing coverage.out was created

---

### Screenshot 4: Open Coverage HTML Report
**Command:**
```powershell
cd c:\Users\HP\OneDrive\Desktop\swe302_assignments-master\golang-gin-realworld-example-app
go tool cover -html=coverage.out -o coverage.html
Start-Process coverage.html
```
**What to capture:** Browser window showing the HTML coverage report with green/red/gray code highlighting

---

### Screenshot 5: Articles Package Tests Only
**Command:**
```powershell
cd c:\Users\HP\OneDrive\Desktop\swe302_assignments-master\golang-gin-realworld-example-app
$env:CGO_ENABLED=1; go test ./articles -v
```
**What to capture:** Terminal output showing all 20 articles package tests passing

---

### Screenshot 6: Articles Package Coverage
**Command:**
```powershell
cd c:\Users\HP\OneDrive\Desktop\swe302_assignments-master\golang-gin-realworld-example-app
$env:CGO_ENABLED=1; go test ./articles -cover
```
**What to capture:** Terminal output showing articles package coverage (~70-75%)

---

### Screenshot 7: Common Package Tests
**Command:**
```powershell
cd c:\Users\HP\OneDrive\Desktop\swe302_assignments-master\golang-gin-realworld-example-app
$env:CGO_ENABLED=1; go test ./common -v
```
**What to capture:** Terminal output showing all common package tests passing

---

### Screenshot 8: Users Package Tests
**Command:**
```powershell
cd c:\Users\HP\OneDrive\Desktop\swe302_assignments-master\golang-gin-realworld-example-app
$env:CGO_ENABLED=1; go test ./users -v
```
**What to capture:** Terminal output showing all users package tests passing

---

### Screenshot 9: Integration Tests
**Command:**
```powershell
cd c:\Users\HP\OneDrive\Desktop\swe302_assignments-master\golang-gin-realworld-example-app
$env:CGO_ENABLED=1; go test -v integration_test.go
```
**What to capture:** Terminal output showing all 17 integration tests passing

---

### Screenshot 10: Backend Test Summary with Count
**Command:**
```powershell
cd c:\Users\HP\OneDrive\Desktop\swe302_assignments-master\golang-gin-realworld-example-app
$env:CGO_ENABLED=1; go test ./... -v | Select-String "PASS|FAIL" | Measure-Object -Line
```
**What to capture:** Terminal showing total number of tests executed

---

## 🔵 Frontend Screenshots

### Screenshot 11: Install Frontend Testing Dependencies
**Command:**
```powershell
cd c:\Users\HP\OneDrive\Desktop\swe302_assignments-master\react-redux-realworld-example-app
npm install --save-dev @testing-library/react @testing-library/jest-dom @testing-library/user-event
```
**What to capture:** Terminal output showing packages being installed (only needed first time)

---

### Screenshot 12: Run All Frontend Tests
**Command:**
```powershell
cd c:\Users\HP\OneDrive\Desktop\swe302_assignments-master\react-redux-realworld-example-app
npm test -- --watchAll=false
```
**What to capture:** Terminal output showing all test suites passing with summary (e.g., "Test Suites: 10 passed, Tests: 65 passed")

---

### Screenshot 13: Frontend Tests with Coverage
**Command:**
```powershell
cd c:\Users\HP\OneDrive\Desktop\swe302_assignments-master\react-redux-realworld-example-app
npm test -- --watchAll=false --coverage
```
**What to capture:** Terminal output showing test results AND coverage table (statements, branches, functions, lines percentages)

---

### Screenshot 14: Frontend Coverage HTML Report
**Command:**
```powershell
cd c:\Users\HP\OneDrive\Desktop\swe302_assignments-master\react-redux-realworld-example-app
npm test -- --watchAll=false --coverage --coverageReporters=html
Start-Process coverage\lcov-report\index.html
```
**What to capture:** Browser window showing the HTML coverage report with file list and percentages

---

### Screenshot 15: Component Tests Only
**Command:**
```powershell
cd c:\Users\HP\OneDrive\Desktop\swe302_assignments-master\react-redux-realworld-example-app
npm test -- --testPathPattern="components" --watchAll=false
```
**What to capture:** Terminal output showing only component tests (ArticleList, ArticlePreview, Login, Header, Editor)

---

### Screenshot 16: Reducer Tests Only
**Command:**
```powershell
cd c:\Users\HP\OneDrive\Desktop\swe302_assignments-master\react-redux-realworld-example-app
npm test -- --testPathPattern="reducers" --watchAll=false
```
**What to capture:** Terminal output showing only reducer tests (auth, articleList, editor)

---

### Screenshot 17: Middleware Tests Only
**Command:**
```powershell
cd c:\Users\HP\OneDrive\Desktop\swe302_assignments-master\react-redux-realworld-example-app
npm test -- --testPathPattern="middleware" --watchAll=false
```
**What to capture:** Terminal output showing middleware tests passing

---

### Screenshot 18: Integration Tests Only
**Command:**
```powershell
cd c:\Users\HP\OneDrive\Desktop\swe302_assignments-master\react-redux-realworld-example-app
npm test -- --testPathPattern="integration" --watchAll=false
```
**What to capture:** Terminal output showing all 9 integration tests passing

---

### Screenshot 19: Verbose Frontend Test Output
**Command:**
```powershell
cd c:\Users\HP\OneDrive\Desktop\swe302_assignments-master\react-redux-realworld-example-app
npm test -- --watchAll=false --verbose
```
**What to capture:** Terminal output showing detailed test names and descriptions

---

### Screenshot 20: Frontend Test Files List
**Command:**
```powershell
cd c:\Users\HP\OneDrive\Desktop\swe302_assignments-master\react-redux-realworld-example-app
Get-ChildItem -Path src -Filter "*.test.js" -Recurse | Select-Object FullName
```
**What to capture:** Terminal output listing all test files created

---

## 📊 Additional Evidence Screenshots

### Screenshot 21: Backend Test Files
**Command:**
```powershell
cd c:\Users\HP\OneDrive\Desktop\swe302_assignments-master\golang-gin-realworld-example-app
Get-ChildItem -Path . -Filter "*test.go" -Recurse | Select-Object FullName
```
**What to capture:** List of all backend test files

---

### Screenshot 22: Project Structure
**Command:**
```powershell
cd c:\Users\HP\OneDrive\Desktop\swe302_assignments-master
tree /F /A
```
**What to capture:** Complete project directory tree showing all files

---

### Screenshot 23: Articles Test File Content (First 50 lines)
**Command:**
```powershell
cd c:\Users\HP\OneDrive\Desktop\swe302_assignments-master\golang-gin-realworld-example-app
Get-Content articles\unit_test.go -Head 50
```
**What to capture:** First 50 lines of articles/unit_test.go showing test structure

---

### Screenshot 24: Integration Test File Content (First 50 lines)
**Command:**
```powershell
cd c:\Users\HP\OneDrive\Desktop\swe302_assignments-master\golang-gin-realworld-example-app
Get-Content integration_test.go -Head 50
```
**What to capture:** First 50 lines of integration_test.go

---

### Screenshot 25: Frontend Test Utils File
**Command:**
```powershell
cd c:\Users\HP\OneDrive\Desktop\swe302_assignments-master\react-redux-realworld-example-app
Get-Content src\test-utils.js
```
**What to capture:** Complete test-utils.js file content

---

## 🎯 Quick All-in-One Commands

### Backend - All Tests in One Go
```powershell
cd c:\Users\HP\OneDrive\Desktop\swe302_assignments-master\golang-gin-realworld-example-app
$env:CGO_ENABLED=1
Write-Host "=== Running All Tests ===" -ForegroundColor Green
go test ./... -v
Write-Host "`n=== Coverage Summary ===" -ForegroundColor Green
go test ./... -cover
Write-Host "`n=== Generating Coverage Report ===" -ForegroundColor Green
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
Write-Host "`nOpening coverage report in browser..." -ForegroundColor Green
Start-Process coverage.html
```

### Frontend - All Tests in One Go
```powershell
cd c:\Users\HP\OneDrive\Desktop\swe302_assignments-master\react-redux-realworld-example-app
Write-Host "=== Running All Frontend Tests ===" -ForegroundColor Green
npm test -- --watchAll=false --coverage
Write-Host "`n=== Generating HTML Coverage Report ===" -ForegroundColor Green
npm test -- --watchAll=false --coverage --coverageReporters=html
Write-Host "`nOpening coverage report in browser..." -ForegroundColor Green
Start-Process coverage\lcov-report\index.html
```

---

## 📝 Screenshot Checklist

Use this checklist to track your progress:

### Backend Screenshots (10)
- [ ] Screenshot 1: All backend tests verbose
- [ ] Screenshot 2: Coverage summary
- [ ] Screenshot 3: Coverage profile generation
- [ ] Screenshot 4: HTML coverage report
- [ ] Screenshot 5: Articles tests verbose
- [ ] Screenshot 6: Articles coverage
- [ ] Screenshot 7: Common tests
- [ ] Screenshot 8: Users tests
- [ ] Screenshot 9: Integration tests
- [ ] Screenshot 10: Test count summary

### Frontend Screenshots (10)
- [ ] Screenshot 11: Install dependencies
- [ ] Screenshot 12: All frontend tests
- [ ] Screenshot 13: Tests with coverage
- [ ] Screenshot 14: HTML coverage report
- [ ] Screenshot 15: Component tests only
- [ ] Screenshot 16: Reducer tests only
- [ ] Screenshot 17: Middleware tests only
- [ ] Screenshot 18: Integration tests only
- [ ] Screenshot 19: Verbose test output
- [ ] Screenshot 20: Test files list

### Additional Evidence (5)
- [ ] Screenshot 21: Backend test files list
- [ ] Screenshot 22: Project structure tree
- [ ] Screenshot 23: Articles test file content
- [ ] Screenshot 24: Integration test file content
- [ ] Screenshot 25: Test utils file content

---

## 💡 Tips for Taking Screenshots

1. **Terminal Screenshots**: Make sure the terminal window is wide enough to show full output without wrapping
2. **Browser Screenshots**: Capture the entire browser window including the URL bar
3. **Code Screenshots**: Use a readable font size (12-14pt recommended)
4. **Highlight Important Parts**: Use a red box or arrow to highlight key metrics (coverage %, test counts)
5. **Label Each Screenshot**: Save with descriptive names like "backend_coverage_summary.png"

---

## ⚠️ Troubleshooting

### If Backend Tests Fail:
```powershell
# Make sure you're in the right directory
cd c:\Users\HP\OneDrive\Desktop\swe302_assignments-master\golang-gin-realworld-example-app

# Check Go version (should be 1.16+)
go version

# Enable CGO (required for SQLite)
$env:CGO_ENABLED=1

# Clean and rebuild
go clean -testcache
go test ./... -v
```

### If Frontend Tests Fail:
```powershell
# Make sure dependencies are installed
cd c:\Users\HP\OneDrive\Desktop\swe302_assignments-master\react-redux-realworld-example-app
npm install

# Clear cache and run tests
npm test -- --clearCache
npm test -- --watchAll=false
```

### If Coverage Report Doesn't Open:
```powershell
# Manually open the file
# Backend:
Start-Process "c:\Users\HP\OneDrive\Desktop\swe302_assignments-master\golang-gin-realworld-example-app\coverage.html"

# Frontend:
Start-Process "c:\Users\HP\OneDrive\Desktop\swe302_assignments-master\react-redux-realworld-example-app\coverage\lcov-report\index.html"
```

---

## ✅ Summary

**Total Screenshots Needed:** 25
- Backend: 10 screenshots
- Frontend: 10 screenshots  
- Additional Evidence: 5 screenshots

**Estimated Time:** 15-20 minutes to run all commands and capture screenshots

**Note:** You can use the "Quick All-in-One Commands" section to run multiple commands sequentially, then just capture the terminal output at the end of each section.

Good luck with your screenshots! 📸
