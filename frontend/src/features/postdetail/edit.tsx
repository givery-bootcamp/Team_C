import {
  Box,
  Button,
  FormControl,
  FormLabel,
  Input,
  Text,
  Textarea,
  VStack,
  useToast,
} from '@chakra-ui/react';
import React, { useEffect, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { useAppDispatch, useAppSelector } from 'shared/hooks';
import { APIService } from 'shared/services';
import { RootState } from 'shared/store';

export function EditPost() {
  const { postId } = useParams<{ postId: string }>();
  const navigate = useNavigate();
  const dispatch = useAppDispatch();
  const toast = useToast();
  const { postdetail, status, error } = useAppSelector(
    (state: RootState) => state.postdetail,
  );

  const [title, setTitle] = useState('');
  const [body, setBody] = useState('');
  const [isSubmitting, setIsSubmitting] = useState(false);

  useEffect(() => {
    if (postId) {
      dispatch(APIService.getPostDetail({ id: parseInt(postId, 10) }));
    }
  }, [dispatch, postId]);

  useEffect(() => {
    if (postdetail) {
      setTitle(postdetail.title || '');
      setBody(postdetail.body || '');
    }
  }, [postdetail]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (postId) {
      setIsSubmitting(true);
      try {
        await dispatch(
          APIService.editPost({
            id: parseInt(postId, 10),
            param: { title, body },
          }),
        ).unwrap();

        toast({
          title: '投稿が更新されました',
          status: 'success',
          duration: 3000,
          isClosable: true,
        });

        navigate(`/posts/${postId}`);
      } catch (error) {
        toast({
          title: '更新に失敗しました',
          description: error as string,
          status: 'error',
          duration: 3000,
          isClosable: true,
        });
      } finally {
        setIsSubmitting(false);
      }
    }
  };

  if (status === 'failed') {
    return (
      <Box>
        <Text color="red.500">Failed to load post: {error}</Text>
        <Button
          onClick={() =>
            dispatch(APIService.getPostDetail({ id: parseInt(postId!, 10) }))
          }
        >
          再読み込み
        </Button>
      </Box>
    );
  }

  return (
    <Box>
      <form onSubmit={handleSubmit}>
        <VStack spacing={4}>
          <FormControl isRequired>
            <FormLabel>タイトル</FormLabel>
            <Input
              value={title}
              onChange={(e) => setTitle(e.target.value)}
              placeholder="タイトル変えるの？"
            />
          </FormControl>
          <FormControl isRequired>
            <FormLabel>Content</FormLabel>
            <Textarea
              value={body}
              onChange={(e) => setBody(e.target.value)}
              placeholder="変なこと書くなよ！"
            />
          </FormControl>
          <Button
            type="submit"
            colorScheme="blue"
            isLoading={isSubmitting}
            loadingText="更新中"
          >
            更新
          </Button>
        </VStack>
      </form>
    </Box>
  );
}
