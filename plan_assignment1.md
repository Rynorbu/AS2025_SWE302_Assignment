# Assignment 1 Implementation Plan
## Unit Testing, Integration Testing & Test Coverage

**Created:** November 29, 2025  
**Deadline:** November 30, 2025, 11:59 PM

---

## 📋 Overview

This plan outlines the implementation strategy for Assignment 1, covering both **Backend (Go/Gin)** and **Frontend (React/Redux)** testing requirements.

### Assignment Breakdown
- **Backend Testing:** 70 points (Unit: 40, Integration: 30)
- **Frontend Testing:** 100 points (Unit: 40, Redux: 30, Integration: 30)
- **Documentation & Quality:** 10 points

---

## 🎯 Part A: Backend Testing (Go/Gin)

### Current State Analysis
✅ **Existing Tests:**
- `users/unit_test.go` - Comprehensive user tests (login, registration, following, profile)
- `common/unit_test.go` - Database, JWT, validation tests

❌ **Missing Tests:**
- `articles/unit_test.go` - **NO TESTS** (0% coverage)
- `integration_test.go` - No integration tests in root

### Task 1: Unit Testing (40 points)

#### 1.1 Analyze Existing Tests
**Command to run:**
```powershell
cd golang-gin-realworld-example-app
go test ./... -v
```

**Deliverable:** `testing-analysis.md`
- Document which packages have tests
- Identify failing tests
- Explain reasons for failures

#### 1.2 Create Articles Unit Tests
**File to create:** `golang-gin-realworld-example-app/articles/unit_test.go`

**Test Categories (Minimum 15 tests):**

1. **Model Tests (5 tests)**
   - `TestArticleCreation` - Create article with valid data
   - `TestArticleValidation` - Empty title/body validation
   - `TestArticleFavorite` - Favorite functionality
   - `TestArticleUnfavorite` - Unfavorite functionality
   - `TestArticleTagAssociation` - Tag relationships

2. **Serializer Tests (5 tests)**
   - `TestArticleSerializer` - Single article JSON output
   - `TestArticleSerializerWithAuthor` - Author information included
   - `TestArticleListSerializer` - Multiple articles serialization
   - `TestCommentSerializer` - Comment JSON structure
   - `TestCommentSerializerWithAuthor` - Comment with author info

3. **Validator Tests (5 tests)**
   - `TestArticleModelValidatorValid` - Valid article data
   - `TestArticleModelValidatorMissingTitle` - Missing title error
   - `TestArticleModelValidatorMissingBody` - Missing body error
   - `TestCommentModelValidatorValid` - Valid comment data
   - `TestCommentModelValidatorMissingBody` - Missing comment body

**Command to test:**
```powershell
cd golang-gin-realworld-example-app
go test ./articles -v
```

#### 1.3 Enhance Common Package Tests
**File to modify:** `golang-gin-realworld-example-app/common/unit_test.go`

**Additional Tests (5 tests):**
- `TestGenTokenWithDifferentUserIDs` - Token generation for multiple users
- `TestGenTokenUniqueness` - Tokens are unique per user
- `TestJWTTokenExpiration` - Token expiration handling
- `TestDatabaseConnectionError` - Connection failure handling
- `TestRandStringUniqueness` - Random string uniqueness

**Command to test:**
```powershell
go test ./common -v
```

### Task 2: Integration Testing (30 points)

**File to create:** `golang-gin-realworld-example-app/integration_test.go`

#### 2.1 Authentication Integration Tests (5 tests)
- `TestUserRegistrationFlow` - POST /api/users
- `TestUserLoginFlow` - POST /api/users/login
- `TestGetCurrentUserAuthenticated` - GET /api/user (with token)
- `TestGetCurrentUserUnauthenticated` - GET /api/user (no token → 401)
- `TestLoginInvalidCredentials` - Login with wrong password

#### 2.2 Article CRUD Integration Tests (5 tests)
- `TestCreateArticleAuthenticated` - POST /api/articles (with auth)
- `TestCreateArticleUnauthenticated` - POST /api/articles (no auth → 401)
- `TestListArticles` - GET /api/articles
- `TestGetSingleArticle` - GET /api/articles/:slug
- `TestUpdateArticleByAuthor` - PUT /api/articles/:slug (by author)
- `TestUpdateArticleByNonAuthor` - PUT /api/articles/:slug (unauthorized)
- `TestDeleteArticleByAuthor` - DELETE /api/articles/:slug

