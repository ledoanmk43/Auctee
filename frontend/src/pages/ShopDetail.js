import { useState, useEffect, lazy, Suspense } from 'react';
import { Link as RouterLink, useNavigate, useLocation, useSearchParams } from 'react-router-dom';
import TimeAgo from 'javascript-time-ago';
import vi from 'javascript-time-ago/locale/vi';
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
TimeAgo.addLocale(vi);
export default function ShopDetail() {
  const timeAgo = new TimeAgo('vi-VN');
  const theme = useTheme();
  const navigate = useNavigate();
  const [searchParams, setSearchParams] = useSearchParams();
  const userId = searchParams.get('user');

  const [userData, setUserData] = useState();

  const handleFetchUserData = async () => {
    await fetch(`http://localhost:8080/auctee/user?id=${userId}`, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
      mode: 'cors',
    }).then((res) => {
      if (res.status === 200) {
        res.json().then((data) => {
          setUserData(data);
        });
      }
      if (res.status === 401) {
        alert('You need to login first');

        navigate('/auctee/login', { replace: true });
      }
    });
  };
  const [auctionsData, setAuctionData] = useState();

  const handleFetchAuctionData = async () => {
    await fetch(`http://localhost:8080/auctee/user/auctions?id=${userId}`, {
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
  useEffect(() => {
    // eslint-disable-next-line no-unused-expressions
    !userData && handleFetchUserData();
  }, [userData]);
  return (
    <Suspense startTransition callback={<></>}>
      <Page sx={{ minHeight: 580 }} title={userData?.shopname}>
        <RootStyle sx={{ px: 3, py: 2 }}>
          {/* Heading */}
          <Stack direction="row">
            <Avatar sx={{ width: 67, height: 67 }} src={userData?.avatar} alt="photoURL" />
            <Stack sx={{ mx: 2 }} alignItems="flex-start">
              <Typography fontSize={'1.2rem'} variant="body2" sx={{ color: 'black' }}>
                {userData?.shopname}
              </Typography>
              <Typography variant="caption" sx={{ color: 'green', position: 'relative' }}>
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
                Shop yêu thích
              </Typography>
            </Stack>
            {/* Auctions */}
            <Stack sx={{ ml: 1, borderLeft: '1px solid #f0f0f1' }} alignItems="flex-start">
              {/* Auctions count */}
              <Stack direction="row" alignItems="center">
                <Typography fontSize={'0.9rem'} variant="body2" sx={{ mx: 2, color: 'black' }}>
                  Phiên đấu giá:
                </Typography>
                <Typography fontSize={'0.9rem'} variant="body2" sx={{ color: theme.palette.background.main }}>
                  {auctionsData?.length}
                </Typography>
              </Stack>
              {/* Reply rate */}
              <Stack direction="row" alignItems="center">
                <Typography fontSize={'0.9rem'} variant="body2" sx={{ mx: 2, color: 'black' }}>
                  Đánh giá:
                </Typography>
                <Typography fontSize={'0.9rem'} variant="body2" sx={{ color: theme.palette.background.main }}>
                  4.8/5 ({userData?.id * 43} đánh giá)
                </Typography>
              </Stack>
              {/* Date Join */}
              <Stack direction="row" alignItems="center">
                <Typography fontSize={'0.9rem'} variant="body2" sx={{ mx: 2, color: 'black' }}>
                  Tham gia:
                </Typography>
                <Typography fontSize={'0.9rem'} variant="body2" sx={{ color: theme.palette.background.main }}>
                  {userData?.created_at && timeAgo.format(new Date(userData.created_at))}
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
          <Typography variant="h5" sx={{ mb: 2 }}>
            Các phiên đấu giá gần đây
          </Typography>
          {auctionsData && <ProductList auctions={auctionsData} />}
        </Container>
      </Page>
    </Suspense>
  );
}
