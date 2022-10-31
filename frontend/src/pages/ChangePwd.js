import { useState, useEffect, lazy, Suspense } from 'react';
import { useNavigate, useLocation } from 'react-router-dom';

import { LoadingButton } from '@mui/lab';
// material
import { Container, Typography, Stack, Divider, IconButton, InputAdornment } from '@mui/material';

import { styled, useTheme } from '@mui/material/styles';

import Iconify from '../components/Iconify';
import { FormProvider, RHFTextField } from '../components/hook-form';
// import ChangePasswordForm from '../sections/password';

const ChangePasswordForm = lazy(() => import('../sections/update-password'));
const Page = lazy(() => import('../components/Page'));

const RootStyle = styled('div')(({ theme }) => ({
  [theme.breakpoints.up('md')]: {
    display: 'flex',
    flexDirection: 'column',
    mx: '176x',
    backgroundColor: 'white',
    height: '100%',
    minHeight: '580px',
  },
}));

export default function ChangePwd() {
  const navigate = useNavigate();
  const location = useLocation();

  return (
    <Suspense startTransition callback={<></>}>
      <Page title="Mật khẩu">
        <RootStyle sx={{ px: 3, py: 2 }}>
          <Stack>
            <Typography fontSize={'1.2rem'} variant="body2" sx={{ color: 'black' }}>
              Đổi mật khẩu
            </Typography>
            <Typography variant="body2" sx={{ color: 'black', opacity: 0.8, pb: 2 }}>
              Để bảo mật tài khoản, vui lòng không chia sẻ mật khẩu cho người khác
            </Typography>
            <Divider />
          </Stack>
          {/* Main */}
          <ChangePasswordForm />
        </RootStyle>
      </Page>
    </Suspense>
  );
}