#### 2.3 Article Interaction Tests (5 tests)
- `TestFavoriteArticle` - POST /api/articles/:slug/favorite
- `TestUnfavoriteArticle` - DELETE /api/articles/:slug/favorite
- `TestCreateComment` - POST /api/articles/:slug/comments
- `TestListComments` - GET /api/articles/:slug/comments
- `TestDeleteComment` - DELETE /api/articles/:slug/comments/:id

**Command to test:**
```powershell
cd golang-gin-realworld-example-app
go test -v integration_test.go
```

### Task 3: Test Coverage Analysis (30 points)

#### 3.1 Generate Coverage Reports

**Commands:**
```powershell
cd golang-gin-realworld-example-app

# Run all tests with coverage
go test ./... -cover

# Generate detailed coverage profile
go test ./... -coverprofile=coverage.out

# Generate HTML coverage report
go tool cover -html=coverage.out -o coverage.html

# View coverage per package
go test ./common -coverprofile=common_coverage.out
go test ./users -coverprofile=users_coverage.out
go test ./articles -coverprofile=articles_coverage.out

# Generate individual HTML reports
go tool cover -html=common_coverage.out -o common_coverage.html
go tool cover -html=users_coverage.out -o users_coverage.html
go tool cover -html=articles_coverage.out -o articles_coverage.html
```

#### 3.2 Coverage Requirements
- ✅ `common/` package: **minimum 70%**
- ✅ `users/` package: **minimum 70%**
- ✅ `articles/` package: **minimum 70%**
- ✅ Overall project: **minimum 70%**

#### 3.3 Create Coverage Report
**File to create:** `golang-gin-realworld-example-app/coverage-report.md`

**Content:**
1. **Current Coverage Statistics**
   - Coverage % per package
   - Overall project coverage %
   - Screenshots of coverage.html

2. **Identified Gaps**
   - Uncovered functions/methods
   - Reasons for gaps
   - Critical code to test

3. **Improvement Plan**
   - Tests to add for 80% coverage
   - High-value test cases

**Screenshot Commands:**
```powershell
# Open coverage reports in browser for screenshots
Start-Process coverage.html
Start-Process common_coverage.html
Start-Process users_coverage.html
Start-Process articles_coverage.html
```

---

## 🎯 Part B: Frontend Testing (React/Redux)

### Current State Analysis
❌ **No existing tests found** - Need to set up testing infrastructure

### Prerequisites Setup

**Install Testing Dependencies:**
```powershell
cd react-redux-realworld-example-app
npm install --save-dev @testing-library/react @testing-library/jest-dom @testing-library/user-event redux-mock-store
```

### Task 4: Component Unit Tests (40 points)

#### 4.1 Setup Test Utils
**File to create:** `src/test-utils.js`
- Redux store mock configuration
- Router wrapper utilities
- Common test helpers

#### 4.2 Component Test Files (Minimum 20 tests across 5 files)

**1. ArticleList Component** - `src/components/ArticleList.test.js` (4 tests)
- Test empty articles array rendering
- Test multiple articles rendering
- Test loading state display
- Test article click navigation

**2. ArticlePreview Component** - `src/components/ArticlePreview.test.js` (4 tests)
- Test article data rendering (title, description, author)
- Test favorite button rendering
- Test tag list rendering
- Test author profile link

**3. Login Component** - `src/components/Login.test.js` (5 tests)
- Test form rendering (email & password fields)
- Test email input updates
- Test password input updates
- Test form submission
- Test error message display

**4. Header Component** - `src/components/Header.test.js` (4 tests)
- Test navigation links for authenticated user
- Test navigation links for guest user
- Test "New Article" link visibility (logged in only)
- Test "Sign In" link visibility (guest only)

**5. Editor Component** - `src/components/Editor.test.js` (3 tests)
- Test article form fields rendering
- Test tag input functionality
- Test form submission

**Command to test:**
```powershell
cd react-redux-realworld-example-app
npm test
```

**For CI/non-interactive:**
```powershell
npm test -- --watchAll=false
```

### Task 5: Redux Integration Tests (30 points)

#### 5.1 Reducer Tests (3 files, ~9 tests)

