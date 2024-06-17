import { createSlice } from '@reduxjs/toolkit';
import { APIService } from '../services';
import { model_Post } from 'api';

export type PostsState = {
  posts?: model_Post[];
};

export const initialState: PostsState = {};

export const postsSlice = createSlice({
  name: 'posts',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(APIService.getPosts.fulfilled, (state, action) => {
      state.posts = action.payload;
    });
  },
});

export default postsSlice.reducer;
