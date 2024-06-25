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
  useToast,
} from '@chakra-ui/react';
import { useNavigate } from 'react-router-dom'

function PostsHeader() {
  return (
    <Flex fontSize={'4xl'} p={3}>
      投稿一覧
    </Flex>
  );
}

export function Posts() {
  const navigate = useNavigate()
  const { posts } = useAppSelector((state) => state.post);
  const dispatch = useAppDispatch();
  const toast = useToast();

  useEffect(() => {
    const loginSuccess = localStorage.getItem('loginSuccess');
    if (loginSuccess == 'true') {
      toast({
        title: 'ログイン成功',
        description: 'ログインに成功しました。',
        status: 'success',
        duration: 3000,
        isClosable: true,
      });

      localStorage.removeItem('loginSuccess');
    }
  }, [toast]);

  useEffect(() => {
    dispatch(APIService.getPosts());
  }, [dispatch]);

  return (
    <div className="posts-container">
      <PostsHeader />
      {posts?.map((post: model_Post) => (
        <Card key={post.id} margin={2} onClick={() => {navigate(`${post.id}`)}} cursor={'pointer'} _hover={{ bg: "gray.100" }}>
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
