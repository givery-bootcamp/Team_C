import { AppRoute } from './AppRoute';

import {
  ChakraProvider,
  Container,
  useColorModeValue,
  Box,
} from '@chakra-ui/react';
import { Header } from 'components/header';

function App() {
  const bgColor = useColorModeValue('gray.50', 'gray.800');
  return (
    <ChakraProvider>
      <Box
        height="100vh"
        width="100vw"
        display="flex"
        flexDirection="column"
        bg={bgColor}
      >
        <Header />
        <Container
          as="main"
          maxW="container.xl"
          flex="1"
          display="flex"
          flexDirection="column"
          overflow="auto"
          px={12}
        >
          <AppRoute />
        </Container>
      </Box>
    </ChakraProvider>
  );
}

export default App;
