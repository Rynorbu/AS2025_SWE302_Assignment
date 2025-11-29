import { promiseMiddleware, localStorageMiddleware } from './middleware';
import { ASYNC_START, ASYNC_END, LOGIN, LOGOUT, REGISTER } from './constants/actionTypes';

describe('promiseMiddleware', () => {
  let store;
  let next;

  beforeEach(() => {
    store = {
      getState: jest.fn(() => ({ viewChangeCounter: 0 })),
      dispatch: jest.fn()
    };
    next = jest.fn();
  });

  test('calls next for non-promise actions', () => {
    const action = { type: 'TEST_ACTION', payload: 'test' };
    
    promiseMiddleware(store)(next)(action);
    
    expect(next).toHaveBeenCalledWith(action);
  });

  test('dispatches ASYNC_START for promise actions', () => {
    const action = {
      type: 'TEST_ACTION',
      payload: Promise.resolve({ data: 'test' })
    };
    
    promiseMiddleware(store)(next)(action);
    
    expect(store.dispatch).toHaveBeenCalledWith({
      type: ASYNC_START,
      subtype: 'TEST_ACTION'
    });
  });

  test('handles successful promise resolution', async () => {
    const resolvedData = { data: 'success' };
    const action = {
      type: 'TEST_ACTION',
      payload: Promise.resolve(resolvedData)
    };
    
    promiseMiddleware(store)(next)(action);
    
    // Wait for promise to resolve
    await new Promise(resolve => setTimeout(resolve, 10));
    
    expect(store.dispatch).toHaveBeenCalledWith(
      expect.objectContaining({
        type: ASYNC_END
      })
    );
  });

  test('handles promise rejection with error', async () => {
    const error = new Error('Test error');
    error.response = { body: { errors: { message: ['Error occurred'] } } };
    
    const action = {
      type: 'TEST_ACTION',
      payload: Promise.reject(error)
    };
    
    promiseMiddleware(store)(next)(action);
    
    // Wait for promise to reject
    await new Promise(resolve => setTimeout(resolve, 10));
    
    expect(store.dispatch).toHaveBeenCalled();
  });

  test('skips outdated requests when view changes', async () => {
    let viewCounter = 0;
    store.getState = jest.fn(() => ({ viewChangeCounter: viewCounter }));
    
    const action = {
      type: 'TEST_ACTION',
      payload: new Promise(resolve => {
        setTimeout(() => {
          viewCounter = 5; // Simulate view change
          resolve({ data: 'test' });
        }, 10);
      })
    };
    
    promiseMiddleware(store)(next)(action);
    
    await new Promise(resolve => setTimeout(resolve, 20));
    
    // Should have dispatched ASYNC_START but not final action due to view change
    expect(store.dispatch).toHaveBeenCalledWith({
      type: ASYNC_START,
      subtype: 'TEST_ACTION'
    });
  });
});

describe('localStorageMiddleware', () => {
  let store;
  let next;
  let mockLocalStorage;

  beforeEach(() => {
    store = {
      getState: jest.fn(),
      dispatch: jest.fn()
    };
    next = jest.fn();
    
    // Mock localStorage
    mockLocalStorage = {
      getItem: jest.fn(),
      setItem: jest.fn(),
      removeItem: jest.fn(),
      clear: jest.fn()
    };
    global.window = { localStorage: mockLocalStorage };
  });

  test('saves JWT token to localStorage on LOGIN', () => {
    const action = {
      type: LOGIN,
      error: false,
      payload: {
        user: {
          email: 'test@example.com',
          token: 'jwt-token-12345',
          username: 'testuser'
        }
      }
    };
    
    localStorageMiddleware(store)(next)(action);
    
    expect(mockLocalStorage.setItem).toHaveBeenCalledWith('jwt', 'jwt-token-12345');
    expect(next).toHaveBeenCalledWith(action);
  });

  test('saves JWT token to localStorage on REGISTER', () => {
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
    
    localStorageMiddleware(store)(next)(action);
    
    expect(mockLocalStorage.setItem).toHaveBeenCalledWith('jwt', 'new-jwt-token');
    expect(next).toHaveBeenCalledWith(action);
  });

  test('does not save token on LOGIN with error', () => {
    const action = {
      type: LOGIN,
      error: true,
      payload: {
        errors: { 'email or password': ['is invalid'] }
      }
    };
    
    localStorageMiddleware(store)(next)(action);
    
    expect(mockLocalStorage.setItem).not.toHaveBeenCalled();
    expect(next).toHaveBeenCalledWith(action);
  });

  test('clears JWT token on LOGOUT', () => {
    const action = { type: LOGOUT };
    
    localStorageMiddleware(store)(next)(action);
    
    expect(mockLocalStorage.setItem).toHaveBeenCalledWith('jwt', '');
    expect(next).toHaveBeenCalledWith(action);
  });

  test('passes through unrelated actions', () => {
    const action = { type: 'UNRELATED_ACTION', payload: 'data' };
    
    localStorageMiddleware(store)(next)(action);
    
    expect(mockLocalStorage.setItem).not.toHaveBeenCalled();
    expect(next).toHaveBeenCalledWith(action);
  });
});
