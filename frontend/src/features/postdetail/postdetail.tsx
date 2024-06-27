import {
  AlertDialog,
  AlertDialogBody,
  AlertDialogContent,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogOverlay,
  Avatar,
  Box,
  Button,
  ButtonGroup,
  Divider,
  HStack,
  Heading,
  Text,
  useToast,
} from '@chakra-ui/react';
import { useEffect, useRef, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { useAppDispatch, useAppSelector } from 'shared/hooks';
import { APIService, DateService } from 'shared/services';
import { RootState } from 'shared/store';

export function PostDetail() {
  const params = useParams();
  const { postdetail } = useAppSelector((state: RootState) => state.postdetail);
  const dispatch = useAppDispatch();
  const navigate = useNavigate();
  const toast = useToast();
  const postId = Number(params.postId);
  const [isDeleteDialogOpen, setIsDeleteDialogOpen] = useState(false);
  const [isDeleting, setIsDeleting] = useState(false);
  const [isSecondDialogOpen, setIsSecondDialogOpen] = useState(false);
  const [isThirdDialogOpen, setIsThirdDialogOpen] = useState(false);
  // AlertDialogのcancelRef用のuseRef
  const cancelRef = useRef<HTMLButtonElement>(null);

  useEffect(() => {
    dispatch(APIService.getPostDetail({ id: postId }));
  }, [dispatch, postId]);

  const handleDelete = async () => {
    setIsDeleting(true);
    try {
      const resultAction = await dispatch(
        APIService.deletePost({ id: postId }),
      );
      if (APIService.deletePost.fulfilled.match(resultAction)) {
        toast({
          title: '投稿が削除されました',
          status: 'success',
          duration: 3000,
          isClosable: true,
        });
        navigate('/posts'); // 投稿一覧ページへリダイレクト
      } else {
        throw new Error('投稿の削除に失敗しました');
      }
    } catch (error) {
      toast({
        title: 'エラー',
        description: '投稿の削除に失敗しました',
        status: 'error',
        duration: 3000,
        isClosable: true,
      });
    } finally {
      setIsDeleting(false);
      setIsDeleteDialogOpen(false);
    }
  };

  return (
    <Box>
      <Heading as={'h1'}>{postdetail?.title}</Heading>
      <HStack>
        <Avatar size="sm" name={postdetail?.user?.name} />
        <Text as={'span'}>{postdetail?.user?.name}</Text>
      </HStack>
      <Text paddingTop={5}>{postdetail?.body}</Text>
      <Divider />
      <HStack spacing={10} fontSize={'small'} color={'gray'}>
        <Text as={'span'}>
          作成日時 {DateService.formatDate(postdetail?.created_at)}
        </Text>
        <Text as={'span'}>
          投稿日時 {DateService.formatDate(postdetail?.updated_at)}
        </Text>
      </HStack>
      <ButtonGroup colorScheme="blue" paddingTop={5}>
        <Button>編集</Button>
        <Button onClick={() => setIsDeleteDialogOpen(true)} colorScheme="red">
          削除
        </Button>
      </ButtonGroup>

      <AlertDialog
        isOpen={isDeleteDialogOpen}
        leastDestructiveRef={cancelRef}
        onClose={() => setIsDeleteDialogOpen(false)}
      >
        <AlertDialogOverlay>
          <AlertDialogContent>
            <AlertDialogHeader fontSize="lg" fontWeight="bold">
              投稿を削除してやろうか？
            </AlertDialogHeader>

            <AlertDialogBody>
              削除していいの？本当に削除していいの？
            </AlertDialogBody>

            <AlertDialogFooter>
              <Button
                ref={cancelRef}
                onClick={() => setIsDeleteDialogOpen(false)}
              >
                削除しねーよ
              </Button>
              <Button
                colorScheme="red"
                onClick={() => {
                  setIsDeleteDialogOpen(false);
                  setIsSecondDialogOpen(true);
                }}
                ml={3}
              >
                削除しちゃうの？
              </Button>
            </AlertDialogFooter>
          </AlertDialogContent>
        </AlertDialogOverlay>
      </AlertDialog>
      <AlertDialog
        isOpen={isSecondDialogOpen}
        leastDestructiveRef={cancelRef}
        onClose={() => setIsSecondDialogOpen(false)}
      >
        <AlertDialogOverlay>
          <AlertDialogContent>
            <AlertDialogHeader fontSize="lg" fontWeight="bold">
              え？本当に削除するんですか？
            </AlertDialogHeader>
            <AlertDialogBody>
              さっき「削除」を押しましたよね？本当に本当に削除しちゃっていいんですか？
            </AlertDialogBody>
            <AlertDialogFooter>
              <Button
                ref={cancelRef}
                onClick={() => setIsSecondDialogOpen(false)}
              >
                やっぱりやめます
              </Button>
              <Button
                colorScheme="red"
                onClick={() => {
                  setIsSecondDialogOpen(false);
                  setIsThirdDialogOpen(true);
                }}
                ml={3}
                isLoading={isDeleting}
              >
                本当に削除します
              </Button>
            </AlertDialogFooter>
          </AlertDialogContent>
        </AlertDialogOverlay>
      </AlertDialog>
      <AlertDialog
        isOpen={isThirdDialogOpen}
        leastDestructiveRef={cancelRef}
        onClose={() => setIsThirdDialogOpen(false)}
      >
        <AlertDialogOverlay>
          <AlertDialogContent>
            <AlertDialogHeader fontSize="lg" fontWeight="bold">
              なんで削除するんですか？
            </AlertDialogHeader>
            <AlertDialogBody>削除したら泣いちゃうよ？</AlertDialogBody>
            <AlertDialogFooter>
              <Button
                ref={cancelRef}
                onClick={() => setIsThirdDialogOpen(false)}
              >
                慰める
              </Button>
              <Button
                colorScheme="red"
                onClick={handleDelete}
                ml={3}
                isLoading={isDeleting}
              >
                じゃあね
              </Button>
            </AlertDialogFooter>
          </AlertDialogContent>
        </AlertDialogOverlay>
      </AlertDialog>
    </Box>
  );
}
