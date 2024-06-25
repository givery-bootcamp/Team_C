import { createAsyncThunk } from '@reduxjs/toolkit';
import { Hello } from 'shared/models';
import axios, { AxiosError } from 'axios';
import {
  AuthApi,
  Configuration,
  ModelPost,
  ModelUserSigninParam,
  PostApi,
} from 'api';
import { RootState } from 'shared/store';
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
  ModelPost[],
  { limit: number; offset: number },
  {
    rejectValue: string;
    state: RootState;
  }
>('getPosts', async ({ limit, offset }, { rejectWithValue }) => {
  try {
    const api = new PostApi(configuration, API_ENDPOINT_PATH, axiosInstance);
    const getPostsResponse = await api.apiPostsGet(limit, offset);
    return getPostsResponse.data;
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
