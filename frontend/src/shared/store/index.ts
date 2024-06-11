import { configureStore } from '@reduxjs/toolkit';

import helloReducer, { helloSlice } from './HelloSlice';
import postReducer, { postsSlice } from './postsSlice';
export const store = configureStore({
  reducer: {
    hello: helloReducer,
    post: postReducer,
  },
});

export const actions = {
  ...helloSlice.actions,
  ...postsSlice.actions,
};

export type RootState = ReturnType<typeof store.getState>;

export type AppDispatch = typeof store.dispatch;
