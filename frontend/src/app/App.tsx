import { AppRoute } from './AppRoute';

import './App.scss';
import { ChakraProvider } from '@chakra-ui/react';
import { Link } from 'react-router-dom';

function App() {
  return (
    <ChakraProvider>
      <div className="app-root">
        <Link to="/">
          <header className="app-header">サンプルアプリケーション</header>
        </Link>
        <main className="app-body container">
          <AppRoute />
        </main>
      </div>
    </ChakraProvider>
  );
}

export default App;
