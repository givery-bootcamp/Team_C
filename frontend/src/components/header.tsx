import { Box, Button, Flex, Heading } from '@chakra-ui/react';
import { Link, useLocation, useNavigate } from 'react-router-dom';

export function Header() {
  const navigate = useNavigate();
  const location = useLocation();

  const navigateSignIn = () => {
    navigate('/login');
  };

  const isLoginPage = location.pathname === '/login';

  return (
    <Box as="header" py={4} bg="gray.100">
      <Flex
        maxW="container.xl"
        mx="auto"
        px={4}
        justifyContent="space-between"
        alignItems="center"
      >
        <Heading as="h1" size="lg">
          <Link to="/">team3 掲示板</Link>
        </Heading>
        {!isLoginPage && (
          <Button colorScheme="blue" onClick={navigateSignIn}>
            サインイン
          </Button>
        )}
      </Flex>
    </Box>
  );
}
