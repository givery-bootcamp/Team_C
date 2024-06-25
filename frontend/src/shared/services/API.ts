import { createAsyncThunk } from '@reduxjs/toolkit';
import { Hello } from 'shared/models';
import axios from 'axios';
import {
  AuthApi,
  Configuration,
  ModelPost,
  ModelUserSigninParam,
  PostApi,
} from 'api';
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

export const getPosts = createAsyncThunk<ModelPost[]>('getPosts', async () => {
  const api = new PostApi(configuration, API_ENDPOINT_PATH, axiosInstance);
  const getPostsResponse = await api.apiPostsGet(20, 0);
  return await getPostsResponse.data;
});

export const postSignin = createAsyncThunk<
  ModelUserSigninParam,
  ModelUserSigninParam
>('postSignin', async (userSigninParam: ModelUserSigninParam) => {
  const api = new AuthApi(configuration, API_ENDPOINT_PATH, axiosInstance);
  const postSignupResponse = await api.apiSigninPost(userSigninParam);

  return await postSignupResponse.data;
});
