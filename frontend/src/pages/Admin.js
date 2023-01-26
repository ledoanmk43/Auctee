import { useState, useEffect, lazy, Suspense } from 'react';
import { Link as RouterLink, useNavigate, useLocation, useSearchParams, useOutletContext } from 'react-router-dom';

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
  LinearProgress,
} from '@mui/material';

import { styled, useTheme } from '@mui/material/styles';
import { Icon } from '@iconify/react';
import { ProductSort, ProductList, ProductFilterSidebar } from '../sections/@dashboard/products';
import DashboardApp from './DashboardApp';

const Page = lazy(() => import('../components/Page'));
const UpdateProfileForm = lazy(() => import('../sections/update-profile'));

const RootStyle = styled('div')(({ theme }) => ({
  [theme.breakpoints.up('md')]: {
    display: 'flex',
    flexDirection: 'column',
    mx: '176x',
    backgroundColor: 'white',
    height: '100%',
  },
}));

export default function Admin() {
  const userData = useOutletContext();
  const theme = useTheme();
  const navigate = useNavigate();
  const [searchParams, setSearchParams] = useSearchParams();
  const userId = searchParams.get('id');

  const [auctionsData, setAuctionData] = useState();

  const handleFetchAuctionData = async () => {
    await fetch(`http://localhost:8080/auctee/auctions?page=${1}`, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
      mode: 'cors',
    }).then((res) => {
      if (res.status === 200) {
        res.json().then((data) => {
          setAuctionData(data);
        });
      }
      if (res.status === 401) {
        alert('You need to login first');
        navigate('/auctee/login', { replace: true });
      }
    });
  };
  useEffect(() => {
    // eslint-disable-next-line no-unused-expressions
    !auctionsData && handleFetchAuctionData();
  }, [auctionsData]);

  return (
    <Suspense startTransition callback={<></>}>
      <Page sx={{ minHeight: 580, maxWidth: 980 }} title="Admin Dashboard">
        <RootStyle sx={{ px: 3, py: 2 }}>
          {/* Heading */}
          <Stack direction="row">
            <Avatar sx={{ width: 67, height: 67 }} src={userData?.avatar} alt="photoURL" />
            <Stack sx={{ mx: 2 }} alignItems="flex-start">
              <Typography fontSize={'1.2rem'} variant="body2" sx={{ color: 'black' }}>
                {userData?.lastname} &nbsp;
                {userData?.firstname}
              </Typography>
              <Typography variant="caption" sx={{ color: 'black', position: 'relative' }}>
                Điểm uy tín : {userData?.honor_point} &nbsp;
              </Typography>
              <Typography
                variant="button"
                sx={{
                  textTransform: 'none',
                  bgcolor: '#F62217',
                  color: 'white',
                  borderRadius: 0.5,
                  fontSize: '0.7rem',
                  px: 0.5,
                  mr: 1.5,
                }}
              >
                Quản trị viên
              </Typography>
            </Stack>
            {/* Auctions */}
            <Stack sx={{ ml: 1, borderLeft: '1px solid #f0f0f1' }} alignItems="flex-start">
              {/* Auctions count */}
              <Stack direction="row" alignItems="center">
                <Typography fontSize={'0.9rem'} variant="body2" sx={{ mx: 2, color: 'black' }}>
                  Tồng thu nhập ước tính:
                </Typography>
                <Typography fontSize={'0.9rem'} variant="body2" sx={{ color: theme.palette.background.main }}>
                  {userData?.system_balance.toLocaleString('tr-TR', {
                    style: 'currency',
                    currency: 'VND',
                  })}
                </Typography>
              </Stack>
              {/* Reply rate */}
              <Stack direction="row" alignItems="center">
                <Typography fontSize={'0.9rem'} variant="body2" sx={{ mx: 2, color: 'black' }}>
                  Tổng số lượng người dùng:
                </Typography>
                <Typography fontSize={'0.9rem'} variant="body2" sx={{ color: theme.palette.background.main }}>
                  {userData?.total_user} tài khoản
                </Typography>
              </Stack>
              {/* Date Join */}
              <Stack direction="row" alignItems="center">
                <Typography fontSize={'0.9rem'} variant="body2" sx={{ mx: 2, color: 'black' }}>
                  Tham gia:
                </Typography>
                <Typography fontSize={'0.9rem'} variant="body2" sx={{ color: theme.palette.background.main }}>
                  {userData?.id} tháng trước
                </Typography>
              </Stack>
            </Stack>
            {/* Contact */}
            <Stack sx={{ ml: 4, borderLeft: '1px solid #f0f0f1' }} alignItems="flex-start">
              {/* Auctions count */}
              <Stack direction="row" alignItems="center">
                <Typography fontSize={'0.9rem'} variant="body2" sx={{ mx: 2, color: 'black' }}>
                  Liên hệ:
                </Typography>
                <Typography fontSize={'0.9rem'} variant="body2" sx={{ color: theme.palette.background.main }}>
                  {userData?.phone}
                </Typography>
              </Stack>
              {/* Auctions count */}
              <Stack direction="row" alignItems="center">
                <Typography fontSize={'0.9rem'} variant="body2" sx={{ mx: 2, color: 'black' }}>
                  Email:
                </Typography>
                <Typography fontSize={'0.9rem'} variant="body2" sx={{ color: theme.palette.background.main }}>
                  {userData?.email}
                </Typography>
              </Stack>
              {/* Reply rate */}
              <Stack direction="row" alignItems="center">
                <Typography fontSize={'0.9rem'} variant="body2" sx={{ mx: 2, color: 'black' }}>
                  Địa chỉ:
                </Typography>
                <Typography fontSize={'0.9rem'} variant="body2" sx={{ color: theme.palette.background.main }}>
                  ********, HCM
                </Typography>
              </Stack>
            </Stack>
          </Stack>
        </RootStyle>
        <Container sx={{ my: 3 }}>
          <DashboardApp />
          {/* {auctionsData && <ProductList auctions={auctionsData} />} */}
        </Container>
      </Page>
    </Suspense>
  );
}
