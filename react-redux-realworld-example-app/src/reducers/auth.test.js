import authReducer from './auth';
import {
  LOGIN,
  LOGOUT,
  REGISTER,
  LOGIN_PAGE_UNLOADED,
  REGISTER_PAGE_UNLOADED,
  ASYNC_START,
  UPDATE_FIELD_AUTH
} from '../constants/actionTypes';

describe('auth reducer', () => {
  test('should return initial state', () => {
    expect(authReducer(undefined, {})).toEqual({});
  });

  test('should handle LOGIN with success', () => {
    const action = {
      type: LOGIN,
      error: false,
      payload: {
        user: {
          email: 'test@example.com',
          token: 'jwt-token',
          username: 'testuser'
        }
      }
    };

    const expectedState = {
      inProgress: false,
      errors: null
    };

    expect(authReducer({}, action)).toEqual(expectedState);
  });

  test('should handle LOGIN with error', () => {
    const action = {
      type: LOGIN,
      error: true,
      payload: {
        errors: {
          'email or password': ['is invalid']
        }
      }
    };

    const expectedState = {
      inProgress: false,
      errors: {
        'email or password': ['is invalid']
      }
    };

    expect(authReducer({}, action)).toEqual(expectedState);
  });

  test('should handle REGISTER action', () => {
    const action = {
      type: REGISTER,
      error: false,
      payload: {
        user: {
          email: 'new@example.com',
          token: 'new-jwt-token',
          username: 'newuser'
        }
      }
    };

    const expectedState = {
      inProgress: false,
      errors: null
    };

    expect(authReducer({}, action)).toEqual(expectedState);
  });

  test('should handle REGISTER with validation errors', () => {
    const action = {
      type: REGISTER,
      error: true,
      payload: {
        errors: {
          username: ['is too short'],
          email: ['is invalid']
        }
      }
    };

    const expectedState = {
      inProgress: false,
      errors: {
        username: ['is too short'],
        email: ['is invalid']
      }
    };

    expect(authReducer({}, action)).toEqual(expectedState);
  });

  test('should handle LOGIN_PAGE_UNLOADED', () => {
    const initialState = {
      email: 'test@example.com',
      password: 'password123',
      errors: { test: 'error' }
    };

    const action = { type: LOGIN_PAGE_UNLOADED };

    expect(authReducer(initialState, action)).toEqual({});
  });

  test('should handle REGISTER_PAGE_UNLOADED', () => {
    const initialState = {
      username: 'testuser',
      email: 'test@example.com',
      password: 'password123'
    };

    const action = { type: REGISTER_PAGE_UNLOADED };

    expect(authReducer(initialState, action)).toEqual({});
  });

  test('should handle UPDATE_FIELD_AUTH for email', () => {
    const initialState = {
      email: '',
      password: ''
    };

    const action = {
      type: UPDATE_FIELD_AUTH,
      key: 'email',
      value: 'test@example.com'
    };

    const expectedState = {
      email: 'test@example.com',
      password: ''
    };

    expect(authReducer(initialState, action)).toEqual(expectedState);
  });

  test('should handle UPDATE_FIELD_AUTH for password', () => {
    const initialState = {
      email: 'test@example.com',
      password: ''
    };

    const action = {
      type: UPDATE_FIELD_AUTH,
      key: 'password',
      value: 'secretpassword'
    };

    const expectedState = {
      email: 'test@example.com',
      password: 'secretpassword'
    };

    expect(authReducer(initialState, action)).toEqual(expectedState);
  });

  test('should handle ASYNC_START for LOGIN', () => {
    const initialState = {
      email: 'test@example.com',
      password: 'password123'
    };

    const action = {
      type: ASYNC_START,
      subtype: LOGIN
    };

    const expectedState = {
      email: 'test@example.com',
      password: 'password123',
      inProgress: true
    };

    expect(authReducer(initialState, action)).toEqual(expectedState);
  });

  test('should handle ASYNC_START for REGISTER', () => {
    const initialState = {
      username: 'newuser',
      email: 'new@example.com',
      password: 'password123'
    };

    const action = {
      type: ASYNC_START,
      subtype: REGISTER
    };

    const expectedState = {
      username: 'newuser',
      email: 'new@example.com',
      password: 'password123',
      inProgress: true
    };

    expect(authReducer(initialState, action)).toEqual(expectedState);
  });
});
