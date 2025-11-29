import React from 'react';
import { render, screen } from '@testing-library/react';
import { BrowserRouter } from 'react-router-dom';
import Header from './Header';
import '@testing-library/jest-dom';

describe('Header Component', () => {
  const appName = 'Conduit';

  test('displays navigation links for guest users', () => {
    render(
      <BrowserRouter>
        <Header appName={appName} currentUser={null} />
      </BrowserRouter>
    );

    expect(screen.getByText('conduit')).toBeInTheDocument();
    expect(screen.getByText('Home')).toBeInTheDocument();
    expect(screen.getByText('Sign in')).toBeInTheDocument();
    expect(screen.getByText('Sign up')).toBeInTheDocument();
  });

  test('does not display New Post link for guest users', () => {
    render(
      <BrowserRouter>
        <Header appName={appName} currentUser={null} />
      </BrowserRouter>
    );

    expect(screen.queryByText(/New Post/i)).not.toBeInTheDocument();
    expect(screen.queryByText(/Settings/i)).not.toBeInTheDocument();
  });

  test('displays navigation links for logged-in users', () => {
    const currentUser = {
      username: 'testuser',
      email: 'test@example.com',
      bio: 'Test bio',
      image: 'https://example.com/image.jpg'
    };

    render(
      <BrowserRouter>
        <Header appName={appName} currentUser={currentUser} />
      </BrowserRouter>
    );

    expect(screen.getByText('Home')).toBeInTheDocument();
    expect(screen.getByText(/New Post/i)).toBeInTheDocument();
    expect(screen.getByText(/Settings/i)).toBeInTheDocument();
    expect(screen.getByText('testuser')).toBeInTheDocument();
  });

  test('does not display Sign in/Sign up links for logged-in users', () => {
    const currentUser = {
      username: 'testuser',
      email: 'test@example.com',
      image: 'https://example.com/image.jpg'
    };

    render(
      <BrowserRouter>
        <Header appName={appName} currentUser={currentUser} />
      </BrowserRouter>
    );

    expect(screen.queryByText('Sign in')).not.toBeInTheDocument();
    expect(screen.queryByText('Sign up')).not.toBeInTheDocument();
  });

  test('displays user profile image for logged-in users', () => {
    const currentUser = {
      username: 'testuser',
      email: 'test@example.com',
      image: 'https://example.com/user-image.jpg'
    };

    render(
      <BrowserRouter>
        <Header appName={appName} currentUser={currentUser} />
      </BrowserRouter>
    );

    const userImage = screen.getByAltText('testuser');
    expect(userImage).toBeInTheDocument();
    expect(userImage).toHaveAttribute('src', 'https://example.com/user-image.jpg');
  });

  test('profile link points to correct user profile', () => {
    const currentUser = {
      username: 'johndoe',
      email: 'john@example.com',
      image: 'https://example.com/john.jpg'
    };

    render(
      <BrowserRouter>
        <Header appName={appName} currentUser={currentUser} />
      </BrowserRouter>
    );

    const profileLink = screen.getByText('johndoe').closest('a');
    expect(profileLink).toHaveAttribute('href', '/@johndoe');
  });
});
