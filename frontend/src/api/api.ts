/* tslint:disable */
/* eslint-disable */
/**
 * 掲示板アプリ
 * 3班の掲示板アプリのAPI仕様書
 *
 * The version of the OpenAPI document: バージョン(1.0)
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import type { Configuration } from './configuration';
import type { AxiosPromise, AxiosInstance, RawAxiosRequestConfig } from 'axios';
import globalAxios from 'axios';
// Some imports not used depending on template conditions
// @ts-ignore
import { DUMMY_BASE_URL, assertParamExists, setApiKeyToObject, setBasicAuthToObject, setBearerAuthToObject, setOAuthToObject, setSearchParams, serializeDataIfNeeded, toPathString, createRequestFunction } from './common';
import type { RequestArgs } from './base';
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, BaseAPI, RequiredError, operationServerMap } from './base';

/**
 * 
 * @export
 * @interface ModelCreatePostParam
 */
export interface ModelCreatePostParam {
    /**
     * 
     * @type {string}
     * @memberof ModelCreatePostParam
     */
    'body'?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelCreatePostParam
     */
    'title'?: string;
}
/**
 * 
 * @export
 * @interface ModelCreateUserParam
 */
export interface ModelCreateUserParam {
    /**
     * 
     * @type {string}
     * @memberof ModelCreateUserParam
     */
    'name'?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelCreateUserParam
     */
    'password'?: string;
}
/**
 * 
 * @export
 * @interface ModelHelloWorld
 */
export interface ModelHelloWorld {
    /**
     * 
     * @type {string}
     * @memberof ModelHelloWorld
     */
    'lang'?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelHelloWorld
     */
    'message'?: string;
}
/**
 * 
 * @export
 * @interface ModelPost
 */
export interface ModelPost {
    /**
     * 
     * @type {string}
     * @memberof ModelPost
     */
    'body'?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelPost
     */
    'created_at'?: string;
    /**
     * 
     * @type {number}
     * @memberof ModelPost
     */
    'id'?: number;
    /**
     * 
     * @type {string}
     * @memberof ModelPost
     */
    'title'?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelPost
     */
    'updated_at'?: string;
    /**
     * 
     * @type {ModelUser}
     * @memberof ModelPost
     */
    'user'?: ModelUser;
}
/**
 * 
 * @export
 * @interface ModelUpdatePostParam
 */
export interface ModelUpdatePostParam {
    /**
     * 
     * @type {string}
     * @memberof ModelUpdatePostParam
     */
    'body'?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelUpdatePostParam
     */
    'title'?: string;
}
/**
 * 
 * @export
 * @interface ModelUser
 */
export interface ModelUser {
    /**
     * 
     * @type {number}
     * @memberof ModelUser
     */
    'id'?: number;
    /**
     * 
     * @type {string}
     * @memberof ModelUser
     */
    'name'?: string;
}
/**
 * 
 * @export
 * @interface ModelUserSigninParam
 */
export interface ModelUserSigninParam {
    /**
     * 
     * @type {string}
     * @memberof ModelUserSigninParam
     */
    'name'?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelUserSigninParam
     */
    'password'?: string;
}

/**
 * AuthApi - axios parameter creator
 * @export
 */
