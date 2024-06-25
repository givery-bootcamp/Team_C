import React from 'react';
import { useAppDispatch, useAppSelector } from 'shared/hooks';
import { useFormik } from 'formik';
import { useNavigate } from 'react-router-dom';
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
  useToast,
} from '@chakra-ui/react';
import { APIService } from 'shared/services';
import * as Yup from 'yup';
import { ModelUserSigninParam } from 'api';

export function SignInForm() {
  const dispatch = useAppDispatch();
  const navigate = useNavigate();
  const toast = useToast();
  const [showPassword, setShowPassword] = React.useState(false);

  const formik = useFormik<ModelUserSigninParam>({
    initialValues: {
      name: '',
      password: '',
    },
    validationSchema: Yup.object({
      name: Yup.string()
        .max(15, 'Must be 15 characters or less')
        .required('Required'),
      password: Yup.string()
        .max(10, 'Must be 10 characters or less')
        .required('Required'),
    }),
    onSubmit: async (value) => {
      try {
        const result = await dispatch(APIService.postSignin({ param: value }));
        if (APIService.postSignin.fulfilled.match(result)) {
          localStorage.setItem('loginSuccess', 'true');
          navigate('/posts');
        } else if (APIService.postSignin.rejected.match(result)) {
          toast({
            title: 'ログイン失敗',
            description: 'ユーザー名またはパスワードが正しくありません。',
            status: 'error',
            duration: 5000,
            isClosable: true,
          });
        }
      } catch (error) {
        console.error('サインインエラー:', error);
        toast({
          title: 'エラー',
          description:
            'ログイン中にエラーが発生しました。後でもう一度お試しください。',
          status: 'error',
          duration: 5000,
          isClosable: true,
        });
      }
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
        p={12}
        width={['90%', '80%', '60%', '50%']}
        maxWidth="700px"
        borderWidth={1}
        borderRadius="lg"
        boxShadow={useColorModeValue('md', 'dark-lg')}
        bg={useColorModeValue('white', 'gray.700')}
      >
        <VStack spacing={8} align="stretch" w="100%">
          <Heading as="h2" size="xl" textAlign="center" fontSize="4xl">
            ログイン
          </Heading>
          <form onSubmit={formik.handleSubmit} style={{ width: '100%' }}>
            <VStack spacing={6} align="stretch" w="100%">
              <FormControl isRequired>
                <FormLabel fontSize="xl">名前</FormLabel>
                <Input
                  id="name"
                  type="text"
                  placeholder="名前を入力しろ"
                  {...formik.getFieldProps('name')}
                  height="60px"
                  fontSize="lg"
                />
              </FormControl>

              <FormControl isRequired>
                <FormLabel fontSize="xl">パスワード</FormLabel>
                <InputGroup size="lg">
                  <Input
                    id="password"
                    type={showPassword ? 'text' : 'password'}
                    placeholder="パスワードを入力しろ"
                    {...formik.getFieldProps('password')}
                    height="60px"
                    fontSize="lg"
                  />
                  <InputRightElement width="5.5rem" height="60px">
                    <Button
                      h="2rem"
                      size="md"
                      onClick={togglePasswordVisibility}
                    >
                      {showPassword ? '隠す' : '表示'}
                    </Button>
                  </InputRightElement>
                </InputGroup>
              </FormControl>

              <Button
                mt={8}
                colorScheme="teal"
                isLoading={formik.isSubmitting}
                type="submit"
                w="100%"
                height="60px"
                fontSize="xl"
              >
                ログイン
              </Button>
            </VStack>
          </form>
        </VStack>
      </Box>
    </Box>
  );
}
