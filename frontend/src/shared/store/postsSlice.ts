import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { APIService } from '../services';
import { ModelPost } from 'api';

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
        },
      )
      .addCase(APIService.getPosts.rejected, (state, action) => {
        state.status = 'failed';
        state.error = action.error.message ?? 'unknown error';
      });
  },
});

export default postsSlice.reducer;
