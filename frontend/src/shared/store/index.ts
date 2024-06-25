import { configureStore } from '@reduxjs/toolkit';

import helloReducer, { helloSlice } from './HelloSlice';
import postReducer, { postsSlice } from './postsSlice';
import postDetailReducer, { postDetailSlice } from './postdetailSlice'
import signinReducer, { signinSlice } from './signinSlice';
export const store = configureStore({
  reducer: {
    hello: helloReducer,
    post: postReducer,
    postdetail: postDetailReducer,
    signin: signinReducer,
  },
});

export const actions = {
  ...helloSlice.actions,
  ...postsSlice.actions,
  ...postDetailSlice.actions,
  ...signinSlice.actions,
};

export type RootState = ReturnType<typeof store.getState>;

export type AppDispatch = typeof store.dispatch;
