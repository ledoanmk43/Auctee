import { useState, useEffect, lazy, Suspense, useContext } from 'react';
import { Link as RouterLink, useNavigate, useLocation, useSearchParams } from 'react-router-dom';

// material
import {
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle,
  TextField,
  Typography,
  Stack,
  Button,
  Divider,
  Box,
} from '@mui/material';

import { styled, useTheme } from '@mui/material/styles';
import { Icon } from '@iconify/react';
import { ReloadContext } from '../utils/Context';

const Page = lazy(() => import('../components/Page'));
const CreateAddressForm = lazy(() => import('../sections/update-address/CreateAddressForm'));
const AddressList = lazy(() => import('../sections/update-address'));

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

export default function Addess() {
  const theme = useTheme();
  const { isReloading, setIsReloading } = useContext(ReloadContext);

  return (
    <Suspense startTransition callback={<></>}>
      <Page title="Địa chỉ">
        <RootStyle sx={{ px: 3, py: 2 }}>
          {/* Heading */}
          <Stack sx={{ pb: 0 }}>
            <Typography fontSize={'1.2rem'} variant="body2" sx={{ color: 'black' }}>
              Địa chỉ giao hàng
            </Typography>
            <Typography variant="body2" sx={{ color: 'black', opacity: 0.8, pb: 2 }}>
              Địa chỉ càng chính xác giao hàng càng nhanh
            </Typography>
            <Divider />
          </Stack>
          {/* Main */}
          <CreateAddressForm />
          <AddressList />
        </RootStyle>
      </Page>
    </Suspense>
  );
}
