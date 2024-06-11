import { useEffect } from 'react';

import { useAppDispatch, useAppSelector } from '../../shared/hooks';
import { APIService } from '../../shared/services';
import { Post } from '../../shared/models/post';
import './posts.scss';
export function Posts() {
  const { posts } = useAppSelector((state) => state.post);
  const dispatch = useAppDispatch();

  useEffect(() => {
    dispatch(APIService.getPosts());
  }, [dispatch]);

return (
    <div className="posts-container">
        {posts?.map((post: Post) => (
            <div key={post.id} className="post">
                <h2 className="post-title">{post.title}</h2>
                <p className="post-body">{post.body}</p>
            </div>
        ))}
    </div>
);
}