**1. Auth Reducer** - `src/reducers/auth.test.js` (3 tests)
- Test `LOGIN` action updates token and user
- Test `LOGOUT` action clears state
- Test `REGISTER` action sets user data

**2. Article List Reducer** - `src/reducers/articleList.test.js` (3 tests)
- Test `ARTICLE_PAGE_LOADED` updates articles
- Test pagination state updates
- Test `APPLY_TAG_FILTER` changes active filter

**3. Editor Reducer** - `src/reducers/editor.test.js` (3 tests)
- Test `UPDATE_FIELD_EDITOR` updates form fields
- Test `EDITOR_PAGE_LOADED` initializes form
- Test `ADD_TAG` adds tag to list

#### 5.2 Action Creator Tests
**File to create:** `src/actions.test.js` (3 tests)
- Test action creators return correct action types
- Test action creators include correct payloads
- Test async action structure

#### 5.3 Middleware Tests
**File to create:** `src/middleware.test.js` (3 tests)
- Test promise middleware unwraps promises
- Test localStorage middleware saves token
- Test request cancellation for outdated requests

**Command to test:**
```powershell
npm test -- --testPathPattern="reducers|actions|middleware"
```

### Task 6: Frontend Integration Tests (30 points)

**File to create:** `src/integration.test.js`

#### Integration Test Cases (Minimum 5 tests)

**1. Login Flow Test**
- Render login form
- Fill email and password fields
- Submit form
- Verify Redux state updates
- Verify localStorage contains token
- Verify navigation to home page

**2. Registration Flow Test**
- Render registration form
- Fill all required fields
- Submit form
- Verify user is created
- Verify automatic login

**3. Article Creation Flow Test**
- User authenticated
- Navigate to editor
- Fill article form (title, description, body, tags)
- Submit form
- Verify article created
- Verify redirect to article page

**4. Article Favorite Flow Test**
- Render article preview
- Click favorite button
- Verify API call made
- Verify Redux state updates
- Verify button style changes
- Verify favorite count increments

**5. Logout Flow Test**
- User authenticated
- Click logout button
- Verify localStorage cleared
- Verify Redux state cleared
- Verify redirect to home page

**Command to test:**
```powershell
npm test -- --testPathPattern="integration"
```

---

## 📸 Evidence Collection Commands

### Backend Evidence

**1. Run All Backend Tests:**
```powershell
cd golang-gin-realworld-example-app
go test ./... -v > test_results.txt
```

**2. Generate Coverage:**
```powershell
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
Start-Process coverage.html
# Take screenshot of coverage.html
```

**3. Run Specific Package Tests:**
```powershell
# Articles tests
go test ./articles -v

# Common tests
go test ./common -v

# Users tests
go test ./users -v

# Integration tests
go test -v integration_test.go
```

**4. Coverage by Package:**
```powershell
go test ./common -cover
go test ./users -cover
go test ./articles -cover
```

### Frontend Evidence

**1. Run All Frontend Tests:**
```powershell
cd react-redux-realworld-example-app
npm test -- --watchAll=false --coverage
```

**2. Run Specific Test Suites:**
```powershell
# Component tests
npm test -- --testPathPattern="components" --watchAll=false

# Reducer tests
npm test -- --testPathPattern="reducers" --watchAll=false

# Integration tests
npm test -- --testPathPattern="integration" --watchAll=false
```

**3. Generate Coverage Report:**
```powershell
npm test -- --watchAll=false --coverage --coverageReporters=html
# Coverage report will be in coverage/lcov-report/index.html
Start-Process coverage/lcov-report/index.html
# Take screenshot
```

---

## 📦 Deliverables Checklist

### Backend Files
- [ ] `golang-gin-realworld-example-app/testing-analysis.md`
- [ ] `golang-gin-realworld-example-app/articles/unit_test.go`
- [ ] `golang-gin-realworld-example-app/common/unit_test.go` (enhanced)
- [ ] `golang-gin-realworld-example-app/integration_test.go`
- [ ] `golang-gin-realworld-example-app/coverage.out`
- [ ] `golang-gin-realworld-example-app/coverage.html`
- [ ] `golang-gin-realworld-example-app/coverage-report.md`

