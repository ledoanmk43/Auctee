import { useState, useEffect } from 'react';
import { Outlet, useNavigate } from 'react-router-dom';
// material
import { styled } from '@mui/material/styles';
import { Container } from '@mui/material';
//
import MainNavbar from './MainNavbar';
import Sidebar from './Sidebar';

// ----------------------------------------------------------------------

const APP_BAR_MOBILE = 64;
const APP_BAR_DESKTOP = 92;

const RootStyle = styled('div')({
  display: 'flex',
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
  paddingBottom: theme.spacing(10),
  [theme.breakpoints.up('lg')]: {
    paddingTop: APP_BAR_DESKTOP + 24,
    paddingLeft: theme.spacing(2),
    paddingRight: theme.spacing(2),
  },
}));

// ----------------------------------------------------------------------

export default function MainLayout() {
  const navigate = useNavigate();

  const [open, setOpen] = useState(false);

  const [isFetching, setIsFetching] = useState(true);
  const [userData, setUserData] = useState();

  const handleFetchUserData = async () => {
    await fetch('http://localhost:1001/auctee/user/profile', {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
    }).then((res) => {
      if (res.status === 200) {
        res.json().then((data) => {
          setUserData(data);
          setIsFetching(false);
        });
      }
      if (res.status === 401) {
        alert('You need to login first');
        setIsFetching(true);
        navigate('/auctee/login', { replace: true });
      }
    });
  };
  useEffect(() => {
    handleFetchUserData();
  }, [isFetching]);

  return !isFetching ? (
    <RootStyle>
      <MainNavbar userData={userData} onOpenSidebar={() => setOpen(true)} />
      <div style={{ margin: 'auto', display: 'flex', flexDirection: 'row' }}>
        <Sidebar userData={userData} />
        <MainStyle userData={userData}>
          <Outlet />
        </MainStyle>
      </div>
    </RootStyle>
  ) : (
    <></>
  );
}
