import { AppRoute } from './AppRoute';

import './App.scss';
import { ChakraProvider } from '@chakra-ui/react';
import { Header } from 'components/header';

function App() {
  return (
    <ChakraProvider>
      <div className="app-root">
        <Header />
        <main className="app-body container">
          <AppRoute />
        </main>
      </div>
    </ChakraProvider>
  );
}

export default App;
