import { useState, useEffect, lazy, Suspense } from 'react';
import { Link as RouterLink, useNavigate, useLocation, useSearchParams } from 'react-router-dom';

// material
import {
  Container,
  Box,
  Avatar,
  FormGroup,
  TextField,
  RadioGroup,
  Radio,
  FormControlLabel,
  Typography,
  Stack,
  Button,
  Select,
  MenuItem,
  Divider,
} from '@mui/material';

import { styled, useTheme } from '@mui/material/styles';

const Page = lazy(() => import('../components/Page'));
const UpdateProfileForm = lazy(() => import('../sections/update-profile'));

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

export default function User() {
  return (
    <Suspense startTransition callback={<></>}>
      <Page title="Hồ sơ">
        <RootStyle sx={{ px: 3, py: 2 }}>
          {/* Heading */}
          <Stack>
            <Typography fontSize={'1.2rem'} variant="body2" sx={{ color: 'black' }}>
              Hồ Sơ Của Tôi
            </Typography>
            <Typography variant="body2" sx={{ color: 'black', opacity: 0.8, pb: 2 }}>
              Quản lý thông tin cá nhân để bảo mật tài khoản
            </Typography>
            <Divider />
          </Stack>
          {/* Main */}
          <UpdateProfileForm />
        </RootStyle>
      </Page>
    </Suspense>
  );
}
