import { Routes, Route } from 'react-router-dom';

import { HelloWorld } from 'features/helloworld';
import { Posts } from 'features/posts';
import { PostDetail } from 'features/postdetail/postdetail';
export const AppRoute = () => {
  return (
    <Routes>
      <Route path="/" element={<HelloWorld />} />
      <Route path="/posts/" element={<Posts />} />
        <Route path=":postId" element={ <PostDetail /> } />{/*post/のみで飛ぶように要改善*/}
    </Routes>
  );
};
