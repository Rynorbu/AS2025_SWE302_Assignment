// Test utilities for React/Redux testing
import React from 'react';
import { render } from '@testing-library/react';
import { Provider } from 'react-redux';
import { createStore, applyMiddleware } from 'redux';
import { BrowserRouter } from 'react-router-dom';
import { promiseMiddleware } from './middleware';
import reducer from './reducer';

// Create a test store with middleware
export function createMockStore(initialState = {}) {
  return createStore(
    reducer,
    initialState,
    applyMiddleware(promiseMiddleware)
  );
}

// Render component with Redux Provider and Router
export function renderWithProviders(
  ui,
  {
    initialState = {},
    store = createMockStore(initialState),
    ...renderOptions
  } = {}
) {
  function Wrapper({ children }) {
    return (
      <Provider store={store}>
        <BrowserRouter>
          {children}
        </BrowserRouter>
      </Provider>
    );
  }

  return {
    ...render(ui, { wrapper: Wrapper, ...renderOptions }),
    store
  };
}

// Create mock article data
export const mockArticle = {
  slug: 'test-article',
  title: 'Test Article',
  description: 'Test Description',
  body: 'Test Body',
  tagList: ['test', 'react'],
  createdAt: '2025-11-29T00:00:00.000Z',
  updatedAt: '2025-11-29T00:00:00.000Z',
  favorited: false,
  favoritesCount: 0,
  author: {
    username: 'testuser',
    bio: 'Test Bio',
    image: 'https://example.com/image.jpg',
    following: false
  }
};

// Create mock user data
export const mockUser = {
  email: 'test@example.com',
  token: 'mock-jwt-token',
  username: 'testuser',
  bio: 'Test Bio',
  image: 'https://example.com/image.jpg'
};

// Create mock store with common data
export function createMockStoreWithData() {
  return createMockStore({
    common: {
      appName: 'Conduit',
      token: mockUser.token,
      currentUser: mockUser
    },
    articleList: {
      articles: [mockArticle],
      articlesCount: 1,
      currentPage: 0
    },
    auth: {
      email: '',
      password: ''
    }
  });
}

// Mock localStorage
export const localStorageMock = {
  getItem: jest.fn(),
  setItem: jest.fn(),
  removeItem: jest.fn(),
  clear: jest.fn(),
};

global.localStorage = localStorageMock;

export default {
  renderWithProviders,
  createMockStore,
  createMockStoreWithData,
  mockArticle,
  mockUser,
  localStorageMock
};
