import { useEffect } from 'react';

import { useAppDispatch, useAppSelector } from '../../shared/hooks';
import { APIService } from '../../shared/services';
import { Post } from '../../shared/models/post';

export function Posts() {
  const { posts } = useAppSelector((state) => state.post);
  const dispatch = useAppDispatch();

  useEffect(() => {
    dispatch(APIService.getPosts());
  }, [dispatch]);

  return (
    <div>
      {posts?.map((post: Post) => (
        <div key={post.id}>
          <h2>{post.title}</h2>
          <p>{post.body}</p>
        </div>
      ))}
    </div>
  );
}
