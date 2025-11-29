import React from 'react';
import { render, screen } from '@testing-library/react';
import { BrowserRouter } from 'react-router-dom';
import ArticleList from './ArticleList';
import '@testing-library/jest-dom';

describe('ArticleList Component', () => {
  const mockArticles = [
    {
      slug: 'article-1',
      title: 'Article 1',
      description: 'Description 1',
      body: 'Body 1',
      tagList: ['tag1'],
      createdAt: '2025-11-29T00:00:00.000Z',
      favorited: false,
      favoritesCount: 5,
      author: {
        username: 'author1',
        bio: 'Bio 1',
        image: 'https://example.com/image1.jpg',
        following: false
      }
    },
    {
      slug: 'article-2',
      title: 'Article 2',
      description: 'Description 2',
      body: 'Body 2',
      tagList: ['tag2'],
      createdAt: '2025-11-29T00:00:00.000Z',
      favorited: true,
      favoritesCount: 10,
      author: {
        username: 'author2',
        bio: 'Bio 2',
        image: 'https://example.com/image2.jpg',
        following: true
      }
    }
  ];

  test('renders loading state when articles is null/undefined', () => {
    render(
      <BrowserRouter>
        <ArticleList articles={null} />
      </BrowserRouter>
    );
    
    expect(screen.getByText('Loading...')).toBeInTheDocument();
  });

  test('renders empty state when articles array is empty', () => {
    render(
      <BrowserRouter>
        <ArticleList articles={[]} />
      </BrowserRouter>
    );
    
    expect(screen.getByText('No articles are here... yet.')).toBeInTheDocument();
  });

  test('renders multiple articles when articles array has items', () => {
    render(
      <BrowserRouter>
        <ArticleList articles={mockArticles} />
      </BrowserRouter>
    );
    
    expect(screen.getByText('Article 1')).toBeInTheDocument();
    expect(screen.getByText('Article 2')).toBeInTheDocument();
    expect(screen.getByText('Description 1')).toBeInTheDocument();
    expect(screen.getByText('Description 2')).toBeInTheDocument();
  });

  test('renders correct number of articles', () => {
    const { container } = render(
      <BrowserRouter>
        <ArticleList articles={mockArticles} />
      </BrowserRouter>
    );
    
    const articlePreviews = container.querySelectorAll('.article-preview');
    expect(articlePreviews.length).toBe(2);
  });
});
