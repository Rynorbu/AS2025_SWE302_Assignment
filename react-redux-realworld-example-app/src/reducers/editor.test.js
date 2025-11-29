import editorReducer from './editor';
import {
  EDITOR_PAGE_LOADED,
  EDITOR_PAGE_UNLOADED,
  ARTICLE_SUBMITTED,
  ASYNC_START,
  ADD_TAG,
  REMOVE_TAG,
  UPDATE_FIELD_EDITOR
} from '../constants/actionTypes';

describe('editor reducer', () => {
  test('should return initial state', () => {
    expect(editorReducer(undefined, {})).toEqual({});
  });

  test('should handle EDITOR_PAGE_LOADED for new article', () => {
    const action = {
      type: EDITOR_PAGE_LOADED,
      payload: null
    };

    const expectedState = {
      articleSlug: '',
      title: '',
      description: '',
      body: '',
      tagInput: '',
      tagList: []
    };

    expect(editorReducer({}, action)).toEqual(expectedState);
  });

  test('should handle EDITOR_PAGE_LOADED for editing existing article', () => {
    const action = {
      type: EDITOR_PAGE_LOADED,
      payload: {
        article: {
          slug: 'existing-article',
          title: 'Existing Article',
          description: 'Existing Description',
          body: 'Existing Body',
          tagList: ['react', 'testing']
        }
      }
    };

    const expectedState = {
      articleSlug: 'existing-article',
      title: 'Existing Article',
      description: 'Existing Description',
      body: 'Existing Body',
      tagInput: '',
      tagList: ['react', 'testing']
    };

    expect(editorReducer({}, action)).toEqual(expectedState);
  });

  test('should handle EDITOR_PAGE_UNLOADED', () => {
    const initialState = {
      articleSlug: 'test-article',
      title: 'Test',
      description: 'Test Description',
      body: 'Test Body',
      tagList: ['test']
    };

    const action = { type: EDITOR_PAGE_UNLOADED };

    expect(editorReducer(initialState, action)).toEqual({});
  });

  test('should handle UPDATE_FIELD_EDITOR for title', () => {
    const initialState = {
      title: '',
      description: '',
      body: ''
    };

    const action = {
      type: UPDATE_FIELD_EDITOR,
      key: 'title',
      value: 'New Article Title'
    };

    const expectedState = {
      title: 'New Article Title',
      description: '',
      body: ''
    };

    expect(editorReducer(initialState, action)).toEqual(expectedState);
  });

  test('should handle UPDATE_FIELD_EDITOR for description', () => {
    const initialState = {
      title: 'Title',
      description: '',
      body: ''
    };

    const action = {
      type: UPDATE_FIELD_EDITOR,
      key: 'description',
      value: 'Article description'
    };

    const expectedState = {
      title: 'Title',
      description: 'Article description',
      body: ''
    };

    expect(editorReducer(initialState, action)).toEqual(expectedState);
  });

  test('should handle UPDATE_FIELD_EDITOR for body', () => {
    const initialState = {
      title: 'Title',
      description: 'Description',
      body: ''
    };

    const action = {
      type: UPDATE_FIELD_EDITOR,
      key: 'body',
      value: 'Article body content'
    };

    const expectedState = {
      title: 'Title',
      description: 'Description',
      body: 'Article body content'
    };

    expect(editorReducer(initialState, action)).toEqual(expectedState);
  });

  test('should handle ADD_TAG', () => {
    const initialState = {
      tagList: ['react', 'javascript'],
      tagInput: 'testing'
    };

    const action = { type: ADD_TAG };

    const expectedState = {
      tagList: ['react', 'javascript', 'testing'],
      tagInput: ''
    };

    expect(editorReducer(initialState, action)).toEqual(expectedState);
  });

  test('should handle ADD_TAG to empty tag list', () => {
    const initialState = {
      tagList: [],
      tagInput: 'firsttag'
    };

    const action = { type: ADD_TAG };

    const expectedState = {
      tagList: ['firsttag'],
      tagInput: ''
    };

    expect(editorReducer(initialState, action)).toEqual(expectedState);
  });

  test('should handle REMOVE_TAG', () => {
    const initialState = {
      tagList: ['react', 'javascript', 'testing', 'redux']
    };

    const action = {
      type: REMOVE_TAG,
      tag: 'javascript'
    };

    const expectedState = {
      tagList: ['react', 'testing', 'redux']
    };

    expect(editorReducer(initialState, action)).toEqual(expectedState);
  });

  test('should handle REMOVE_TAG that does not exist', () => {
    const initialState = {
      tagList: ['react', 'redux']
    };

    const action = {
      type: REMOVE_TAG,
      tag: 'nonexistent'
    };

    const expectedState = {
      tagList: ['react', 'redux']
    };

    expect(editorReducer(initialState, action)).toEqual(expectedState);
  });

  test('should handle ARTICLE_SUBMITTED successfully', () => {
    const initialState = {
      title: 'Test Article',
      inProgress: true
    };

    const action = {
      type: ARTICLE_SUBMITTED,
      error: false,
      payload: {}
    };

    const expectedState = {
      title: 'Test Article',
      inProgress: null,
      errors: null
    };

    expect(editorReducer(initialState, action)).toEqual(expectedState);
  });

  test('should handle ARTICLE_SUBMITTED with errors', () => {
    const initialState = {
      title: 'Test Article',
      inProgress: true
    };

    const action = {
      type: ARTICLE_SUBMITTED,
      error: true,
      payload: {
        errors: {
          title: ['is too short'],
          body: ['is required']
        }
      }
    };

    const expectedState = {
      title: 'Test Article',
      inProgress: null,
      errors: {
        title: ['is too short'],
        body: ['is required']
      }
    };

    expect(editorReducer(initialState, action)).toEqual(expectedState);
  });

  test('should handle ASYNC_START for ARTICLE_SUBMITTED', () => {
    const initialState = {
      title: 'Test Article',
      description: 'Description'
    };

    const action = {
      type: ASYNC_START,
      subtype: ARTICLE_SUBMITTED
    };

    const expectedState = {
      title: 'Test Article',
      description: 'Description',
      inProgress: true
    };

    expect(editorReducer(initialState, action)).toEqual(expectedState);
  });
});
