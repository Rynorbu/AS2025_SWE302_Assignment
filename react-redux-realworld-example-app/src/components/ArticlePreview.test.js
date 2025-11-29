import React from 'react';
import { render, screen, fireEvent } from '@testing-library/react';
import { Provider } from 'react-redux';
import { BrowserRouter } from 'react-router-dom';
import ArticlePreview from './ArticlePreview';
import { createMockStore, mockArticle } from '../test-utils';
import '@testing-library/jest-dom';

describe('ArticlePreview Component', () => {
  const store = createMockStore({
    common: { currentUser: null }
  });

  test('renders article title, description, and author', () => {
    render(
      <Provider store={store}>
        <BrowserRouter>
          <ArticlePreview article={mockArticle} />
        </BrowserRouter>
      </Provider>
    );

    expect(screen.getByText('Test Article')).toBeInTheDocument();
    expect(screen.getByText('Test Description')).toBeInTheDocument();
    expect(screen.getByText('testuser')).toBeInTheDocument();
  });

  test('renders favorite button with correct count', () => {
    render(
      <Provider store={store}>
        <BrowserRouter>
          <ArticlePreview article={mockArticle} />
        </BrowserRouter>
      </Provider>
    );

    const favoriteButton = screen.getByRole('button');
    expect(favoriteButton).toBeInTheDocument();
    expect(favoriteButton).toHaveTextContent('0');
  });

  test('renders tag list correctly', () => {
    render(
      <Provider store={store}>
        <BrowserRouter>
          <ArticlePreview article={mockArticle} />
        </BrowserRouter>
      </Provider>
    );

    expect(screen.getByText('test')).toBeInTheDocument();
    expect(screen.getByText('react')).toBeInTheDocument();
  });

  test('renders author profile link', () => {
    render(
      <Provider store={store}>
        <BrowserRouter>
          <ArticlePreview article={mockArticle} />
        </BrowserRouter>
      </Provider>
    );

    const authorLinks = screen.getAllByRole('link', { name: /testuser/i });
    expect(authorLinks.length).toBeGreaterThan(0);
    expect(authorLinks[0]).toHaveAttribute('href', '/@testuser');
  });

  test('favorite button has correct class when not favorited', () => {
    const unfavoritedArticle = { ...mockArticle, favorited: false };
    
    render(
      <Provider store={store}>
        <BrowserRouter>
          <ArticlePreview article={unfavoritedArticle} />
        </BrowserRouter>
      </Provider>
    );

    const favoriteButton = screen.getByRole('button');
    expect(favoriteButton).toHaveClass('btn-outline-primary');
  });

  test('favorite button has correct class when favorited', () => {
    const favoritedArticle = { ...mockArticle, favorited: true };
    
    render(
      <Provider store={store}>
        <BrowserRouter>
          <ArticlePreview article={favoritedArticle} />
        </BrowserRouter>
      </Provider>
    );

    const favoriteButton = screen.getByRole('button');
    expect(favoriteButton).toHaveClass('btn-primary');
  });
});
