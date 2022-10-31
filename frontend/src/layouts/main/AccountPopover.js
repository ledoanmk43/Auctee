import { useRef, useState, useContext } from 'react';
import { Link as RouterLink, useNavigate } from 'react-router-dom';
// @mui
import { alpha } from '@mui/material/styles';
import { Box, Divider, Typography, Stack, MenuItem, Avatar, IconButton } from '@mui/material';
// components
import MenuPopover from '../../components/MenuPopover';
// mocks_
import account from '../../API/account';
import { LoginContext } from '../../utils/Context';

// ----------------------------------------------------------------------

const MENU_OPTIONS = [
  {
    label: 'Trang chủ',
    icon: 'eva:home-fill',
    linkTo: '/auctee/home',
  },
  {
    label: 'Hồ sơ',
    icon: 'eva:person-fill',
    linkTo: '/auctee/user/profile',
  },
];

// ----------------------------------------------------------------------

export default function AccountPopover({ userData }) {
  const navigate = useNavigate();
  const anchorRef = useRef(null);

  const { loggedIn, setLoggedIn } = useContext(LoginContext);

  const [open, setOpen] = useState(null);

  const handleOpen = (event) => {
    setOpen(event.currentTarget);
  };
  const handleLogout = async () => {
    if (loggedIn) {
      await fetch('http://localhost:1001/auctee/logout', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
      }).then((res) => {
        if (res.status === 200) {
          setLoggedIn(false);
          navigate('/auctee/login');
        }
      });
    } else {
      navigate('/auctee/login');
    }
    handleClose();
  };

  const handleClose = () => {
    setOpen(null);
  };

  return (
    <>
      <IconButton
        ref={anchorRef}
        onClick={handleOpen}
        sx={{
          p: 0,
          ...(open && {
            '&:before': {
              zIndex: 1,
              content: "''",
              width: '100%',
              height: '100%',
              borderRadius: '50%',
              position: 'absolute',
              bgcolor: (theme) => alpha(theme.palette.grey[900], 0.8),
            },
          }),
        }}
      >
        <Avatar src={userData.avatar} alt="photoURL" />
      </IconButton>

      <MenuPopover
        open={Boolean(open)}
        anchorEl={open}
        onClose={handleClose}
        sx={{
          p: 0,
          mt: 1.5,
          ml: 0.75,
          '& .MuiMenuItem-root': {
            typography: 'body2',
            borderRadius: 0.75,
          },
        }}
      >
        <Box sx={{ my: 1.5, px: 2.5 }}>
          <Typography variant="subtitle2" noWrap>
            {userData.nickname}
          </Typography>
          <Typography variant="body2" sx={{ color: 'text.secondary' }} noWrap>
            {userData.email}
          </Typography>
        </Box>

        <Divider sx={{ borderStyle: 'dashed' }} />

        <Stack sx={{ p: 1 }}>
          {MENU_OPTIONS.map((option) => (
            <MenuItem key={option.label} to={option.linkTo} component={RouterLink} onClick={handleClose}>
              {option.label}
            </MenuItem>
          ))}
        </Stack>

        <Divider sx={{ borderStyle: 'dashed' }} />

        <MenuItem onClick={handleLogout} sx={{ m: 1 }}>
          {loggedIn ? 'Đăng xuất' : 'Đăng nhập'}
        </MenuItem>
      </MenuPopover>
    </>
  );
}
