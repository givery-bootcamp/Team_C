import { Routes, Route } from 'react-router-dom';

import { HelloWorld } from 'features/helloworld';
import { Posts } from 'features/posts';
export const AppRoute = () => {
  return (
    <Routes>
      <Route path="/" element={<HelloWorld />} />
      <Route path="/posts/" element={<Posts />} />
    </Routes>
  );
};
