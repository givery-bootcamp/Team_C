import { createAsyncThunk } from '@reduxjs/toolkit';

import { Hello } from '../models';
import { Post } from '../models/post';

const API_ENDPOINT_PATH = import.meta.env.VITE_API_ENDPOINT_PATH ?? '';

export const getHello = createAsyncThunk<Hello>('getHello', async () => {
  const response = await fetch(`${API_ENDPOINT_PATH}/hello`);
  return await response.json();
});

export const getPosts = createAsyncThunk<Post[]>('getPosts', async() => {
  const response = await fetch(`${API_ENDPOINT_PATH}/api/posts/`);
  return await response.json()
})
