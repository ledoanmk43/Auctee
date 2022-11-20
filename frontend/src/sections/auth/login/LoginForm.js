import * as Yup from 'yup';
import { useState, useContext } from 'react';
import { useNavigate } from 'react-router-dom';
// form
import { useForm } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers/yup';
// @mui
import { Link, Stack, IconButton, InputAdornment, Typography } from '@mui/material';
import { LoadingButton } from '@mui/lab';
// components
import Iconify from '../../../components/Iconify';
import { FormProvider, RHFTextField, RHFCheckbox } from '../../../components/hook-form';

import { LoginContext } from '../../../utils/Context';

// ----------------------------------------------------------------------

export default function LoginForm() {
  const navigate = useNavigate();
  const { loggedIn, setLoggedIn } = useContext(LoginContext);

  const [showPassword, setShowPassword] = useState(false);

  const LoginSchema = Yup.object().shape({
    email: Yup.string().email('Email không chính xác').required('Email is required'),
    password: Yup.string().required('Password is required'),
  });

  const defaultValues = {
    email: '',
    password: '',
    remember: true,
  };

  const methods = useForm({
    resolver: yupResolver(LoginSchema),
    defaultValues,
  });

  const {
    handleSubmit,
    formState: { isSubmitting },
  } = methods;

  const [isBadReq, setIsBadReq] = useState(false);
  const [errorLogin, setErrorLogin] = useState('');

  const onSubmit = async (data) => {
    const user = {
      username: data.email,
      password: data.password,
    };
    await fetch('http://localhost:1001/auctee/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(user),
      mode: 'cors',
      credentials: 'include',
    }).then((res) => {
      if (res.status === 200) {
        setLoggedIn(true);
        navigate('/auctee/home', { replace: true });
      }
      if (res.status === 401) {
        setLoggedIn(false);
        setIsBadReq(true);
        res.json().then((data) => {
          setErrorLogin(data.message);
        });
      }
    });
  };

  return (
    <FormProvider methods={methods} onSubmit={handleSubmit(onSubmit)}>
      <Stack spacing={3}>
        <RHFTextField name="email" label="Email" />

        <RHFTextField
          name="password"
          label="Mật khẩu"
          type={showPassword ? 'text' : 'password'}
          InputProps={{
            endAdornment: (
              <InputAdornment position="end">
                <IconButton onClick={() => setShowPassword(!showPassword)} edge="end">
                  <Iconify icon={showPassword ? 'eva:eye-fill' : 'eva:eye-off-fill'} />
                </IconButton>
              </InputAdornment>
            ),
          }}
        />
      </Stack>
      {isBadReq && (
        <Stack item>
          <Typography color="red" variant="subtitle1">
            {errorLogin}
          </Typography>
        </Stack>
      )}
      <Stack direction="row" alignItems="center" justifyContent="space-between" sx={{ my: 2 }}>
        <RHFCheckbox name="remember" label="Remember me" />
        <Link variant="subtitle2" underline="hover">
          Forgot password?
        </Link>
      </Stack>

      <LoadingButton fullWidth size="large" type="submit" variant="contained" loading={isSubmitting}>
        Đăng nhập
      </LoadingButton>
    </FormProvider>
  );
}
