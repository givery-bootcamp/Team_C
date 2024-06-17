import { useEffect } from 'react';

import { useAppDispatch, useAppSelector } from 'shared/hooks';
import { APIService } from 'shared/services';
import { model_Post } from 'api';
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

export function Posts() {
  const { posts } = useAppSelector((state) => state.post);
  const dispatch = useAppDispatch();

  useEffect(() => {
    dispatch(APIService.getPosts());
  }, [dispatch]);

  return (
    <div className="posts-container">
      {posts?.map((post: model_Post) => (
        <Card key={post.id}>
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