export const AuthApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * Signin
         * @summary User signin
         * @param {ModelUserSigninParam} body リクエスト
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        apiSigninPost: async (body: ModelUserSigninParam, options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'body' is not null or undefined
            assertParamExists('apiSigninPost', 'body', body)
            const localVarPath = `/api/signin`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(body, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * signout
         * @summary user signout
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        apiSignoutPost: async (options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/signout`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * Create User
         * @summary Signup User
         * @param {ModelCreateUserParam} body リクエスト
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        apiSignupPost: async (body: ModelCreateUserParam, options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'body' is not null or undefined
            assertParamExists('apiSignupPost', 'body', body)
            const localVarPath = `/api/signup`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(body, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * AuthApi - functional programming interface
 * @export
 */
export const AuthApiFp = function(configuration?: Configuration) {
    const localVarAxiosParamCreator = AuthApiAxiosParamCreator(configuration)
    return {
        /**
         * Signin
         * @summary User signin
         * @param {ModelUserSigninParam} body リクエスト
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async apiSigninPost(body: ModelUserSigninParam, options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<ModelUser>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.apiSigninPost(body, options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['AuthApi.apiSigninPost']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
        /**
         * signout
         * @summary user signout
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async apiSignoutPost(options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<object>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.apiSignoutPost(options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['AuthApi.apiSignoutPost']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
        /**
         * Create User
         * @summary Signup User
         * @param {ModelCreateUserParam} body リクエスト
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async apiSignupPost(body: ModelCreateUserParam, options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<ModelUser>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.apiSignupPost(body, options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['AuthApi.apiSignupPost']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
    }
};

/**
 * AuthApi - factory interface
 * @export
 */
export const AuthApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    const localVarFp = AuthApiFp(configuration)
    return {
        /**
         * Signin
         * @summary User signin
         * @param {ModelUserSigninParam} body リクエスト
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        apiSigninPost(body: ModelUserSigninParam, options?: any): AxiosPromise<ModelUser> {
            return localVarFp.apiSigninPost(body, options).then((request) => request(axios, basePath));
        },
        /**
         * signout
         * @summary user signout
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        apiSignoutPost(options?: any): AxiosPromise<object> {
            return localVarFp.apiSignoutPost(options).then((request) => request(axios, basePath));
        },
        /**
         * Create User
         * @summary Signup User
         * @param {ModelCreateUserParam} body リクエスト
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        apiSignupPost(body: ModelCreateUserParam, options?: any): AxiosPromise<ModelUser> {
            return localVarFp.apiSignupPost(body, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * AuthApi - object-oriented interface
 * @export
 * @class AuthApi
 * @extends {BaseAPI}
 */
export class AuthApi extends BaseAPI {
    /**
     * Signin
     * @summary User signin
     * @param {ModelUserSigninParam} body リクエスト
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof AuthApi
     */
    public apiSigninPost(body: ModelUserSigninParam, options?: RawAxiosRequestConfig) {
        return AuthApiFp(this.configuration).apiSigninPost(body, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * signout
     * @summary user signout
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof AuthApi
     */
    public apiSignoutPost(options?: RawAxiosRequestConfig) {
        return AuthApiFp(this.configuration).apiSignoutPost(options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * Create User
     * @summary Signup User
     * @param {ModelCreateUserParam} body リクエスト
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof AuthApi
     */
    public apiSignupPost(body: ModelCreateUserParam, options?: RawAxiosRequestConfig) {
        return AuthApiFp(this.configuration).apiSignupPost(body, options).then((request) => request(this.axios, this.basePath));
    }
}



/**
 * HelloWorldApi - axios parameter creator
 * @export
 */
export const HelloWorldApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * hello world
         * @summary hello world
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        helloGet: async (options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/hello`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * HelloWorldApi - functional programming interface
 * @export
 */
export const HelloWorldApiFp = function(configuration?: Configuration) {
    const localVarAxiosParamCreator = HelloWorldApiAxiosParamCreator(configuration)
    return {
        /**
         * hello world
         * @summary hello world
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async helloGet(options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<ModelHelloWorld>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.helloGet(options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['HelloWorldApi.helloGet']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
    }
};

/**
 * HelloWorldApi - factory interface
 * @export
 */
export const HelloWorldApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    const localVarFp = HelloWorldApiFp(configuration)
    return {
        /**
         * hello world
         * @summary hello world
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        helloGet(options?: any): AxiosPromise<ModelHelloWorld> {
            return localVarFp.helloGet(options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * HelloWorldApi - object-oriented interface
 * @export
 * @class HelloWorldApi
 * @extends {BaseAPI}
 */
export class HelloWorldApi extends BaseAPI {
    /**
     * hello world
     * @summary hello world
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof HelloWorldApi
     */
    public helloGet(options?: RawAxiosRequestConfig) {
        return HelloWorldApiFp(this.configuration).helloGet(options).then((request) => request(this.axios, this.basePath));
    }
}



/**
 * PostApi - axios parameter creator
 * @export
 */
export const PostApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * get posts
         * @summary get posts
         * @param {number} [limit] Limit
         * @param {number} [offset] Offset
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        apiPostsGet: async (limit?: number, offset?: number, options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/posts`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            if (limit !== undefined) {
                localVarQueryParameter['limit'] = limit;
            }

            if (offset !== undefined) {
                localVarQueryParameter['offset'] = offset;
            }


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * delete post
         * @summary delete post
         * @param {number} id PostID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        apiPostsIdDelete: async (id: number, options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            assertParamExists('apiPostsIdDelete', 'id', id)
            const localVarPath = `/api/posts/{id}`
                .replace(`{${"id"}}`, encodeURIComponent(String(id)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'DELETE', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * get post by id
         * @summary get post by id
         * @param {number} id PostID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        apiPostsIdGet: async (id: number, options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            assertParamExists('apiPostsIdGet', 'id', id)
            const localVarPath = `/api/posts/{id}`
                .replace(`{${"id"}}`, encodeURIComponent(String(id)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * update post
         * @summary update post
         * @param {number} id PostID
         * @param {ModelUpdatePostParam} body リクエスト
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        apiPostsIdPut: async (id: number, body: ModelUpdatePostParam, options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            assertParamExists('apiPostsIdPut', 'id', id)
            // verify required parameter 'body' is not null or undefined
            assertParamExists('apiPostsIdPut', 'body', body)
            const localVarPath = `/api/posts/{id}`
                .replace(`{${"id"}}`, encodeURIComponent(String(id)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'PUT', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(body, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * create post
         * @summary create post
         * @param {ModelCreatePostParam} body リクエスト
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        apiPostsPost: async (body: ModelCreatePostParam, options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'body' is not null or undefined
            assertParamExists('apiPostsPost', 'body', body)
            const localVarPath = `/api/posts`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(body, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * PostApi - functional programming interface
 * @export
 */
export const PostApiFp = function(configuration?: Configuration) {
    const localVarAxiosParamCreator = PostApiAxiosParamCreator(configuration)
    return {
        /**
         * get posts
         * @summary get posts
         * @param {number} [limit] Limit
         * @param {number} [offset] Offset
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async apiPostsGet(limit?: number, offset?: number, options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelPost>>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.apiPostsGet(limit, offset, options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['PostApi.apiPostsGet']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
        /**
         * delete post
         * @summary delete post
         * @param {number} id PostID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async apiPostsIdDelete(id: number, options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<void>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.apiPostsIdDelete(id, options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['PostApi.apiPostsIdDelete']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
        /**
         * get post by id
         * @summary get post by id
         * @param {number} id PostID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async apiPostsIdGet(id: number, options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<ModelPost>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.apiPostsIdGet(id, options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['PostApi.apiPostsIdGet']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
        /**
         * update post
         * @summary update post
         * @param {number} id PostID
         * @param {ModelUpdatePostParam} body リクエスト
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async apiPostsIdPut(id: number, body: ModelUpdatePostParam, options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<ModelPost>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.apiPostsIdPut(id, body, options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['PostApi.apiPostsIdPut']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
        /**
         * create post
         * @summary create post
         * @param {ModelCreatePostParam} body リクエスト
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async apiPostsPost(body: ModelCreatePostParam, options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<ModelPost>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.apiPostsPost(body, options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['PostApi.apiPostsPost']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
    }
};

/**
 * PostApi - factory interface
 * @export
 */
export const PostApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    const localVarFp = PostApiFp(configuration)
    return {
        /**
         * get posts
         * @summary get posts
         * @param {number} [limit] Limit
         * @param {number} [offset] Offset
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        apiPostsGet(limit?: number, offset?: number, options?: any): AxiosPromise<Array<ModelPost>> {
            return localVarFp.apiPostsGet(limit, offset, options).then((request) => request(axios, basePath));
        },
        /**
         * delete post
         * @summary delete post
         * @param {number} id PostID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        apiPostsIdDelete(id: number, options?: any): AxiosPromise<void> {
            return localVarFp.apiPostsIdDelete(id, options).then((request) => request(axios, basePath));
        },
        /**
         * get post by id
         * @summary get post by id
         * @param {number} id PostID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        apiPostsIdGet(id: number, options?: any): AxiosPromise<ModelPost> {
            return localVarFp.apiPostsIdGet(id, options).then((request) => request(axios, basePath));
        },
        /**
         * update post
         * @summary update post
         * @param {number} id PostID
         * @param {ModelUpdatePostParam} body リクエスト
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        apiPostsIdPut(id: number, body: ModelUpdatePostParam, options?: any): AxiosPromise<ModelPost> {
            return localVarFp.apiPostsIdPut(id, body, options).then((request) => request(axios, basePath));
        },
        /**
         * create post
         * @summary create post
         * @param {ModelCreatePostParam} body リクエスト
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        apiPostsPost(body: ModelCreatePostParam, options?: any): AxiosPromise<ModelPost> {
            return localVarFp.apiPostsPost(body, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * PostApi - object-oriented interface
 * @export
 * @class PostApi
 * @extends {BaseAPI}
 */
export class PostApi extends BaseAPI {
    /**
     * get posts
     * @summary get posts
     * @param {number} [limit] Limit
     * @param {number} [offset] Offset
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PostApi
     */
    public apiPostsGet(limit?: number, offset?: number, options?: RawAxiosRequestConfig) {
        return PostApiFp(this.configuration).apiPostsGet(limit, offset, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * delete post
     * @summary delete post
     * @param {number} id PostID
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PostApi
     */
    public apiPostsIdDelete(id: number, options?: RawAxiosRequestConfig) {
        return PostApiFp(this.configuration).apiPostsIdDelete(id, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * get post by id
     * @summary get post by id
     * @param {number} id PostID
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PostApi
     */
    public apiPostsIdGet(id: number, options?: RawAxiosRequestConfig) {
        return PostApiFp(this.configuration).apiPostsIdGet(id, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * update post
     * @summary update post
     * @param {number} id PostID
     * @param {ModelUpdatePostParam} body リクエスト
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PostApi
     */
    public apiPostsIdPut(id: number, body: ModelUpdatePostParam, options?: RawAxiosRequestConfig) {
        return PostApiFp(this.configuration).apiPostsIdPut(id, body, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * create post
     * @summary create post
     * @param {ModelCreatePostParam} body リクエスト
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PostApi
     */
    public apiPostsPost(body: ModelCreatePostParam, options?: RawAxiosRequestConfig) {
        return PostApiFp(this.configuration).apiPostsPost(body, options).then((request) => request(this.axios, this.basePath));
    }
}



/**
 * UserApi - axios parameter creator
 * @export
 */
export const UserApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * get login user
         * @summary get login user
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        apiUsersGet: async (options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/users`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * UserApi - functional programming interface
 * @export
 */
export const UserApiFp = function(configuration?: Configuration) {
    const localVarAxiosParamCreator = UserApiAxiosParamCreator(configuration)
    return {
        /**
         * get login user
         * @summary get login user
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async apiUsersGet(options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<ModelUser>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.apiUsersGet(options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['UserApi.apiUsersGet']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
    }
};

/**
 * UserApi - factory interface
 * @export
 */
export const UserApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    const localVarFp = UserApiFp(configuration)
    return {
        /**
         * get login user
         * @summary get login user
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        apiUsersGet(options?: any): AxiosPromise<ModelUser> {
            return localVarFp.apiUsersGet(options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * UserApi - object-oriented interface
 * @export
 * @class UserApi
 * @extends {BaseAPI}
 */
export class UserApi extends BaseAPI {
    /**
     * get login user
     * @summary get login user
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof UserApi
     */
    public apiUsersGet(options?: RawAxiosRequestConfig) {
        return UserApiFp(this.configuration).apiUsersGet(options).then((request) => request(this.axios, this.basePath));
    }
}



