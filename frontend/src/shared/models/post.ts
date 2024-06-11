import { User } from "./user";

export type Post = {
    id: number;
    title: string;
    body: string;
    user: User;
    created_at: string;
    updated_at: string;
}