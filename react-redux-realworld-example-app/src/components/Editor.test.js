import React from 'react';
import { render, screen, fireEvent } from '@testing-library/react';
import { Provider } from 'react-redux';
import { BrowserRouter, Route } from 'react-router-dom';
import Editor from './Editor';
import { createMockStore } from '../test-utils';
import '@testing-library/jest-dom';

describe('Editor Component', () => {
  const mockMatch = {
    params: {},
    isExact: true,
    path: '/editor',
    url: '/editor'
  };

  test('renders article form fields', () => {
    const store = createMockStore({
      editor: {
        title: '',
        description: '',
        body: '',
        tagInput: '',
        tagList: []
      }
    });

    render(
      <Provider store={store}>
        <BrowserRouter>
          <Editor match={mockMatch} />
        </BrowserRouter>
      </Provider>
    );

    expect(screen.getByPlaceholderText('Article Title')).toBeInTheDocument();
    expect(screen.getByPlaceholderText("What's this article about?")).toBeInTheDocument();
    expect(screen.getByPlaceholderText('Write your article (in markdown)')).toBeInTheDocument();
    expect(screen.getByPlaceholderText('Enter tags')).toBeInTheDocument();
  });

  test('renders publish button', () => {
    const store = createMockStore({
      editor: {
        title: '',
        description: '',
        body: '',
        tagInput: '',
        tagList: []
      }
    });

    render(
      <Provider store={store}>
        <BrowserRouter>
          <Editor match={mockMatch} />
        </BrowserRouter>
      </Provider>
    );

    const publishButton = screen.getByRole('button', { name: /publish article/i });
    expect(publishButton).toBeInTheDocument();
  });

  test('displays existing tags', () => {
    const store = createMockStore({
      editor: {
        title: 'Test Article',
        description: 'Test Description',
        body: 'Test Body',
        tagInput: '',
        tagList: ['react', 'testing', 'javascript']
      }
    });

    render(
      <Provider store={store}>
        <BrowserRouter>
          <Editor match={mockMatch} />
        </BrowserRouter>
      </Provider>
    );

    expect(screen.getByText('react')).toBeInTheDocument();
    expect(screen.getByText('testing')).toBeInTheDocument();
    expect(screen.getByText('javascript')).toBeInTheDocument();
  });

  test('form fields display current values from store', () => {
    const store = createMockStore({
      editor: {
        title: 'My Article Title',
        description: 'My Description',
        body: 'My Article Body',
        tagInput: 'newtag',
        tagList: []
      }
    });

    render(
      <Provider store={store}>
        <BrowserRouter>
          <Editor match={mockMatch} />
        </BrowserRouter>
      </Provider>
    );

    expect(screen.getByDisplayValue('My Article Title')).toBeInTheDocument();
    expect(screen.getByDisplayValue('My Description')).toBeInTheDocument();
    expect(screen.getByDisplayValue('My Article Body')).toBeInTheDocument();
    expect(screen.getByDisplayValue('newtag')).toBeInTheDocument();
  });

  test('publish button is disabled when inProgress is true', () => {
    const store = createMockStore({
      editor: {
        title: 'Test',
        description: 'Test',
        body: 'Test',
        tagInput: '',
        tagList: [],
        inProgress: true
      }
    });

    render(
      <Provider store={store}>
        <BrowserRouter>
          <Editor match={mockMatch} />
        </BrowserRouter>
      </Provider>
    );

    const publishButton = screen.getByRole('button', { name: /publish article/i });
    expect(publishButton).toBeDisabled();
  });
});
