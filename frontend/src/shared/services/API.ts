import { createAsyncThunk } from '@reduxjs/toolkit';
import { ApiClient } from 'api/ApiClient'; // Replace 'path/to/ApiClient' with the actual path to the ApiClient file

import { Hello } from 'shared/models';
import { model_Post } from 'api/models/model_Post';

const API_ENDPOINT_PATH = import.meta.env.VITE_API_ENDPOINT_PATH ?? '';
const api = new ApiClient({ BASE: API_ENDPOINT_PATH })

export const getHello = createAsyncThunk<Hello>('getHello', async () => {
  const response = await fetch(`${API_ENDPOINT_PATH}/hello`);
  return await response.json();
});

export const getPosts = createAsyncThunk<model_Post[]>('getPosts', async () => {
  const getPostsResponse = await api.post.getApiPosts({ limit: 20, offset: 0 });
  return await getPostsResponse;
});

export const getPostDetail = createAsyncThunk<model_Post, number>('getPostDetail', async (post_id: number) => {
  const getPostDetailResponse = await api.post.getApiPosts1({id: post_id})
  return await getPostDetailResponse
})
