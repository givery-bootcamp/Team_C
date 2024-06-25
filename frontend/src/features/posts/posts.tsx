import { useEffect, useState } from 'react';
import { useAppDispatch, useAppSelector } from 'shared/hooks';
import { APIService } from 'shared/services';
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
import { useQuery } from 'shared/hooks/usequery';
import { ModelPost } from 'api';
import { RootState } from 'shared/store';

function PostsHeader() {
  return (
    <Flex fontSize={'4xl'} p={3}>
      投稿一覧
    </Flex>
  );
}

export function Posts() {
  const { posts, status, error } = useAppSelector(
    (state: RootState) => state.post,
  );
  const dispatch = useAppDispatch();
  const query = useQuery();
  const toast = useToast();
  const [fetchParams, setFetchParams] = useState({ limit: 20, offset: 0 });

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
    const limit = parseInt(query.get('limit') ?? '20', 10);
    const offset = parseInt(query.get('offset') ?? '0', 10);
    setFetchParams({ limit, offset });
  }, [query]);

  useEffect(() => {
    dispatch(APIService.getPosts(fetchParams));
  }, [dispatch, fetchParams]);

  if (status === 'loading') {
    return <div>loading...</div>;
  }
  if (status === 'failed') {
    return <div>failed to fetch posts: {error}</div>;
  }

  return (
    <div className="posts-container">
      <PostsHeader />
      {posts?.map((post: ModelPost) => (
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
