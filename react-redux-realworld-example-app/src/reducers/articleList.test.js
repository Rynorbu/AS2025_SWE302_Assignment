import articleListReducer from './articleList';
import {
  ARTICLE_FAVORITED,
  ARTICLE_UNFAVORITED,
  SET_PAGE,
  APPLY_TAG_FILTER,
  HOME_PAGE_LOADED,
  HOME_PAGE_UNLOADED,
  CHANGE_TAB
} from '../constants/actionTypes';

describe('articleList reducer', () => {
  test('should return initial state', () => {
    expect(articleListReducer(undefined, {})).toEqual({});
  });

  test('should handle ARTICLE_FAVORITED', () => {
    const initialState = {
      articles: [
        { slug: 'article-1', title: 'Article 1', favorited: false, favoritesCount: 5 },
        { slug: 'article-2', title: 'Article 2', favorited: false, favoritesCount: 10 }
      ]
    };

    const action = {
      type: ARTICLE_FAVORITED,
      payload: {
        article: {
          slug: 'article-1',
          favorited: true,
          favoritesCount: 6
        }
      }
    };

    const newState = articleListReducer(initialState, action);

    expect(newState.articles[0].favorited).toBe(true);
    expect(newState.articles[0].favoritesCount).toBe(6);
    expect(newState.articles[1].favorited).toBe(false); // Other articles unchanged
  });

  test('should handle ARTICLE_UNFAVORITED', () => {
    const initialState = {
      articles: [
        { slug: 'article-1', title: 'Article 1', favorited: true, favoritesCount: 6 },
        { slug: 'article-2', title: 'Article 2', favorited: true, favoritesCount: 10 }
      ]
    };

    const action = {
      type: ARTICLE_UNFAVORITED,
      payload: {
        article: {
          slug: 'article-1',
          favorited: false,
          favoritesCount: 5
        }
      }
    };

    const newState = articleListReducer(initialState, action);

    expect(newState.articles[0].favorited).toBe(false);
    expect(newState.articles[0].favoritesCount).toBe(5);
  });

  test('should handle SET_PAGE', () => {
    const initialState = {
      articles: [],
      articlesCount: 0,
      currentPage: 0
    };

    const newArticles = [
      { slug: 'page-2-article-1', title: 'Page 2 Article 1' },
      { slug: 'page-2-article-2', title: 'Page 2 Article 2' }
    ];

    const action = {
      type: SET_PAGE,
      page: 2,
      payload: {
        articles: newArticles,
        articlesCount: 50
      }
    };

    const newState = articleListReducer(initialState, action);

    expect(newState.articles).toEqual(newArticles);
    expect(newState.articlesCount).toBe(50);
    expect(newState.currentPage).toBe(2);
  });

  test('should handle APPLY_TAG_FILTER', () => {
    const initialState = {
      articles: [],
      tab: 'all',
      tag: null,
      currentPage: 5
    };

    const filteredArticles = [
      { slug: 'react-article', title: 'React Article', tagList: ['react'] }
    ];

    const action = {
      type: APPLY_TAG_FILTER,
      tag: 'react',
      pager: () => {},
      payload: {
        articles: filteredArticles,
        articlesCount: 15
      }
    };

    const newState = articleListReducer(initialState, action);

    expect(newState.articles).toEqual(filteredArticles);
    expect(newState.articlesCount).toBe(15);
    expect(newState.tag).toBe('react');
    expect(newState.tab).toBe(null);
    expect(newState.currentPage).toBe(0); // Reset to first page
  });

  test('should handle HOME_PAGE_LOADED', () => {
    const action = {
      type: HOME_PAGE_LOADED,
      tab: 'all',
      pager: () => {},
      payload: [
        { tags: ['react', 'javascript', 'testing'] },
        {
          articles: [
            { slug: 'article-1', title: 'Article 1' },
            { slug: 'article-2', title: 'Article 2' }
          ],
          articlesCount: 2
        }
      ]
    };

    const newState = articleListReducer({}, action);

    expect(newState.tags).toEqual(['react', 'javascript', 'testing']);
    expect(newState.articles.length).toBe(2);
    expect(newState.articlesCount).toBe(2);
    expect(newState.currentPage).toBe(0);
    expect(newState.tab).toBe('all');
  });

  test('should handle HOME_PAGE_UNLOADED', () => {
    const initialState = {
      articles: [{ slug: 'article-1' }],
      articlesCount: 1,
      tags: ['react'],
      currentPage: 2
    };

    const action = { type: HOME_PAGE_UNLOADED };

    const newState = articleListReducer(initialState, action);

    expect(newState).toEqual({});
  });

  test('should handle CHANGE_TAB', () => {
    const initialState = {
      articles: [],
      tab: 'all',
      tag: 'react',
      currentPage: 3
    };

    const newArticles = [
      { slug: 'feed-article', title: 'Feed Article' }
    ];

    const action = {
      type: CHANGE_TAB,
      tab: 'feed',
      pager: () => {},
      payload: {
        articles: newArticles,
        articlesCount: 5
      }
    };

    const newState = articleListReducer(initialState, action);

    expect(newState.articles).toEqual(newArticles);
    expect(newState.articlesCount).toBe(5);
    expect(newState.tab).toBe('feed');
    expect(newState.tag).toBe(null); // Tag cleared when changing tabs
    expect(newState.currentPage).toBe(0); // Reset to first page
  });

  test('should handle HOME_PAGE_LOADED with null payload', () => {
    const action = {
      type: HOME_PAGE_LOADED,
      tab: 'all',
      pager: () => {},
      payload: null
    };

    const newState = articleListReducer({}, action);

    expect(newState.tags).toEqual([]);
    expect(newState.articles).toEqual([]);
    expect(newState.articlesCount).toBe(0);
  });
});
