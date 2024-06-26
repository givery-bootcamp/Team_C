import { Routes, Route } from 'react-router-dom';

import { HelloWorld } from 'features/helloworld';
import { Posts } from 'features/posts';
import { SignInForm } from 'features/login';
import { PostDetail } from 'features/postdetail';

export const AppRoute = () => {
  return (
    <Routes>
      <Route path="/" element={<HelloWorld />} />
      <Route path="/posts/" element={<Posts />} />
      <Route path="/posts/:postId" element={ <PostDetail /> } />
      <Route path="/login" element={<SignInForm />} />
    </Routes>
  );
};
