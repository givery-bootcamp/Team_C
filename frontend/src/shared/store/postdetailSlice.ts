import { createSlice } from '@reduxjs/toolkit';
import { APIService } from '../services';
import { ModelPost } from 'api';

export type PostDetailState = {
  postdetail?: ModelPost;
};

export const initialState: PostDetailState = {};

export const postDetailSlice = createSlice({
  name: 'postdetail',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(APIService.getPostDetail.fulfilled, (state, action) => {
      state.postdetail = action.payload;
    });
  },
});

export default postDetailSlice.reducer;
