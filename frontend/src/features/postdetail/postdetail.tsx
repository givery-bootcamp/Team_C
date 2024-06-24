import { useEffect } from 'react';

import { useAppDispatch, useAppSelector } from 'shared/hooks';
import { APIService } from 'shared/services';
import { model_Post } from 'api';
import { Box } from '@chakra-ui/react';
import { useParams } from 'react-router-dom'

export function PostDetail() {
  const post_id = useParams()
  const { postdetail } = useAppSelector((state) => state.postdetail);
  const dispatch = useAppDispatch();

  useEffect(() => {
    dispatch(APIService.getPostDetail());
  }, [dispatch]);

  console.log(post_id)

  return(
    <Box>
        <h1>{postdetail?.id}</h1>
        <h1>{postdetail?.created_at}</h1>
        
    </Box>
  )
}