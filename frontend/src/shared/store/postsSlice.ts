import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { ModelPost } from 'api';
import { APIService } from '../services';

interface PostsState {
  posts: ModelPost[] | null;
  status: 'idle' | 'loading' | 'succeeded' | 'failed';
  error: string | null;
}
export const initialState: PostsState = {
  posts: null,
  status: 'idle',
  error: null,
};

export const postsSlice = createSlice({
  name: 'posts',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(APIService.getPosts.pending, (state) => {
        state.status = 'loading';
      })
      .addCase(
        APIService.getPosts.fulfilled,
        (state, action: PayloadAction<ModelPost[]>) => {
          state.posts = action.payload;
          state.status = 'succeeded';
          state.error = null;
        },
      )
      .addCase(APIService.getPosts.rejected, (state, action) => {
        state.status = 'failed';
        state.error = action.error.message ?? 'unknown error';
      })
      .addCase(APIService.createPost.pending, (state) => {
        state.status = 'loading';
      })
      .addCase(APIService.createPost.fulfilled, (state, action) => {
        state.posts?.unshift(action.payload);
        state.status = 'succeeded';
        state.error = null;
      })
      .addCase(APIService.createPost.rejected, (state, action) => {
        state.status = 'failed';
        state.error = action.error.message ?? 'failed to create post';
      });
  },
});

export const selectPosts = (state: { post: PostsState }) => state.post;
export default postsSlice.reducer;
