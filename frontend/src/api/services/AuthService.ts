/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { model_User } from '../models/model_User';
import type { model_UserSigninParam } from '../models/model_UserSigninParam';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class AuthService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * User signin
     * Signin
     * @returns model_User OK
     * @throws ApiError
     */
    public postApiSignin({
        body,
    }: {
        /**
         * リクエスト
         */
        body: model_UserSigninParam,
    }): CancelablePromise<model_User> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/api/signin',
            body: body,
        });
    }
    /**
     * user signout
     * signout
     * @returns any OK
     * @throws ApiError
     */
    public postApiSignout(): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/api/signout',
        });
    }
}
