import { AppRoute } from './AppRoute';

import './App.scss';
import { ChakraProvider } from '@chakra-ui/react';
import SideMenu from '../components/sidemenu'

function App() {
  return (
    <ChakraProvider>
      <div className="app-root">
        <header className="app-header">サンプルアプリケーション</header>
        <main className="app-body container">
          <SideMenu />
          <AppRoute />
        </main>
      </div>
    </ChakraProvider>
  );
}

export default App;
