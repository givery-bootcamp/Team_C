import { Route, Routes } from 'react-router-dom';

import { HelloWorld } from 'features/helloworld';
import { SignInForm } from 'features/login';
import { PostDetail } from 'features/postdetail';
import { EditPost } from 'features/postdetail/edit';
import { Posts } from 'features/posts';

export const AppRoute = () => {
  return (
    <Routes>
      <Route path="/" element={<HelloWorld />} />
      <Route path="/posts/" element={<Posts />} />
      <Route path="/posts/:postId" element={<PostDetail />} />
      <Route path="/posts/:postId/edit" element={<EditPost />} />
      <Route path="/login" element={<SignInForm />} />
    </Routes>
  );
};
