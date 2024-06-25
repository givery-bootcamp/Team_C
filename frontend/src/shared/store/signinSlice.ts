import { createSlice } from '@reduxjs/toolkit';
import { APIService } from '../services';
import { ModelUserSigninParam } from 'api';

interface SigninState {
  SignInForm: ModelUserSigninParam | null;
  status: 'idle' | 'loading' | 'succeeded' | 'failed';
  error: string | null;
}
export const initialState: SigninState = {
  SignInForm: null,
  status: 'idle',
  error: null,
};

export const signinSlice = createSlice({
  name: 'signin',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(APIService.postSignin.pending, (state) => {
        state.status = 'loading';
      })
      .addCase(APIService.postSignin.fulfilled, (state, action) => {
        state.SignInForm = action.payload;
        state.error = null;
        state.status = 'succeeded';
      })
      .addCase(APIService.postSignin.rejected, (state, action) => {
        state.status = 'failed';
        state.error = action.error.message ?? 'unknown error';
      });
  },
});

export default signinSlice.reducer;
