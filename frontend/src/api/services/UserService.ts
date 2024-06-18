/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { model_User } from '../models/model_User';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class UserService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * get login user
     * get login user
     * @returns model_User OK
     * @throws ApiError
     */
    public getApiUsers(): CancelablePromise<model_User> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/api/users',
        });
    }
}
