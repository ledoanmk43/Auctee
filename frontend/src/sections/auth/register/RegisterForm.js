import * as Yup from 'yup';
import { useState, useContext } from 'react';
import { useNavigate } from 'react-router-dom';
// form
import { useForm } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers/yup';
// @mui
import { Stack, IconButton, InputAdornment, Typography } from '@mui/material';
import { LoadingButton } from '@mui/lab';
// components
import Iconify from '../../../components/Iconify';
import { FormProvider, RHFTextField } from '../../../components/hook-form';
import { LoginContext } from '../../../utils/Context';
// ----------------------------------------------------------------------

export default function RegisterForm() {
  const navigate = useNavigate();
  const { loggedIn, setLoggedIn } = useContext(LoginContext);
  const [showPassword, setShowPassword] = useState(false);

  const RegisterSchema = Yup.object().shape({
    firstname: Yup.string().required('First name required'),
    lastname: Yup.string().required('Last name required'),
    email: Yup.string().email('Email không chính xác').required('Email is required'),
    password: Yup.string().required('Password is required'),
  });

  const defaultValues = {
    firstName: '',
    lastName: '',
    email: '',
    password: '',
  };

  const methods = useForm({
    resolver: yupResolver(RegisterSchema),
    defaultValues,
  });

  const {
    handleSubmit,
    formState: { isSubmitting },
  } = methods;

  const [isConflictUserName, setIsConflictUserName] = useState(false);

  const [errorRegister, setErrorRegister] = useState('');
  const onSubmit = async (data) => {
    const user = {
      firstname: data.firstname,
      lastname: data.lastname,
      username: data.email,
      password: data.password,
    };
    await fetch('http://localhost:1001/auctee/register', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(user),
      credentials: 'include',
    }).then((res) => {
      if (res.status === 201) {
        setLoggedIn(true);
        navigate('/auctee/home', { replace: true });
      }
      if (res.status === 401) {
        setIsConflictUserName(true);
        setLoggedIn(false);
        res.json().then((data) => {
          setErrorRegister(data.message);
        });
      }
      if (res.status === 409) {
        setIsConflictUserName(true);
        setLoggedIn(false);
        res.json().then((data) => {
          setErrorRegister(data.message);
        });
      }
    });
  };

  return (
    <FormProvider methods={methods} onSubmit={handleSubmit(onSubmit)}>
      <Stack spacing={3}>
        <Stack direction={{ xs: 'column', sm: 'row' }} spacing={2}>
          <RHFTextField name="firstname" label="Tên" />
          <RHFTextField name="lastname" label="Họ" />
        </Stack>

        <RHFTextField name="email" label="Email" />

        <RHFTextField
          name="password"
          label="Mật khẩu"
          type={showPassword ? 'text' : 'password'}
          InputProps={{
            endAdornment: (
              <InputAdornment position="end">
                <IconButton edge="end" onClick={() => setShowPassword(!showPassword)}>
                  <Iconify icon={showPassword ? 'eva:eye-fill' : 'eva:eye-off-fill'} />
                </IconButton>
              </InputAdornment>
            ),
          }}
        />
        {isConflictUserName && (
          <Stack item>
            <Typography color="red" variant="subtitle1">
              {errorRegister}
            </Typography>
          </Stack>
        )}
        <LoadingButton fullWidth size="large" type="submit" variant="contained" loading={isSubmitting}>
          Đăng ký
        </LoadingButton>
      </Stack>
    </FormProvider>
  );
}
