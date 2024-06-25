import { createAsyncThunk } from '@reduxjs/toolkit';
import { ApiClient } from 'api/ApiClient';
import { Hello } from 'shared/models';
import { model_Post } from 'api/models/model_Post';
import { model_UserSigninParam } from 'api/models/model_UserSigninParam';
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

export const postSignin = createAsyncThunk<
  model_UserSigninParam,
  model_UserSigninParam
>('postSignin', async (userSigninParam: model_UserSigninParam) => {
  const postSignupResponse = await api.auth.postApiSignin({
    body: userSigninParam,
  });

  return await postSignupResponse;
});
