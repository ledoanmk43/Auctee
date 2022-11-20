import PropTypes from 'prop-types';
import { useEffect, useState } from 'react';
import { Link as RouterLink, useLocation } from 'react-router-dom';
// material
import { styled } from '@mui/material/styles';
import { Box, Link, Drawer, Typography, Avatar, Divider, useTheme, LinearProgress, Stack } from '@mui/material';
import { Icon } from '@iconify/react';
// mock
import account from '../../API/account';
// hooks
import useResponsive from '../../hooks/useResponsive';
// components
import Logo from '../../components/Logo';
import Scrollbar from '../../components/Scrollbar';
import NavSection from '../../components/NavSection';
//
import navConfig from './NavConfig';

// ----------------------------------------------------------------------

const DRAWER_WIDTH = 280;

const RootStyle = styled('div')(({ theme }) => ({
  flexShrink: 0,
  position: 'relative',
}));

const AccountStyle = styled('div')(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  padding: theme.spacing(1, 2.5),
  borderRadius: Number(theme.shape.borderRadius) * 1.5,
}));

// ----------------------------------------------------------------------

export default function DashboardSidebar({ userData }) {
  const theme = useTheme();
  const location = useLocation();

  const isDesktop = useResponsive('up', 'lg');

  const [isDisplay, setIsDislay] = useState(false);
  useEffect(() => {
    if (location.pathname.includes('user')) {
      setIsDislay(true);
    } else {
      setIsDislay(false);
    }
  }, [location, isDisplay]);

  const renderContent = (
    <Scrollbar>
      <Box sx={{ mt: 14.5, mb: 2, mx: 2.5 }}>
        <Link underline="none" component={RouterLink} to="#">
          <AccountStyle>
            <Avatar sx={{ width: 56, height: 56 }} src={userData.avatar} alt="photoURL" />
            <Box sx={{ ml: 2 }}>
              <Typography variant="subtitle1" sx={{ color: 'text.primary' }}>
                {userData.nickname}
              </Typography>
              <LinearProgress color="success" variant="determinate" value={userData.honor_point} />

              <Typography variant="caption" sx={{ color: 'black', position: 'relative' }}>
                Điểm uy tín : {userData.honor_point} &nbsp;
              </Typography>

              <Box sx={{ display: 'flex' }}>
                <Icon fontSize={'1rem'} icon="bx:edit" color={theme.palette.text.secondary} />
                <Typography variant="caption" sx={{ ml: 0.5, color: 'text.secondary' }}>
                  Sửa Hồ Sơ
                </Typography>
              </Box>
            </Box>
          </AccountStyle>
        </Link>
      </Box>
      <Divider sx={{ ml: '15%', width: '70%' }} />
      <NavSection navConfig={navConfig} />
    </Scrollbar>
  );

  return (
    <>
      {isDisplay && (
        <RootStyle>
          {isDesktop && (
            <Drawer
              open
              variant="permanent"
              PaperProps={{
                sx: {
                  width: DRAWER_WIDTH,
                  bgcolor: 'background.default',
                  borderRightStyle: 'none',
                  zIndex: 0,
                  position: 'relative',
                },
              }}
            >
              {renderContent}
            </Drawer>
          )}
        </RootStyle>
      )}
    </>
  );
}
