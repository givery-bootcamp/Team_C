import { useEffect } from 'react';

import { useAppDispatch, useAppSelector } from 'shared/hooks';
import { APIService } from 'shared/services';
import { Box, Divider, HStack, Heading, Text, Avatar, Button, ButtonGroup } from '@chakra-ui/react';
import { useParams } from 'react-router-dom'

function format_date(date_string:any) {
  const date = new Date(date_string)
  return date.toLocaleString('ja-JP')
}

export function PostDetail() {
  const {postId} = useParams()
  const { postdetail } = useAppSelector((state) => state.postdetail);
  const dispatch = useAppDispatch();

  useEffect(() => {
    dispatch(APIService.getPostDetail(Number(postId)));
  }, [dispatch]);

  // console.log(postId)

  return(
    <Box>
        <Heading as={'h1'} >{postdetail?.title}</Heading>
        <HStack >
          <Avatar size="sm" name={postdetail?.user?.name} />
          <Text as={'span'}>{postdetail?.user?.name}</Text>
        </HStack>
        
        <HStack spacing={10}>
          <Text as={'span'}>作成日時 {format_date(postdetail?.created_at)}</Text>
          <Text as={'span'}>投稿日時 {format_date(postdetail?.updated_at)}</Text>
        </HStack>
        <Divider />
        <Text>{postdetail?.body}</Text>
        <ButtonGroup colorScheme='blue'>
          <Button>編集</Button>
          <Button>削除</Button>
        </ButtonGroup>
        
    </Box>
  )
}
