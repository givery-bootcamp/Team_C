import React from 'react';
import { useAppDispatch } from 'shared/hooks';
import { useFormik } from 'formik';
import {
  Box,
  VStack,
  FormControl,
  FormLabel,
  Input,
  InputGroup,
  InputRightElement,
  Button,
  Heading,
  useColorModeValue,
} from '@chakra-ui/react';
import { APIService } from 'shared/services';

interface SignInFormValues {
  name: string;
  password: string;
}

export const SignInForm: React.FC = () => {
  const dispatch = useAppDispatch();
  const [showPassword, setShowPassword] = React.useState(false);

  const formik = useFormik<SignInFormValues>({
    initialValues: {
      name: '',
      password: '',
    },
    onSubmit: (values) => {
      dispatch(APIService.postSignin(values));
    },
  });

  const togglePasswordVisibility = () => setShowPassword(!showPassword);

  return (
    <Box
      minHeight="100vh"
      width="100%"
      display="flex"
      alignItems="center"
      justifyContent="center"
      bg={useColorModeValue('gray.50', 'gray.800')}
    >
      <Box
        p={8}
        maxWidth="400px"
        borderWidth={1}
        borderRadius="lg"
        boxShadow="lg"
        bg={useColorModeValue('white', 'gray.700')}
      >
        <VStack spacing={4} align="flex-start" w="100%">
          <Heading as="h2" size="xl" textAlign="center" w="100%">
            ログイン
          </Heading>
          <form onSubmit={formik.handleSubmit} style={{ width: '100%' }}>
            <VStack spacing={4} align="flex-start" w="100%">
              <FormControl isRequired>
                <FormLabel>名前</FormLabel>
                <Input
                  id="name"
                  type="text"
                  placeholder="名前を入力しろ"
                  {...formik.getFieldProps('name')}
                />
              </FormControl>

              <FormControl isRequired>
                <FormLabel>パスワード</FormLabel>
                <InputGroup>
                  <Input
                    id="password"
                    type={showPassword ? 'text' : 'password'}
                    placeholder="パスワードを入力しろ"
                    {...formik.getFieldProps('password')}
                  />
                  <InputRightElement width="4.5rem">
                    <Button
                      h="1.75rem"
                      size="sm"
                      onClick={togglePasswordVisibility}
                    >
                      {showPassword ? '隠す' : '表示'}
                    </Button>
                  </InputRightElement>
                </InputGroup>
              </FormControl>

              <Button
                mt={4}
                colorScheme="teal"
                isLoading={formik.isSubmitting}
                type="submit"
                w="100%"
              >
                ログイン
              </Button>
            </VStack>
          </form>
        </VStack>
      </Box>
    </Box>
  );
};
