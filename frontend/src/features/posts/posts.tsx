import { useEffect, useState } from 'react';
import { useAppDispatch, useAppSelector } from 'shared/hooks';
import { APIService } from 'shared/services';
import {
  Avatar,
  Box,
  Button,
  Card,
  CardBody,
  CardFooter,
  CardHeader,
  Flex,
  FormControl,
  FormLabel,
  Heading,
  Icon,
  IconButton,
  Input,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay,
  Textarea,
  useDisclosure,
  useToast,
  VStack,
} from '@chakra-ui/react';
import { useQuery } from 'shared/hooks/usequery';
import { ModelCreatePostParam } from 'api';
import { RootState } from 'shared/store';

export function Posts() {
  const { isOpen, onOpen, onClose } = useDisclosure();
  const { posts, status, error } = useAppSelector(
    (state: RootState) => state.post,
  );
  const dispatch = useAppDispatch();
  const query = useQuery();
  const toast = useToast();
  const [isInitialLoad, setIsInitialLoad] = useState(true);
  const [title, setTitle] = useState('');
  const [body, setBody] = useState('');

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

  const handleCreatePost = async () => {
    const postData: ModelCreatePostParam = { title, body };
    const resultAction = await dispatch(
      APIService.createPost({ param: postData }),
    );
    if (APIService.createPost.fulfilled.match(resultAction)) {
      onClose();
      setTitle('');
      setBody('');
      toast({
        title: 'Post created.',
        description: 'Your new post has been successfully created.',
        status: 'success',
        duration: 3000,
        isClosable: true,
      });
    } else {
      toast({
        title: 'Error',
        description: 'Failed to create post. Please try again.',
        status: 'error',
        duration: 3000,
        isClosable: true,
      });
    }
  };

  useEffect(() => {
    if (isInitialLoad) {
      const limit = parseInt(query.get('limit') ?? '20', 10);
      const offset = parseInt(query.get('offset') ?? '0', 10);

      dispatch(APIService.getPosts({ limit, offset }));
      setIsInitialLoad(false);
    }
  }, [dispatch, query, isInitialLoad]);

  if (status === 'loading' && isInitialLoad) {
    return <div>loading...</div>;
  }
  if (status === 'failed') {
    return <div>failed to fetch posts: {error}</div>;
  }

  return (
    <Box>
      <Button colorScheme="blue" onClick={onOpen} mb={4}>
        Create New Post
      </Button>

      <VStack spacing={4} align="stretch">
        {posts?.map((post) => (
          <Box key={post.id} p={4} borderWidth={1} borderRadius="md">
            <h2>{post.title}</h2>
            <p>{post.body}</p>
          </Box>
        ))}
      </VStack>

      <Modal isOpen={isOpen} onClose={onClose}>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>Create New Post</ModalHeader>
          <ModalCloseButton />
          <ModalBody>
            <VStack spacing={4}>
              <FormControl>
                <FormLabel>Title</FormLabel>
                <Input
                  value={title}
                  onChange={(e) => setTitle(e.target.value)}
                  placeholder="Enter post title"
                />
              </FormControl>
              <FormControl>
                <FormLabel>Content</FormLabel>
                <Textarea
                  value={body}
                  onChange={(e) => setBody(e.target.value)}
                  placeholder="Enter post content"
                />
              </FormControl>
            </VStack>
          </ModalBody>

          <ModalFooter>
            <Button
              colorScheme="blue"
              mr={3}
              onClick={handleCreatePost}
              isLoading={status === 'loading'}
            >
              Create
            </Button>
            <Button variant="ghost" onClick={onClose}>
              Cancel
            </Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </Box>
  );
}
