import { useState, useEffect, lazy, Suspense } from 'react';
import { Outlet, useNavigate } from 'react-router-dom';
// material
import { styled, useTheme } from '@mui/material/styles';
import { Stack, Typography, CardMedia } from '@mui/material';
import { Icon } from '@iconify/react';
//
// import MainNavbar from './MainNavbar';
// import Sidebar from './Sidebar';

const ProductCartWidget = lazy(() =>
  import('../../sections/@dashboard/products').then((module) => ({
    default: module.ProductCartWidget,
  }))
);

const MainNavbar = lazy(() => import('./MainNavbar'));
const Sidebar = lazy(() => import('./Sidebar'));

// ----------------------------------------------------------------------

const APP_BAR_MOBILE = 64;
const APP_BAR_DESKTOP = 92;

const RootStyle = styled('div')({
  justifyContent: 'space-between',
  display: 'flex',
  flexDirection: 'column',
  minHeight: '100%',
  overflow: 'hidden',
});

const MainStyle = styled('div')(({ theme }) => ({
  minHeight: '100%',
  minWidth: '1074px',
  zIndex: 0,
  overflow: 'hidden',
  backgroundColor: theme.palette.background.default,
  paddingTop: APP_BAR_MOBILE + 24,
  paddingBottom: theme.spacing(5),
  [theme.breakpoints.up('lg')]: {
    paddingTop: APP_BAR_DESKTOP + 24,
    paddingLeft: theme.spacing(2),
    paddingRight: theme.spacing(2),
  },
}));

// ----------------------------------------------------------------------

export default function MainLayout() {
  const theme = useTheme();
  const navigate = useNavigate();

  const [open, setOpen] = useState(false);

  const [userData, setUserData] = useState();

  const handleFetchUserData = async () => {
    await fetch('http://localhost:8080/auctee/user/profile', {
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
  useEffect(() => {
    // eslint-disable-next-line no-unused-expressions
    !userData && handleFetchUserData();
  }, [userData]);

  return userData ? (
    <Suspense fallback={<></>}>
      <RootStyle>
        <MainNavbar userData={userData} onOpenSidebar={() => setOpen(true)} />
        <div style={{ margin: 'auto', display: 'flex', flexDirection: 'row' }}>
          <Sidebar userData={userData} />
          <MainStyle userData={userData}>
            <Outlet context={userData} />
            <ProductCartWidget />
          </MainStyle>
        </div>
        <Stack sx={{ bgcolor: 'white', pb: 1 }} alignItems="center">
          <Stack
            direction="row"
            justifyContent="space-between"
            sx={{
              px: 30,
              py: 2,
              bgcolor: 'white',
              minHeight: 200,
              borderTop: `5px solid ${theme.palette.background.main}`,
              minWidth: '100%',
            }}
          >
            {/* One */}
            <Stack>
              <Typography sx={{ mb: 1.5 }} fontWeight={600} variant="body1">
                CHĂM SÓC KHÁCH HÀNG
              </Typography>
              <Typography sx={{ my: 0.5 }} variant="body2">
                Trung Tâm Trợ Giúp
              </Typography>
              <Typography sx={{ my: 0.5 }} variant="body2">
                Hướng Dẫn Bán hàng
              </Typography>
              <Typography sx={{ my: 0.5 }} variant="body2">
                Hướng Dẫn Mua hàng
              </Typography>
              <Typography sx={{ my: 0.5 }} variant="body2">
                Chăm Sóc Khách Hàng
              </Typography>
            </Stack>
            {/* Two */}
            <Stack alignItems="flex-start">
              <Typography fontWeight={600} variant="body1">
                THANH TOÁN
              </Typography>
              <Stack direction="row" alignItems="center">
                <CardMedia component="img" height="100" image="/static/momo.png" alt="MoMo" />
              </Stack>
            </Stack>
            {/* Three */}
            <Stack>
              <Typography sx={{ mb: 1.5 }} fontWeight={600} variant="body1">
                THEO DÕI CHÚNG TÔI TRÊN
              </Typography>
              <Stack sx={{ my: 0.5 }} direction="row" alignItems="center">
                <Icon icon="ri:facebook-box-fill" />
                <Typography sx={{ mx: 1 }} variant="body2">
                  Facebook
                </Typography>
              </Stack>
              <Stack sx={{ my: 0.5 }} direction="row" alignItems="center">
                <Icon icon="ri:instagram-fill" />
                <Typography sx={{ mx: 1 }} variant="body2">
                  Instagram
                </Typography>
              </Stack>
              <Stack sx={{ my: 0.5 }} direction="row" alignItems="center">
                <Icon icon="mdi:linkedin" />
                <Typography sx={{ mx: 1 }} variant="body2">
                  LinkedIn
                </Typography>
              </Stack>
            </Stack>
          </Stack>
          <Typography>Đấu giá Trực tuyến</Typography>
          <Typography>© 2022 Auctee</Typography>
        </Stack>
      </RootStyle>
    </Suspense>
  ) : (
    <></>
  );
}
