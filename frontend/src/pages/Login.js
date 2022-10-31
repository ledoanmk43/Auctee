import { useEffect, Suspense, useState, useContext } from 'react';
import { Link as RouterLink, useNavigate } from 'react-router-dom';

// @mui
import { styled } from '@mui/material/styles';
import { Card, Link, Container, Typography } from '@mui/material';
// hooks
import useResponsive from '../hooks/useResponsive';
// components
import Page from '../components/Page';
import Logo from '../components/Logo';
// sections
import { LoginForm } from '../sections/auth/login';
import AuthSocial from '../sections/auth/AuthSocial';
import { LoginContext } from '../utils/Context';
// ----------------------------------------------------------------------

const RootStyle = styled('div')(({ theme }) => ({
  [theme.breakpoints.up('md')]: {
    display: 'flex',
  },
}));

const HeaderStyle = styled('header')(({ theme }) => ({
  top: 0,
  zIndex: 9,
  lineHeight: 0,
  width: '100%',
  display: 'flex',
  alignItems: 'center',
  position: 'absolute',
  padding: theme.spacing(3),
  justifyContent: 'space-between',
  [theme.breakpoints.up('md')]: {
    alignItems: 'flex-start',
    padding: theme.spacing(7, 5, 0, 7),
  },
}));

const SectionStyle = styled(Card)(({ theme }) => ({
  width: '100%',
  maxWidth: 464,
  display: 'flex',
  flexDirection: 'column',
  justifyContent: 'center',
  margin: theme.spacing(2, 0, 2, 2),
}));

const ContentStyle = styled('div')(({ theme }) => ({
  maxWidth: 480,
  margin: 'auto',
  minHeight: '100vh',
  display: 'flex',
  justifyContent: 'center',
  flexDirection: 'column',
  padding: theme.spacing(12, 0),
}));

// ----------------------------------------------------------------------

export default function Login() {
  const navigate = useNavigate();
  const { loggedIn, setLoggedIn } = useContext(LoginContext);

  const smUp = useResponsive('up', 'sm');

  const mdUp = useResponsive('up', 'md');

  const [isShown, setIsShown] = useState(false);

  useEffect(() => {
    const interval = setInterval(() => {
      if (loggedIn) {
        navigate('/auctee/home');
      } else {
        setIsShown(true);
      }
    }, 100);

    return () => clearInterval(interval);
  }, [loggedIn, isShown]);

  return (
    <>
      {isShown && (
        <Suspense fallback={<></>}>
          <Page title="Đăng nhập">
            <RootStyle>
              <HeaderStyle>
                <Logo />

                {smUp && (
                  <Typography variant="body2" sx={{ mt: { md: -2 } }}>
                    Chưa có tài khoản?
                    <Link variant="subtitle2" component={RouterLink} to="/auctee/register">
                      Đăng ký ngay
                    </Link>
                  </Typography>
                )}
              </HeaderStyle>

              {mdUp && (
                <SectionStyle>
                  <Typography variant="h3" sx={{ px: 5, mt: 10, mb: 5 }}>
                    Hi, Welcome Back
                  </Typography>
                  <img src="/static/illustrations/illustration_login.png" alt="login" />
                </SectionStyle>
              )}

              <Container maxWidth="sm">
                <ContentStyle>
                  <Typography variant="h4" gutterBottom>
                    Đăng nhập vào Auctee
                  </Typography>

                  <Typography sx={{ color: 'text.secondary', mb: 5 }}>Vui lòng điền thông tin</Typography>

                  <AuthSocial />

                  <LoginForm />

                  {!smUp && (
                    <Typography variant="body2" align="center" sx={{ mt: 3 }}>
                      Chưa có tài khoản?
                      <Link variant="subtitle2" component={RouterLink} to="/register">
                        Đăng ký ngay
                      </Link>
                    </Typography>
                  )}
                </ContentStyle>
              </Container>
            </RootStyle>
          </Page>
        </Suspense>
      )}
    </>
  );
}
