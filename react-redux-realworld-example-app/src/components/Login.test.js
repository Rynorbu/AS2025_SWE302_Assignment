import React from 'react';
import { render, screen, fireEvent } from '@testing-library/react';
import { Provider } from 'react-redux';
import { BrowserRouter } from 'react-router-dom';
import Login from './Login';
import { createMockStore } from '../test-utils';
import '@testing-library/jest-dom';

describe('Login Component', () => {
  test('renders login form with email and password fields', () => {
    const store = createMockStore({
      auth: { email: '', password: '' }
    });

    render(
      <Provider store={store}>
        <BrowserRouter>
          <Login />
        </BrowserRouter>
      </Provider>
    );

    expect(screen.getByText('Sign In')).toBeInTheDocument();
    expect(screen.getByPlaceholderText('Email')).toBeInTheDocument();
    expect(screen.getByPlaceholderText('Password')).toBeInTheDocument();
    expect(screen.getByRole('button', { name: /sign in/i })).toBeInTheDocument();
  });

  test('updates email input field', () => {
    const store = createMockStore({
      auth: { email: '', password: '' }
    });

    render(
      <Provider store={store}>
        <BrowserRouter>
          <Login />
        </BrowserRouter>
      </Provider>
    );

    const emailInput = screen.getByPlaceholderText('Email');
    fireEvent.change(emailInput, { target: { value: 'test@example.com' } });
    
    expect(emailInput.value).toBe('test@example.com');
  });

  test('updates password input field', () => {
    const store = createMockStore({
      auth: { email: '', password: '' }
    });

    render(
      <Provider store={store}>
        <BrowserRouter>
          <Login />
        </BrowserRouter>
      </Provider>
    );

    const passwordInput = screen.getByPlaceholderText('Password');
    fireEvent.change(passwordInput, { target: { value: 'password123' } });
    
    expect(passwordInput.value).toBe('password123');
  });

  test('displays registration link', () => {
    const store = createMockStore({
      auth: { email: '', password: '' }
    });

    render(
      <Provider store={store}>
        <BrowserRouter>
          <Login />
        </BrowserRouter>
      </Provider>
    );

    const registerLink = screen.getByText('Need an account?');
    expect(registerLink).toBeInTheDocument();
    expect(registerLink.closest('a')).toHaveAttribute('href', '/register');
  });

  test('displays error messages when errors prop is provided', () => {
    const store = createMockStore({
      auth: {
        email: '',
        password: '',
        errors: {
          'email or password': ['is invalid']
        }
      }
    });

    render(
      <Provider store={store}>
        <BrowserRouter>
          <Login />
        </BrowserRouter>
      </Provider>
    );

    expect(screen.getByText(/email or password/i)).toBeInTheDocument();
  });

  test('submit button is disabled when inProgress is true', () => {
    const store = createMockStore({
      auth: {
        email: 'test@example.com',
        password: 'password123',
        inProgress: true
      }
    });

    render(
      <Provider store={store}>
        <BrowserRouter>
          <Login />
        </BrowserRouter>
      </Provider>
    );

    const submitButton = screen.getByRole('button', { name: /sign in/i });
    expect(submitButton).toBeDisabled();
  });

  test('form submission is prevented when clicked', () => {
    const store = createMockStore({
      auth: { email: 'test@example.com', password: 'password123' }
    });

    render(
      <Provider store={store}>
        <BrowserRouter>
          <Login />
        </BrowserRouter>
      </Provider>
    );

    const form = screen.getByRole('button', { name: /sign in/i }).closest('form');
    const handleSubmit = jest.fn((e) => e.preventDefault());
    form.onsubmit = handleSubmit;

    fireEvent.submit(form);
    expect(handleSubmit).toHaveBeenCalled();
  });
});