### Frontend Files
- [ ] `react-redux-realworld-example-app/src/test-utils.js`
- [ ] `react-redux-realworld-example-app/src/components/ArticleList.test.js`
- [ ] `react-redux-realworld-example-app/src/components/ArticlePreview.test.js`
- [ ] `react-redux-realworld-example-app/src/components/Login.test.js`
- [ ] `react-redux-realworld-example-app/src/components/Header.test.js`
- [ ] `react-redux-realworld-example-app/src/components/Editor.test.js`
- [ ] `react-redux-realworld-example-app/src/reducers/auth.test.js`
- [ ] `react-redux-realworld-example-app/src/reducers/articleList.test.js`
- [ ] `react-redux-realworld-example-app/src/reducers/editor.test.js`
- [ ] `react-redux-realworld-example-app/src/actions.test.js`
- [ ] `react-redux-realworld-example-app/src/middleware.test.js`
- [ ] `react-redux-realworld-example-app/src/integration.test.js`

### Documentation
- [ ] `ASSIGNMENT_1_REPORT.md` (Summary of testing approach, test cases, coverage)

### Evidence (Screenshots)
- [ ] Backend test execution (all passing)
- [ ] Backend coverage report (HTML)
- [ ] Backend coverage by package
- [ ] Frontend test execution (all passing)
- [ ] Frontend coverage report (HTML)
- [ ] Integration tests passing

---

## 🚀 Execution Order

### Phase 1: Backend Setup & Analysis (30 min)
1. Run existing tests to understand current state
2. Create `testing-analysis.md`
3. Identify articles package structure

### Phase 2: Backend Unit Tests (2 hours)
1. Create `articles/unit_test.go` with 15+ tests
2. Enhance `common/unit_test.go` with 5+ tests
3. Run tests and verify all pass

### Phase 3: Backend Integration Tests (2 hours)
1. Create `integration_test.go` in root
2. Implement 15+ integration test cases
3. Test authentication, CRUD, and interactions

### Phase 4: Backend Coverage (1 hour)
1. Generate coverage reports
2. Create `coverage-report.md`
3. Take screenshots
4. Analyze gaps and improvements

### Phase 5: Frontend Setup (30 min)
1. Install testing dependencies
2. Create `test-utils.js`
3. Set up mock store configuration

### Phase 6: Frontend Component Tests (2 hours)
1. Create 5 component test files
2. Implement 20+ component tests
3. Verify all tests pass

### Phase 7: Frontend Redux Tests (1.5 hours)
1. Create reducer test files
2. Create action and middleware tests
3. Verify Redux integration

### Phase 8: Frontend Integration Tests (1.5 hours)
1. Create `integration.test.js`
2. Implement 5+ end-to-end flow tests
3. Test with real Redux store

### Phase 9: Documentation & Evidence (1 hour)
1. Create `ASSIGNMENT_1_REPORT.md`
2. Collect all screenshots
3. Verify all deliverables
4. Final testing run

**Total Estimated Time:** 12-14 hours

---

## ⚠️ Important Notes

1. **Database Setup:** Integration tests will need a test database. Use SQLite in-memory or test database.
2. **Mock Data:** Create reusable fixtures for consistent testing.
3. **Test Independence:** Each test should be independent and not rely on others.
4. **Clean Up:** Always clean up test data after each test.
5. **Descriptive Names:** Use clear, descriptive test names that explain what's being tested.
6. **Edge Cases:** Don't forget to test error conditions and edge cases.
7. **Coverage Goal:** Aim for 70%+ but prioritize meaningful tests over just coverage percentage.

---

## 📚 Resources

- [Go Testing Package](https://golang.org/pkg/testing/)
- [Testify Assert](https://github.com/stretchr/testify)
- [Gin Testing](https://github.com/gin-gonic/gin#testing)
- [React Testing Library](https://testing-library.com/docs/react-testing-library/intro/)
- [Jest Documentation](https://jestjs.io/docs/getting-started)
- [Redux Testing](https://redux.js.org/usage/writing-tests)

---

## ✅ Ready to Start?

Once you review and approve this plan, I will:
1. ✅ Execute all tasks systematically
2. ✅ Create all required test files
3. ✅ Generate coverage reports
4. ✅ Provide you with commands for screenshots
5. ✅ Create comprehensive documentation

**Please confirm to proceed with implementation!**
