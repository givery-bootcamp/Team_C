import { useEffect, useState } from 'react';

import { useAppDispatch, useAppSelector } from 'shared/hooks';
import { APIService } from 'shared/services';
import { model_UserSigninParam } from '../../api/models/model_UserSigninParam';
import {
  Button,
  FormControl,
  FormLabel,
  Input,
  InputGroup,
  InputRightElement,
} from '@chakra-ui/react';
import { useFormik } from 'formik';

export function Login() {
  const { signinParam } = useAppSelector((state) => state.signin);
  const [show, setShow] = useState(false);
  const dispatch = useAppDispatch();

  //   useEffect(() => {
  //     dispatch(APIService.postSignin(signinParam as model_UserSigninParam));
  //   }, [dispatch, signinParam]);

  const handleClick = () => setShow(!show);

  const formik = useFormik({
    initialValues: {
      name: '',
      password: '',
    } as model_UserSigninParam,
    onSubmit: (values) => {
      dispatch(APIService.postSignin(values));
    },
  });

  return (
    <form
      style={{
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        justifyContent: 'center',
        height: '100vh',
        width: '100vw',
      }}
    >
      <FormControl isRequired>
        <FormLabel>name</FormLabel>
        <InputGroup
          size="md"
          style={{
            marginBottom: '20px',
            width: '300px',
          }}
        >
          <Input
            pr="4.5rem"
            type="text"
            placeholder="Enter your name"
            value={signinParam?.name}
          />
        </InputGroup>
      </FormControl>
      <FormControl isRequired>
        <FormLabel>password</FormLabel>
        <InputGroup
          size="md"
          style={{
            marginBottom: '20px',
            width: '300px',
          }}
        >
          <Input
            pr="4.5rem"
            type="password"
            placeholder="Enter your password"
            value={signinParam?.password}
          />
          <InputRightElement width="4.5rem">
            <Button h="1.75rem" size="sm" onClick={handleClick}>
              {show ? 'Hide' : 'Show'}
            </Button>
          </InputRightElement>
        </InputGroup>
      </FormControl>
      <Button
        mt={4}
        colorScheme="teal"
        isLoading={formik.isSubmitting}
        type="submit"
      >
        Submit
      </Button>
    </form>
  );
}
