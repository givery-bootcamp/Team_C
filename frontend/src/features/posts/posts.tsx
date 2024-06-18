import { useEffect } from 'react';

import { useAppDispatch, useAppSelector } from '../../shared/hooks';
import { APIService } from '../../shared/services';
import { Post } from '../../shared/models/post';
import {
  Avatar,
  Box,
  Card,
  CardBody,
  CardFooter,
  CardHeader,
  Flex,
  Heading,
  Icon,
  IconButton,
} from '@chakra-ui/react';

function PostsHeader() {
  return (
    <Flex fontSize={'4xl'} p={3}>
      投稿一覧
    </Flex>
  )
}

export function Posts() {
  const { posts } = useAppSelector((state) => state.post);
  const dispatch = useAppDispatch();

  useEffect(() => {
    dispatch(APIService.getPosts());
  }, [dispatch]);

  return (
    <div className="posts-container">
      <PostsHeader />
      {posts?.map((post: Post) => (
        <Card key={post.id} margin={2}>
          <CardHeader>
            <Flex>
              <Flex flex="1" gap="4" alignItems="center" flexWrap="wrap">
                <Avatar size="sm" name={post.user?.name} />
                <Box>
                  <Heading size="sm">{post.user?.name}</Heading>
                  <text>@{post.user?.id}</text>
                </Box>
              </Flex>
              <IconButton
                variant={'ghost'}
                colorScheme="gray"
                aria-label="icon"
                icon={<Icon />}
              />
            </Flex>
          </CardHeader>
          <CardBody>
            <Heading size="md">{post.title}</Heading>
            <text>{post.body}</text>
          </CardBody>
          <CardFooter>
            {post.created_at} - {post.updated_at}
          </CardFooter>
        </Card>
      ))}
    </div>
  );
}
