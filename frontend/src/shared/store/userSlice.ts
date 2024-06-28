import { createSlice } from "@reduxjs/toolkit";
import { ModelUser } from "api";
import { APIService } from "shared/services";

interface UserState {
    user: ModelUser | null;
}

export const initialState: UserState = {
    user: null,
}

export const userSlice = createSlice({
    name: 'user',
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder
            .addCase(APIService.getUser.fulfilled,  (state, action) => {
                state.user = action.payload;
            });
    }
})

export default userSlice.reducer;
