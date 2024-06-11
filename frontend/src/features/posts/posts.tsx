import { useEffect } from 'react';

import { useAppDispatch, useAppSelector } from '../../shared/hooks';
import { APIService } from '../../shared/services';
import { Post } from '../../shared/models/post';
import { Box, Grid } from '@chakra-ui/react';
export function Posts() {
  const { posts } = useAppSelector((state) => state.post);
  const dispatch = useAppDispatch();

  useEffect(() => {
    dispatch(APIService.getPosts());
  }, [dispatch]);

  return (
    <div className="posts-container">
      <Grid templateColumns="repeat(3, 1fr)" gap={6}>
        {posts?.map((post: Post) => (
          <Box key={post.id} borderWidth="1px" borderRadius="lg" p={4}>
            <h2 className="post-title">{post.title}</h2>
            <p className="post-body">{post.body}</p>
          </Box>
        ))}
      </Grid>
    </div>
  );
}
