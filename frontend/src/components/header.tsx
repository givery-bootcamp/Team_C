import { Box, Button, Flex, Heading } from '@chakra-ui/react';
import { Link, useLocation, useNavigate } from 'react-router-dom';
import { useAppSelector } from 'shared/hooks';

export function Header() {
  const navigate = useNavigate();
  const location = useLocation();

  const navigateSignIn = () => {
    navigate('/login');
  };

  const isLoginPage = location.pathname === '/login';

  const {user} = useAppSelector((state) => state.user)

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
          <Link to="/posts">team3 掲示板</Link>
        </Heading>
        {!isLoginPage && (
          user ? (
            <Button colorScheme='green'>サインアウト</Button>
           ) : (
            <Button colorScheme="blue" onClick={navigateSignIn}>
              サインイン
            </Button>
          )
        )}
      </Flex>
    </Box>
  );
}
