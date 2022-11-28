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
const CreateAuctionForm = lazy(() => import('../sections/update-auction/CreateAuctionForm'));
const AuctionList = lazy(() => import('../sections/update-auction'));

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

export default function Auction() {
  const theme = useTheme();
  const { isReloading, setIsReloading } = useContext(ReloadContext);

  return (
    <Suspense startTransition callback={<></>}>
      <Page title="Tất cả phiên đấu giá">
        <RootStyle sx={{ px: 3, py: 2, maxWidth: '980px !important' }}>
          {/* Heading */}
          <Stack sx={{ pb: 0 }}>
            <Typography fontSize={'1.2rem'} variant="body2" sx={{ color: 'black' }}>
              Quản lý sàn đấu giá
            </Typography>
            <Typography variant="body2" sx={{ color: 'black', opacity: 0.8, pb: 2 }}>
              Tạo phiên đấu giá sản phẩm của bạn
            </Typography>
            <Divider />
          </Stack>
          {/* Main */}
          <CreateAuctionForm />
          <AuctionList />
        </RootStyle>
      </Page>
    </Suspense>
  );
}
