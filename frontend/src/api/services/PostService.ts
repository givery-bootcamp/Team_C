/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { model_Post } from '../models/model_Post';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class PostService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * get posts
     * get posts
     * @returns model_Post OK
     * @throws ApiError
     */
    public getApiPosts({
        limit,
        offset,
    }: {
        /**
         * Limit
         */
        limit?: number,
        /**
         * Offset
         */
        offset?: number,
    }): CancelablePromise<Array<model_Post>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/api/posts',
            query: {
                'limit': limit,
                'offset': offset,
            },
        });
    }
    /**
     * get post by id
     * get post by id
     * @returns model_Post OK
     * @throws ApiError
     */
    public getApiPosts1({
        id,
    }: {
        /**
         * PostID
         */
        id: number,
    }): CancelablePromise<model_Post> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/api/posts/{id}',
            path: {
                'id': id,
            },
        });
    }
}
