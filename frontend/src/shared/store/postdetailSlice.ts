import { createSlice } from '@reduxjs/toolkit';
import { ModelPost } from 'api';
import { APIService } from '../services';

export type PostDetailState = {
  postdetail?: ModelPost;
  status: 'idle' | 'loading' | 'succeeded' | 'failed';
  error: string | null;
};

export const initialState: PostDetailState = {
  postdetail: undefined,
  status: 'idle',
  error: null,
};

export const postDetailSlice = createSlice({
  name: 'postdetail',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(APIService.getPostDetail.fulfilled, (state, action) => {
        state.postdetail = action.payload;
      })
      .addCase(APIService.getPostDetail.pending, (state) => {
        state.status = 'loading';
      })
      .addCase(APIService.getPostDetail.rejected, (state, action) => {
        state.status = 'failed';
        state.error = action.error.message ?? 'unknown error';
      })
      .addCase(APIService.deletePost.fulfilled, (state) => {
        state.status = 'succeeded';
        state.postdetail = undefined;
      })
      .addCase(APIService.deletePost.rejected, (state, action) => {
        state.status = 'failed';
        state.error = action.error.message ?? 'unknown error';
      })
      .addCase(APIService.editPost.fulfilled, (state, action) => {
        state.status = 'succeeded';
        state.postdetail = action.payload;
      })
      .addCase(APIService.editPost.rejected, (state, action) => {
        state.status = 'failed';
        state.error = action.error.message ?? 'unknown error';
      });
  },
});

export default postDetailSlice.reducer;
