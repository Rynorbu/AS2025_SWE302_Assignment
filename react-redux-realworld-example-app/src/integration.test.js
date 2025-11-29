/**
 * Integration Tests for React/Redux Application
 * 
 * These tests verify end-to-end flows combining components,
 * Redux state management, and user interactions.
 */

import React from 'react';
import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import { Provider } from 'react-redux';
import { BrowserRouter } from 'react-router-dom';
import { createStore, applyMiddleware } from 'redux';
import reducer from './reducer';
import { promiseMiddleware, localStorageMiddleware } from './middleware';
import Login from './components/Login';
import Header from './components/Header';
import ArticleList from './components/ArticleList';
import ArticlePreview from './components/ArticlePreview';
import { LOGIN, LOGOUT, UPDATE_FIELD_AUTH, ARTICLE_FAVORITED } from './constants/actionTypes';
import '@testing-library/jest-dom';

describe('Integration Tests', () => {
  let store;
  let mockLocalStorage;

  beforeEach(() => {
    // Create store with middleware
    store = createStore(
      reducer,
      applyMiddleware(promiseMiddleware, localStorageMiddleware)
    );

    // Mock localStorage
    mockLocalStorage = {
      getItem: jest.fn(),
      setItem: jest.fn(),
      removeItem: jest.fn(),
      clear: jest.fn()
    };
    global.window = { localStorage: mockLocalStorage };
  });

  test('Login Flow: Updates Redux state and localStorage', async () => {
    const mockUser = {
      email: 'test@example.com',
      token: 'test-jwt-token',
      username: 'testuser',
      bio: 'Test bio',
      image: null
    };

    // Simulate successful login
    store.dispatch({
      type: LOGIN,
      error: false,
      payload: { user: mockUser }
    });

    const state = store.getState();

    // Verify auth state is updated
    expect(state.auth.inProgress).toBe(false);
    expect(state.auth.errors).toBe(null);

    // Verify localStorage was called
    expect(mockLocalStorage.setItem).toHaveBeenCalledWith('jwt', 'test-jwt-token');
  });

  test('Login Flow: Displays error on failed login', () => {
    store.dispatch({
      type: LOGIN,
      error: true,
      payload: {
        errors: {
          'email or password': ['is invalid']
        }
      }
    });

    const state = store.getState();

    expect(state.auth.errors).toEqual({
      'email or password': ['is invalid']
    });
    expect(mockLocalStorage.setItem).not.toHaveBeenCalled();
  });

  test('Logout Flow: Clears localStorage and Redux state', () => {
    // Set initial logged-in state
    store.dispatch({
      type: LOGIN,
      error: false,
      payload: {
        user: {
          email: 'test@example.com',
          token: 'test-token',
          username: 'testuser'
        }
      }
    });

    // Logout
    store.dispatch({ type: LOGOUT });

    // Verify localStorage was cleared
    expect(mockLocalStorage.setItem).toHaveBeenCalledWith('jwt', '');
  });

  test('Article Favorite Flow: Updates article state in Redux', () => {
    const mockArticles = [
      {
        slug: 'test-article',
        title: 'Test Article',
        description: 'Description',
        body: 'Body',
        tagList: [],
        favorited: false,
        favoritesCount: 5,
        author: { username: 'author', image: null }
      }
    ];

    // Set initial article list state
    store.dispatch({
      type: 'HOME_PAGE_LOADED',
      tab: 'all',
      pager: () => {},
      payload: [
        { tags: [] },
        { articles: mockArticles, articlesCount: 1 }
      ]
    });

    // Favorite the article
    store.dispatch({
      type: ARTICLE_FAVORITED,
      payload: {
        article: {
          slug: 'test-article',
          favorited: true,
          favoritesCount: 6
        }
      }
    });

    const state = store.getState();
    const article = state.articleList.articles.find(a => a.slug === 'test-article');

    expect(article.favorited).toBe(true);
    expect(article.favoritesCount).toBe(6);
  });

  test('Header Component: Shows correct links based on auth state', () => {
    // Test without logged-in user
    const { rerender } = render(
      <Provider store={store}>
        <BrowserRouter>
          <Header appName="Conduit" currentUser={null} />
        </BrowserRouter>
      </Provider>
    );

    expect(screen.getByText('Sign in')).toBeInTheDocument();
    expect(screen.getByText('Sign up')).toBeInTheDocument();
    expect(screen.queryByText(/New Post/i)).not.toBeInTheDocument();

    // Simulate login
    const currentUser = {
      username: 'testuser',
      email: 'test@example.com',
      image: null
    };

    rerender(
      <Provider store={store}>
        <BrowserRouter>
          <Header appName="Conduit" currentUser={currentUser} />
        </BrowserRouter>
      </Provider>
    );

    expect(screen.queryByText('Sign in')).not.toBeInTheDocument();
    expect(screen.getByText(/New Post/i)).toBeInTheDocument();
    expect(screen.getByText('testuser')).toBeInTheDocument();
  });

  test('Article List: Renders articles from Redux state', () => {
    const mockArticles = [
      {
        slug: 'article-1',
        title: 'First Article',
        description: 'First Description',
        body: 'Body',
        tagList: ['react'],
        favorited: false,
        favoritesCount: 3,
        createdAt: '2025-11-29T00:00:00.000Z',
        author: {
          username: 'author1',
          image: null,
          following: false
        }
      },
      {
        slug: 'article-2',
        title: 'Second Article',
        description: 'Second Description',
        body: 'Body',
        tagList: ['javascript'],
        favorited: true,
        favoritesCount: 10,
        createdAt: '2025-11-29T00:00:00.000Z',
        author: {
          username: 'author2',
          image: null,
          following: false
        }
      }
    ];

    // Dispatch articles to store
    store.dispatch({
      type: 'HOME_PAGE_LOADED',
      tab: 'all',
      pager: () => {},
      payload: [
        { tags: [] },
        { articles: mockArticles, articlesCount: 2 }
      ]
    });

    render(
      <Provider store={store}>
        <BrowserRouter>
          <ArticleList
            articles={store.getState().articleList.articles}
            articlesCount={2}
            currentPage={0}
          />
        </BrowserRouter>
      </Provider>
    );

    expect(screen.getByText('First Article')).toBeInTheDocument();
    expect(screen.getByText('Second Article')).toBeInTheDocument();
    expect(screen.getByText('First Description')).toBeInTheDocument();
    expect(screen.getByText('Second Description')).toBeInTheDocument();
  });

  test('Form Input Updates: Login form fields update Redux state', () => {
    render(
      <Provider store={store}>
        <BrowserRouter>
          <Login />
        </BrowserRouter>
      </Provider>
    );

    const emailInput = screen.getByPlaceholderText('Email');
    const passwordInput = screen.getByPlaceholderText('Password');

    // Type in email field
    fireEvent.change(emailInput, { target: { value: 'user@example.com' } });

    // Check Redux state
    let state = store.getState();
    expect(state.auth.email).toBe('user@example.com');

    // Type in password field
    fireEvent.change(passwordInput, { target: { value: 'password123' } });

    // Check Redux state again
    state = store.getState();
    expect(state.auth.password).toBe('password123');
  });

  test('Complete User Flow: Registration -> Login -> View Articles', () => {
    // Step 1: Simulate registration
    store.dispatch({
      type: 'REGISTER',
      error: false,
      payload: {
        user: {
          email: 'newuser@example.com',
          token: 'new-user-token',
          username: 'newuser',
          bio: '',
          image: null
        }
      }
    });

    let state = store.getState();
    expect(state.common.token).toBe('new-user-token');
    expect(state.common.currentUser.username).toBe('newuser');

    // Step 2: Load articles
    const mockArticles = [
      {
        slug: 'welcome-article',
        title: 'Welcome Article',
        description: 'Welcome to the platform',
        body: 'Content',
        tagList: ['welcome'],
        favorited: false,
        favoritesCount: 0,
        createdAt: '2025-11-29T00:00:00.000Z',
        author: {
          username: 'admin',
          image: null,
          following: false
        }
      }
    ];

    store.dispatch({
      type: 'HOME_PAGE_LOADED',
      tab: 'all',
      pager: () => {},
      payload: [
        { tags: ['welcome'] },
        { articles: mockArticles, articlesCount: 1 }
      ]
    });

    state = store.getState();
    expect(state.articleList.articles.length).toBe(1);
    expect(state.articleList.articles[0].title).toBe('Welcome Article');

    // Verify localStorage was updated during registration
    expect(mockLocalStorage.setItem).toHaveBeenCalledWith('jwt', 'new-user-token');
  });
});
