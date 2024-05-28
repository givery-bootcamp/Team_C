import { user } from "./user";

export type post = {
    id: number;
    title: string;
    body: string;
    user: user;
    created_at: string;
    updated_at: string;
}