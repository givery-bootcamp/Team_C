import { createSlice } from '@reduxjs/toolkit';
import { APIService } from '../services';
import { model_UserSigninParam } from 'api';

export type SigninState = {
  signinParam?: model_UserSigninParam;
};

export const initialState: SigninState = {};

export const signinSlice = createSlice({
  name: 'signin',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(APIService.postSignin.fulfilled, (state, action) => {
      state.signinParam = action.payload;
    });
  },
});

export default signinSlice.reducer;
