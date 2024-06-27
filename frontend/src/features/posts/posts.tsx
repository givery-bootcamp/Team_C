import {
  Avatar,
  Box,
  Button,
  Flex,
  FormControl,
  FormLabel,
  Heading,
  HStack,
  Input,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay,
  Spacer,
  Spinner,
  Text,
  Textarea,
  useColorModeValue,
  useDisclosure,
  useToast,
  VStack,
} from '@chakra-ui/react';
import { ModelCreatePostParam } from 'api';
import { useCallback, useEffect, useRef, useState } from 'react';
import { useAppDispatch, useAppSelector } from 'shared/hooks';
import { useQuery } from 'shared/hooks/usequery';
import { APIService } from 'shared/services';
import { RootState } from 'shared/store';

export function Posts() {
  const { isOpen, onOpen, onClose } = useDisclosure();
  const { posts, status, error, hasMore } = useAppSelector(
    (state: RootState) => state.post,
  );
  const dispatch = useAppDispatch();
  const query = useQuery();
  const toast = useToast();
  const [title, setTitle] = useState('');
  const [body, setBody] = useState('');
  const [offset, setOffset] = useState(0);
  const [isCreating, setIsCreating] = useState(false);

  const loadMoreRef = useRef<HTMLDivElement>(null);
  const scrollPositionRef = useRef(0);
  const bgColor = useColorModeValue('white', 'gray.800');
  const borderColor = useColorModeValue('gray.200', 'gray.700');

  const formatDate = (dateString?: string) => {
    if (!dateString) return '';
    const date = new Date(dateString);
    return date.toLocaleDateString('ja-JP', {
      year: 'numeric',
      month: 'long',
      day: 'numeric',
    });
  };

  const handleScroll = useCallback(() => {
    scrollPositionRef.current = window.scrollY;
  }, []);

  useEffect(() => {
    window.addEventListener('scroll', handleScroll, { passive: true });
    return () => window.removeEventListener('scroll', handleScroll);
  }, [handleScroll]);

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
    setIsCreating(true);
    const postData: ModelCreatePostParam = { title, body };
    try {
      const resultAction = await dispatch(
        APIService.createPost({ param: postData }),
      );
      if (APIService.createPost.fulfilled.match(resultAction)) {
        onClose();
        setTitle('');
        setBody('');
        toast({
          title: '投稿成功',
          description: '新しい投稿が作成されました。',
          status: 'success',
          duration: 3000,
          isClosable: true,
        });
        // 投稿リストを再読み込み
        dispatch(APIService.getPosts({ limit: 20, offset: 0 }));
      } else {
        throw new Error('投稿の作成に失敗しました');
      }
    } catch (error) {
      onClose();
      toast({
        title: 'エラー',
        description: '投稿できませんでした。もう一度お試しください。',
        status: 'error',
        duration: 3000,
        isClosable: true,
      });
      dispatch(APIService.getPosts({ limit: 20, offset: 0 }));
    } finally {
      setIsCreating(false);
    }
  };
  const limit = parseInt(query.get('limit') ?? '20', 10);

  useEffect(() => {
    dispatch(APIService.getPosts({ limit, offset: 0 }));
  }, [dispatch, limit]);

  useEffect(() => {
    const observer = new IntersectionObserver(
      (entries) => {
        if (entries[0].isIntersecting && hasMore && status !== 'loading') {
          setOffset((prevOffset) => prevOffset + limit);
        }
      },
      { threshold: 0.3 },
    );

    if (loadMoreRef.current) observer.observe(loadMoreRef.current);

    return () => observer.disconnect();
  }, [hasMore, limit, status]);

  useEffect(() => {
    if (offset > 0) {
      dispatch(APIService.getPosts({ limit, offset }));
    }
  }, [dispatch, limit, offset]);

  useEffect(() => {
    if (status === 'succeeded' && offset > 0) {
      window.scrollTo({
        top: scrollPositionRef.current,
        behavior: 'auto',
      });
    }
  }, [status, offset, posts]);

  if (status === 'failed') {
    return <div>failed to fetch posts: {error}</div>;
  }
  const Loading = () => {
    if (status === 'loading') {
      return <Spinner />;
    }
  };

  return (
    <Box>
      <Button colorScheme="blue" onClick={onOpen} mb={4}>
        新規投稿を作成
      </Button>
      {status === 'succeeded' && (
        <Box>
          <VStack spacing={6} align="stretch" width="100%">
            {posts?.map((post) => (
              <Box
                key={post.id}
                p={5}
                shadow="md"
                borderWidth={1}
                borderRadius="lg"
                bg={bgColor}
                borderColor={borderColor}
                _hover={{ shadow: 'lg' }}
                transition="all 0.3s"
              >
                <Flex align="center" mb={4}>
                  <Avatar size="sm" name={post.user?.name} mr={2} />
                  <Text fontWeight="bold">{post.user?.name}</Text>
                  <Spacer />
                  <Text fontSize="sm" color="gray.500">
                    {formatDate(post.created_at)}
                  </Text>
                </Flex>

                <Heading as="h3" size="md" mb={2}>
                  {post.title}
                </Heading>

                <Text noOfLines={3} mb={4}>
                  {post.body}
                </Text>

                <HStack spacing={4} fontSize="sm" color="gray.500">
                  <Text>作成日: {formatDate(post.created_at)}</Text>
                  <Text>更新日: {formatDate(post.updated_at)}</Text>
                </HStack>
              </Box>
            ))}
          </VStack>
          <Box ref={loadMoreRef} h="20px" mt={4}>
            {Loading()}
          </Box>
        </Box>
      )}

      <Modal isOpen={isOpen} onClose={onClose}>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>投稿する</ModalHeader>
          <ModalCloseButton />
          <ModalBody>
            <VStack spacing={4}>
              <FormControl>
                <FormLabel>タイトル</FormLabel>
                <Input
                  value={title}
                  onChange={(e) => setTitle(e.target.value)}
                  placeholder="タイトルは？？"
                />
              </FormControl>
              <FormControl>
                <FormLabel>Content</FormLabel>
                <Textarea
                  value={body}
                  onChange={(e) => setBody(e.target.value)}
                  placeholder="中身ないような内容を書くな！"
                />
              </FormControl>
            </VStack>
          </ModalBody>

          <ModalFooter>
            <Button
              colorScheme="blue"
              mr={3}
              onClick={handleCreatePost}
              isLoading={isCreating}
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
