import { createAsyncThunk } from '@reduxjs/toolkit';
import {
  AuthApi,
  Configuration,
  ModelPost,
  ModelUserSigninParam,
  PostApi,
} from 'api';
import axios, { AxiosError } from 'axios';
import { Hello } from 'shared/models';
import { RootState } from 'shared/store';
import { ModelCreatePostParam } from '../../api/api';
const API_ENDPOINT_PATH = import.meta.env.VITE_API_ENDPOINT_PATH ?? '';

const configuration = new Configuration({
  basePath: API_ENDPOINT_PATH,
});
const axiosInstance = axios.create({
  baseURL: API_ENDPOINT_PATH,
  withCredentials: true,
});

export const getHello = createAsyncThunk<Hello>('getHello', async () => {
  const response = await fetch(`${API_ENDPOINT_PATH}/hello`);
  return await response.json();
});

export const getPosts = createAsyncThunk<
  { posts: ModelPost[]; hasMore: boolean; offset: number },
  { limit: number; offset: number },
  {
    rejectValue: string;
    state: RootState;
  }
>('getPosts', async ({ limit, offset }, { rejectWithValue }) => {
  try {
    const api = new PostApi(configuration, API_ENDPOINT_PATH, axiosInstance);
    const getPostsResponse = await api.apiPostsGet(limit, offset);
    return {
      posts: getPostsResponse.data,
      hasMore: getPostsResponse.data.length === limit,
      offset: offset,
    };
  } catch (error) {
    const err = error as AxiosError;
    return rejectWithValue(err.message ?? 'failed to fetch posts');
  }
});

export const postSignin = createAsyncThunk<
  ModelUserSigninParam,
  { param: ModelUserSigninParam },
  {
    rejectValue: string;
    state: RootState;
  }
>('postSignin', async (param, { rejectWithValue }) => {
  try {
    const api = new AuthApi(configuration, API_ENDPOINT_PATH, axiosInstance);
    const postSigninResponse = await api.apiSigninPost(param.param);
    return postSigninResponse.data;
  } catch (error) {
    const err = error as AxiosError;
    return rejectWithValue(err.message ?? 'failed to fetch posts');
  }
});

export const createPost = createAsyncThunk<
  ModelPost,
  { param: ModelCreatePostParam },
  {
    rejectValue: string;
    state: RootState;
  }
>('createPost', async (param, { rejectWithValue }) => {
  try {
    const api = new PostApi(configuration, API_ENDPOINT_PATH, axiosInstance);
    const postSigninResponse = await api.apiPostsPost(param.param);
    return postSigninResponse.data;
  } catch (error) {
    const err = error as AxiosError;
    return rejectWithValue(err.message ?? 'failed to fetch posts');
  }
});

export const getPostDetail = createAsyncThunk<
  ModelPost,
  { id: number },
  {
    rejectValue: string;
    state: RootState;
  }
>('getPostDetail', async ({ id }, { rejectWithValue }) => {
  try {
    const api = new PostApi(configuration, API_ENDPOINT_PATH, axiosInstance);
    const getPostDetailResponse = await api.apiPostsIdGet(id);
    return getPostDetailResponse.data;
  } catch (error) {
    const err = error as AxiosError;
    return rejectWithValue(err.message ?? 'failed to fetch postdetail');
  }
});

export const deletePost = createAsyncThunk(
  'deletePost',
  async (id: number, { rejectWithValue }) => {
    try {
      const api = new PostApi(configuration, API_ENDPOINT_PATH, axiosInstance);
      const deletePostResponse = await api.apiPostsIdDelete(id);
      return deletePostResponse.data;
    } catch (error) {
      return rejectWithValue((error as Error).message);
    }
  },
);
