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
const CreateProductForm = lazy(() => import('../sections/update-product/CreateProductForm'));
const ProductList = lazy(() => import('../sections/update-product'));

const RootStyle = styled('div')(({ theme }) => ({
  [theme.breakpoints.up('md')]: {
    display: 'flex',
    flexDirection: 'column',
    mx: '176x',
    backgroundColor: 'white',
    height: '100%',
    minHeight: '580px',
    maxWidth: '1042px',
  },
}));

export default function Inventory() {
  const theme = useTheme();
  const { isReloading, setIsReloading } = useContext(ReloadContext);

  return (
    <Suspense startTransition callback={<></>}>
      <Page title="Tất cả sản phẩm">
        <RootStyle sx={{ px: 3, py: 2 }}>
          {/* Heading */}
          <Stack sx={{ pb: 0 }}>
            <Typography fontSize={'1.2rem'} variant="body2" sx={{ color: 'black' }}>
              Quản lý sản phẩm
            </Typography>
            <Typography variant="body2" sx={{ color: 'black', opacity: 0.8, pb: 2 }}>
              Đăng sản phẩm và mở phiên đấu giá
            </Typography>
            <Divider />
          </Stack>
          {/* Main */}
          <CreateProductForm />
          <ProductList />
        </RootStyle>
      </Page>
    </Suspense>
  );
}
