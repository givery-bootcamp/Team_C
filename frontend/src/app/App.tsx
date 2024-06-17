import { AppRoute } from './AppRoute';

import './App.scss';
import { ChakraProvider } from '@chakra-ui/react';

function App() {
  return (
    <ChakraProvider>
      <div className="app-root">
        <header className="app-header">サンプルアプリケーション</header>
        <main className="app-body container">
          <AppRoute />
        </main>
      </div>
    </ChakraProvider>
  );
}

export default App;
